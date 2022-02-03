<div align="center">
<h1>Personal Wallet on Solana using Go ♾️ </h1>
</div>


## Setting up environment

* Installation of [Cobra](https://pkg.go.dev/github.com/spf13/cobra) Go package

```bash
go get -u github.com/spf13/cobra/cobra@latest
```

* Installation of Solana-go-sdk package

```bash
go get -u github.com/portto/solana-go-sdk@v1.8.1
```

--- 

## Getting Started

* Clone this repository

```bash
git clone https://github.com/RiteshPuvvada/Personal-Solana-Wallet.git

cd Personal-Solana-Wallet
```

* Available CLI commands

```bash
┌──(root💀kali)-[~/Desktop/solana/wallet]
└─# go run main.go

A CLI wallet application created in Go that interacts with the Solana blockchain.

Usage:
  personal-wallet [command]

Available Commands:
  completion     Generate the autocompletion script for the specified shell
  createWallet   Creates a new wallet
  help           Help about any command
  importWallet   Imports and existing wallet
  requestAirdrop Request airdrop in Solana
  transfer       Transfer SOL

Flags:
  -h, --help     help for personal-wallet
  -t, --toggle   Help message for toggle

Use "personal-wallet [command] --help" for more information about a command.

```

---

## Features and Usage

- [x] Creating personal Wallet
- [x] Fetching the personal Wallet
- [x] Airdrop SOL
- [x] Transferring SOL

---

* Create your personal Wallet

```bash
go run main.go createWallet
```

* Fetch the personal Wallet

```bash
go run main.go importWallet
```

* Airdrop SOL to the Wallet - **Maximum 5 SOL in one transaction**

```bash
go run main.go requestAirdrop 3
```
**We can check the transaction hash on the [Solana explorer](https://explorer.solana.com/?cluster=devnet) to confirm the transaction.**

* Transferring SOL to another address

```
go run main.go transfer <recepient public address> <amount in SOL>
```
