import { StandardMerkleTree } from "@openzeppelin/merkle-tree";
import { HttpCachingChain, HttpChainClient } from "drand-client";
import { BigNumber, ethers } from "ethers";
import { timelockDecrypt } from "tlock-js";
import TokenABI from "./abi/GovernanceToken.json" assert { type: "json" };
import GovernorABI from "./abi/GovernorContract.json" assert { type: "json" };

export const timelockDecryption = async (ciphertext) => {
  const fastestNodeClient = await getFastestNode();
  const result = await timelockDecrypt(ciphertext, fastestNodeClient);
  return result;
};

const addresses = {
  token: "0xd61801Fa782da077873394E75cc7740EaA2809A5",
  governor: "0xcF53746114e9b1B0Dd0900CDd76b39D77d14bb86",
  registry: "0xc799246Ec502b51e34c975cAd3CD7541a85DA9F6",
};

const getVotes = async (proposalId) => {
  // TODO implement your own logic to get the votes from IPFS or any other source
  return [
    {
      account: "0x000000",
      cyphertext: "0x000000",
      signature: "0x000000",
    },
  ];
};

const testnetUnchainedUrl =
  "https://pl-eu.testnet.drand.sh/7672797f548f3f4748ac4bf3352fc6c6b6468c9ad40ad456a397545c6e2df5bf";

const getFastestNode = async () => {
  const chain = new HttpCachingChain(testnetUnchainedUrl);
  const client = new HttpChainClient(chain);

  return client;
};

async function run() {
  const proposalId = process.env.PROPOSAL_ID;
  const nodeUrl = process.env.NODE_URL;

  const provider = new ethers.providers.JsonRpcProvider(nodeUrl);

  const governorContract = new ethers.Contract(
    addresses.governor,
    GovernorABI,
    provider
  );

  const proposalStruct = await governorContract.proposals(proposalId);

  const votes = await getVotes(proposalStruct.id);

  const decryptedVotes = [];

  // Decrypt timelocked votes using drand
  for (const vote of votes) {
    let message;
    try {
      message = await timelockDecryption(vote.cyphertext);
    } catch (e) {}
    decryptedVotes.push({
      ...vote,
      message,
    });
  }

  // Verify signatures
  const signedVotes = [];
  for (const vote of decryptedVotes) {
    const recoveredAddress = ethers.utils.verifyMessage(
      vote.message,
      vote.signature
    );
    if (recoveredAddress === vote.account) {
      signedVotes.push(vote);
    }
  }

  const tokenContract = new ethers.Contract(
    addresses.token,
    TokenABI,
    provider
  );

  // Count votes
  let forVotes = BigNumber.from(0);
  let againstVotes = BigNumber.from(0);
  let abstainVotes = BigNumber.from(0);
  for (const vote of signedVotes) {
    const balance = await tokenContract.balanceOfAt(
      vote.account,
      proposalStruct.snapshotId.toString()
    );
    const voteObj = JSON.parse(vote.message);

    if (voteObj.vote === "for") {
      forVotes = forVotes.add(balance);
    }
    if (voteObj.vote === "against") {
      againstVotes = againstVotes.add(balance);
    }
  }

  const totalSupplyAtVote = await tokenContract.totalSupplyAt(
    proposalStruct.snapshotId.toString()
  );
  abstainVotes = totalSupplyAtVote.sub(forVotes.add(againstVotes));

  // Reward calculation
  const quorumPercentage = await governorContract.quorumPercentage();
  const majority = forVotes.gt(againstVotes) ? "for" : "against";
  const majorityValue = majority === "for" ? forVotes : againstVotes;
  const arrivedToConsensus =
    !forVotes.eq(againstVotes) &&
    majorityValue.gte(
      totalSupplyAtVote.mul(quorumPercentage).div(ethers.utils.parseEther("1"))
    );

  const toReward = [];
  const emissionPerVote = await governorContract.emissionPerVote();
  if (arrivedToConsensus) {
    for (const vote of signedVotes) {
      const balance = await tokenContract.balanceOfAt(
        vote.account,
        proposalStruct.snapshotId.toString()
      );
      const voteObj = JSON.parse(vote.message);
      if (voteObj.vote == majority) {
        toReward.push([
          vote.account,
          balance
            .mul(emissionPerVote)
            .div(ethers.utils.parseEther("1"))
            .toString(),
        ]);
      }
    }
  } else {
    toReward.push([ethers.constants.AddressZero, "0"]);
  }

  const tree = StandardMerkleTree.of(toReward, ["address", "uint256"]);

  const proofData = {};
  for (const [i, v] of tree.entries()) {
    // (3)
    const proof = tree.getProof(i);

    proofData[v[0]] = {
      amount: v[1],
      proof,
    };
  }
  const data = JSON.stringify(proofData);

  // Encode calldata
  const calldata = ethers.utils.defaultAbiCoder.encode(
    [
      {
        type: "tuple",
        components: [
          { name: "forVotes", type: "uint256" },
          { name: "againstVotes", type: "uint256" },
          { name: "abstainVotes", type: "uint256" },
          { name: "voteMerkleRoot", type: "bytes32" },
          { name: "data", type: "string" },
        ],
      },
    ],
    [
      {
        forVotes,
        againstVotes,
        abstainVotes,
        voteMerkleRoot: tree.root,
        data,
      },
    ]
  );

  // This is the only log that should be done as this will be send to lilypad by the bacalhaul operator
  // by calling the returnLilypadResults function on the lilypad contract with the calldata as parameter
  console.log("calldata", calldata);
}

run();
