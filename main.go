package main

import (
	"flag"
	"fmt"
	"time"

	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	play_sound("assets/bell.wav")
	minutes := flag.Int("minutes", 22, "# minutes")
	flag.Parse()
	pomodoro(*minutes)
	play_sound("assets/bell.wav")
}

func pomodoro(minutes int) {
	seconds := minutes * 60
	ui_bar_and_timer(seconds)
	fmt.Println("Done!")
}

func play_sound(file string) {
	f, _ := os.Open(file)
	s, format, _ := wav.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)
}

func ui_bar_and_timer(seconds int) {
	bar := pb.StartNew(seconds)
	fmt.Println("- Time to focus!")
	for i := 0; i < seconds; i++ {
		time.Sleep(time.Second * time.Duration(1))
		bar.Add(1)
	}
	bar.Finish()
}
