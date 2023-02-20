package playlist

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc() string {
	if pl.current.currentElem == nil {
		pl.RequestChan <- SongProcessing{Exist: false}
		pl.current.currentElem = pl.current.currentElem.Prev()
	} else {
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "next"
	}
	return ""
}

// prevChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) prevChannelsProc() string {
	if pl.current.currentElem == nil {
		pl.RequestChan <- SongProcessing{Exist: false}
	} else {
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true}

		return "prev"
	}
	return ""
}
