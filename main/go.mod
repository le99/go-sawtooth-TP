module example.com/main

go 1.16

replace example.com/handler => ../handler

require (
	example.com/handler v0.0.0-00010101000000-000000000000
	github.com/hyperledger/sawtooth-sdk-go v0.1.4 // indirect
)
