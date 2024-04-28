package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	apiv1 "github.com/di-dao/core/api/did/v1"
	"github.com/di-dao/core/crypto/core/curves"
	"github.com/di-dao/core/crypto/core/protocol"
	"github.com/di-dao/core/crypto/tecdsa/dklsv1"
	"github.com/di-dao/core/x/did/types"
)

// vault is the global vault instance
var vault vaultStore

// defaultCurve is the default curve used for key generation

// Keeper defines the middleware keeper.
type Keeper struct {
	cdc codec.BinaryCodec

	logger log.Logger

	// state management
	OrmDB  apiv1.StateStore
	Params collections.Item[types.Params]
	Schema collections.Schema

	authority string
}

// NewKeeper creates a new poa Keeper instance
func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService, logger log.Logger, authority string) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)
	sb := collections.NewSchemaBuilder(storeService)
	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}
	db, err := ormdb.NewModuleDB(&types.ORMModuleSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}
	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}
	k := Keeper{
		cdc:       cdc,
		logger:    logger,
		Params:    collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		authority: authority,
		OrmDB:     store,
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	return k
}

// GenerateKeyshares generates a new keyshare set. First step
func (k Keeper) GenerateKeyshares(ctx sdk.Context) (*ValidatorKeyshare, *UserKeyshare, error) {
	defaultCurve := curves.P256()
	bob := dklsv1.NewBobDkg(defaultCurve, protocol.Version1)
	alice := dklsv1.NewAliceDkg(defaultCurve, protocol.Version1)
	err := StartKsProtocol(bob, alice)
	if err != nil {
		return nil, nil, err
	}
	aliceRes, err := alice.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	bobRes, err := bob.Result(protocol.Version1)
	if err != nil {
		return nil, nil, err
	}
	valKs := createValidatorKeyshare(aliceRes)
	usrKs := createUserKeyshare(bobRes)
	return valKs, usrKs, nil
}

// LinkController links a user identifier to a kss pair creating a controller. Second step
func (k Keeper) LinkController(ctx sdk.Context, usrKs *UserKeyshare, valKs *ValidatorKeyshare, identifier string) (string, error) {
	c, err := CreateController(usrKs, valKs)
	if err != nil {
		return "", err
	}
	return c.Link("email", identifier)
}

// AssignVault assigns a vault to a controller. Third step
func (k Keeper) AssignVault(ctx sdk.Context, c Controller) error {
	return nil
}
