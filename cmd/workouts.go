package cmd

import (
	"fmt"
	"strings"

	"github.com/ajbosco/rides/peloton"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newWorkoutsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "workouts",
		Short: "Lists workouts you've completed",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: workouts,
	}

	flags := cmd.Flags()
	flags.String("category", "cycling", "Workout type to display, defaults to 'cycling'")
	flags.String("limit", "10", "Maximum number of workouts to display")

	cmd.AddCommand(
		newDownloadCmd(),
	)

	return cmd
}

func workouts(cmd *cobra.Command, args []string) error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	client, err := getAuthenticatedClient(username, password)
	if err != nil {
		return fmt.Errorf("failed to get authenticated client, %w", err)
	}

	limit := viper.GetInt("limit")
	workoutType := strings.ToLower(viper.GetString("category"))

	user, err := client.GetUser()
	if err != nil {
		return fmt.Errorf("failed to get user, %w", err)
	}

	workouts, err := client.GetUserWorkoutsCSV(user.ID)
	if err != nil {
		return fmt.Errorf("failed to get user workout csv, %w", err)
	}
	if workoutType != "" {
		var filterWorkouts []peloton.WorkoutCSV
		for _, w := range workouts {
			if strings.ToLower(w.FitnessDiscipline) == workoutType {
				filterWorkouts = append(filterWorkouts, w)
			}
		}
		workouts = filterWorkouts
	}
	if len(workouts) > limit {
		workouts = workouts[len(workouts)-limit:]
	}
	table := newTable("Timestamp", "Workout", "Instructor", "type", "Total Output", "Avg Watts", "Distance (mi)")
	for _, w := range workouts {
		table.Append([]string{w.WorkoutTimestamp, w.Title, w.InstructorName, w.FitnessDiscipline, w.TotalOutput, w.AvgWatts, w.DistanceMiles})
	}
	fmt.Printf("Your last %d %s workouts:\n", limit, workoutType)
	table.Render()
	return nil
}
