package sonr

import (
	"context"
	"encoding/json"
	"fmt"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/sonrhq/core/internal/local"
	"google.golang.org/grpc"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Tendermint Node RPC                              ||
// ! ||--------------------------------------------------------------------------------||

var kBaseAPIUrl = fmt.Sprintf("http://%s", local.NodeAPIHost())

// The `EncodeTx` function is a method of the `LocalContext` struct. It takes a transaction (`tx`) as input and returns the encoded transaction bytes as output.
func EncodeCosmosTx(tx *txtypes.Tx) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/cosmos/tx/v1beta1/encode", kBaseAPIUrl)
	req := &txtypes.TxEncodeRequest{
		Tx: tx,
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	bz, err := local.PostJSON(endpoint, body)
	if err != nil {
		return nil, err
	}
	resp := new(txtypes.TxEncodeResponse)
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return resp.TxBytes, nil
}

// BroadcastTx broadcasts a transaction on the Sonr blockchain network
func BroadcastCosmosTx(rawTx []byte) (*txtypes.BroadcastTxResponse, error) {
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		local.NodeGrpcHost(), // Or your gRPC server address.
		grpc.WithInsecure(),  // The Cosmos SDK doesn't support any transport security mechanism.
	)
	if err != nil {
		return nil, err
	}
	defer grpcConn.Close()

	// Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// service.

	txClient := txtypes.NewServiceClient(grpcConn)
	// We then call the BroadcastTx method on this client.
	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: rawTx, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		return nil, err
	}
	if grpcRes == nil {
		return nil, fmt.Errorf("no response from broadcast tx")
	}
	if grpcRes.GetTxResponse() != nil && grpcRes.GetTxResponse().Code != 0 {
		return nil, fmt.Errorf("failed to broadcast transaction: %s", grpcRes.GetTxResponse().RawLog)
	}
	return grpcRes, nil
}
