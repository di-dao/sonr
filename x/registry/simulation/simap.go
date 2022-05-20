package simulation

import (
	"crypto/ed25519"
	cryptrand "crypto/rand"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
)

// FindAccount find a specific address from an account list
func FindAccount(accs []simtypes.Account, address string) (simtypes.Account, bool) {
	creator, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		panic(err)
	}
	return simtypes.FindAccount(accs, creator)
}

// Mock Credential object from webauthn test bench https://github.com/psteniusubi/webauthn-tester
func CreateMockCredential() (*did.Credential, error) {

	return &did.Credential{
		ID:              []byte("ktIQAlFosR9OMGnyJnGthmKcIodPb323F3UqPVe9kvB-eOYrE-pNchsSuiN4ZE0ICyAaRiCb6vfF-7Y5nrvcoD-D42KQsXzhJd14ciqzibA"),
		AttestationType: "platform",
		PublicKey:       []byte("o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjULNeTz6C0GMu_DqhSIoYH2el7Mz1NsKQQF3Zq9ruMdVFFAAAAAK3OAAI1vMYKZIsLJfHwVQMAUJLSEAJRaLEfTjBp8iZxrYZinCKHT299txd1Kj1XvZLwfnjmKxPqTXIbErojeGRNCAsgGkYgm"),
		Authenticator: did.Authenticator{
			AAGUID:    []byte("eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoiOHhBM2t3dUVCM0xtc2UxMkJyT2FrSDlZUWlrIiwib3JpZ2luIjoiaHR0cHM6Ly9wc3Rlbml1c3ViaS5naXRodWIuaW8iLCJjcm9zc09yaWdpbiI6ZmFsc2V9"),
			SignCount: 1,
		},
	}, nil
}

// CreateMockDidDocument creates a mock did document for testing
func CreateMockDidDocument(simAccount simtypes.Account) (*did.Document, error) {
	//webauthncred := CreateMockCredential()
	// idStr := "ktIQAlFosR9OMGnyJnGthmKcIodPb323F3UqPVe9kvB-eOYrE-pNchsSuiN4ZE0ICyAaRiCb6vfF-7Y5nrvcoD-D42KQsXzhJd14ciqzibA"
	pubKey, _, err := ed25519.GenerateKey(cryptrand.Reader)
	if err != nil {
		return nil, err
	}

	didUrl, err := did.ParseDID(fmt.Sprintf("did:snr:%s", simAccount.Address.String()))
	if err != nil {
		return nil, err
	}
	didController, err := did.ParseDID(fmt.Sprintf("did:snr:%s#test", simAccount.Address.String()))
	if err != nil {
		return nil, err
	}

	vm, err := did.NewVerificationMethod(*didUrl, ssi.JsonWebKey2020, *didController, pubKey)
	if err != nil {
		return nil, err
	}
	ctxUri, err := ssi.ParseURI("https://www.w3.org/ns/did/v1")
	if err != nil {
		return nil, err
	}

	doc := did.Document{
		ID:      *didUrl,
		Context: []ssi.URI{*ctxUri},
	}
	doc.AddAuthenticationMethod(vm)
	return &doc, nil
}
