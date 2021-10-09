package main

import (
	"time"
)

type SimpleAction struct {
	Kind      uint8
	Code      uint16
	KeyString string
}

// func (a SimpleAction) String() string {
// 	return fmt.Sprintf()
// }

type OffsetAction struct {
	Offset time.Duration
	Action SimpleAction
}
