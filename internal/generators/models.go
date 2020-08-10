package generators

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

// generateModelsCmd represents the generateModels command
var generateModelsCmd = &cobra.Command{
	Use:   "generate:user",
	Short: "Generates a generic user model",
	Long: `Generates a generic user model implementing Authable, 
	Recoverable, Confirmable, Lockable, and the OAuthable module, 
	into the specified folder`,
	Run: func(cmd *cobra.Command, args []string) {
		err := generateModels(cmd)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			os.Exit(-1)
		}
	},
}

func generateModels(cmd *cobra.Command) error {
	dst, _ := cmd.Flags().GetString("destination")

	if len(dst) == 0 {
		return fmt.Errorf("You must specify a destination path")
	}
	_, c, _, _ := runtime.Caller(0)
	tplDir := path.Join(c, "../../../examples/models/user.go")
	err := CopyDir(tplDir, dst)
	if err != nil {
		return fmt.Errorf("Could not create the directory: %w", err)
	}
	fmt.Println("User model generated into: " + dst)

	return nil
}

func init() {
	rootCmd.AddCommand(generateModelsCmd)
	generateModelsCmd.Flags().StringP("destination", "d", "", "the destination path of the generated user model")
}
