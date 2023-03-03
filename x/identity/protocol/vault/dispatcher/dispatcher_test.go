package dispatcher_test

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/x/identity/protocol/vault/dispatcher"
	"github.com/stretchr/testify/assert"
)

type accTest struct {
	name     string
	coinType crypto.CoinType
}

var accTests = []accTest{
	{"Ethereum", crypto.ETHCoinType},
	{"Bitcoin", crypto.BTCCoinType},
	{"Filecoin", crypto.FILCoinType},
	{"Handshake", crypto.HNSCoinType},
	{"Litecoin", crypto.LTCCoinType},
	{"Cosmos", crypto.COSMOSCoinType},
}

func usedNetworks() string {
	var nets []string
	for _, test := range accTests {
		nets = append(nets, test.coinType.Symbol())
	}
	return strings.Join(nets, ", ")
}

func TestDispatcherAccounts(t *testing.T) {
	t.Logf("Initialize new DID Controller...")
	startTime := time.Now()
	d := dispatcher.New()

	w, err := d.BuildNewDIDController()
	checkErr(t, err)
	t.Logf("(%s) - Root Identifier", w.ID())
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	un := usedNetworks()
	t.Logf("\nCreate %s accounts....", un)
	startTime = time.Now()
	for i, test := range accTests {
		vm, err := w.CreateAccount(test.name, test.coinType)
		checkErr(t, err)
		t.Logf("\t* [%d] #%s", i, vm.IDFragmentSuffix())
		t.Logf("\t\t↪ Address: %s", vm.BlockchainAccountId)
		t.Logf("\t\t↪ Controller: %s", vm.Controller)
		t.Logf("\t\t↪ Type: %s", vm.Type)
		t.Logf("\t\t↪ Multibase PubKey: %s", vm.PublicKeyMultibase)
	}
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	t.Logf("\nGet %s accounts....", un)
	startTime = time.Now()
	for i, test := range accTests {
		acc, err := w.GetAccount(test.name)
		checkErr(t, err)
		t.Logf("\t↪ [%d]<%s> %s", i, test.coinType.Symbol(), acc.Name())
	}
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	msg := []byte("Hello World!")
	t.Logf("\nSign msg (%v) with DID Controller....", msg)
	startTime = time.Now()
	sig, err := w.Sign(msg)
	checkErr(t, err)
	t.Logf("\t↪ %s", sig)
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	t.Logf("\nVerify msg (%v) with DID Controller....", msg)
	startTime = time.Now()
	ok, err := w.Verify(msg, sig)
	checkErr(t, err)
	t.Logf("\t↪ %t", ok)
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))
}

func TestDispatcherSignature(t *testing.T) {
	t.Logf("Initialize new DID Controller...")
	startTime := time.Now()
	d := dispatcher.New()

	w, err := d.BuildNewDIDController()
	checkErr(t, err)
	t.Logf("(%s) - Root Identifier", w.ID())
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	msg := []byte("Hello World!")
	t.Logf("\nSign msg with DID Controller....")
	startTime = time.Now()
	sig, err := w.Sign(msg)
	checkErr(t, err)
	t.Logf("\t↪ %s", base64.RawStdEncoding.EncodeToString(sig))
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))

	t.Logf("\nVerify msg with DID Controller....")
	startTime = time.Now()
	ok, err := w.Verify(msg, sig)
	checkErr(t, err)
	assert.True(t, ok)
	t.Logf("DONE! Time elapsed: %s\n", time.Since(startTime))
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}