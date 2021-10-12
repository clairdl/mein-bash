package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type audioPlayer interface {
	initPlayer()
	play()
}

type audioSource struct {
	streamer beep.StreamSeekCloser
	format   beep.Format
	done     chan bool
}

func (s audioSource) play() {
	fmt.Println("play hit")
	speaker.Play(beep.Seq(s.streamer, beep.Callback(func() {
		s.done <- true
	})))
}

func initPlayer(f *os.File) (*audioSource, error) {
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return nil, err
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	avs := audioSource{
		streamer,
		format,
		make(chan bool),
	}

	return &avs, nil
}

func main() {
	f, err := os.Open("notif.mp3")
	if err != nil {
		log.Fatal(err)
	}
	avs, err := initPlayer(f)
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]

	interval, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("\nERR invalid input: interval argument (%v) is not an integer", args[0])
	}

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	speaker.Init(avs.format.SampleRate, avs.format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, avs.streamer), Paused: false}
	speaker.Play(ctrl)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("play invoked")
				speaker.Clear()
				speaker.Play(beep.Seq(avs.streamer))
			}

		}
	}()

	fmt.Println("zzz")
	time.Sleep(time.Hour * 2)
}
