package spacexapi

import (
	"time"
)

type NextLaunch struct {
	FlightNumber          int           `json:"flight_number"`
	MissionName           string        `json:"mission_name"`
	MissionID             []interface{} `json:"mission_id"`
	LaunchYear            string        `json:"launch_year"`
	LaunchDateUnix        int           `json:"launch_date_unix"`
	LaunchDateUtc         time.Time     `json:"launch_date_utc"`
	LaunchDateLocal       string        `json:"launch_date_local"`
	IsTentative           bool          `json:"is_tentative"`
	TentativeMaxPrecision string        `json:"tentative_max_precision"`
	Tbd                   bool          `json:"tbd"`
	LaunchWindow          interface{}   `json:"launch_window"`
	Rocket                struct {
		RocketID   string `json:"rocket_id"`
		RocketName string `json:"rocket_name"`
		RocketType string `json:"rocket_type"`
		FirstStage struct {
			Cores []struct {
				CoreSerial     string      `json:"core_serial"`
				Flight         int         `json:"flight"`
				Block          int         `json:"block"`
				Gridfins       bool        `json:"gridfins"`
				Legs           bool        `json:"legs"`
				Reused         bool        `json:"reused"`
				LandSuccess    interface{} `json:"land_success"`
				LandingIntent  bool        `json:"landing_intent"`
				LandingType    string      `json:"landing_type"`
				LandingVehicle string      `json:"landing_vehicle"`
			} `json:"cores"`
		} `json:"first_stage"`
		SecondStage struct {
			Block    int `json:"block"`
			Payloads []struct {
				PayloadID      string        `json:"payload_id"`
				NoradID        []interface{} `json:"norad_id"`
				Reused         bool          `json:"reused"`
				Customers      []string      `json:"customers"`
				Nationality    string        `json:"nationality"`
				Manufacturer   interface{}   `json:"manufacturer"`
				PayloadType    string        `json:"payload_type"`
				PayloadMassKg  interface{}   `json:"payload_mass_kg"`
				PayloadMassLbs interface{}   `json:"payload_mass_lbs"`
				Orbit          string        `json:"orbit"`
				OrbitParams    struct {
					ReferenceSystem string      `json:"reference_system"`
					Regime          string      `json:"regime"`
					Longitude       interface{} `json:"longitude"`
					SemiMajorAxisKm interface{} `json:"semi_major_axis_km"`
					Eccentricity    interface{} `json:"eccentricity"`
					PeriapsisKm     interface{} `json:"periapsis_km"`
					ApoapsisKm      interface{} `json:"apoapsis_km"`
					InclinationDeg  interface{} `json:"inclination_deg"`
					PeriodMin       interface{} `json:"period_min"`
					LifespanYears   interface{} `json:"lifespan_years"`
					Epoch           interface{} `json:"epoch"`
					MeanMotion      interface{} `json:"mean_motion"`
					Raan            interface{} `json:"raan"`
					ArgOfPericenter interface{} `json:"arg_of_pericenter"`
					MeanAnomaly     interface{} `json:"mean_anomaly"`
				} `json:"orbit_params"`
			} `json:"payloads"`
		} `json:"second_stage"`
		Fairings struct {
			Reused          bool        `json:"reused"`
			RecoveryAttempt interface{} `json:"recovery_attempt"`
			Recovered       interface{} `json:"recovered"`
			Ship            interface{} `json:"ship"`
		} `json:"fairings"`
	} `json:"rocket"`
	Ships     []interface{} `json:"ships"`
	Telemetry struct {
		FlightClub interface{} `json:"flight_club"`
	} `json:"telemetry"`
	LaunchSite struct {
		SiteID       string `json:"site_id"`
		SiteName     string `json:"site_name"`
		SiteNameLong string `json:"site_name_long"`
	} `json:"launch_site"`
	LaunchSuccess interface{} `json:"launch_success"`
	Links         struct {
		MissionPatch      interface{}   `json:"mission_patch"`
		MissionPatchSmall interface{}   `json:"mission_patch_small"`
		RedditCampaign    interface{}   `json:"reddit_campaign"`
		RedditLaunch      interface{}   `json:"reddit_launch"`
		RedditRecovery    interface{}   `json:"reddit_recovery"`
		RedditMedia       interface{}   `json:"reddit_media"`
		Presskit          interface{}   `json:"presskit"`
		ArticleLink       interface{}   `json:"article_link"`
		Wikipedia         interface{}   `json:"wikipedia"`
		VideoLink         interface{}   `json:"video_link"`
		YoutubeID         interface{}   `json:"youtube_id"`
		FlickrImages      []interface{} `json:"flickr_images"`
	} `json:"links"`
	Details            string      `json:"details"`
	Upcoming           bool        `json:"upcoming"`
	StaticFireDateUtc  interface{} `json:"static_fire_date_utc"`
	StaticFireDateUnix interface{} `json:"static_fire_date_unix"`
	Timeline           interface{} `json:"timeline"`
}
