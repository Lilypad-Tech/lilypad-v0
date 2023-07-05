// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Snapshot.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/utils/math/SafeCast.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import "./DataGovernanceToken.sol";
import "./LilypadCallerInterface.sol";
import "./LilypadEventsUpgradeable.sol";

contract GovernanceToken is ERC20Snapshot, Ownable {
    constructor() ERC20("GovernanceToken", "GOV") {
        _mint(msg.sender, 1000000000000000000000000);
    }

    function snapshot() public onlyOwner returns (uint256) {
        return _snapshot();
    }

    function mint(address wallet, uint256 amount) public onlyOwner {
        _mint(wallet, amount);
    }
}

// This sample contract is a timelocked off-chain voting Governor contract that uses Lilypad to
// resolve the timelock encrypted off-chain votes with drand, calculate the vote resolution and rewards off-chain
// and publish the results on-chain.
// This contract works together with the docker image in examples/drand-timelock-voting
// The contract is based on the Governor contract from DeFiKicks(https://github.com/md0x/defikicks)

contract GovernorContract is Context, LilypadCallerInterface, Ownable {
    GovernanceToken public token;

    uint256 public votingPeriod;
    uint256 public quorumPercentage;
    uint256 public emissionPerVote = 1 * 10 ** 16;
    string public dockerImage = "your_drand_timelock_voting_image";

    LilypadEventsUpgradeable bridge;

    event ProposalExecuted(bytes32 indexed proposalId);

    event ProposalCreated(
        bytes32 indexed proposalId,
        address proposer,
        address[] targets,
        uint256[] values,
        bytes[] calldatas,
        uint256 voteStart,
        uint256 voteEnd,
        string description,
        bytes32 descriptionHash
    );

    event VoteResolutionRequested(bytes32 indexed proposalId, uint256 bridgeId);

    event ProposalUpdated(
        bytes32 indexed proposalId,
        bytes32 voteMerkleRoot,
        uint256 forVotes,
        uint256 againstVotes,
        uint256 abstainVotes,
        string data
    );

    event ClaimedReward(
        address indexed user,
        uint256 amount,
        bytes32 indexed proposalId
    );

    enum ProposalState {
        Pending,
        Active,
        Canceled,
        Defeated,
        Succeeded,
        Queued,
        Expired,
        Executed,
        ResolutionToRequest,
        ResolutionRequested
    }

    struct ProposalCore {
        uint64 voteStart;
        address proposer;
        uint64 voteEnd;
        bool executed;
        bool canceled;
        uint256 forVotes;
        uint256 againstVotes;
        uint256 abstainVotes;
        uint256 bridgeId;
        bytes32 voteMerkleRoot;
        uint256 snapshotId;
    }

    struct ResolutionResponse {
        uint256 forVotes;
        uint256 againstVotes;
        uint256 abstainVotes;
        bytes32 voteMerkleRoot;
        string data;
    }

    string constant specStart =
        "{"
        '"Engine": "docker",'
        '"Verifier": "noop",'
        '"PublisherSpec": {"Type": "estuary"},'
        '"Docker": {'
        '"Image": "';

    string constant specMiddle =
        '",'
        '"EnvironmentVariables": ["PROPOSAL_ID=';

    string constant specEnd =
        '","NODE_URL=https://api.calibration.node.glif.io/rpc/v0"]},'
        '"Language":{"JobContext":{}},'
        '"Wasm":{"EntryModule":{}},'
        '"Resources":{"GPU":""},'
        '"Deal": {"Concurrency": 1}'
        "}";

    mapping(bytes32 => ProposalCore) public proposals;
    mapping(uint256 => bytes32) public jobIdToProposal;
    mapping(bytes32 => mapping(address => bool)) public alreadyClaimed;

    constructor(address _token, address bridgeContract) {
        bridge = LilypadEventsUpgradeable(bridgeContract);
        token = GovernanceToken(_token);
    }

    // Setters with only owner

    function setEmissionPerVote(uint256 _emissionPerVote) external onlyOwner {
        emissionPerVote = _emissionPerVote;
    }

    function setVotingPeriod(uint256 period) external onlyOwner {
        votingPeriod = period;
    }

    function setQuorumPercentage(uint256 percentage) external onlyOwner {
        quorumPercentage = percentage;
    }

    function setDockerImage(string memory image) external onlyOwner {
        dockerImage = image;
    }

    function claimReward(
        bytes32 proposalId,
        uint256 amount,
        bytes32[] memory merkleProof
    ) external {
        bytes32 leaf = keccak256(
            bytes.concat(keccak256(abi.encode(msg.sender, amount)))
        );

        require(
            MerkleProof.verify(
                merkleProof,
                proposals[proposalId].voteMerkleRoot,
                leaf
            ),
            "Invalid proof"
        );

        require(
            !alreadyClaimed[proposalId][msg.sender],
            "Reward already claimed"
        );

        alreadyClaimed[proposalId][msg.sender] = true;

        token.mint(msg.sender, amount);

        emit ClaimedReward(msg.sender, amount, proposalId);
    }

    function requestVoteResolution(bytes32 proposalId) public payable virtual {
        require(
            state(proposalId) == ProposalState.ResolutionToRequest,
            "Governor: vote not in ResolutionToRequest state"
        );
        uint256 lilypadFee = bridge.getLilypadFee();
        require(msg.value >= lilypadFee, "Governor: insufficient fee");

        string memory spec = getSpecForProposalId(proposalId);

        uint256 id = bridge.runLilypadJob{value: lilypadFee}(
            address(this),
            spec,
            uint8(LilypadResultType.StdOut)
        );
        require(id > 0, "job didn't return a value");
        proposals[proposalId].bridgeId = id;
        jobIdToProposal[id] = proposalId;
        emit VoteResolutionRequested(proposalId, id);
    }

    function propose(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        string memory description
    ) public virtual returns (bytes32) {
        address proposer = _msgSender();
        require(
            _isValidDescriptionForProposer(proposer, description),
            "Governor: proposer restricted"
        );

        uint256 currentTimepoint = clock();

        bytes32 proposalId = hashProposal(
            targets,
            values,
            calldatas,
            keccak256(bytes(description))
        );

        require(
            targets.length == values.length,
            "Governor: invalid proposal length"
        );
        require(
            targets.length == calldatas.length,
            "Governor: invalid proposal length"
        );
        require(targets.length > 0, "Governor: empty proposal");
        require(
            proposals[proposalId].voteStart == 0,
            "Governor: proposal already exists"
        );

        uint256 snapshot = currentTimepoint;
        uint256 deadline = snapshot + votingPeriod;

        proposals[proposalId] = ProposalCore({
            proposer: proposer,
            voteStart: SafeCast.toUint64(snapshot),
            voteEnd: SafeCast.toUint64(deadline),
            executed: false,
            canceled: false,
            forVotes: 0,
            againstVotes: 0,
            abstainVotes: 0,
            bridgeId: 0,
            voteMerkleRoot: bytes32(0),
            snapshotId: token.snapshot()
        });

        emit ProposalCreated(
            proposalId,
            proposer,
            targets,
            values,
            calldatas,
            snapshot,
            deadline,
            description,
            keccak256(bytes(description))
        );

        return proposalId;
    }

    function execute(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 descriptionHash
    ) public payable virtual returns (bytes32) {
        bytes32 proposalId = hashProposal(
            targets,
            values,
            calldatas,
            descriptionHash
        );

        ProposalState currentState = state(proposalId);

        require(
            currentState == ProposalState.Succeeded ||
                currentState == ProposalState.Queued,
            "Governor: proposal not successful"
        );

        proposals[proposalId].executed = true;

        emit ProposalExecuted(proposalId);

        _execute(proposalId, targets, values, calldatas, descriptionHash);

        return proposalId;
    }

    // Lilypad callbacks
    function lilypadFulfilled(
        address _from,
        uint _jobId,
        LilypadResultType _resultType,
        string calldata _result
    ) external override {
        require(_resultType == LilypadResultType.StdOut);
        require(msg.sender == address(bridge));

        ResolutionResponse memory resolutionResponse = abi.decode(
            hexStrToBytes(_result),
            (ResolutionResponse)
        );
        ProposalCore storage proposal = proposals[jobIdToProposal[_jobId]];
        proposal.voteMerkleRoot = resolutionResponse.voteMerkleRoot;
        proposal.forVotes = resolutionResponse.forVotes;
        proposal.againstVotes = resolutionResponse.againstVotes;
        proposal.abstainVotes = resolutionResponse.abstainVotes;

        emit ProposalUpdated(
            jobIdToProposal[_jobId],
            resolutionResponse.voteMerkleRoot,
            resolutionResponse.forVotes,
            resolutionResponse.againstVotes,
            resolutionResponse.abstainVotes,
            resolutionResponse.data
        );
    }

    function lilypadCancelled(
        address _from,
        uint256 _jobId,
        string calldata _errorMsg
    ) external override {
        require(_from == address(bridge));
        proposals[jobIdToProposal[_jobId]].bridgeId = 0;
    }

    // Getters
    function getSpecForProposalId(
        bytes32 proposalId
    ) public view virtual returns (string memory) {
        return
            string.concat(
                specStart,
                dockerImage,
                specMiddle,
                Strings.toHexString(uint256(proposalId)),
                specEnd
            );
    }

    function clock() public view virtual returns (uint48) {
        return SafeCast.toUint48(block.timestamp);
    }

    function getLilypadFee() public view virtual returns (uint256) {
        return bridge.getLilypadFee();
    }

    function proposalSnapshot(
        bytes32 proposalId
    ) public view virtual returns (uint256) {
        return proposals[proposalId].voteStart;
    }

    function proposalDeadline(
        bytes32 proposalId
    ) public view virtual returns (uint256) {
        return proposals[proposalId].voteEnd;
    }

    function state(
        bytes32 proposalId
    ) public view virtual returns (ProposalState) {
        ProposalCore storage proposal = proposals[proposalId];

        if (proposal.executed) {
            return ProposalState.Executed;
        }

        if (proposal.canceled) {
            return ProposalState.Canceled;
        }

        uint256 snapshot = proposalSnapshot(proposalId);

        if (snapshot == 0) {
            revert("Governor: unknown proposal id");
        }

        uint256 currentTimepoint = clock();

        if (snapshot >= currentTimepoint) {
            return ProposalState.Pending;
        }

        uint256 deadline = proposalDeadline(proposalId);

        if (deadline >= currentTimepoint) {
            return ProposalState.Active;
        }

        if (proposals[proposalId].bridgeId == 0) {
            return ProposalState.ResolutionToRequest;
        }

        if (proposals[proposalId].voteMerkleRoot == bytes32(0)) {
            return ProposalState.ResolutionRequested;
        }

        if (_quorumReached(proposalId) && _voteSucceeded(proposalId)) {
            return ProposalState.Succeeded;
        } else {
            return ProposalState.Defeated;
        }
    }

    function hashProposal(
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 descriptionHash
    ) public pure virtual returns (bytes32) {
        return
            keccak256(abi.encode(targets, values, calldatas, descriptionHash));
    }

    // Internal
    function _quorumReached(
        bytes32 proposalId
    ) internal view virtual returns (bool) {
        return
            proposals[proposalId].forVotes >=
            (quorumPercentage *
                token.totalSupplyAt(proposals[proposalId].snapshotId)) /
                1 ether;
    }

    function _voteSucceeded(
        bytes32 proposalId
    ) internal view virtual returns (bool) {
        return
            proposals[proposalId].forVotes > proposals[proposalId].againstVotes;
    }

    function _execute(
        bytes32 /* proposalId */,
        address[] memory targets,
        uint256[] memory values,
        bytes[] memory calldatas,
        bytes32 /*descriptionHash*/
    ) internal virtual {
        string memory errorMessage = "Governor: call reverted without message";
        for (uint256 i = 0; i < targets.length; ++i) {
            (bool success, bytes memory returndata) = targets[i].call{
                value: values[i]
            }(calldatas[i]);
            Address.verifyCallResult(success, returndata, errorMessage);
        }
    }

    function _tryHexToUint(bytes1 char) private pure returns (bool, uint8) {
        uint8 c = uint8(char);
        unchecked {
            // Case 0-9
            if (47 < c && c < 58) {
                return (true, c - 48);
            }
            // Case A-F
            else if (64 < c && c < 71) {
                return (true, c - 55);
            }
            // Case a-f
            else if (96 < c && c < 103) {
                return (true, c - 87);
            }
            // Else: not a hex char
            else {
                return (false, 0);
            }
        }
    }

    function _isValidDescriptionForProposer(
        address proposer,
        string memory description
    ) internal view virtual returns (bool) {
        uint256 len = bytes(description).length;

        // Length is too short to contain a valid proposer suffix
        if (len < 52) {
            return true;
        }

        // Extract what would be the `#proposer=0x` marker beginning the suffix
        bytes12 marker;
        assembly {
            // - Start of the string contents in memory = description + 32
            // - First character of the marker = len - 52
            //   - Length of "#proposer=0x0000000000000000000000000000000000000000" = 52
            // - We read the memory word starting at the first character of the marker:
            //   - (description + 32) + (len - 52) = description + (len - 20)
            // - Note: Solidity will ignore anything past the first 12 bytes
            marker := mload(add(description, sub(len, 20)))
        }

        // If the marker is not found, there is no proposer suffix to check
        if (marker != bytes12("#proposer=0x")) {
            return true;
        }

        // Parse the 40 characters following the marker as uint160
        uint160 recovered = 0;
        for (uint256 i = len - 40; i < len; ++i) {
            (bool isHex, uint8 value) = _tryHexToUint(bytes(description)[i]);
            // If any of the characters is not a hex digit, ignore the suffix entirely
            if (!isHex) {
                return true;
            }
            recovered = (recovered << 4) | value;
        }

        return recovered == uint160(proposer);
    }

    // Helper functions
    function hexStrToBytes(
        string memory _hexStr
    ) public pure returns (bytes memory) {
        bytes memory strBytes = bytes(_hexStr);

        // Check for '0x' prefix
        uint offset;
        if (
            strBytes.length >= 2 &&
            strBytes[0] == bytes1("0") &&
            strBytes[1] == bytes1("x")
        ) {
            offset = 2;
        }

        require(
            (strBytes.length - offset) % 2 == 0,
            "Invalid hex string length!"
        );

        bytes memory result = new bytes((strBytes.length - offset) / 2);

        for (uint i = offset; i < strBytes.length; i += 2) {
            uint8 upper = charToUint8(uint8(strBytes[i]));
            uint8 lower = charToUint8(uint8(strBytes[i + 1]));

            result[(i - offset) / 2] = bytes1((upper << 4) | lower);
        }

        return result;
    }

    function charToUint8(uint8 c) private pure returns (uint8) {
        if (c >= 48 && c <= 57) {
            return c - 48;
        }
        if (c >= 97 && c <= 102) {
            return 10 + c - 97;
        }
        if (c >= 65 && c <= 70) {
            return 10 + c - 65;
        }

        revert("Invalid hex char!");
    }
}
