// SPDX-License-Identifier: MIT

// This is a FVM smart contract which allows users to check if a particular EthereumAccount is Fraudulent

// Layout of Contract:
// version
// imports
// errors
// interfaces, libraries, contracts
// Type declarations
// State variables
// Events
// Modifiers
// Functions

// Layout of Functions:
// constructor
// receive function (if exists)
// fallback function (if exists)
// external
// public
// internal
// private
// view & pure functions
pragma solidity >=0.8.10 <0.9.0;
import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "https://github.com/bacalhau-project/lilypad/blob/main/hardhat/contracts/LilypadEventsUpgradeable.sol";
import "https://github.com/bacalhau-project/lilypad/blob/main/hardhat/contracts/LilypadCallerInterface.sol";

/*
 * @title FraudDetector
 * @author Kazeem Hakeem
 *
 * This is the contract meant to check if particular Ethereum Account is Fraudulent using Machine Learning
 * The Ethereum Account is passes through Lilypad and bacalhau for computation
 */
contract FraudDetector is LilypadCallerInterface, Ownable {
    ///////////////////
    // State Variables
    ///////////////////
    uint256 public lilypadFee;
    string private constant _TABLE_PREFIX = "fraud_detect";
    string constant specStart =
        "{"
        '"Engine": "docker",'
        '"Verifier": "noop",'
        '"Network": {'
        '"Type": "HTTP",'
        '"Domains": ["api.etherscan.com"]},'
        '"PublisherSpec": {"Type": "estuary"},'
        '"Docker": {'
        '"Image": "hakymulla/hackfs:inference",'
        '"Entrypoint": ["python", "predict.py",  "--address", "';

    string constant specEnd =
        '"]},'
        '"Resources": {"CPU": "1"},'
        '"Outputs": [{"Name": "outputs", "Path": "/outputs"}],'
        '"Deal": {"Concurrency": 1}'
        "}";

    address public bridgeAddress;

    LilypadEventsUpgradeable bridge;

    struct Results {
        address caller;
        string prompt;
        string standardOutput;
    }
    Results[] public result;

    mapping(uint256 => string) prompts;

    ///////////////////
    // Events
    ///////////////////
    event NewResult(Results image);

    ///////////////////
    // Modifiers
    ///////////////////
    modifier onlyBridge() {
        require(msg.sender == bridgeAddress, "Not owner");
        _;
    }

    ///////////////////
    // Functions
    ///////////////////
    /*
     * constructor
     * @_bridgeContractAddress : The LilypadEvents Contract Address
     * @notice This function will also create a TableLand Table on the current network
     */
    constructor(address _bridgeContractAddress) {
        console.log("Deploying Fraud Detection contract");
        bridgeAddress = _bridgeContractAddress;
        bridge = LilypadEventsUpgradeable(_bridgeContractAddress);
        uint fee = bridge.getLilypadFee();
        lilypadFee = fee;
    }

    ///////////////////
    // External Functions
    ///////////////////
    /*
     * @param _prompt: The ethereum address to run through the ML algorith
     * @notice Requires a fee payment to cover gas cost of the callback
     */
    function IsFraudDetector(string memory _prompt) external payable {
        require(msg.value >= lilypadFee, "Not enough to run Lilypad job");
        string memory spec = string(
            abi.encodePacked(specStart, _prompt, specEnd)
        );

        uint id = bridge.runLilypadJob{value: lilypadFee}(
            address(this),
            spec,
            uint8(LilypadResultType.StdOut)
        );

        require(id > 0, "job didn't return a value");
        prompts[id] = _prompt;
    }

    function getLilypadFee() external {
        uint fee = bridge.getLilypadFee();
        console.log("fee", fee);
        lilypadFee = fee;
    }

    function lilypadFulfilled(
        address _from,
        uint _jobId,
        LilypadResultType _resultType,
        string calldata _result
    ) external override {
        //need some checks here that it a legitimate result
        require(_from == address(bridge)); //really not secure
        require(_resultType == LilypadResultType.StdOut);

        Results memory res = Results({
            caller: msg.sender,
            standardOutput: _result,
            prompt: prompts[_jobId]
        });
        result.push(res);
        emit NewResult(res);
        delete prompts[_jobId];
    }

    function lilypadCancelled(
        address _from,
        uint256 _jobId,
        string calldata _errorMsg
    ) external override {
        require(_from == address(bridge));
        console.log(_errorMsg);
        delete prompts[_jobId];
    }

    ///////////////////
    // Public Functions
    ///////////////////
    function setBridgeAddress(address _newAddress) public onlyOwner {
        bridgeAddress = _newAddress;
    }

    function setLPEventsAddress(address _eventsAddress) public onlyOwner {
        bridge = LilypadEventsUpgradeable(_eventsAddress);
    }

    function setLilypadFee(uint256 _fee) public onlyOwner {
        require(_fee > 0, "Lilypad fee must be greater than 0");
        lilypadFee = _fee;
    }

    ////////////////////////////////////////////////////////////////////////////
    ////////////////////////////////////////////////////////////////////////////
    // External & Public View & Pure Functions
    ////////////////////////////////////////////////////////////////////////////
    ////////////////////////////////////////////////////////////////////////////
    function allResult() public view returns (Results[] memory) {
        return result;
    }
}
