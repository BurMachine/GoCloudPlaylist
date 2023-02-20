package httpHandlers

type processingResponse struct {
	Name     string `json:"song_name"`
	Duration string `json:"song_duration"`
	Status   string `json:"playback_status"`
}
