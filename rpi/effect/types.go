package effect

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/color"
)

type State struct {
	Time   time.Duration
	Colors []color.RGBA
}

type Context struct {
	LedCount  int
	StartTime time.Duration
	Duration  time.Duration
}

type Effect func(state State) State
