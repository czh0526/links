package global

import (
	"fmt"
	"github.com/czh0526/links/account"
	"github.com/czh0526/links/links"
	libp2p_crypto "github.com/libp2p/go-libp2p/core/crypto"
)

func GetMyAccount(nickname string) (*account.Account, error) {
	return &account.Account{
		Id:       "account-12345",
		Nickname: nickname,
		Phone:    "13520746670",
	}, nil
}

func GetPrivateKey(id string) (libp2p_crypto.PrivKey, error) {
	privateKey, _, err := libp2p_crypto.GenerateKeyPair(
		libp2p_crypto.ECDSA, 2048)
	if err != nil {
		return nil, fmt.Errorf("create priv key failed, err = %v", err)
	}
	return privateKey, nil
}

func GetMyFriends(id string) (map[string]*links.Friend, error) {
	return links.GetMyFriends(id)
}

func GetMyGroups(id string) (map[string]*links.Group, error) {
	return links.GetMyGroups(id)
}
