package local

import (
	"github.com/kazdevl/twimec/domain"
)

type ConfigRepository struct {
	StorageLocation string
}

func NewConfigRepository(sl string) *ConfigRepository {
	return &ConfigRepository{
		StorageLocation: sl,
	}
}

func (c *ConfigRepository) All() []domain.ConfigContentAcquisition {
	// TODO impl
	return nil
}

func (c *ConfigRepository) Store(config domain.ConfigContentAcquisition) error {
	// TODO impl
	return nil
}

func (c *ConfigRepository) Update(config domain.ConfigContentAcquisition) error {
	// TODO impl
	return nil
}
