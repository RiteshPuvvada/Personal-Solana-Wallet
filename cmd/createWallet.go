/*
Copyright © 2022 Ritesh Puvvada riteshpuvvada@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet", 
	Short: "Creates a new wallet", 
	Long: "Creates a new wallet and provides wallet address and private key.",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new wallet.")
        wallet := CreateNewWallet(rpc.DevnetRPCEndpoint)
        fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
        fmt.Println("Private Key Saved in 'key_data' file")
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
