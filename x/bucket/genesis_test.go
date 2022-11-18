package bucket_test

import (
	"testing"

	keepertest "github.com/sonr-io/sonr/testutil/keeper"
	"github.com/sonr-io/sonr/testutil/nullify"
	"github.com/sonr-io/sonr/x/bucket"
	"github.com/sonr-io/sonr/x/bucket/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		BucketDefinitions: []types.BucketConfig{
			{
				Uuid: "did:sonr:1",
			},
			{
				Uuid: "did:sonr:2",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BucketKeeper(t)
	bucket.InitGenesis(ctx, *k, genesisState)
	got := bucket.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.BucketDefinitions, got.BucketDefinitions)
	// this line is used by starport scaffolding # genesis/test/assert
}