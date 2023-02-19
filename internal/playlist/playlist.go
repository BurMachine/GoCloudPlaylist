package playlist

import (
	"container/list"
	_ "container/list"
	"fmt"
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
		playing:     true,
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
			continue // ??
		}

		if pl.playing {
			el, ok := elem.Value.(Song)
			if !ok {
				println(123)
			}
			for i := 0; i < el.Duration; i++ {
				action := pl.playingProc(elem, el, i)
				if action == "next" {
					break
				} else if action == "prev" {

				}
			}
			elem = elem.Next()
			continue
		} else {
			el, ok := elem.Value.(Song)
			if !ok {
				println(123)
			}
			action := pl.pausedProc(elem, el)
			if action == "next" {
				elem = elem.Next()
			}
			continue
		}
	}
}

func (pl *Playlist) playingProc(elem *list.Element, el Song, i int) string {
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
			pl.Logger.Info().Msg(fmt.Sprintf("%s continued on %d/%d", el.Name, i, el.Duration))
			break
		case <-pl.NextChan:
			return pl.nextChannelsProc(elem)
		}
	case <-pl.NextChan:
		return pl.nextChannelsProc(elem)
	default:
		time.Sleep(time.Second)
	}
	return ""
}

func (pl *Playlist) pausedProc(elem *list.Element, el Song) string {
	select {
	case <-pl.PlayChan:
		pl.RequestChan <- SongProcessing{
			name:        el.Name,
			currentTime: 0,
			duration:    el.Duration,
		}
		pl.Logger.Info().Msg(fmt.Sprintf("%s continued on %d/%d", el.Name, 0, el.Duration))
		pl.playing = true
		break
	case <-pl.NextChan:
		return pl.nextChannelsProc(elem)
	}
	return ""
}
