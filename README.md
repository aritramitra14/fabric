
# TimeFabric
This is a fork of Hyperledger Fabric 1.4 (https://github.com/hyperledger/fabric) repository and contains the implementation of TimeFabric.
TimeFabric is a proof of concept to run time based smart contracts in Fabric. Instructions are provided below to run TimeFabic using binaries.


# Instructions 

- The Hyperledger Fabric 1.4.10 prerequisites should be installed
- Run `make peer` command to generate the peer binary from Makefile. The binary will be generated in `./build/bin` folder in the local repo
- Generate a new docker image from `fabric-ccenv` for chaincode compilation
	- A sample dockerfile for generating new docker image for `fabric-ccenv` has been provided
	- Run `docker build .` command to generate the new docker image and rename it to `hyperledger/fabric-ccenv:latest`
- Sample chaincodes are provided in `/timefabric/chaincodes`
- Sample client app for connecting to network are provided in `/timefabric/javaclient` 


