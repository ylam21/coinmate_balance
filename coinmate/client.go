package coinmate

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

type Client struct {
	ClientID   string
	PublicKey  string
	PrivateKey string
}

func NewClientFromEnv() (*Client, error) {
	clientID := os.Getenv("COINMATE_CLIENT_ID")
	publicKey := os.Getenv("COINMATE_PUBLIC_KEY")
	privateKey := os.Getenv("COINMATE_PRIVATE_KEY")

	if clientID == "" {
		return nil, EnvVarMissing("COINMATE_CLIENT_ID")
	}
	if publicKey == "" {
		return nil, EnvVarMissing("COINMATE_PUBLIC_KEY")
	}
	if privateKey == "" {
		return nil, EnvVarMissing("COINMATE_PRIVATE_KEY")
	}

	return &Client{clientID, publicKey, privateKey}, nil
}

func (c *Client) computeSignature(nonce string) string {
	message := nonce + c.ClientID + c.PublicKey
	mac := hmac.New(sha256.New, []byte(c.PrivateKey))
	mac.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}
