package types_test

import (
	"testing"

	"github.com/sonrhq/core/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PrimaryIdentities: []types.DidDocument{
					{
						Id: "0",
					},
					{
						Id: "1",
					},
				},
				BlockchainIdentities: []types.DidDocument{
					{
						Id: "0",
					},
					{
						Id: "1",
					},
				},
				Relationships: []types.VerificationRelationship{
					{
						Reference: "0",
					},
					{
						Reference: "1",
					},
				},
				ServiceList: []types.Service{
					{
						Id: "0",
					},
					{
						Id: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				// require.Error(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
