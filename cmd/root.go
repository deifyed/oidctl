package cmd

import (
	"fmt"
	"os"

	"github.com/deifyed/oidctl/pkg/core"
	"github.com/spf13/cobra"
)

var (
	config            core.Config
	discoveryDocument core.DiscoveryDocument
)

var rootCmd = &cobra.Command{
	Use:   "oidctl",
	Short: "oidctl simplifies acquiring tokens from an OIDC authentication provider",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		config = core.LoadConfig()
		if err = config.Validate(); err != nil {
			return fmt.Errorf("validating config: %w", err)
		}

		discoveryDocument, err = core.GetDiscoveryDocument(config.DiscoveryURL)
		if err != nil {
			return fmt.Errorf("fetching discovery document: %w", err)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
