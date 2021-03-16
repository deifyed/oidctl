package cmd

import (
	"fmt"
	"os"

	"github.com/deifyed/oidctl/pkg/flows/clientcredentials"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCredentialsCmd)
}

var clientCredentialsCmd = &cobra.Command{
	Use:     "client-credentials AUDIENCE",
	Short:   "Authenticate using the client credentials flow",
	Aliases: []string{"cc"},
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		flow := clientcredentials.NewClientCredentials(args[0])

		token, err := flow.Authenticate(discoveryDocument, config.ClientID, config.ClientSecret)
		if err != nil {
			return fmt.Errorf("authenticating with client credentials flow: %w", err)
		}

		fmt.Fprint(os.Stdout, token.Data)

		return nil
	},
}
