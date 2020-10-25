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
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download our workouts to a CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		password := viper.GetString("password")
		client, err := getAuthenticatedClient(username, password)
		if err != nil {
			log.Fatal(err)
		}

		location := viper.GetString("location")

		user, err := client.GetUser()
		if err != nil {
			log.Fatal(err)
		}

		workouts, err := client.GetUserWorkoutsCSV(user.ID)
		f, err := os.Create(location)
		if err != nil {
			log.Fatal(err)
		}
		err = gocsv.MarshalFile(workouts, f)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	workoutsCmd.AddCommand(downloadCmd)
	flags := downloadCmd.Flags()
	flags.String("location", "my-workouts.csv", "File location to save csv")
	viper.BindPFlag("location", flags.Lookup("location"))
}
