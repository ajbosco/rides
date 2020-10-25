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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a workout to your schedule",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client, err := getAuthenticatedClient(username, password)
		if err != nil {
			log.Fatal(err)
		}

		workoutID := viper.GetString("workout-id")
		if workoutID == "" {
			log.Fatal(fmt.Errorf("workout-id is required"))
		}

		user, err := client.GetUser()
		if err != nil {
			log.Fatal(err)
		}

		err = client.CreateReservation(user.ID, workoutID)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	scheduleCmd.AddCommand(addCmd)
	flags := addCmd.Flags()
	flags.String("workout-id", "", "Workout ID to add to your schedule")
	viper.BindPFlag("workout-id", flags.Lookup("workout-id"))
}