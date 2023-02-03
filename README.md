# NFT MetaData API with Cloud Functions
## Cloud Functions
Since this api only provides metadata of nft, it's better to use cloud functions or other serverless service.

This repo is using existing tokenURI to search the metadata, but in practiacal case, backend should store metadata for all nfts or save them into the database. The target function should filter if the token id is valid before getting the metadata

During the initialization, we get the enviornment variables by [Viper](https://github.com/spf13/viper), using a struct to wrap all env variables.

In the target function, firstly, retrieving the token id from request url, then checking if the token id is valid or not. To obtain the on-chain data, we use go-ethereum without doubt.

### Local Test
You should test locally before deploying on gcp. First, we need to set up the target function name with below command:
```bash
export FUNCTION_TARGET="YOUR FUNCTION NAME"
```
Then, run `go mod tidy` to update dependencies
Lastly, you can run the local server
```bash
$ go run cmd/main.go
```
Test your code with curl on new terminal
```bash
$ curl localhost:8080/{tokenId} -X GET
```

### Deploy
Usually, using CI / CD to manage deployment is the best implementation, but in this case, we deploy cloud functions locally, so it saves works by writing a bash script.

If you counter permission denied issue, change the permission of deploy.sh by below command
```bash
$ sudo chmod +x deploy.sh
```

## Go-Ethereum
We only read the on-chain data, so there's no need to set private key on env variables. The nfts of CloneX are all minted, so we use `TotalSupply()`. However, if the project is during sale stage, you should replace this function with `CurrentTokenId()` or other function to check the latest minted token id to prevent nft metadata leaking issue.

If you don't know how to generate wrapped contract in go, please refer to [contract binding](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings#what-is-an-abi).
