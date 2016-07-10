package ufc

type Fighter struct {
	Id                 int    `json:"id"`
	Nickname           string `json:"nickname"`
	Wins               int    `json:"wins"`
	StatId             int    `json:"statid"`
	Losses             int    `json:"losses"`
	LastName           string `json:"last_name"`
	WeightClass        string `json:"weight_class"`
	TitleHolder        bool   `json:"title_holder"`
	Draws              int    `json:"draws"`
	FirstName          string `json:"first_name"`
	FighterStatus      string `json:"fighter_status"`
	Rank               string `json:"rank"`
	PoundForPoundRank  string `json:"pound_for_pound_rank"`
	Thumbnail          string `json:"thumbnail"`
	BeltThumbnail      string `json:"belt_thumbnail"`
	LeftFullBodyImage  string `json:"left_full_body_image"`
	RightFullBodyImage string `json:"right_full_body_image"`
	ProfileImage       string `json:"profile_image"`
	Link               string `json:"link"`

	// attributes below are only present on single fighter
	CityResiding             string  `json:"city_residing"`
	Dob                      string  `json:"dob"`
	StoreLink                string  `json:"store_link"`
	Height                   int     `json:"height"`
	CountryResiding          string  `json:"country_residing"`
	KoTkoWins                int     `json:"ko_tko_wins"`
	SubmissionWins           int     `json:"submission_wins"`
	DecisionWins             int     `json:"decision_wins"`
	WeightKg                 int     `json:"weight_kg"`
	AverageFightTime         float32 `json:"AverageFightTime"`
	AverageFightTime_Seconds float32 `json:"AverageFightTime_Seconds"`
	KDAverage                float32 `json:"KDAverage"`
	SLpM                     float32 `json:"SLpM"`
	StrikingAccuracy         float32 `json:"StrikingAccuracy"`
	SApM                     float32 `json:"SApM"`
	StrikingDefense          float32 `json:"StrikingDefense"`
	TakedownAverage          float32 `json:"TakedownAverage"`
	TakedownAccuracy         float32 `json:"TakedownAccuracy"`
	TakedownDefense          float32 `json:"TakedownDefense"`
	SubmissionsAverage       float32 `json:"SubmissionsAverage"`
	AvgSLpM                  float32 `json:"avg_SLpM"`
	MaxSLpM                  float32 `json:"max_SLpM"`
	AvgSApM                  float32 `json:"avg_SApM"`
	MaxSApM                  float32 `json:"max_SApM"`
	AvgTD                    float32 `json:"avg_TD"`
	MaxTD                    float32 `json:"max_TD"`
	AvgSubmissionAttempts    float32 `json:"avg_submission_attempts"`
	MaxSubmissionAttempts    float32 `json:"max_submission_attempts"`
}
