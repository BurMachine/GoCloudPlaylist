package playlist

import "container/list"

// nextChannelsProc - отправка данных в каналы, повторяющийся код
func (pl *Playlist) nextChannelsProc(elem *list.Element) string {
	if elem.Next() == nil {
		pl.RequestChan <- SongProcessing{exist: false}
	} else {
		newElem := elem.Next()
		el, _ := newElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration, exist: true}
		return "next"
	}
	return ""
}
