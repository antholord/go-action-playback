package main

import "fmt"

var defaultTimelineKey = "1"
type TimelineService struct {
	timelines map[string]Timeline
}

func NewTimelineService(recordedActionsChannel chan []OffsetAction) *TimelineService {
	s := TimelineService{
		timelines: make(map[string]Timeline),
	}

	go func() {
		actions := <-recordedActionsChannel
		timeline := Timeline{
			Name:    defaultTimelineKey,
			Actions: actions,
		}
		s.addTimeline(timeline)
		fmt.Println(timeline)
	}()
	return &s
}

func (s *TimelineService) addTimeline(t Timeline) {
	s.timelines[t.Name] = t
}

func (s *TimelineService) getTimelines() map[string]Timeline {
	return s.timelines
}

func (s *TimelineService) getTimeline() Timeline {
	if timeline, ok := s.timelines[defaultTimelineKey]; ok {
		return timeline
	} else {
		return GetDummyTimeline()
	}
}
