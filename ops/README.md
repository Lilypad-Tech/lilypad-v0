# Deploying Lilypad

`lilypad` just runs on a normal Google Compute Cloud VM as a systemd service. To deploy a new version of `lilypad` or change the contract that it is listening to, you only need to upload a new version of the binary and environment variables and then restart the service.

If you want to deploy lilypad to a new VM, there is a Terraform file you can use to provision a VM, and then follow the deploy instructions for updating lilypad versions.

## Updating the Lilypad version

Make sure you are logged in to `gcloud` in your terminal, and then just run `make deploy` from the root of this repository!

The `deploy.sh` script will connect to a Google Cloud instance called `lilypad-vm-0` and copy upwards the new binary version, the lilypad.service file, and the `hardhat/.env` file which should contain the private key and deployed contract address to listen to (`make` will complain if these are not present). It will then start or restart the lilypad service.

## Provisioning a new Lilypad VM

Note that Lilypad is only currently designed to have one listener per contract. Jobs and return calls will be duplicated if multiple lilypads are listening to the same contract.

```bash
cd ops
gcloud auth login
terraform apply
```

## Checking that lilypad is running

```bash
gcloud compute ssh --zone=us-central1-a --project=bacalhau-production lilypad-vm-0 -- journalctl --unit=lilypad -f
```
