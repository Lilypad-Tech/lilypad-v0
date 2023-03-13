# 🌺🐸 lilypad
A simple and lightweight FVM "oracle" using Bacalhau. We're enabling calling Bacalhau jobs from smart contracts! 

This project harnesses the power of onchain compute with off-chain computations including tradional Oracle jobs, AI, ML and well any Docker or WASM job you can run on Bacalhau.

The work is primarily initially focussed on FEVM contract target.


![Project Lilypad in 3 mins (2)](https://user-images.githubusercontent.com/12529822/223378567-91e08ae4-9859-441d-bbfe-d1b7516c6543.png)


**YouTube Video Link**


[![YouTube Video](http://img.youtube.com/vi/9lF7omNEK-c/0.jpg)](https://www.youtube.com/watch?v=9lF7omNEK-c "Project Lilypad")


### We're live on FVM Hyperspace

Contract Address: `0x75B01cAeFF294f10d21ff3429C683230e3d8C9B6`

See it on Block Explorer [here](https://fvm.starboard.ventures/contracts/0x75B01cAeFF294f10d21ff3429C683230e3d8C9B6) 

Open it in Remix [here](https://remix.ethereum.org/bacalhau-project/lilypad/blob/main/hardhat/contracts/LilypadEvents.sol)

See the examples folder for how to use this in your own contracts!

[Read more here](https://bit.ly/project-lilypad)

Get help: [FilecoinProject Slack](https://filecoinproject.slack.com/) #bacalhau or #bacalhau-lilypad channel 

### How Stuff Works
![image (15)](https://user-images.githubusercontent.com/12529822/224299570-366bde1c-1f48-4af9-9d7c-0d4f8a0fc1fc.png)


### How do I get started on using Lilypad in my project?🧑‍💻

1.  Create a contract that implements [`LilypadCallerInterface`](./hardhat/contracts/LilypadCallerInterface.sol). As part of this interface you need to implement 2 functions:

    *   `lilypadFulfilled` - a callback function that will be called when the job completes successfully 
    *   `lilypadCancelled` - a callback function that will be called when the job fails

2.  To trigger a job from your contract, you need to call our `LilypadEvents` contract which the bridge is listening to. You will connect to Bacalhau network via this bridge. Create an instance of [`LilypadEvents`](./hardhat/contracts/LilypadEvents.sol) by passing the public contract address above to the `LilypadEvents` constructor. See our [example](./examples/contracts/StableDiffusionCaller.sol#L29). 
3.  To make a call to Bacalhau, call `runBacalhauJob` from your function. You need to pass the following parameters: 
    
    | Name | Type | Purpose |
    |:---:|:---:|:---:|
    | `_from` | `address` | The address of the calling contract, to which success or failure will be passed back. You should probably use address(this) from your contract. |
    | `_spec` | `string` | A Bacalhau job spec in JSON format. See below for more information on creating a job spec. |
    | `_resultType` | [`LilypadResultType`](./hardhat/contracts/LilypadCallerInterface.sol#L4-L9) | The type of result that you want to be returned. If you specify CID, the result tree will come back as a retrievable IPFS CID. If you specify StdOut, StdErr or ExitCode, those raw values output from the job will be returned. |

### What do I need to know to run Bacalhau?

*   Bacalhau is language-agnostic, and supports [Docker](https://docs.bacalhau.org/getting-started/docker-workload-onboarding) or [WASM](https://docs.bacalhau.org/getting-started/wasm-workload-onboarding) workloads. As long as you can run your executable in a container, you can run it in Bacalhau.
*   You need to supply a Bacalhau job spec. To create a job spec, you can:
    *   Run a Bacalhau job successfully, and then get the job spec back using `bacalhau describe <job_id> --format=json`.
    *   Generate a job spec without running anything, using `bacalhau docker run --dry-run`.
    *   Writing a job spec by hand, by using our [schema](https://schema.bacalhau.org) as a guide.
*   What can I do with Bacalhau now? You can:
    *   read from IPFS, Filecoin, or URLs
    *   write into Estuary or IPFS   


### Any Example Jobs?

We have a full complement of example jobs you can leverage on the [Bacalhau Docs Site](https://docs.bacalhau.org/)

Try out
- YOLO
- OCR
- Video Editing
- and many, many more!

### More Resources

See a video of this project in action [here](https://youtu.be/B0l0gFYxADY)

Read more and see the presentation slides [here](https://bit.ly/project-lilypad)

### See our Example Project

We've created Waterlily - an AI-Art generator that pays royalties to artists (or donations to art foundations for public works).

See it live in action [waterlily.ai](https://www.waterlily.ai/)

![Screenshot 2023-03-14 at 10 14 23 am](https://user-images.githubusercontent.com/12529822/224852799-594fd941-be82-4b7e-b7cd-2ba306857243.png)
 

### Get in touch

We'd also LOVE to hear about what use cases you have - contact us in the [FilecoinProject Slack](https://filecoinproject.slack.com/) #bacalhau or #bacalhau-lilypad channel

Contributions to this repo would also earn you many 5/5 Stable Diffusion Rainbow Unicorns (and be gratefully recieved)!  

### Thanks for your interest!

❤️ Bacalhau team: off-chain filecoin-native decentralised compute ❤️

#buildwithbacalhau

![image](https://user-images.githubusercontent.com/12529822/220625332-b0e6a08a-b77d-41f7-90a8-248852a353c8.png)

