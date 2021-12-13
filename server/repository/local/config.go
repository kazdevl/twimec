package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kazdevl/twimec/domain"
)

type ConfigRepository struct {
	storagePath string
}

func NewConfigRepository(sp string) *ConfigRepository {
	return &ConfigRepository{
		storagePath: sp,
	}
}

func (c *ConfigRepository) Get(contentID string) *domain.ConfigContentAcquisition {
	contentDir := filepath.Join(c.storagePath, contentID)
	f, err := os.Open(fmt.Sprintf("%s/acquisition.json", contentDir))
	defer f.Close()
	if err != nil {
		return nil
	}
	var config domain.ConfigContentAcquisition
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil
	}
	return &config
}

func (c *ConfigRepository) All() []domain.ConfigContentAcquisition {
	des, err := os.ReadDir(c.storagePath)
	if err != nil {
		return nil
	}
	result := make([]domain.ConfigContentAcquisition, len(des))
	for index, de := range des {
		contentDir := filepath.Join(c.storagePath, de.Name())
		f, err := os.Open(fmt.Sprintf("%s/acquisition.json", contentDir))
		defer f.Close()
		if err != nil {
			return nil
		}
		var config domain.ConfigContentAcquisition
		if err := json.NewDecoder(f).Decode(&config); err != nil {
			return nil
		}

		result[index] = config
	}
	return result
}

func (c *ConfigRepository) Store(config domain.ConfigContentAcquisition) error {
	contentDir := filepath.Join(c.storagePath, config.ContentID)
	f, err := os.Create(fmt.Sprintf("%s/acquisition.json", contentDir))
	defer f.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(config)
	return err
}

func (c *ConfigRepository) Update(config domain.ConfigContentAcquisition) error {
	contentDir := filepath.Join(c.storagePath, config.ContentID)
	f, err := os.Create(fmt.Sprintf("%s/acquisition.json", contentDir))
	defer f.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(config)
	return err
}
