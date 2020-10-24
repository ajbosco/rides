package peloton

// Auth defines the struct for the authentication
type auth struct {
	User     string `json:"username_or_email,omitempty"`
	Password string `json:"password,omitempty"`
}

// User defines the struct for the user object
type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Username       string    `json:"username"`
	Ftp            int       `json:"cycling_workout_ftp"`
	CustomMaxHr    int       `json:"customized_max_heart_rate"`
	TotalWorkOuts  int       `json:"total_workouts"`
	DefaultMaxHr   int       `json:"default_max_heart_rate"`
	EstimatedFtp   int       `json:"estimated_cycling_ftp"`
	DefaultHrZones []float64 `json:"default_heart_rate_zones"`
	Birthday       int       `json:"birthday"`
	Location       string    `json:"location"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	Gender         string    `json:"gender"`
}

// Workouts defines the struct for the workouts object
type Workouts struct {
	Count int `json:"count"`
	Data  []Workout
}

// Workout contains information about a specific workout
type Workout struct {
	FitnessDiscipline string `json:"fitness_discipline"`
	Created           int    `json:"created_at"`
	EndTime           int    `json:"end_time"`
	ID                string `json:"id"`
	StartTime         int    `json:"start_time"`
	UserID            string `json:"user_id"`
	Status            string `json:"status"`
	PersonalRecord    bool   `json:"is_total_work_personal_record"`
}

// Schedule defines the struct for the schedule object
type Schedule struct {
	Total     int                `json:"total"`
	Count     int                `json:"count"`
	SortBy    string             `json:"sort_by"`
	Data      []ScheduledWorkout `json:"data"`
	Equipment []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"equipment"`
	Instructors []struct {
		ID                 string   `json:"id"`
		Name               string   `json:"name"`
		FirstName          string   `json:"first_name"`
		LastName           string   `json:"last_name"`
		UserID             string   `json:"user_id"`
		FitnessDisciplines []string `json:"fitness_disciplines"`
	} `json:"instructors"`
	RideTypes []struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		DisplayName       string `json:"display_name"`
		FitnessDiscipline string `json:"fitness_discipline"`
		IsActive          bool   `json:"is_active"`
		ListOrder         int    `json:"list_order"`
	} `json:"ride_types"`
	ClassTypes []struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		DisplayName       string `json:"display_name"`
		FitnessDiscipline string `json:"fitness_discipline"`
		IsActive          bool   `json:"is_active"`
		ListOrder         int    `json:"list_order"`
	} `json:"class_types"`
	Rides []struct {
		ClassTypeIds                 []string      `json:"class_type_ids"`
		ContentProvider              string        `json:"content_provider"`
		ContentFormat                string        `json:"content_format"`
		Description                  string        `json:"description"`
		DifficultyEstimate           float64       `json:"difficulty_estimate"`
		OverallEstimate              float64       `json:"overall_estimate"`
		DifficultyRatingAvg          float64       `json:"difficulty_rating_avg"`
		DifficultyRatingCount        int           `json:"difficulty_rating_count"`
		DifficultyLevel              interface{}   `json:"difficulty_level"`
		Duration                     int           `json:"duration"`
		EquipmentIds                 []string      `json:"equipment_ids"`
		EquipmentTags                []interface{} `json:"equipment_tags"`
		ExtraImages                  []interface{} `json:"extra_images"`
		FitnessDiscipline            string        `json:"fitness_discipline"`
		FitnessDisciplineDisplayName string        `json:"fitness_discipline_display_name"`
		HasClosedCaptions            bool          `json:"has_closed_captions"`
		HasPedalingMetrics           bool          `json:"has_pedaling_metrics"`
		HomePelotonID                string        `json:"home_peloton_id"`
		ID                           string        `json:"id"`
		ImageURL                     string        `json:"image_url"`
		InstructorID                 string        `json:"instructor_id"`
		IsArchived                   bool          `json:"is_archived"`
		IsClosedCaptionShown         bool          `json:"is_closed_caption_shown"`
		IsExplicit                   bool          `json:"is_explicit"`
		HasFreeMode                  bool          `json:"has_free_mode"`
		IsLiveInStudioOnly           bool          `json:"is_live_in_studio_only"`
		Language                     string        `json:"language"`
		OriginLocale                 string        `json:"origin_locale"`
		Length                       int           `json:"length"`
		LiveStreamID                 string        `json:"live_stream_id"`
		LiveStreamURL                interface{}   `json:"live_stream_url"`
		Location                     string        `json:"location"`
		Metrics                      []string      `json:"metrics"`
		OriginalAirTime              int           `json:"original_air_time"`
		OverallRatingAvg             float64       `json:"overall_rating_avg"`
		OverallRatingCount           int           `json:"overall_rating_count"`
		PedalingStartOffset          int           `json:"pedaling_start_offset"`
		PedalingEndOffset            int           `json:"pedaling_end_offset"`
		PedalingDuration             int           `json:"pedaling_duration"`
		Rating                       int           `json:"rating"`
		RideTypeID                   string        `json:"ride_type_id"`
		RideTypeIds                  []string      `json:"ride_type_ids"`
		SampleVodStreamURL           interface{}   `json:"sample_vod_stream_url"`
		ScheduledStartTime           int           `json:"scheduled_start_time"`
		SeriesID                     string        `json:"series_id"`
		SoldOut                      bool          `json:"sold_out"`
		StudioPelotonID              string        `json:"studio_peloton_id"`
		Title                        string        `json:"title"`
		TotalRatings                 int           `json:"total_ratings"`
		TotalInProgressWorkouts      int           `json:"total_in_progress_workouts"`
		TotalWorkouts                int           `json:"total_workouts"`
		VodStreamURL                 interface{}   `json:"vod_stream_url"`
		VodStreamID                  string        `json:"vod_stream_id"`
		Captions                     []string      `json:"captions"`
	} `json:"rides"`
	Limit            interface{} `json:"limit"`
	BrowseCategories []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		ListOrder      int    `json:"list_order"`
		IconURL        string `json:"icon_url"`
		PortalImageURL string `json:"portal_image_url"`
	} `json:"browse_categories"`
	FitnessDisciplines []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"fitness_disciplines"`
}

// ScheduledWorkout contains data for a scheduled workout
type ScheduledWorkout struct {
	ID                      string      `json:"id"`
	RideID                  string      `json:"ride_id"`
	ServerTime              int         `json:"server_time"`
	ScheduledStartTime      int         `json:"scheduled_start_time"`
	StartTime               int         `json:"start_time"`
	EndTime                 int         `json:"end_time"`
	PedalingStartTime       int         `json:"pedaling_start_time"`
	PedalingEndTime         int         `json:"pedaling_end_time"`
	Status                  string      `json:"status"`
	Countdown               interface{} `json:"countdown"`
	IsLive                  bool        `json:"is_live"`
	IsStudio                bool        `json:"is_studio"`
	IsEncore                bool        `json:"is_encore"`
	SecondsSinceStart       interface{} `json:"seconds_since_start"`
	CreatedAt               int         `json:"created_at"`
	TotalWorkouts           int         `json:"total_workouts"`
	IsComplete              bool        `json:"is_complete"`
	TotalHomeReservations   int         `json:"total_home_reservations"`
	AuthedUserReservationID interface{} `json:"authed_user_reservation_id"`
}

// ScheduleRequest contains data needed to request schedule
type ScheduleRequest struct {
	UserID    string `json:"user_id"`
	WorkoutID string `json:"peloton_id"`
}

// FavoritesRequest contains data needed to request favorite
type FavoritesRequest struct {
	RideID string `json:"ride_id"`
}
