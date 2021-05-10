package main

import (
	"example.com/handler"
	"github.com/hyperledger/sawtooth-sdk-go/processor"

	"fmt"
	"syscall"
)

func main() {

	handler := &handler.TPHandler{}
	fmt.Println(handler.FamilyVersions())
	fmt.Println("1")
	endpoint := "tcp://127.0.0.1:4004"
	// // In docker, endpoint would be the validator's container name
	// // with port 4004
	processor := processor.NewTransactionProcessor(endpoint)

	processor.AddHandler(handler)
	processor.ShutdownOnSignal(syscall.SIGINT, syscall.SIGTERM)

	processor.Start()
}
