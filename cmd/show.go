package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/ajbosco/rides/peloton"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show your upcoming workout schedule",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
		RunE: show,
	}

	flags := cmd.Flags()
	flags.String("category", "cycling", "Workout type to display, defaults to 'cycling'")
	flags.String("start", time.Now().Format("2006-01-02"), "Start Date to fetch upcoming workouts from, defaults to Today")
	flags.String("end", time.Now().AddDate(0, 0, 2).Format("2006-01-02"), "End Date to fetch upcoming workouts until, defaults to Tomorrow")

	return cmd
}

func show(cmd *cobra.Command, args []string) error {
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

	userSchedule, err := client.GetUserSchedule(workoutType, start, end)
	if err != nil {
		return fmt.Errorf("failed to get user schedule, %w", err)
	}

	table := newTable("Workout ID", "Workout", "Instructor", "Start Time")

	for _, u := range userSchedule {
		scheduledStart := u.ScheduledStartTime
		classStartTime := time.Unix(int64(scheduledStart), 0).Format("Mon, 02 Jan 2006 15:04:05 MST")
		var title string
		var instructor peloton.Instructor
		for _, r := range schedule.Rides {
			if u.RideID == r.ID {
				title = r.Title
				instructor, err = client.GetInstructorByID(r.InstructorID)
				if err != nil {
					return fmt.Errorf("failed to get instructor name, %w", err)
				}
			}
		}
		workoutID := fmt.Sprintf("%v", u.AuthedUserReservationID)
		table.Append([]string{workoutID, title, instructor.FirstName + " " + instructor.LastName, classStartTime})
	}
	fmt.Printf("Your scheduled %s workouts from %s to %s:\n", workoutType, startDate, endDate)
	table.Render()

	return nil
}
