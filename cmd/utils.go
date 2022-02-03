package cmd

import ( 
   "context"	
   "io/ioutil"
   "log"

   "github.com/portto/solana-go-sdk/client"
   "github.com/portto/solana-go-sdk/client/rpc"
   "github.com/portto/solana-go-sdk/common"
   "github.com/portto/solana-go-sdk/program/sysprog"
   "github.com/portto/solana-go-sdk/types" 
)

type Wallet struct { 
	account   types.Account 
    c 	    *client.Client 

}

func CreateNewWallet(RPCEndpoint string) Wallet {
	newAccount := types.NewAccount()
	data := []byte(newAccount.PrivateKey) // convert the private key to byte array for storage
    err := ioutil.WriteFile("key_data", data, 0644) 
    if err != nil { 
        log.Fatal(err) 
    }
	return Wallet{
		newAccount,
		client.NewClient(RPCEndpoint),
	}
}

func ImportOldWallet(RPCEndpoint string) (Wallet, error) {
	contents, _ := ioutil.ReadFile("key_data")
	privateKey := []byte(string(contents))
	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func RequestAirdrop(amount uint64) (string, error) { 
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    amount = amount * 1e9 // turning SOL into lamports 
    txhash, err := wallet.c.RequestAirdrop( 
        context.TODO(),                      // request context
        wallet.account.PublicKey.ToBase58(), // wallet address requesting airdrop 
		amount,                              // amount of SOL in lamport 
    ) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}


func GetBalance() (uint64, error) {
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    balance, err := wallet.c.GetBalance( 
        context.TODO(),                      // request context
        wallet.account.PublicKey.ToBase58(), // wallet to fetch balance for 
    ) 
    if err != nil { 
        return 0, nil 
    }
    return balance, nil
}


func Transfer(receiver string, amount uint64) (string, error) { 
    wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint) 
    response, err := wallet.c.GetRecentBlockhash(context.TODO()) 
    if err != nil { 
        return "", err 
    } 


    message := types.NewMessage( 
        wallet.account.PublicKey, // public key of the transaction signer
        []types.Instruction{
            sysprog.Transfer( 
                wallet.account.PublicKey, // public key of the transaction sender 
                common.PublicKeyFromString(receiver), // wallet address of the transaction receiver 
                amount,                               // transaction amount in lamport 
            ), 
        }, 
        response.Blockhash, // recent block hash 
    ) 

    // create a transaction with the message and TX signer 
    tx, err := types.NewTransaction(message, []types.Account{wallet.account, wallet.account}) 
    if err != nil { 
        return "", err 
    } 

    // send the transaction to the blockchain 
    txhash, err := wallet.c.SendTransaction2(context.TODO(), tx) 
    if err != nil { 
        return "", err 
    } 
    return txhash, nil 
}