package weather

import "github.com/michelpessoa/desafioCloudRun/internal/entity"

type WeatherProviderInterface interface {
	GetWeather(city string) (entity.Weather, error)
}
