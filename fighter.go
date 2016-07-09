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
}
