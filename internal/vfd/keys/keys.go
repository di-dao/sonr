package keys

import (
	"github.com/di-dao/sonr/crypto/mpc"
)

func Generate() {
	set, err := mpc.GenerateKss()
	set.Encrypt()
}
