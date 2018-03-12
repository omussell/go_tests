package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	Long: `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
	},
}

func init() {
	//nameCmd.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//nameCmd.Flags().IntVarP(&echoSeed, "seed", "s", 0, "seed for entropy")
	rootCmd.PersistentFlags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
        viper.BindPFlag("times", rootCmd.PersistentFlags().Lookup("times"))
	rootCmd.PersistentFlags().IntVarP(&echoSeed, "seed", "s", 0, "seed for entropy")
        viper.BindPFlag("seed", rootCmd.PersistentFlags().Lookup("seed"))
	passwordCmd.Flags().IntVarP(&pwLength, "length", "l", 32, "length of password")
        viper.BindPFlag("length", passwordCmd.PersistentFlags().Lookup("length"))

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

var nameCmd = &cobra.Command{
  Use:   "name",
  Short: "Print a name",
  Long:  `Prints a name, default random, use --seed to
make it less random`,
  Run: func(cmd *cobra.Command, args []string) {
    //gofakeit.Seed(int64(echoSeed))
    gofakeit.Seed(int64(viper.GetInt("seed")))
    //for i := 0; i < echoTimes; i++ {
    for i := 0; i < viper.GetInt("times"); i++ {
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
    gofakeit.Seed(int64(viper.GetInt("seed")))
    //for i := 0; i < echoTimes; i++ {
    for i := 0; i < viper.GetInt("times"); i++ {
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
    gofakeit.Seed(int64(viper.GetInt("seed")))
    //for i := 0; i < echoTimes; i++ {
    for i := 0; i < viper.GetInt("times"); i++ {
      //fmt.Println(gofakeit.Password(pwLower, pwUpper, pwNumeric, pwSpecial, pwSpace, pwLength))
      fmt.Println(gofakeit.Password(true, true, true, true, false, pwLength))
    }
  },
}
