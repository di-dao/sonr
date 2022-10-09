package motor

import (
	"fmt"
	"testing"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/stretchr/testify/assert"
)

func (suite *MotorTestSuite) Test_LoginWithKeys() {
	suite.T().Run("with password and psk", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			Did:       suite.motorWithKeys.Address,
			Password:  "password123",
			AesPskKey: pskKey,
		}

		_, err := suite.motor.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", suite.motor.GetBalance())
			fmt.Println("address: ", suite.motor.Address)
		}
	})

	suite.T().Run("with DSC and PSK", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		fmt.Printf("aes: %x\n", aesKey)
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := mt.LoginWithKeysRequest{
			Did:       suite.motorWithKeys.Address,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		}

		_, err := suite.motor.LoginWithKeys(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", suite.motor.GetBalance())
			fmt.Println("address: ", suite.motor.Address)
		}
	})
}

func (suite *MotorTestSuite) Test_LoginWithKeyring() {
	req := mt.LoginRequest{
		Did:      suite.motor.Address,
		Password: "password123",
	}

	fmt.Println("Empty Motor generated")
	_, err := suite.motor.Login(req)
	assert.NoError(suite.T(), err, "login succeeds")

	if err == nil {
		fmt.Println("balance: ", suite.motor.GetBalance())
		fmt.Println("address: ", suite.motor.Address)
	}
}

func (suite *MotorTestSuite) Test_LoginAndMakeRequest() {
	aesKey := loadKey("aes.key")
	fmt.Printf("aes: %x\n", aesKey)
	if aesKey == nil || len(aesKey) != 32 {
		suite.T().Errorf("could not load key.")
		return
	}

	pskKey := loadKey(fmt.Sprintf("psk%s", suite.motorWithKeys.Address))
	if pskKey == nil || len(pskKey) != 32 {
		suite.T().Errorf("could not load psk key")
		return
	}

	req := mt.LoginWithKeysRequest{
		Did:       suite.motorWithKeys.Address,
		AesDscKey: aesKey,
		AesPskKey: pskKey,
	}

	_, err := suite.motorWithKeys.LoginWithKeys(req)
	assert.NoError(suite.T(), err, "login succeeds")

	// do something with the logged in account
	suite.motorWithKeys.DIDDocument.AddAlias("gotest.snr")
	_, err = updateWhoIs(suite.motorWithKeys)
	assert.NoError(suite.T(), err, "updates successfully")
}
