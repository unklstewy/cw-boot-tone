package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"go-morse-code-app/src/morse"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

var noteFreqs = map[int]float64{
	1: 110.0, 2: 116.5, 3: 123.5, 4: 130.8, 5: 138.5, 6: 146.8, 7: 155.5, 8: 164.8, 9: 174.6, 10: 185.0, 11: 196.0, 12: 207.6,
	13: 220.0, 14: 223.0, 15: 247.0, 16: 261.0, 17: 277.0, 18: 294.0, 19: 311.0, 20: 329.6, 21: 349.0, 22: 370.0, 23: 392.0, 24: 415.3,
	25: 440.0, 26: 466.0, 27: 494.0, 28: 523.3, 29: 554.3, 30: 587.3, 31: 622.3, 32: 659.3, 33: 698.5, 34: 740.0, 35: 784.0, 36: 830.6,
	37: 880.0, 38: 932.3, 39: 987.8, 40: 1046.5, 41: 1108.7, 42: 1174.7, 43: 1244.5, 44: 1318.5, 45: 1397.0,
}

func calcUnitDurationMs(wpm int) int {
	return int(1200 / wpm)
}

func playTones(tones string, freq float64, unitDurationMs int) {
	toneList := strings.Split(tones, ",")
	sampleRate := beep.SampleRate(44100)
	var phase float64

	for i := 0; i < len(toneList)-1; i += 2 {
		pitch := toneList[i]
		delay := toneList[i+1]
		units, err := strconv.Atoi(delay)
		if err != nil {
			continue
		}
		dur := time.Duration(units*unitDurationMs) * time.Millisecond
		numSamples := int(sampleRate.N(dur))

		if pitch == "0" {
			// Silence
			silence := beep.Take(numSamples, beep.Silence(-1))
			done := make(chan bool)
			speaker.Play(beep.Seq(silence, beep.Callback(func() { done <- true })))
			<-done
		} else {
			// Tone (sine wave) with phase continuity
			phaseInc := 2 * math.Pi * freq / float64(sampleRate)
			streamer := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
				for i := range samples {
					sample := 0.3 * math.Sin(phase)
					samples[i][0] = sample
					samples[i][1] = sample
					phase += phaseInc
					if phase > 2*math.Pi {
						phase -= 2 * math.Pi
					}
				}
				return len(samples), true
			})
			tone := beep.Take(numSamples, streamer)
			done := make(chan bool)
			speaker.Play(beep.Seq(tone, beep.Callback(func() { done <- true })))
			<-done
		}
	}
}

func main() {
	// Optional flags with defaults
	wpm := flag.Int("wpm", 25, "Morse code speed in words per minute (default 25)")
	tone := flag.Int("tone", 38, "Note value for Morse tone (see OpenGD77 table, default 38 for A# 932Hz)")
	flag.Parse()

	unitDurationMs := calcUnitDurationMs(*wpm)
	freq, ok := noteFreqs[*tone]
	if !ok {
		log.Fatalf("Invalid tone value: %d", *tone)
	}

	// Update Morse code constants for selected tone
	morse.Dot = fmt.Sprintf("%d,2", *tone)
	morse.Dash = fmt.Sprintf("%d,6", *tone)

	if err := speaker.Init(beep.SampleRate(44100), 4410); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to convert to Morse code: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	morseCode := morse.EncodeToMorse(input)
	fmt.Println("Morse Code:", morseCode)

	fmt.Print("Would you like to hear the Morse code? (y/n): ")
	answer, _ := reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(answer)) == "y" {
		playTones(morseCode, freq, unitDurationMs)
	}
}
