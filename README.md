# rides

[![Github Actions](https://github.com/ajbosco/rides/workflows/build/badge.svg?branch=master&event=push)](https://github.com/ajbosco/rides/actions?workflow=build)
[![Go Report Card](https://goreportcard.com/badge/github.com/ajbosco/rides?style=flat-square)](https://goreportcard.com/report/github.com/ajbosco/rides)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/ajbosco/rides/peloton)

The [Peloton](https://www.onepeloton.com) CLI you never wanted.

- [Authentication](#authentication)
- [Usage](#usage)
  * [View Upcoming Workouts](#view-upcoming-workouts)
  * [Show Scheduled Workouts](#show-scheduled-workouts)
  * [Add Workout to Schedule](#add-workout-to-schedule)
  * [Remove Workout from Schedule](#remove-workout-from-schedule)
  * [View Completed Workouts](#view-completed-workouts)
  * [Download Completed Workouts](#download-completed-workouts)


## Authentication

1. Create a `.rides.yaml` file in your home directory
```console
username: your-username
password: your-password
```
 
## Usage

```console
$ rides -h
A CLI for intereacting with the Peloton API

Usage:
  rides [command]

Available Commands:
  help        Help about any command
  schedule    Show your upcoming workout schedule and add/remove to it
  upcoming    Show upcoming workouts you can add to your schedule
  user        Display information about authenticated user
  workouts    Lists workouts you've completed

Flags:
  -h, --help              help for rides
      --password string   Password for Peloton
      --username string   Username for Peloton

Use "rides [command] --help" for more information about a command.
```

### View Upcoming Workouts

```console
$ rides upcoming -h
Show upcoming workouts you can add to your schedule

Usage:
  rides upcoming [flags]

Flags:
      --category string   Workout type to display (default "cycling")
      --end string        End Date to fetch upcoming workouts until (default "2020-10-27")
  -h, --help              help for upcoming
      --start string      Start Date to fetch upcoming workouts from (default "2020-10-25")

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton
```

### Show Scheduled Workouts

```console
$ rides schedule show -h
Show your upcoming workout schedule

Usage:
  rides schedule show [flags]

Flags:
      --category string   Workout type to display (default "cycling")
      --end string        End Date to fetch upcoming workouts until (default "2020-10-27")
  -h, --help              help for show
      --start string      Start Date to fetch upcoming workouts from (default "2020-10-25")

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton
```

### Add Workout to Schedule

```console
$ rides schedule add -h
Add a workout to your schedule

Usage:
  rides schedule add [flags]

Flags:
  -h, --help                help for add
      --workout-id string   Workout ID to add to your schedule

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton
```

### Remove Workout from Schedule

```console
$ rides schedule remove -h
Remove a workout from your schedule

Usage:
  rides schedule remove [flags]

Flags:
  -h, --help                help for remove
      --workout-id string   Workout ID to remove from your schedule

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton
```

### View Completed Workouts

```console
$ rides workouts -h
Lists workouts you've completed

Usage:
  rides workouts [flags]
  rides workouts [command]

Available Commands:
  download    Download our workouts to a CSV file

Flags:
      --category string   Workout type to display (default "")
  -h, --help              help for workouts
      --limit string      Maximum number of workouts to display (default "10")

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton

Use "rides workouts [command] --help" for more information about a command.
```

### Download Completed Workouts

```console
$ rides workouts download -h
Download your workouts to a CSV file

Usage:
  rides workouts download [flags]

Flags:
  -h, --help              help for download
      --location string   File location to save csv (default "my-workouts.csv")

Global Flags:
      --password string   Password for Peloton
      --username string   Username for Peloton
```