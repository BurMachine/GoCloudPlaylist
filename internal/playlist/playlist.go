package playlist

import (
	"container/list"
	_ "container/list"
	"fmt"
	"github.com/rs/zerolog"
	"time"

	"sync"
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

				select {
				case <-pl.StopChan:
					pl.RequestChan <- SongProcessing{name: el.Name, currentTime: i, duration: el.Duration}
					pl.Logger.Info().Msg(fmt.Sprintf("%s paused on %d/%d", el.Name, i, el.Duration))
					select {
					case <-pl.PlayChan:
						pl.RequestChan <- SongProcessing{
							name:        el.Name,
							currentTime: i,
							duration:    el.Duration,
						}
						pl.Logger.Info().Msg(fmt.Sprintf("%s continued on %d/%d", el.Name, i, el.Duration))
						break
					}
				default:
					time.Sleep(time.Second)
				}

			}
			elem = elem.Next()
			continue
		} else {
			el, ok := elem.Value.(Song)
			if !ok {
				println(123)
			}
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
			}
			continue

		}
	}
}
