package coinmate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

type Client struct {
	ClientID   string
	PublicKey  string
	PrivateKey string
}

func EnvVarMissing(envVarName string) error {
	return fmt.Errorf("Missing enviroment variable: %s", envVarName)
}

func (c *Client) NewClientFromEnv() error {
	clientID := os.Getenv("COINMATE_CLIENT_ID")
	publicKey := os.Getenv("COINMATE_PUBLIC_KEY")
	privateKey := os.Getenv("COINMATE_PRIVATE_KEY")

	if clientID == "" {
		return EnvVarMissing("COINMATE_CLIENT_ID")
	}
	if publicKey == "" {
		return EnvVarMissing("COINMATE_PUBLIC_KEY")
	}
	if privateKey == "" {
		return EnvVarMissing("COINMATE_PRIVATE_KEY")
	}

	c.ClientID = clientID
	c.PublicKey = publicKey
	c.PrivateKey = privateKey

	return nil
}

func (c *Client) computeSignature(nonce string) string {
	message := nonce + c.ClientID + c.PublicKey
	mac := hmac.New(sha256.New, []byte(c.PrivateKey))
	mac.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}
