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
	current *list.Element
	mutex   *sync.RWMutex

	playing  bool
	PlayChan chan struct{}
	StopChan chan struct{}
	NextChan chan struct{}
	PrevChan chan struct{}

	// Каналы ответа
	RequestChan chan SongProcessing

	Logger *zerolog.Logger
}

func Init() *Playlist {
	return &Playlist{
		list:        list.New(),
		current:     nil,
		mutex:       &sync.RWMutex{},
		playing:     false,
		PlayChan:    make(chan struct{}),
		StopChan:    make(chan struct{}),
		NextChan:    make(chan struct{}),
		PrevChan:    make(chan struct{}),
		RequestChan: make(chan SongProcessing),
	}
}

func (pl Playlist) Run() {
	// Временно
	pl.list.PushBack(Song{Name: "Run Free", Duration: 10})
	pl.list.PushBack(Song{Name: "Demolisher", Duration: 11})
	// Временно

	elem := pl.list.Front()
	for {
		if elem == nil {
			println("finished")
			break // ??
		}

		if pl.playing {
			el, ok := elem.Value.(Song)
			if !ok {
				println(123)
			}
			for i := 0; i < el.Duration; i++ {
				action := pl.playingProc(elem, i)
				if action == "next" {
					break
				} else if action == "prev" {
					break
				} else if i == el.Duration-1 {
					elem = elem.Next()
					break
				}
			}
			continue
		} else {
			pl.pausedProc(elem)
			continue
		}
	}
}

func (pl *Playlist) playingProc(elem *list.Element, i int) string {
	el, _ := elem.Value.(Song)
	select {
	case <-pl.StopChan:
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: i, duration: el.Duration}

		select {
		case <-pl.PlayChan:
			pl.RequestChan <- SongProcessing{
				name:        el.Name,
				currentTime: i,
				duration:    el.Duration,
				exist:       true,
			}
			break
		case <-pl.NextChan:
			return pl.nextChannelsProc(elem)
		case <-pl.PrevChan:
			return pl.prevChannelsProc(elem)
		}
	case <-pl.NextChan:
		return pl.nextChannelsProc(elem)
	case <-pl.PrevChan:
		return pl.prevChannelsProc(elem)
	default:
		time.Sleep(time.Second)
	}
	return ""
}

func (pl *Playlist) pausedProc(elem *list.Element) string {
	el, _ := elem.Value.(Song)
	select {
	case <-pl.PlayChan:
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration}
		break
	case <-pl.StopChan:
		pl.RequestChan <- SongProcessing{name: el.Name, currentTime: 0, duration: el.Duration}
		break
	case <-pl.NextChan:
		return pl.nextChannelsProc(elem)
	case <-pl.PrevChan:
		return pl.prevChannelsProc(elem)
	}
	return ""
}
