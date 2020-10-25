package cmd

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download our workouts to a CSV file",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: download,
	}

	flags := cmd.Flags()
	flags.String("location", "my-workouts.csv", "File location to save csv")

	return cmd
}

func download(cmd *cobra.Command, args []string) error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	client, err := getAuthenticatedClient(username, password)
	if err != nil {
		return fmt.Errorf("failed to get authenticated client, %w", err)
	}

	location := viper.GetString("location")

	user, err := client.GetUser()
	if err != nil {
		return fmt.Errorf("failed to get user, %w", err)
	}

	workouts, err := client.GetUserWorkoutsCSV(user.ID)
	if err != nil {
		return fmt.Errorf("failed to get workouts, %w", err)
	}

	f, err := os.Create(location)
	if err != nil {
		return fmt.Errorf("failed to create csv file, %w", err)
	}

	err = gocsv.MarshalFile(workouts, f)
	if err != nil {
		return fmt.Errorf("failed to download csv file, %w", err)
	}
	fmt.Printf("Saved your workouts to %q\n", location)

	return nil
}
