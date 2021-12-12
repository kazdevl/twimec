package local

import "github.com/kazdevl/twimec/domain"

type ContentInfoRepository struct {
	StorageLocation string
}

func NewContentInfoRepository(sl string) *ContentInfoRepository {
	return &ContentInfoRepository{
		StorageLocation: sl,
	}
}

func (c *ContentInfoRepository) GetAuthorName(contentID string) string {
	// TODO impl
	return ""
}

func (c *ContentInfoRepository) All() []domain.ContentInfo {
	// TODO impl
	return nil
}

func (c *ContentInfoRepository) Store(content domain.ContentInfo) error {
	// TODO impl
	return nil
}

func (c *ContentInfoRepository) Update(content domain.ContentInfo) error {
	// TODO impl
	return nil
}
