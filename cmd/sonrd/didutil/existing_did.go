package didutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/sonr-io/sonr/pkg/did"
	"github.com/sonr-io/sonr/pkg/did/ssi"
	"github.com/sonr-io/sonr/x/object/types"
)

func surveyExistingDid() error {
	file := ""
	prompt := &survey.Input{
		Message: "Enter the path to your DID Json file:",
		Suggest: func(toComplete string) []string {
			files, _ := filepath.Glob(toComplete + "*")
			return files
		},
	}
	survey.AskOne(prompt, &file)

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	doc := &did.Document{}
	err = doc.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	opt := ""
	// Present public key types
	selectOpt := &survey.Select{
		Message: "What action would you like to perform?",
		Options: []string{
			"Add Assertion Method",
			"Add Invocation Method",
			"Add Verification Method",
			"Encrypt JWE",
			"Decrypt JWE",
		},
		Default: "Add Assertion Method",
	}

	// perform the questions
	survey.AskOne(selectOpt, &opt, nil)

	// Handle opt
	switch opt {
	case "Add Assertion Method":
		// Add Assertion Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddAssertionMethod(vm)
	case "Add Invocation Method":
		// Add Invocation Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddCapabilityInvocation(vm)
	case "Add Verification Method":
		// Add Verification Method
		controller := fmt.Sprintf("%s#test-%d", doc.ID, len(doc.Controller))
		didController, err := did.ParseDID(controller)
		if err != nil {
			return err
		}
		priv := secp256k1.GenPrivKey()
		pub := priv.PubKey()
		vm, _ := did.NewVerificationMethod(doc.ID, ssi.ECDSASECP256K1VerificationKey2019, *didController, pub)
		doc.AddAuthenticationMethod(vm)
	case "Encrypt JWE":
		// Generate JWE
		genIpld := false
		prompt := &survey.Confirm{
			Message: "Would you like to generate Object from IPLD schema?",
		}
		survey.AskOne(prompt, &genIpld)
		if genIpld {
			// Generate IPLD
			obj := types.ObjectDoc{
				Label: "test",
				Fields: []*types.TypeField{
					{
						Name: "test",
						Kind: types.TypeKind_TypeKind_String,
					},
					{
						Name: "test2",
						Kind: types.TypeKind_TypeKind_Float,
					},
					{
						Name: "test3",
						Kind: types.TypeKind_TypeKind_Bool,
					},
				},
			}

			// Marshal Object
			buf, err := obj.Marshal()
			if err != nil {
				return err
			}

			// Encrypt JWE with the DIDDoc
			str, err := doc.EncryptJWE(doc.Controller[0], buf)
			if err != nil {
				return err
			}
			fmt.Printf("Object Serial: %s\n", str)
		} else {
			fmt.Println("Pasting objects is WIP...")
		}
	case "Decrypt JWE":
		text := ""
		prompt := &survey.Multiline{
			Message: "Paste your JWE object here:",
		}
		survey.AskOne(prompt, &text)
		// Decrypt JWE
		buf, err := doc.DecryptJWE(doc.Controller[0], text)
		if err != nil {
			return err
		}
		fmt.Println(string(buf))
	default:
		return errors.New("Invalid option")
	}

	return nil
}
