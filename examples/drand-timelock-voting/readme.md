### Description

---

This sample container, which offers timelock encryption and off-chain voting features, was taken from the HackFS 2023 finalist [DeFiKicks](https://defikicks.xyz) project. Please visit their [GitHub](https://github.com/md0x/defikicks) repository or the [project page on EthGlobal](https://ethglobal.com/showcase/defikicks-b1wo0) for more details.

### How to build this container

---

This container is built using the following command:

```bash
docker build -t yourName/yourImageName:latest .
```

### How to run this container

---

This container is run using the following command:

```bash
docker run yourName/yourImageName:latest
```

### How to deploy this image to Docker Hub

---

```bash
docker login
export IMAGE=yourName/yourImageName:latest
docker build -t {$IMAGE} .
docker image push {$IMAGE}
```

### How to run with bacalhau

---

```bash
bacalhau docker run \
   --env PROPOSAL_ID={$PROPOSAL_ID} \
   --env NODE_URL={$NODE_URL} \
    {$IMAGE}
```

### How to get the job spec

---

```bash
bacalhau describe <job_id> --json
```
