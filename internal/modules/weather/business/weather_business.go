package weatherBusiness

import (
	"gpt-twitter-integration/internal/modules/weather/dto"
	weatherService "gpt-twitter-integration/internal/modules/weather/services"
)

func Get(name string) (dto.Weather, error) {
	weather, err := weatherService.GetWeatherByCityName(name)
	if err != nil {
		return weather, err
	}

	return weather, nil
}
