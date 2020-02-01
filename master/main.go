package main

import (
	"fmt"
	"net/http"
	"time"
)

type action struct {
	Sequence []int
	Duration time.Duration
}

var (
	commandInterval        = time.Duration(getIntEnv("COMMAND_INTERVAL", 1)) * time.Minute
	rainForecastUrl        = getStrEnv("RAIN_FORECAST_URL", "https://graphdata.buienradar.nl/2.0/forecast/geo/Rain3Hour?lat=52.357&lon=4.94")
	precipitationThreshold = getFloatEnv("PRECIPITAION_THRESHOLD", 0.5)
	rainIntesityThreshold  = getIntEnv("INTENSITY_THRESHOLD", 30)
)

type Forecasts struct {
	Forecasts []Forecast
}

type Forecast struct {
	UTCDateTime   string  `json:"utcdatetime"`
	Precipitation float64 `json:"precipitation"`
	RainIntensity int     `json:"value"`
}

func main() {
	numOfRandomPatterns := len(randomPatterns)

	for {
		willRain, err := willItRain()
		if err != nil {
			fmt.Println(err)
		}
		data := rainPattern
		if !willRain {
			data = randomPatterns[generateRandomPatternNum(numOfRandomPatterns)]
		}

		len := len(data) - 1
		for i := 0; i <= len; i++ {
			entry := data[i]
			sequence := entry.Sequence

			for n, numericalState := range sequence {
				color, state := mapNumberToLEDState(numericalState)
				if i == 0 || data[i-1].Sequence[n] != numericalState {
					resp, err := http.Get(fmt.Sprintf("http://minion-%d:3333/%s/%d", n, color, state))

					if err != nil {
						fmt.Println(err)
						break
					}

					resp.Body.Close()
					resp.Close = true
				}
			}
			time.Sleep(entry.Duration)

		}

		time.Sleep(commandInterval)
	}

}
