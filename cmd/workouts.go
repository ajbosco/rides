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
	"strings"

	"github.com/ajbosco/rides/peloton"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// workoutsCmd represents the workouts command
var workoutsCmd = &cobra.Command{
	Use:   "workouts",
	Short: "Lists workouts you've completed",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client, err := getAuthenticatedClient(username, password)
		if err != nil {
			log.Fatal(err)
		}

		limit := viper.GetInt("limit")
		workoutType := strings.ToLower(viper.GetString("type"))

		user, err := client.GetUser()
		if err != nil {
			log.Fatal(err)
		}

		workouts, err := client.GetUserWorkoutsCSV(user.ID)
		if err != nil {
			log.Fatal(err)
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
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(workoutsCmd)
	flags := workoutsCmd.Flags()
	flags.String("limit", "10", "Maximum number of workouts to display")
	viper.BindPFlag("limit", flags.Lookup("limit"))
	flags.String("type", "", "Workout type to display, if blank all will show")
	viper.BindPFlag("type", flags.Lookup("type"))
}
