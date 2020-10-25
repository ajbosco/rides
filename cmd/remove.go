package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove a workout from your schedule",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: remove,
	}

	flags := cmd.Flags()
	flags.String("workout-id", "", "Workout ID to remove from your schedule")

	return cmd
}

func remove(cmd *cobra.Command, args []string) error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	client, err := getAuthenticatedClient(username, password)
	if err != nil {
		return fmt.Errorf("failed to get authenticated client, %w", err)
	}

	workoutID := viper.GetString("workout-id")
	if workoutID == "" {
		return fmt.Errorf("workout-id is required, %w", err)
	}

	err = client.RemoveReservation(workoutID)
	if err != nil {
		return fmt.Errorf("failed to remove reservation, %w", err)
	}
	fmt.Printf("Removed workout %q from your schedule!\n", workoutID)
	return nil
}
