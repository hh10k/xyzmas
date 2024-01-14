package animation

import (
	"time"

	"kurt.blackwell.id.au/xyzmas/configuration"
)

func toMinutes(hhmm int) float32 {
	return float32((hhmm/100)*60 + hhmm%60)
}

func GetNightFactor(now time.Time, config configuration.Configuration) float32 {
	if !config.NightEnabled {
		return 0
	}

	var t = float32(now.Hour()*60 + now.Minute())

	var sunriseStart = toMinutes(config.SunriseStart)
	if t < sunriseStart {
		return 1.0
	}
	var sunriseEnd = toMinutes(config.SunriseEnd)
	if t < sunriseEnd {
		return 1.0 - (t-sunriseStart)/(sunriseEnd-sunriseStart)
	}
	var sunsetStart = toMinutes(config.SunsetStart)
	if t < sunsetStart {
		return 0.0
	}
	var sunsetEnd = toMinutes(config.SunsetEnd)
	if t < sunsetEnd {
		return (t - sunsetStart) / (sunsetEnd - sunsetStart)
	}

	return 1
}
