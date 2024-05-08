package types

import (
	"encoding/json"

	didv1 "github.com/di-dao/core/api/did/v1"
)

// AssertionList is a list of Assertion
type AssertionList = []*didv1.Assertion

// ByteArray is a list of byte arrays
type ByteArray = [][]byte

// KeyshareList is a list of Keyshare
type KeyshareList = []*didv1.Keyshare

// VerificationList is a list of Verification
type VerificationList = []*didv1.Verification

// ConvertByteArrayToAssertionList converts a list of byte arrays to a list of assertions
func ConvertByteArrayToAssertionList(bzList ByteArray) (AssertionList, error) {
	result := make([]*didv1.Assertion, len(bzList))
	for i, bz := range bzList {
		var a didv1.Assertion
		if err := json.Unmarshal(bz, &a); err != nil {
			return nil, err
		}
		result[i] = &a
	}
	return result, nil
}

// ConvertByteArrayToKeyshareList converts a list of byte arrays to a keyshare set
func ConvertByteArrayToKeyshareList(bzList ByteArray) (KeyshareList, error) {
	result := make([]*didv1.Keyshare, len(bzList))
	for i, bz := range bzList {
		var k didv1.Keyshare
		if err := json.Unmarshal(bz, &k); err != nil {
			return nil, err
		}
		result[i] = &k
	}
	return result, nil
}

// ConvertByteArrayToVerificationList converts a list of byte arrays to a list of verifications
func ConvertByteArrayToVerificationList(bzList ByteArray) (VerificationList, error) {
	result := make([]*didv1.Verification, len(bzList))
	for i, bz := range bzList {
		var v didv1.Verification
		if err := json.Unmarshal(bz, &v); err != nil {
			return nil, err
		}
		result[i] = &v
	}
	return result, nil
}

// ConvertAssertionListToByteArray converts a list of assertions to a list of byte arrays
func ConvertAssertionListToByteArray(assertions AssertionList) (ByteArray, error) {
	result := make([][]byte, len(assertions))
	for i, a := range assertions {
		bz, err := json.Marshal(a)
		if err != nil {
			return nil, err
		}
		result[i] = bz
	}
	return result, nil
}

// ConvertKeyshareSetToByteArray converts a keyshare set to a list of byte arrays
func ConvertKeyshareListToByteArray(kss KeyshareList) (ByteArray, error) {
	result := make([][]byte, len(kss))
	for i, k := range kss {
		bz, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		result[i] = bz
	}
	return result, nil
}

// ConvertVerificationListToByteArray converts a list of verifications to a list of byte arrays
func ConvertVerificationListToByteArray(verifications VerificationList) (ByteArray, error) {
	result := make([][]byte, len(verifications))
	for i, v := range verifications {
		bz, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		result[i] = bz
	}
	return result, nil
}
