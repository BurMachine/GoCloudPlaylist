package playlist

import "fmt"

type SongProcessing struct {
	name        string
	currentTime int
	duration    int
	exist       bool
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
	return data
}

func (pl Playlist) Pause() SongProcessing {
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
	pl.PrevChan <- struct{}{}
	select {
	case data = <-pl.RequestChan:
		break
	}
	pl.mutex.RUnlock()
	pl.Logger.Info().Msg(fmt.Sprintf("prev %v ", data))
	return data
}
