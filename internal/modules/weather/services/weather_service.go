package weatherService

import (
	"fmt"

	"gpt-twitter-integration/internal/modules/weather/dto"
	"gpt-twitter-integration/pkg/env"
	"gpt-twitter-integration/pkg/httpRequest"
)

func GetWeatherByCityName(name string) (dto.Weather, error) {
	key := env.WeatherApiKey
	url := env.WeatherApiUrl

	req := httpRequest.RequestParams{
		Url: fmt.Sprintf("%s?key=%s&q=%s&aqi=no", url, key, name),
	}

	var weather dto.Weather
	err := req.Get(&weather)
	if err != nil {
		return weather, err
	}

	return weather, nil
}
