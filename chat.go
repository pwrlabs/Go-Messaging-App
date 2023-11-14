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
    //privateKeyHex := "0x9d4428c6e0638331b4866b70c831f8ba51c11b031f4b55eed4087bbb8ef0151f"
    //var wallet = pwrgo.FromPrivateKey(privateKeyHex)
    
    // Create new wallet and print address and keys
    var wallet = pwrgo.NewWallet()
    fmt.Println("New wallet address: ", wallet.Address)
    fmt.Println("New wallet private key: ", wallet.PrivateKeyStr)
    fmt.Println("New wallet public key: ", wallet.PublicKey)

    listener := NewListener()
    listener.Listen()

    fmt.Print("Welcome! Type 'quit' to exit\n")
    scanner := bufio.NewScanner(os.Stdin)
    for {
        // TO-DO: Keep track of nonce and increment on successful VM data tx's. Sometimes RPC returns old nonce. Or auto-retry on old nonce
        var nonce = pwrgo.NonceOfUser(wallet.Address)
        fmt.Printf("[%d]> ", nonce)
        scanner.Scan()
        message := scanner.Text()
        if message == "quit" {
            break
        }

        var data = []byte(message)
        pwrgo.ReturnBlockNumberOnTx = true
        var dataTx = pwrgo.SendVMDataTx(appId, data, nonce, wallet.PrivateKey)
        if dataTx.Success {
            fmt.Printf("[Block #%d] Sent tx: %s \n\n> ", dataTx.BlockNumber, dataTx.TxHash)
        } else {
            log.Fatal("[Block #%d] Error sending tx %s: %s", dataTx.BlockNumber, dataTx.TxHash, dataTx.Error)
        }
        
        time.Sleep(1000 * time.Millisecond)
    }
}