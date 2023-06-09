package commands

import (
	"github.com/fredjeck/configserver/pkg/auth"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var RegisterClientCommand = &cobra.Command{
	Use:   "register",
	Short: "Registers a ClientID on the provided repository",
	Long: `Registers a ClientID on the provided repository by generating a dedicated Client Secret
If the ClientID is not provided a new client id will be generated.
	`,
	Run: registerClient,
}

func init() {
	RegisterClientCommand.Flags().StringSliceP("repositories", "r", []string{}, "target repositories - needs to match repositories configured in the configserver.yaml file")
	RegisterClientCommand.Flags().StringP("clientid", "i", "", "client id")
	RootCommand.AddCommand(RegisterClientCommand)
}

func registerClient(cmd *cobra.Command, _ []string) {
	clientId, err := cmd.Flags().GetString("clientid")
	if len(clientId) == 0 || err != nil {
		clientId = uuid.NewString()
	}
	repo, err := cmd.Flags().GetStringSlice("repositories")
	if len(repo) == 0 || err != nil {
		zap.L().Sugar().Fatal("Missing mandatory argument : repositories")
	}

	spec := auth.NewClientSpec(clientId, repo)

	secret, err := spec.ClientSecret(Key)
	if err != nil {
		zap.L().Sugar().Fatalf("Unable to generate client secret for %s : %s", clientId, err.Error())
	}

	zap.L().Sugar().Infof("Repository: %s", repo)
	zap.L().Sugar().Infof("ClientId: %s", clientId)
	zap.L().Sugar().Infof("ClientSecret: %s", secret)
	zap.L().Sugar().Info("Please store the client secret carefully and do not forget to register the ClientID in the configserver.yaml file")
}
