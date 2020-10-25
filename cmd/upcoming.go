/*
Copyright © 2020 Adam Boscarino

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// upcomingCmd represents the upcoming command
var upcomingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "Show upcoming workouts you can add to your schedule",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client, err := getAuthenticatedClient(username, password)
		if err != nil {
			log.Fatal(err)
		}

		workoutType := strings.ToLower(viper.GetString("category"))
		startDate := viper.GetString("start")
		endDate := viper.GetString("end")
		startTime, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			log.Fatal(err)
		}
		endTime, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			log.Fatal(err)
		}
		start := int(startTime.Unix())
		end := int(endTime.Unix())

		schedule, err := client.GetSchedule(workoutType, start, end)
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}
			table.Append([]string{workoutID, w.Title, i.FirstName + " " + i.LastName, classStartTime, isScheduled})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(upcomingCmd)
	flags := upcomingCmd.Flags()
	flags.String("category", "cycling", "Workout type to display, defaults to 'cycling'")
	viper.BindPFlag("category", flags.Lookup("category"))
	flags.String("start", time.Now().Format("2006-01-02"), "Start Date to fetch upcoming workouts from, defaults to Today")
	viper.BindPFlag("start", flags.Lookup("start"))
	flags.String("end", time.Now().AddDate(0, 0, 2).Format("2006-01-02"), "End Date to fetch upcoming workouts until, defaults to Tomorrow")
	viper.BindPFlag("end", flags.Lookup("end"))
}
