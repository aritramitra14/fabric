package com.ClientTransaction;
import java.io.IOException;
//import java.nio.charset.StandardCharsets;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Collection;
//import java.util.concurrent.TimeoutException;
import java.util.concurrent.CompletableFuture;

//import org.hyperledger.fabric.gateway.Contract;
//import org.hyperledger.fabric.gateway.ContractException;
//import org.hyperledger.fabric.gateway.Gateway;
import org.hyperledger.fabric.gateway.Network;
import org.hyperledger.fabric.gateway.Wallet;
import org.hyperledger.fabric.gateway.impl.GatewayImpl;
//import org.hyperledger.fabric.gateway.impl.NetworkImpl;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.TransactionProposalRequest;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.BlockEvent;
import org.hyperledger.fabric.sdk.ChaincodeID;

public final class Client {
    public static void main(String[] args) throws IOException {

        // Load an existing wallet holding identities used to access the network.
        Path walletDirectory = Paths.get("wallet");
        Wallet wallet = Wallet.createFileSystemWallet(walletDirectory);

        // Path to a common connection profile describing the network.
        Path networkConfigFile = Paths.get("connection.json");
        String user = "Admin@ghost.uwaterloo.ca";
        
        GatewayImpl.Builder b1 = new GatewayImpl.Builder().identity(wallet, user).networkConfig(networkConfigFile);
        
        //GatewayImpl g1 = b1.connect();
        
       
        // Configure the gateway connection used to access the network.
       /* Gateway.Builder builder = Gateway.createBuilder()
                .identity(wallet, user)
                .networkConfig(networkConfigFile);*/
        

        // Create a gateway connection
        try (GatewayImpl gateway = b1.connect()) {

            // Obtain a smart contract deployed on the network.
            Network network = gateway.getNetwork("mychannel");
            //Contract contract = network.getContract("sacc");
            ChaincodeID.Builder chcode = ChaincodeID.newBuilder();
            //Chaincode.ChaincodeID.Builder chcode = org.hyperledger.fabric.protos.peer.Chaincode.ChaincodeID.newBuilder();
            chcode = chcode.setName("sacc");
            chcode= chcode.setVersion("1.0");
            //chcode = chcode.setPath("./");
            
             ChaincodeID ch1 = chcode.build();
            
            Channel c1 = network.getChannel();
            //NetworkImpl n1 = new NetworkImpl(c1,gateway);
            
            TransactionProposalRequest t1 = gateway.getClient().newTransactionProposalRequest();
            t1.setChaincodeID(ch1);
            t1.setFcn("executePurchase");
            t1.setArgs("1");
            TransactionProposalRequest t2 = gateway.getClient().newTransactionProposalRequest();
            t2.setChaincodeID(ch1);
            t2.setFcn("executePurchase");
            t2.setArgs("2");
            
            Collection<ProposalResponse> r1 = c1.sendTransactionProposal(t1);
            Collection<ProposalResponse> r2 = c1.sendTransactionProposal(t2);
            ArrayList<Collection<ProposalResponse>> resps = new ArrayList<>();
            resps.add(r1);
            resps.add(r2);
            System.out.printf("The length of array is: %d", resps.size());
            long l1 = System.currentTimeMillis();
            ArrayList<CompletableFuture<?>> results = new ArrayList<CompletableFuture<?>>();
            for (Collection<ProposalResponse> r: resps) {
            	CompletableFuture<BlockEvent.TransactionEvent> result = c1.sendTransaction(r);
            	results.add(result);
            }
            
            for ( CompletableFuture<?> i : results) {
            	i.join();
            	
            }
            long l2 = System.currentTimeMillis();
            System.out.printf("Execution time is: %d",(l2-l1));
     
         
            	

        } catch (ProposalException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (InvalidArgumentException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
    }
}