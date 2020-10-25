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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Display information about authenticated user",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client, err := getAuthenticatedClient(username, password)
		if err != nil {
			log.Fatal(err)
		}

		user, err := client.GetUser()
		if err != nil {
			log.Fatal(err)
		}
		table := newTable("ID", "First Name", "Last Name", "Username", "FTP", "Total Workouts")
		table.Append([]string{user.ID, user.FirstName, user.LastName, user.Username, strconv.Itoa(user.Ftp), strconv.Itoa(user.TotalWorkouts)})
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
