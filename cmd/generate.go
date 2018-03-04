package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/brianvoe/gofakeit"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random data",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(versionCmd)
	generateCmd.AddCommand(nameCmd)
//	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number",
  Long:  `All software has versions. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}

var nameCmd = &cobra.Command{
  Use:   "name",
  Short: "Print a name",
  Long:  `Prints a name, default random, use --seed to
make it less random`,
  Run: func(cmd *cobra.Command, args []string) {
    gofakeit.Seed(0)
    fmt.Println(gofakeit.Name())
  },
}
