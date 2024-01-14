package transition

import (
	"math"
	"time"
)

type State struct {
	From     float32
	FromTime time.Duration
	To       float32
	ToTime   time.Duration
}

func NewState(value float32) State {
	return State{
		From:     value,
		FromTime: 0,
		To:       value,
		ToTime:   0,
	}
}

func (fade *State) ValueAt(now time.Duration) float32 {
	if now >= fade.ToTime {
		return fade.To
	}
	if now <= fade.FromTime {
		return fade.From
	}

	nowT := float64(now-fade.FromTime) / float64(fade.ToTime-fade.FromTime)
	return fade.From + (fade.To-fade.From)*float32(nowT)
}

func (fade *State) FadeTo(value float32, now time.Duration, fullDuration time.Duration) {
	nowValue := fade.ValueAt(now)

	fade.From = nowValue
	fade.FromTime = now
	fade.To = value
	fade.ToTime = now + time.Duration(math.Abs(float64(value-nowValue))*float64(fullDuration))
}

func (fade *State) IsCompleteAt(now time.Duration) bool {
	return now >= fade.ToTime
}
