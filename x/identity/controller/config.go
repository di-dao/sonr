package controller

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/internal/vault"
	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc"
	"github.com/sonrhq/core/x/identity/models"
	"github.com/sonrhq/core/x/identity/types"
)

// ! ||--------------------------------------------------------------------------------||
// ! ||                              Client Facing Options                             ||
// ! ||--------------------------------------------------------------------------------||

type Options struct {
	// The controller's on config generated handler
	OnConfigGenerated []mpc.OnConfigGenerated

	// Credential to authorize the controller
	WebauthnCredential *crypto.WebauthnCredential

	// Disable IPFS
	DisableIPFS bool

	// Broadcast the transaction
	BroadcastTx bool

	// Username for the controller
	Username string
}

type Option func(*Options)

func WithUsername(username string) Option {
	return func(o *Options) {
		o.Username = username
	}
}

func WithConfigHandlers(handlers ...mpc.OnConfigGenerated) Option {
	return func(o *Options) {
		o.OnConfigGenerated = handlers
	}
}

func WithWebauthnCredential(cred *crypto.WebauthnCredential) Option {
	return func(o *Options) {
		o.WebauthnCredential = cred
	}
}

func WithIPFSDisabled() Option {
	return func(o *Options) {
		o.DisableIPFS = true
	}
}

func WithBroadcastTx() Option {
	return func(o *Options) {
		o.BroadcastTx = true
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Helper Methods for Controller                         ||
// ! ||--------------------------------------------------------------------------------||

func generateInitialAccount(ctx context.Context, credential *crypto.WebauthnCredential, doneCh chan models.Account, errChan chan error, opts *Options) {
	shardName := crypto.PartyID(base64.RawStdEncoding.EncodeToString(credential.Id))
	// Call Handler for keygen
	confs, err := mpc.Keygen(shardName, 1, []crypto.PartyID{"vault"}, opts.OnConfigGenerated...)
	if err != nil {
		errChan <- err
	}

	pubKey, err := crypto.NewPubKeyFromCmpConfig(confs[0])
	if err != nil {
		errChan <- err
	}

	rootDid, _ := crypto.SONRCoinType.FormatDID(pubKey)
	var kss []models.KeyShare
	for _, conf := range confs {
		ksb, err := conf.MarshalBinary()
		if err != nil {
			errChan <- err
		}
		ksDid := fmt.Sprintf("%s#%s", rootDid, conf.ID)
		ks, err := models.NewKeyshare(ksDid, ksb, crypto.SONRCoinType, "Primary")
		if err != nil {
			errChan <- err
		}
		kss = append(kss, ks)
	}
	doneCh <- models.NewAccount(kss, crypto.SONRCoinType)
}

func setupController(ctx context.Context, primary models.Account, opts *Options) (Controller, error) {
	if !opts.DisableIPFS {
		err := vault.InsertAccount(primary)
		if err != nil {
			return nil, err
		}
	}

	doc := types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), nil)
	if opts.WebauthnCredential != nil {
		cred, err := types.ValidateWebauthnCredential(opts.WebauthnCredential, primary.Did())
		if err != nil {
			return nil, err
		}
		doc = types.NewPrimaryIdentity(primary.Did(), primary.PubKey(), cred.ToVerificationMethod())
		if !opts.DisableIPFS {
			err = vault.StoreCredential(cred)
			if err != nil {
				return nil, err
			}
		}
	}


	if opts.Username != "" {
		doc.AlsoKnownAs = []string{opts.Username}
	}

	if opts.BroadcastTx {
		resp, err := local.Context().CreatePrimaryIdentity(doc, primary)
		if err != nil {
			return nil, err
		}
		if resp.TxResponse.Code != 0 {
			return nil, fmt.Errorf("failed to broadcast transaction: %s", resp.TxResponse.TxHash)
		}
	}

	cont := &didController{
		primary:     primary,
		blockchain:  []models.Account{},
		primaryDoc:  doc,
		disableIPFS: opts.DisableIPFS,
	}
	return cont, nil
}