package ufc

type Event struct {
	Id                       int     `json:"id"`
	EventDate                string  `json:"event_date"`
	SecondaryFeatureImage    string  `json:"secondary_feature_image"`
	TicketImage              string  `json:"ticket_image"`
	EventTimeZoneText        string  `json:"event_time_zone_text"`
	ShortDescription         string  `json:"short_description"`
	EventDategmt             string  `json:"event_dategmt"`
	EndEventDategmt          string  `json:"end_event_dategmt"`
	TicketUrl                string  `json:"ticketurl"`
	TicketSellerName         string  `json:"ticket_seller_name"`
	BaseTitle                string  `json:"base_title"`
	TitleTagLine             string  `json:"title_tag_line"`
	TwitterHashtag           string  `json:"twitter_hashtag"`
	TicketGeneralSaleDate    string  `json:"ticket_general_sale_date"`
	Statid                   int     `json:"statid"`
	FeatureImage             string  `json:"feature_image"`
	EventTimeText            string  `json:"event_time_text"`
	TicketGeneralSaleText    string  `json:"ticket_general_sale_text"`
	Subtitle                 string  `json:"subtitle"`
	EventStatus              string  `json:"event_status"`
	IsPpvEvent               bool    `json:"isppvevent "`
	CornorAudioAvailable     bool    `json:"corner_audio_available "`
	CornerAudioBlueStreamUrl string  `json:"corner_audio_blue_stream_url"`
	CornerAudioRedStreamUrl  string  `json:"corner_audio_red_stream_url"`
	LastModified             string  `json:"last_modified"`
	UrlName                  string  `json:"url_name"`
	Created                  string  `json:"created"`
	TrailerUrl               string  `json:"trailer_url"`
	Arena                    string  `json"arena"`
	Location                 string  `json:"location"`
	FmFntFeedUrl             string  `json:"fm_fnt_feed_url"`
	MainEventFighter1Id      int     `json:"main_event_fighter1_id"`
	MainEventFighter2Id      int     `json:"main_event_fighter2_id"`
	Latitude                 float32 `json:"latitude"`
	Longitude                float32 `json:"longitude"`
}
