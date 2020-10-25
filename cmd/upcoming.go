package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newUpcomingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upcoming",
		Short: "Show upcoming workouts you can add to your schedule",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: upcoming,
	}

	flags := cmd.Flags()
	flags.String("category", "cycling", "Workout type to display")
	flags.String("start", time.Now().Format("2006-01-02"), "Start Date to fetch upcoming workouts from")
	flags.String("end", time.Now().AddDate(0, 0, 2).Format("2006-01-02"), "End Date to fetch upcoming workouts until")

	return cmd
}

func upcoming(cmd *cobra.Command, args []string) error {
	username := viper.GetString("username")
	password := viper.GetString("password")
	client, err := getAuthenticatedClient(username, password)
	if err != nil {
		return fmt.Errorf("failed to get authenticated client, %w", err)
	}

	workoutType := strings.ToLower(viper.GetString("category"))
	startDate := viper.GetString("start")
	endDate := viper.GetString("end")
	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return fmt.Errorf("failed to parse start date, %w", err)
	}
	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return fmt.Errorf("failed to parse end date, %w", err)
	}
	start := int(startTime.Unix())
	end := int(endTime.Unix())

	schedule, err := client.GetSchedule(workoutType, start, end)
	if err != nil {
		return fmt.Errorf("failed to get schedule, %w", err)
	}

	table := newTable("Workout ID", "Workout", "Instructor", "Start Time", "Scheduled?")
	for _, w := range schedule.Rides {
		var scheduledStart int
		var isScheduled string
		var workoutID string
		for _, d := range schedule.Data {
			if d.RideID == w.ID {
				workoutID = d.ID
				scheduledStart = d.ScheduledStartTime
				isScheduled = strconv.FormatBool(d.AuthedUserReservationID != nil)
			}
		}
		classStartTime := time.Unix(int64(scheduledStart), 0).Format("Mon, 02 Jan 2006 15:04:05 MST")
		i, err := client.GetInstructorByID(w.InstructorID)
		if err != nil {
			return fmt.Errorf("failed to get instructor name, %w", err)
		}
		table.Append([]string{workoutID, w.Title, i.FirstName + " " + i.LastName, classStartTime, isScheduled})
	}
	fmt.Printf("Upcoming %s workouts from %s to %s:\n", workoutType, startDate, endDate)
	table.Render()

	return nil
}
