package playlist

import "GoCloudPlaylist/internal/models"

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc() string {
	if pl.current.currentElem != nil {
		el, _ := pl.current.currentElem.Value.(models.Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "next"
	}
	return ""
}

// prevChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) prevChannelsProc() string {
	if pl.current.currentElem != nil {
		el, _ := pl.current.currentElem.Value.(models.Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "prev"
	}
	return ""
}
