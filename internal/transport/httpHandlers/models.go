package httpHandlers

type ProcessingResponse struct {
	Name     string `json:"song_name"`
	Duration string `json:"song_duration"`
	Status   string `json:"playback_status"`
}

type Song struct {
	Name     string `json:"song_name"`
	Duration string `json:"song_duration"`
}

type SongName struct {
	Name string `json:"song_name"`
}
