package types

import (
	"strings"

	crypto "github.com/sonrhq/core/internal/crypto"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

// NewDIDDocument creates a new DIDDocument from an Identification and optional VerificationRelationships
func NewDIDDocument(primaryAccount vaulttypes.Account, authentication *VerificationMethod, alias string) *DIDDocument {
	params := DefaultParams()
	didDoc := &DIDDocument{
		Id:                   primaryAccount.Did(),
		Context:              []string{params.AccountDidMethodContext, params.DidBaseContext},
		Authentication:       make([]*VerificationRelationship, 0),
		AssertionMethod:      make([]*VerificationRelationship, 0),
		CapabilityInvocation: make([]*VerificationRelationship, 0),
		CapabilityDelegation: make([]*VerificationRelationship, 0),
		Controller:           []string{authentication.Id},
		AlsoKnownAs:          []string{alias},
		Metadata:             "",
	}
	didDoc.LinkAuthenticationMethod(authentication)
	didDoc.LinkCapabilityInvocationFromVaultAccount(primaryAccount)
	return didDoc
}

// SearchRelationshipsByCoinType returns all verification relationships of a given the query options
func (d *DIDDocument) SearchRelationshipsByCoinType(coinType crypto.CoinType) []*VerificationRelationship {
	method := coinType.DidMethod()
	relationships := make([]*VerificationRelationship, 0)
	for _, relationship := range d.Authentication {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.AssertionMethod {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.CapabilityInvocation {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.CapabilityDelegation {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	for _, relationship := range d.KeyAgreement {
		if strings.Contains(relationship.Reference, method) {
			relationships = append(relationships, relationship)
		}
	}
	return relationships
}

// ListAuthenticationVerificationMethods returns all the VerificationMethods for the AuthenticationRelationships
func (d *DIDDocument) ListAuthenticationVerificationMethods() []*VerificationMethod {
	vms := make([]*VerificationMethod, 0)
	for _, relationship := range d.Authentication {
		vms = append(vms, relationship.VerificationMethod)
	}
	return vms
}

// Address returns the address of the DIDDocument
func (d *DIDDocument) Address() string {
	return strings.Split(d.Id, ":")[2]
}

// LinkAuthenticationMethod adds a VerificationMethod to the Authentication list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the Authentication list
func (id *DIDDocument) LinkAuthenticationMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, auth := range id.Authentication {
		if auth.Reference == vm.Id {
			return nil, false
		}
	}
	vm.Controller = id.Id
	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AuthenticationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.Authentication = append(id.Authentication, vr)
	return vr, true
}

// LinkAssertionMethod adds a VerificationMethod to the AssertionMethod list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the AssertionMethod list
func (id *DIDDocument) LinkAssertionMethod(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, assertionMethod := range id.AssertionMethod {
		if assertionMethod.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               AssertionRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.AssertionMethod = append(id.AssertionMethod, vr)
	return vr, true
}

// LinkCapabilityDelegation adds a VerificationMethod to the CapabilityDelegation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityDelegation list
func (id *DIDDocument) LinkCapabilityDelegation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, capabilityDelegation := range id.CapabilityDelegation {
		if capabilityDelegation.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               CapabilityDelegationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.CapabilityDelegation = append(id.CapabilityDelegation, vr)
	return vr, true
}

// LinkCapabilityInvocation adds a VerificationMethod to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the CapabilityInvocation list
func (id *DIDDocument) LinkCapabilityInvocation(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, capabilityInvocation := range id.CapabilityInvocation {
		if capabilityInvocation.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               CapabilityInvocationRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.CapabilityInvocation = append(id.CapabilityInvocation, vr)
	return vr, true
}

// LinkCapabilityInvocationFromVaultAccount adds a Vault Account to the CapabilityInvocation list of the DID Document and returns the VerificationRelationship
func (id *DIDDocument) LinkCapabilityInvocationFromVaultAccount(accounts ...vaulttypes.Account) ([]*VerificationRelationship, bool) {
	vrs := make([]*VerificationRelationship, 0)
	for _, account := range accounts {
		vm := &VerificationMethod{
			Id:                  account.Did(),
			Type:                crypto.Ed25519KeyType.FormatString(),
			Controller:          id.Id,
			PublicKeyMultibase:  account.PubKey().Multibase(),
			BlockchainAccountId: account.Address(),
		}
		vr := &VerificationRelationship{
			Reference:          account.Did(),
			Type:               CapabilityInvocationRelationshipName,
			VerificationMethod: vm,
			Owner:              id.Id,
		}
		id.CapabilityInvocation = append(id.CapabilityInvocation, vr)
		vrs = append(vrs, vr)
	}
	return vrs, true
}

// LinkKeyAgreement adds a VerificationMethod to the KeyAgreement list of the DID Document and returns the VerificationRelationship
// Returns nil if the VerificationMethod is already in the KeyAgreement list
func (id *DIDDocument) LinkKeyAgreement(vm *VerificationMethod) (*VerificationRelationship, bool) {
	for _, keyAgreement := range id.KeyAgreement {
		if keyAgreement.Reference == vm.Id {
			return nil, false
		}
	}

	vr := &VerificationRelationship{
		Reference:          vm.Id,
		Type:               KeyAgreementRelationshipName,
		VerificationMethod: vm,
		Owner:              id.Id,
	}
	id.KeyAgreement = append(id.KeyAgreement, vr)
	return vr, true
}
