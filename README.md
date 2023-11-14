# Go-Messaging-App
A simple Golang-based messaging application on PWR Chain. Built using pwrgo.

# Install pwrgo library

To install the pwrgo library, run

```
go get github.com/pwrlabs/pwrgo@v0.0.2
```

# Run the application

In the repo root, run

```
go run chat.go chatListener.go
```

# Faucet

Use the PWR faucet to get 100 PWR tokens for testing: https://faucet.pwrlabs.io/

# Example output

```
> go run chat.go chatListener.go
New wallet address:  0x3B942C725939B90FF7b0C67c66af11dbe2bd17dC
New wallet private key:  0x0e90aa4226471e0a9e7ffaf95fdc582bff19968589cb547262588b380fbde48c
New wallet public key:  0x04647509f4e394208350435b78b80466472bbcc3e4195c38ddd937d3517fe8cb2d6bbfae89695d67e364d71074f5c27eca1f5319bda43a779086d8b3adfc1cb874
Welcome! Type 'quit' to exit
[0]> hello world
Using nonce:  0
[Block #8622] Sent tx: 0x31fc383af37ea36cedabec4ed692e86c7a4511f550ac1194646f954b96b5a95c

> Message From 0x3b942c725939b90ff7b0c67c66af11dbe2bd17dc: hello world

> [1]>
```


# To-do:

- Usernames or display names for addresses
- Timestamps on messages
- Import wallet from private key