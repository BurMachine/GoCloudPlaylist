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
	NextChan   chan bool
	PrevChan   chan bool
	StatusChan chan struct{}

	// Каналы ответа
	RequestChan chan SongProcessing

	// Graceful shutdown
	ExitChan chan struct{}

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
		NextChan:    make(chan bool),
		PrevChan:    make(chan bool),
		StatusChan:  make(chan struct{}),
		RequestChan: make(chan SongProcessing),
		ExitChan:    make(chan struct{}),
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
				} else if action == "exit" {
					return
				} else if i == el.Duration-1 {
					pl.current.currentElem = pl.current.currentElem.Next()
					break
				}
			}
			continue
		} else {
			action := pl.pausedProc()
			if action == "exit" {
				return
			}
			continue
		}
	}
}

func (pl *Playlist) playingProc(i int) string {

	select {
	case <-pl.StopChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: i, Duration: el.Duration}
		select {
		case <-pl.PlayChan:
			el, _ = pl.current.currentElem.Value.(Song)
			pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: i, Duration: el.Duration, Exist: true, Playing: true}
			break
		case <-pl.StopChan:
			pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: i, Duration: el.Duration, Exist: true, Playing: true}
		case data := <-pl.NextChan:
			if data {
				return pl.nextChannelsProc()
			} else {
				pl.RequestChan <- SongProcessing{Exist: false}
				return "next"
			}
		case data := <-pl.PrevChan:
			if data {
				return pl.prevChannelsProc()
			} else {
				pl.RequestChan <- SongProcessing{Exist: false}
				return "next"
			}
		case <-pl.StatusChan:
			pl.RequestChan <- SongProcessing{Name: el.Name, Duration: el.Duration, CurrentTime: i, Playing: false}
		case <-pl.ExitChan:
			return "exit"
		}
	case <-pl.PlayChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: i, Duration: el.Duration, Exist: true, Playing: true}
	case data := <-pl.NextChan:
		if data {
			return pl.nextChannelsProc()
		} else {
			pl.RequestChan <- SongProcessing{Exist: false}
			return "next"
		}
	case data := <-pl.PrevChan:
		if data {
			return pl.prevChannelsProc()
		} else {
			pl.RequestChan <- SongProcessing{Exist: false}
			return "next"
		}
	case <-pl.StatusChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, Duration: el.Duration, CurrentTime: i, Playing: true}
	case <-pl.ExitChan:
		return "exit"
	default:
		time.Sleep(time.Second)
	}
	return ""
}

func (pl *Playlist) pausedProc() string {
	select {
	case <-pl.PlayChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true, Playing: true}
		pl.playing = true
		break
	case <-pl.StopChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, CurrentTime: 0, Duration: el.Duration, Exist: true, Playing: true}
		break
	case data := <-pl.NextChan:
		if data {
			return pl.nextChannelsProc()
		} else {
			pl.RequestChan <- SongProcessing{Exist: false}
			return "next"
		}
	case data := <-pl.PrevChan:
		if data {
			return pl.prevChannelsProc()
		} else {
			pl.RequestChan <- SongProcessing{Exist: false}
			return "next"
		}
	case <-pl.StatusChan:
		el, _ := pl.current.currentElem.Value.(Song)
		pl.RequestChan <- SongProcessing{Name: el.Name, Duration: el.Duration, CurrentTime: 0, Playing: false}
		break
	case <-pl.ExitChan:
		return "exit"
	}
	return ""
}
