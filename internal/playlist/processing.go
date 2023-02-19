package playlist

import (
	"errors"
	"fmt"
)

type SongProcessing struct {
	name        string
	currentTime int
	duration    int

	playing bool
	exist   bool
}

func (pl *Playlist) Play() SongProcessing {
	var data SongProcessing
	pl.mutex.RLock()
	pl.PlayChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.playing = true
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("playing %v ", data))
	return data
}

func (pl *Playlist) Pause() SongProcessing {
	var data SongProcessing
	pl.mutex.RLock()
	pl.StopChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.playing = false
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("paused %v ", data))
	return data
}

func (pl *Playlist) Next() SongProcessing {
	var data SongProcessing
	pl.mutex.RLock()
	pl.current.currentElem = pl.current.currentElem.Next()
	pl.NextChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("next %v ", data))
	return data
}

func (pl *Playlist) Prev() SongProcessing {
	var data SongProcessing
	pl.mutex.RLock()
	pl.current.currentElem = pl.current.currentElem.Prev()
	pl.PrevChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("prev %v ", data))
	return data
}

func (pl *Playlist) Status() SongProcessing {
	var data SongProcessing
	pl.mutex.RLock()
	pl.StatusChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("status %v ", data))
	return data
}

// Изменение плейлиста

func (pl *Playlist) AddNewSong(song Song) bool {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	el := pl.list.PushBack(song)
	if el == nil {
		return false
	}
	pl.Logger.Info().Msg(fmt.Sprintf("added %v ", song))
	return true
}

func (pl *Playlist) GetList() ([]Song, error) {
	var res []Song
	pl.mutex.Lock()
	defer pl.mutex.Unlock()
	for e := pl.list.Front(); e != nil; e = e.Next() {
		tmp, ok := e.Value.(Song)
		if !ok {
			return res, errors.New("element to Song converting error")
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (pl *Playlist) DeleteSong(name string) error {
	pl.mutex.Lock()
	defer pl.mutex.Unlock()

	pl.StatusChan <- struct{}{}
	var data SongProcessing
	select {
	case data = <-pl.RequestChan:
		break
	}

	el, ok := pl.current.currentElem.Value.(Song)
	if !ok {
		return errors.New("element to Song converting error")
	}
	for e := pl.list.Front(); e != nil; e = e.Next() {
		tmp, ok := e.Value.(Song)
		if !ok {
			return errors.New("element to Song converting error")
		}
		if tmp.Name == name {
			if name == el.Name {
				if data.playing {
					return errors.New("can't delete song while playing")
				}
				pl.list.Remove(e)
				pl.Logger.Info().Msg(fmt.Sprintf("deleted %v ", tmp))
				break
			} else {
				pl.list.Remove(e)
				pl.Logger.Info().Msg(fmt.Sprintf("deleted %v ", tmp))
				break
			}

		}
	}

	return nil
}
