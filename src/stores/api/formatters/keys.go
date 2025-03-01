package formatters

import (
	"encoding/base64"

	"github.com/consensys/quorum-key-manager/src/stores/api/types"
	"github.com/consensys/quorum-key-manager/src/stores/store/entities"
)

func FormatKeyResponse(key *entities.Key) *types.KeyResponse {
	return &types.KeyResponse{
		ID:               key.ID,
		PublicKey:        base64.StdEncoding.EncodeToString(key.PublicKey),
		Curve:            string(key.Algo.EllipticCurve),
		SigningAlgorithm: string(key.Algo.Type),
		Tags:             key.Tags,
		Annotations:      key.Annotations,
		Disabled:         key.Metadata.Disabled,
		CreatedAt:        key.Metadata.CreatedAt,
		UpdatedAt:        key.Metadata.UpdatedAt,
		ExpireAt:         key.Metadata.ExpireAt,
		DeletedAt:        key.Metadata.DeletedAt,
		DestroyedAt:      key.Metadata.DestroyedAt,
	}
}
