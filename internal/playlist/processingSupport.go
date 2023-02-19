package playlist

import "container/list"

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc(elem *list.Element) string {
	if elem.Next() == nil {
		pl.RequestChan <- SongProcessing{exist: false}
	} else {
		elem = elem.Next()
		el, _ := elem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration, exist: true}

		return "next"
	}
	return ""
}

// prevChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) prevChannelsProc(elem *list.Element) string {
	if elem.Prev() == nil {
		pl.RequestChan <- SongProcessing{exist: false}
	} else {
		elem = elem.Prev()
		el, _ := elem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration, exist: true}

		return "prev"
	}
	return ""
}
