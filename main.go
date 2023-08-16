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

const defaultPomodoroMinutes = 22
const bellSoundPath = "assets/bell.wav"

// main é a entrada principal do programa.
func main() {
	minutes := parseCommandLineArguments()
	playSound(bellSoundPath)
	startPomodoro(minutes)
	playSound(bellSoundPath)
}

// parseCommandLineArguments analisa os argumentos da linha de comando e retorna os minutos definidos pelo usuário.
func parseCommandLineArguments() int {
	minutes := flag.Int("minutes", defaultPomodoroMinutes, "Duration of the pomodoro in minutes")
	flag.Parse()
	return *minutes
}

// startPomodoro executa um ciclo de pomodoro pela duração especificada em minutos.
func startPomodoro(minutes int) {
	displayProgressBar(minutes * 60)
	fmt.Println("Done!")
}

// playSound reproduz um arquivo de som WAV do caminho fornecido.
func playSound(file string) {
	f, _ := os.Open(file)
	defer f.Close()

	streamer, format, _ := wav.Decode(f)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
}

// displayProgressBar exibe uma barra de progresso e conta regressivamente a quantidade especificada de segundos.
func displayProgressBar(seconds int) {
	bar := pb.StartNew(seconds)
	fmt.Println("- Time to focus!!")
	
	for i := 0; i < seconds; i++ {
		time.Sleep(time.Second)
		bar.Add(1)
	}
	bar.Finish()
}
