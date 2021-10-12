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

func main() {
	args := os.Args[1:]

	interval, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("\nERR invalid input: interval argument (%v) is not an integer", args[0])
	}

	f, err := os.Open("notif.mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	// defer streamer.Close()

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	// defer ticker.Stop()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("play")
				shot := buffer.Streamer(0, buffer.Len())
				speaker.Play(shot)
			}

		}
	}()

	fmt.Println("zzz")
	time.Sleep(time.Hour * 2)
}
