package controller_test

import (
	"context"
	"testing"

	"github.com/sonr-io/sonr/internal/protocol/packages/controller"
	"github.com/sonr-io/sonr/internal/protocol/packages/resolver"
	"github.com/sonr-io/sonr/pkg/crypto"
	"github.com/sonr-io/sonr/pkg/crypto/mpc"
	"github.com/sonr-io/sonr/pkg/tx/cosmos"
	"github.com/sonr-io/sonr/x/identity/types"
	"github.com/tendermint/tendermint/libs/rand"
)

func TestTree(t *testing.T) {
	randUuid := rand.Str(4)

	// Call Handler for keygen
	confs, err := mpc.Keygen(crypto.PartyID(randUuid), 1, []crypto.PartyID{"vault"})
	if err != nil {
		t.Fatal(err)
	}

	var kss []controller.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			t.Fatal(err)
		}
		ks, err := controller.NewKeyshare(string(conf.ID), ksb, crypto.SONRCoinType, "test")
		if err != nil {
			t.Fatal(err)
		}
		kss = append(kss, ks)
	}

	for _, ks := range kss {
		t.Logf("keyshare key id: %s", ks.Did())
	}
	acc := controller.NewAccount(kss, crypto.SONRCoinType)
	if err != nil {
		t.Fatal(err)
	}
	msg := []byte("hello world")
	sig, err := acc.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature: %x", sig)
	ok, err := acc.Verify(msg, sig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("verify: %v", ok)
}

func TestController(t *testing.T) {
	t.Log("create controller with initial accounts: bitcoin, ethereum")
	controller, _, err := controller.NewController(context.Background(), controller.WithInitialAccounts("bitcoin", "ethereum"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", controller.Did())
	t.Logf("controller doc: %v", controller.PrimaryIdentity().String())
	t.Log("list accounts")
	accs := controller.ListLocalAccounts()
	msg := []byte("hello world")
	for _, acc := range accs {
		t.Logf("did: %s", acc.Did())
		sig, err := acc.Sign(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("signature: %x", sig)

		ok, err := acc.Verify(msg, sig)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("verify: %v", ok)

		doc := acc.DidDocument(controller.Did())
		t.Logf("did doc: %v", doc.String())
	}

	didDoc := controller.PrimaryIdentity()
	t.Logf("did doc: %v", didDoc.String())
}

func TestNewLoad(t *testing.T) {
	cn, _, err := controller.NewController(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
	didDoc := cn.PrimaryIdentity()

	cn2, err := controller.AuthorizeController(context.Background(), didDoc)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("account: %v", cn2.Did())
}

func TestControllerCreateBroadcastTx(t *testing.T) {
	cn, prim, err := controller.NewController(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("controller: %v", cn.Did())
	didDoc := cn.PrimaryIdentity()

	_, err = cn.CreateAccount("ethTest", crypto.ETHCoinType)
	if err != nil {
		t.Fatal(err)
	}

	msg := types.NewMsgCreateDidDocument(cn.Address(), didDoc)
	txBz, err := cosmos.SignAnyTransactions(prim, "/sonrhq.core.identity.MsgCreateDidDocument", msg)
	if err != nil {
		t.Fatal(err)
	}

	res, err := resolver.BroadcastTx(txBz)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("TX Hash: %v\n", res.TxResponse.TxHash)
	t.Logf("TX Log: %v\n", res.TxResponse.Logs.String())
}