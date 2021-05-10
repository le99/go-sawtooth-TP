package handler

import (
	// "fmt"
	// "github.com/hyperledger/sawtooth-sdk-go/examples/xo_go/src/sawtooth_xo/xo_payload"
	// "github.com/hyperledger/sawtooth-sdk-go/examples/xo_go/src/sawtooth_xo/xo_state"
	// "github.com/hyperledger/sawtooth-sdk-go/logging"
	// "github.com/hyperledger/sawtooth-sdk-go/logging"
	"fmt"

	"github.com/hyperledger/sawtooth-sdk-go/examples/xo_go/src/sawtooth_xo/xo_state"
	"github.com/hyperledger/sawtooth-sdk-go/logging"
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2"
	// "github.com/hyperledger/sawtooth-sdk-go/processor"
	// "github.com/hyperledger/sawtooth-sdk-go/protobuf/processor_pb2"
	// "strings"
)

var logger *logging.Logger = logging.Get()

type TPHandler struct {
}

func (self *TPHandler) FamilyName() string {
	return "tp0"
}

func (self *TPHandler) FamilyVersions() []string {
	return []string{"1.0"}
}

func (self *TPHandler) Namespaces() []string {
	return []string{xo_state.Namespace}
}

func (self *TPHandler) Apply(request *processor_pb2.TpProcessRequest, context *processor.Context) error {
	// The xo player is defined as the signer of the transaction, so we unpack
	// the transaction header to obtain the signer's public key, which will be
	// used as the player's identity.
	header := request.GetHeader()
	player := header.GetSignerPublicKey()
	fmt.Println(player)

	return nil
}
