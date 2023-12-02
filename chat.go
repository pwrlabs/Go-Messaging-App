package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/pwrlabs/pwrgo/pwrgo"
    "time"
    "log"
)

const appId = 1337

func main() {
    // // Import wallet from Private key
    privateKeyHex := "0x9d4428c6e0638331b4866b70c831f8ba51c11b031f4b55eed4087bbb8ef0151f"
    var wallet = pwrgo.FromPrivateKey(privateKeyHex)
    
    // Create new wallet and print address and keys
    //var wallet = pwrgo.NewWallet()
    fmt.Println("Wallet address: ", wallet.Address)
    fmt.Println("Wallet private key: ", wallet.PrivateKeyStr)
    fmt.Println("Wallet public key: ", wallet.PublicKey)

    listener := NewListener()
    listener.Listen()

    fmt.Print("Welcome! Type 'quit' to exit\n")
    scanner := bufio.NewScanner(os.Stdin)
    for {
        // TO-DO: Keep track of nonce and increment on successful VM data tx's. Sometimes RPC returns old nonce. 
		//        Replace the "auto-retry" system below with a local nonce tracking system (combined with auto-retry if necessary)
        var nonce = pwrgo.NonceOfUser(wallet.Address)
        fmt.Printf("[%d]> ", nonce)
        scanner.Scan()
        message := scanner.Text()
        if message == "quit" {
            break
        }

        var data = []byte(message)
        pwrgo.ReturnBlockNumberOnTx = true
		
		for { // Detect "Old Nonce" error, increment nonce, and try again
			var dataTx = pwrgo.SendVMDataTx(appId, data, nonce, wallet.PrivateKey)
			if dataTx.Success {
				fmt.Printf("[Block #%d] Sent tx: %s \n\n> ", dataTx.BlockNumber, dataTx.TxHash)
				break // break out of "Old Nonce" error retry loop
			} else {
				if dataTx.Error == "Old Nonce" {
					nonce = nonce + 1 // loop with next nonce
					time.Sleep(100 * time.Millisecond) // 100ms delay to not ddos rpc
				} else {
					log.Fatalf("[Block #%d] Error sending tx %s: %s\n", dataTx.BlockNumber, dataTx.TxHash, dataTx.Error)
					break // unexpected error, don't retry
				}
			}
        }
        time.Sleep(1000 * time.Millisecond) // wait a second to get latest Nonce from RPC
    }
}