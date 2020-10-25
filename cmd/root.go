package cmd

import (
	"fmt"
	"os"

	"github.com/ajbosco/rides/peloton"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rides",
	Short: "A CLI for intereacting with the Peloton API",
}

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "rides",
		Short:        "A CLI for intereacting with the Peloton API",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
				return err
			}
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.String("username", "", "Username for Peloton")
	flags.String("password", "", "Password for Peloton")

	cmd.AddCommand(
		newUserCmd(),
		newUpcomingCmd(),
		newScheduleCmd(),
		newWorkoutsCmd(),
	)

	return cmd
}

//Execute is the main entrypoint function
func Execute() {
	if err := newRootCmd(os.Args[1:]).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rides" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rides")
	}

	viper.SetEnvPrefix("RIDES")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	viper.ReadInConfig()
}

func getAuthenticatedClient(username, password string) (*peloton.Client, error) {
	client := peloton.NewClient(username, password)
	err := client.Authenticate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newTable(headers ...string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetHeader(headers)

	return table
}
