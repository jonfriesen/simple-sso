package memory

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonfriesen/simple-sso/pkg/sso"
)

var PrivateKey *rsa.PrivateKey
var BaseConf *sso.BaseConfig

func setupBaseConfig() {
	var err error
	BaseConf, err = sso.SetupBaseConfig()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyData, err := ioutil.ReadFile(BaseConf.PrivateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		log.Fatal(err)
	}
}
