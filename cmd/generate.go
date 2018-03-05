package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/brianvoe/gofakeit"
)
var (
	echoTimes int
	echoSeed  int
//	pwLower   bool
//	pwUpper   bool
//	pwNumeric bool
//        pwSpecial bool
//        pwSpace   bool
        pwLength  int
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
	//nameCmd.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//nameCmd.Flags().IntVarP(&echoSeed, "seed", "s", 0, "seed for entropy")
	rootCmd.PersistentFlags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	rootCmd.PersistentFlags().IntVarP(&echoSeed, "seed", "s", 0, "seed for entropy")
	passwordCmd.Flags().IntVarP(&pwLength, "length", "l", 32, "length of password")

	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(versionCmd)
	generateCmd.AddCommand(nameCmd)
	generateCmd.AddCommand(emailCmd)
	generateCmd.AddCommand(passwordCmd)

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
    gofakeit.Seed(int64(echoSeed))
    for i := 0; i < echoTimes; i++ {
      fmt.Println(gofakeit.Name())
    }
  },
}

var emailCmd = &cobra.Command{
  Use:   "email",
  Short: "Print an email",
  Long:  `Prints an email, default random, use --seed to
make it less random`,
  Run: func(cmd *cobra.Command, args []string) {
    gofakeit.Seed(int64(echoSeed))
    for i := 0; i < echoTimes; i++ {
      fmt.Println(gofakeit.Email())
    }
  },
}

var passwordCmd = &cobra.Command{
  Use:   "password",
  Short: "Print password",
  Long:  `Prints a password, default random, use --seed to
make it less random`,
  Run: func(cmd *cobra.Command, args []string) {
    gofakeit.Seed(int64(echoSeed))
    for i := 0; i < echoTimes; i++ {
      //fmt.Println(gofakeit.Password(pwLower, pwUpper, pwNumeric, pwSpecial, pwSpace, pwLength))
      fmt.Println(gofakeit.Password(true, true, true, true, false, pwLength))
    }
  },
}
