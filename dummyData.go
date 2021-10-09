package main

import "time"

func GetDummyTimeline() Timeline {
	return Timeline{
		Name: "dummy",
		Actions: []OffsetAction{
			{
				Offset: time.Millisecond * 1000,
				Action: SimpleAction{
					Kind:      4,
					Code:      40,
					KeyString: "up",
				},
			},

			{
				Offset: time.Millisecond * 3000,
				Action: SimpleAction{
					Kind:      5,
					Code:      40,
					KeyString: "up",
				},
			},
		},
	}
}
