package hashicorp

import (
	"context"

	"github.com/consensys/quorum-key-manager/src/infra/hashicorp/client"
	"github.com/consensys/quorum-key-manager/src/infra/hashicorp/token"
	"github.com/consensys/quorum-key-manager/src/infra/log"

	"github.com/consensys/quorum-key-manager/pkg/errors"
	"github.com/consensys/quorum-key-manager/src/stores/store/secrets/hashicorp"
)

// SecretSpecs is the specs format for an Hashicorp Vault secret store
type SecretSpecs struct {
	MountPoint string `json:"mountPoint"`
	Address    string `json:"address"`
	Token      string `json:"token"`
	TokenPath  string `json:"tokenPath"`
	Namespace  string `json:"namespace"`
}

func NewSecretStore(specs *SecretSpecs, logger log.Logger) (*hashicorp.Store, error) {
	cfg := client.NewConfig(specs.Address, specs.Namespace)
	cli, err := client.NewClient(cfg)
	if err != nil {
		errMessage := "failed to instantiate Hashicorp client (secrets)"
		logger.WithError(err).Error(errMessage, "specs", specs)
		return nil, errors.ConfigError(errMessage)
	}

	if specs.Token != "" {
		cli.SetToken(specs.Token)
	} else if specs.TokenPath != "" {
		tokenWatcher, err := token.NewRenewTokenWatcher(cli, specs.TokenPath, logger)
		if err != nil {
			return nil, err
		}

		go func() {
			err = tokenWatcher.Start(context.Background())
			if err != nil {
				logger.WithError(err).Error("token watcher has exited with errors")
			} else {
				logger.Warn("token watcher has exited gracefully")
			}
		}()
	}

	store := hashicorp.New(cli, specs.MountPoint, logger)
	return store, nil
}
