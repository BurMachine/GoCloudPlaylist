package timeConverting

import (
	"GoCloudPlaylist/internal/playlist"
	"fmt"
	"time"
)

func ConvertFromSongProcToString(songProc playlist.SongProcessing) string {
	s1 := ConvertFromSecondsToString(songProc.Duration)
	s2 := ConvertFromSecondsToString(songProc.CurrentTime)
	return fmt.Sprintf("%s of %s", s2, s1)
}

func ConvertFromSecondsToString(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	Hours := int(duration.Hours())
	Minutes := int(duration.Minutes()) % 60
	Seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", Hours, Minutes, Seconds)
}
