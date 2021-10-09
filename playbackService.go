package main

import (
	"fmt"
	"time"

	robot "github.com/go-vgo/robotgo"
)

type PlaybackService struct {
	isPlaying bool
}

func NewPlaybackService() *PlaybackService {
	s := PlaybackService{isPlaying: false}
	// go func() {
	// 	actions := <-recordedActionsChannel
	// 	timeline := Timeline{
	// 		Name:    "1",
	// 		Actions: actions,
	// 	}
	// 	s.addTimeline(timeline)
	// 	fmt.Println(timeline)
	// }()
	return &s
}

func (s *PlaybackService) playTimeline(tl Timeline) {
	fmt.Println("PLAYING TIMELINE")
	s.isPlaying = true
	defer func() { s.isPlaying = false }()
	startTime := time.Now()
	var timer *time.Timer
	for _, offsetAction := range tl.Actions {
		currentOffset := time.Since(startTime)
		fmt.Println("Action offset: ", offsetAction.Offset)
		if currentOffset >= offsetAction.Offset {
			s.executeAction(offsetAction.Action)
		} else {
			fmt.Println("Time to wait until next execute: ", offsetAction.Offset-currentOffset, " for action: ", offsetAction.Action)
			timer = time.NewTimer(offsetAction.Offset-currentOffset)
			<- timer.C
			fmt.Println("Timer is ready, time to execute action: ", offsetAction.Action)
			s.executeAction(offsetAction.Action)
		}
	}
}

func (s *PlaybackService) executeAction(action SimpleAction) {
	fmt.Println("Executing action", action)
	switch kind := action.Kind; kind {
	case 3, 4:
		robot.KeyDown(action.KeyString)
		break

	case 5:
		robot.KeyUp(action.KeyString)
		break
	}
}
