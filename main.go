package main

import (
	"fmt"
	"time"

	hook "github.com/antholord/gohook"
)

var startRecordingChannel = make(chan bool)
var stopRecordingChannel = make(chan bool)
var timelineRecorded = make(chan []OffsetAction)

var isRecording = false
var timelineService = NewTimelineService(timelineRecorded)
var playbackService = NewPlaybackService()

func main() {
	exit := make(chan bool)

	eventChannel := hook.Start()
	defer hook.End()

	fmt.Println(timelineService)
	go startKeyboardListener(eventChannel)

	<-exit
	fmt.Println("EXITING")
}

func startKeyboardListener(eventChannel chan hook.Event) {
	recordingChannel := make(chan hook.Event, 100000)
	fmt.Println("KEYBOARD LISTENER STARTED")

	go func() {
		for range startRecordingChannel {
			if !isRecording {
				go record(recordingChannel)
			}
		}
	}()

	for ev := range eventChannel {
		handleActionKeyDowns(ev)
		if isRecording {
			recordingChannel <- ev
		}
	}
}

func record(eventChannel chan hook.Event) {
	isRecording = true
	fmt.Println("RECORDING STARTED")

	startTime := time.Now()
	keyboardTimeline := make([]OffsetAction, 0)
	go func() {
		for range stopRecordingChannel {
			isRecording = false
			timelineRecorded <- keyboardTimeline
		}
	}()

	for ev := range eventChannel {
		handleActionKeyDowns(ev)
		if ev.Kind == hook.KeyDown || ev.Kind == hook.KeyUp || ev.Kind == hook.KeyHold {
			sa := SimpleAction{Kind: ev.Kind, Code: ev.Rawcode, KeyString: Rawcode2String[ev.Rawcode]}
			if sa.Kind == hook.KeyHold {
				sa.Kind = hook.KeyDown
			}
			fmt.Println(sa)
			keyboardTimeline = append(keyboardTimeline, OffsetAction{Offset: time.Since(startTime), Action: sa})
		}
	}
}

func handleActionKeyDowns(event hook.Event) {
	if event.Kind == hook.KeyUp {
		if event.Rawcode == 27 && isRecording {
			stopRecording()
			return
		}
		if event.Rawcode == 121 {
			if isRecording {
				stopRecording()
			} else {
				fmt.Println("F10 PRESSED :::::: STARTED RECORDING")
				startRecordingChannel <- true
			}
			return
		}
		if event.Rawcode == 113 { //F2
			playbackService.playTimeline(timelineService.getTimeline())
			return
		}
	}
}

func stopRecording() {
	if isRecording {
		isRecording = false
		fmt.Println("STOPPED RECORDING")
		stopRecordingChannel <- true
	}
}
