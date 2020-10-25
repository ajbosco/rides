package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a workout to your schedule",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: add,
	}

	flags := cmd.Flags()
	flags.String("workout-id", "", "Workout ID to remove from your schedule")

	return cmd
}

func add(cmd *cobra.Command, args []string) error {
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

	user, err := client.GetUser()
	if err != nil {
		return fmt.Errorf("failed to get user, %w", err)
	}

	err = client.CreateReservation(user.ID, workoutID)
	if err != nil {
		return fmt.Errorf("failed to add reservation, %w", err)
	}
	fmt.Printf("Added workout %q to your schedule!\n", workoutID)
	return nil
}
