package playlist

import (
	"container/list"
	_ "container/list"
	"github.com/rs/zerolog"
	"sync"
	"time"
)

type Song struct {
	Name     string
	Duration int
}

type Playlist struct {
	list    *list.List
	current struct {
		currentElem *list.Element
		time        int
	}
	mutex *sync.RWMutex

	playing    bool
	PlayChan   chan struct{}
	StopChan   chan struct{}
	NextChan   chan struct{}
	PrevChan   chan struct{}
	StatusChan chan struct{}

	// Каналы ответа
	RequestChan chan SongProcessing

	Logger *zerolog.Logger
}

func Init() *Playlist {
	return &Playlist{
		list: list.New(),
		current: struct {
			currentElem *list.Element
			time        int
		}{},
		mutex:       &sync.RWMutex{},
		playing:     false,
		PlayChan:    make(chan struct{}),
		StopChan:    make(chan struct{}),
		NextChan:    make(chan struct{}),
		PrevChan:    make(chan struct{}),
		StatusChan:  make(chan struct{}),
		RequestChan: make(chan SongProcessing),
	}
}

func (pl *Playlist) Run() {
	// Временно
	pl.list.PushBack(Song{Name: "Run Free", Duration: 10})
	pl.list.PushBack(Song{Name: "Demolisher", Duration: 11})
	// Временно

	pl.current.currentElem = pl.list.Front()
	for {
		if pl.current.currentElem == nil {
			println("finished")
			break
		}
		if pl.playing {
			el, ok := pl.current.currentElem.Value.(Song)
			if !ok {
				println(123)
			}
			for i := 0; i < el.Duration; i++ {
				action := pl.playingProc(i)
				if action == "next" {
					break
				} else if action == "prev" {
					break
				} else if i == el.Duration-1 {
					pl.current.currentElem = pl.current.currentElem.Next()
					break
				}
			}
			continue
		} else {
			pl.pausedProc()
			continue
		}
	}
}

func (pl *Playlist) playingProc(i int) string {

	select {
	case <-pl.StopChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: i, duration: el.Duration}
		select {
		case <-pl.PlayChan:
			el, _ = pl.current.currentElem.Value.(Song)
			pl.RequestChan <- SongProcessing{name: el.Name, currentTime: i, duration: el.Duration, exist: true}
			break
		case <-pl.NextChan:
			return pl.nextChannelsProc()
		case <-pl.PrevChan:
			return pl.prevChannelsProc()
		case <-pl.StatusChan:
			pl.RequestChan <- SongProcessing{name: el.Name, duration: el.Duration, currentTime: i}
		}
	case <-pl.NextChan:
		return pl.nextChannelsProc()
	case <-pl.PrevChan:
		return pl.prevChannelsProc()
	case <-pl.StatusChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, duration: el.Duration, currentTime: i}
	default:
		time.Sleep(time.Second)
	}
	return ""
}

func (pl *Playlist) pausedProc() string {
	select {
	case <-pl.PlayChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration}
		pl.playing = true
		break
	case <-pl.StopChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration}
		break
	case <-pl.NextChan:
		return pl.nextChannelsProc()
	case <-pl.PrevChan:
		return pl.prevChannelsProc()
	case <-pl.StatusChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{name: el.Name, duration: el.Duration, currentTime: 0}
		break
	}
	return ""
}
