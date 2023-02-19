package playlist

type SongProcessing struct {
	name        string
	currentTime int
	duration    int
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
	return data
}
