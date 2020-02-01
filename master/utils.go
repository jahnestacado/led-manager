package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getStrEnv(name, fallback string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		value = fallback
	}
	return value
}

func getIntEnv(name string, fallback int) int {
	strValue, exists := os.LookupEnv(name)
	if !exists {
		return fallback
	} else {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			fmt.Println(err)
			value = fallback
		}
		return value
	}
}

func getFloatEnv(name string, fallback float64) float64 {
	strValue, exists := os.LookupEnv(name)
	if !exists {
		return fallback
	} else {
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			fmt.Println(err)
			value = fallback
		}
		return value
	}
}

func generateRandomPatternNum(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(max)
}

func mapNumberToLEDState(i int) (string, int) {
	switch i {
	case 1:
		{
			return "green", 1
		}
	case 2:
		{
			return "red", 1
		}
	default:
		{
			return "green", 0
		}
	}
}

func willItRain() (bool, error) {
	resp, err := http.Get(rainForecastUrl)
	if err != nil {
		return false, err
	}
	var forecasts Forecasts
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	defer func() {
		resp.Body.Close()
		resp.Close = true
	}()

	err = json.Unmarshal(bodyBytes, &forecasts)
	if err != nil {
		return false, err
	}

	for _, forecast := range forecasts.Forecasts {
		predictionTime, err := time.Parse("2006-01-02T15:04:05", forecast.UTCDateTime)
		if err != nil {
			return false, err
		}
		now := time.Now()
		in30Mins := now.Add(30 * time.Minute)

		if predictionTime.Before(in30Mins) &&
			forecast.Precipitation > precipitationThreshold &&
			forecast.RainIntensity > rainIntesityThreshold {
			return true, nil
		}
	}

	return false, nil
}
