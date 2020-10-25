package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Display information about authenticated user",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: user,
	}

	return cmd
}

func user(cmd *cobra.Command, args []string) error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	client, err := getAuthenticatedClient(username, password)
	if err != nil {
		return fmt.Errorf("failed to get authenticated client, %w", err)
	}

	user, err := client.GetUser()
	if err != nil {
		return fmt.Errorf("failed to get user, %w", err)
	}
	table := newTable("ID", "First Name", "Last Name", "Username", "FTP", "Total Workouts")
	table.Append([]string{user.ID, user.FirstName, user.LastName, user.Username, strconv.Itoa(user.Ftp), strconv.Itoa(user.TotalWorkouts)})
	fmt.Println("Current Authenticated User:")
	table.Render()
	return nil
}
