package place

import "github.com/michelpessoa/desafioCloudRun/internal/entity"

type PlaceProviderInterface interface {
	GetByCep(cep string) (entity.Place, error)
}
