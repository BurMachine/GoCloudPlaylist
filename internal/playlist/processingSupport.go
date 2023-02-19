package playlist

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc() string {
	if pl.current.currentElem == nil {
		pl.RequestChan <- SongProcessing{exist: false}
	} else {
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration, exist: true}

		return "next"
	}
	return ""
}

// prevChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) prevChannelsProc() string {
	if pl.current.currentElem == nil {
		pl.RequestChan <- SongProcessing{exist: false}
	} else {
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration, exist: true}

		return "prev"
	}
	return ""
}
