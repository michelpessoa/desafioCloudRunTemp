package mocks

import (
	"github.com/michelpessoa/desafioCloudRun/internal/entity"
	"github.com/michelpessoa/desafioCloudRun/internal/infra/place"

	"github.com/stretchr/testify/mock"
)

var _ place.PlaceProviderInterface = (*PlaceProviderMock)(nil)

type PlaceProviderMock struct {
	mock.Mock
}

func NewPlaceProviderMock() *PlaceProviderMock {
	return &PlaceProviderMock{}
}

func (lm *PlaceProviderMock) GetByCep(cep string) (entity.Place, error) {
	args := lm.Called(cep)
	return args.Get(0).(entity.Place), args.Error(1)
}
