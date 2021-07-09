package cmd

import (
	"os"

	"github.com/consensys/quorum-key-manager/cmd/flags"
	"github.com/consensys/quorum-key-manager/pkg/common"
	"github.com/consensys/quorum-key-manager/src/infra/log/zap"
	app "github.com/consensys/quorum-key-manager/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newRunCommand() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run application",
		RunE:  runCmd,
		PreRun: func(cmd *cobra.Command, args []string) {
			preRunBindFlags(viper.GetViper(), cmd.Flags(), "key-manager")
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			// TODO: Identify which error code to return
			os.Exit(0)
		},
	}

	flags.HTTPFlags(runCmd.Flags())
	flags.ManifestFlags(runCmd.Flags())
	flags.PGFlags(runCmd.Flags())

	return runCmd
}

func runCmd(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()

	vipr := viper.GetViper()
	cfg := flags.NewAppConfig(vipr)

	logger, err := zap.NewLogger(cfg.Logger)
	if err != nil {
		return err
	}
	defer syncZapLogger(logger)

	appli, err := app.New(cfg, logger)
	if err != nil {
		logger.WithError(err).Error("could not create app")
		return err
	}

	done := make(chan struct{})
	sig := common.NewSignalListener(func(sig os.Signal) {
		logger.With("sig", sig.String()).Warn("signal intercepted")
		if err = appli.Stop(ctx); err != nil {
			logger.WithError(err).Error("application stopped with errors")
		}
		close(done)
	})

	defer sig.Close()

	err = appli.Start(ctx)
	if err != nil {
		logger.WithError(err).Error("application failed to start")
		return err
	}

	<-done

	return nil
}

