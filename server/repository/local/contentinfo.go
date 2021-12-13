package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kazdevl/twimec/domain"
)

type ContentInfoRepository struct {
	storagePath string
}

func NewContentInfoRepository(sp string) *ContentInfoRepository {
	return &ContentInfoRepository{
		storagePath: sp,
	}
}

func (c *ContentInfoRepository) GetAuthorName(contentID string) string {
	contentDir := filepath.Join(c.storagePath, contentID)
	f, err := os.Open(fmt.Sprintf("%s/info.json", contentDir))
	defer f.Close()
	if err != nil {
		return ""
	}
	var config domain.ContentInfo
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return ""
	}
	return config.AuthorName
}

func (c *ContentInfoRepository) All() []domain.ContentInfo {
	des, err := os.ReadDir(c.storagePath)
	if err != nil {
		return nil
	}
	result := make([]domain.ContentInfo, len(des))
	for index, de := range des {
		contentDir := filepath.Join(c.storagePath, de.Name())
		f, err := os.Open(fmt.Sprintf("%s/info.json", contentDir))
		defer f.Close()
		if err != nil {
			return nil
		}
		var info domain.ContentInfo
		if err := json.NewDecoder(f).Decode(&info); err != nil {
			return nil
		}

		result[index] = info
	}
	return result
}

func (c *ContentInfoRepository) Store(content domain.ContentInfo) error {
	contentDir := filepath.Join(c.storagePath, content.ID)
	f, err := os.Create(fmt.Sprintf("%s/info.json", contentDir))
	defer f.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(content)
	return err
}

func (c *ContentInfoRepository) Update(content domain.ContentInfo) error {
	contentDir := filepath.Join(c.storagePath, content.ID)
	f, err := os.Create(fmt.Sprintf("%s/info.json", contentDir))
	defer f.Close()
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(content)
	return err
}
