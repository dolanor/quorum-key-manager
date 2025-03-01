package types

import (
	manifest "github.com/consensys/quorum-key-manager/src/manifests/types"
)

const (
	HashicorpSecrets manifest.Kind = "HashicorpSecrets"
	AKVSecrets       manifest.Kind = "AKVSecrets"
	AWSSecrets       manifest.Kind = "AWSSecrets"
)
