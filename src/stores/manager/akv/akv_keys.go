package akv

import (
	"github.com/consensys/quorum-key-manager/pkg/errors"
	client2 "github.com/consensys/quorum-key-manager/src/infra/akv/client"
	"github.com/consensys/quorum-key-manager/src/infra/log"
	"github.com/consensys/quorum-key-manager/src/stores/store/keys/akv"
)

// KeySpecs is the specs format for an Azure Key Vault key store
type KeySpecs struct {
	VaultName           string `json:"vaultName"`
	SubscriptionID      string `json:"subscriptionID"`
	TenantID            string `json:"tenantID"`
	AuxiliaryTenantIDs  string `json:"auxiliaryTenantIDs"`
	ClientID            string `json:"clientID"`
	ClientSecret        string `json:"clientSecret"`
	CertificatePath     string `json:"certificatePath"`
	CertificatePassword string `json:"certificatePassword"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	EnvironmentName     string `json:"environmentName"`
	Resource            string `json:"resource"`
}

func NewKeyStore(spec *KeySpecs, logger log.Logger) (*akv.Store, error) {
	cfg := client2.NewConfig(spec.VaultName, spec.TenantID, spec.ClientID, spec.ClientSecret)
	cli, err := client2.NewClient(cfg)
	if err != nil {
		errMessage := "failed to instantiate AKV client (keys)"
		logger.WithError(err).Error(errMessage, "specs", spec)
		return nil, errors.ConfigError(errMessage)
	}

	store := akv.New(cli, logger)
	return store, nil
}
