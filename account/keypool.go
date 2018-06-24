package account

import (
	"github.com/pborman/uuid"
	"go-vite/common"
	"go-vite/crypto/ed25519"
)

// Manage keys

type KeyPool struct {
	ks keyStore
}

func newKeyFromEd25519(priv *ed25519.PrivateKey) *Key {
	id := uuid.NewRandom()
	key := &Key{
		Id:         id,
		Address:    common.PrikeyToAddress(*priv),
		PrivateKey: priv,
	}
	return key
}

func (kp *KeyPool) StoreNewKey(pwd string) (*Key, common.Address, error) {
	_, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, common.InvalidAddress, err
	}
	key := newKeyFromEd25519(&priv)

	if err := kp.ks.StoreKey(key, pwd); err != nil {
		return nil, common.InvalidAddress, err
	}
	return key, key.Address, err
}

func (kp *KeyPool) ExtractKey(a common.Address, pwd string) (common.Address, *Key, error) {
	key, err := kp.ks.ExtractKey(a, pwd)
	return a, key, err
}