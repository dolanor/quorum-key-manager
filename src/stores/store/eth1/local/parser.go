package local

import (
	"github.com/consensys/quorum-key-manager/src/stores/store/entities"
	"github.com/ethereum/go-ethereum/crypto"
)

func ParseKey(key *entities.Key) *entities.ETH1Account {
	pubKey, _ := crypto.UnmarshalPubkey(key.PublicKey)
	return &entities.ETH1Account{
		ID:                  key.ID,
		Address:             crypto.PubkeyToAddress(*pubKey),
		Metadata:            key.Metadata,
		Tags:                key.Tags,
		PublicKey:           crypto.FromECDSAPub(pubKey),
		CompressedPublicKey: crypto.CompressPubkey(pubKey),
	}
}
