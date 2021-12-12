package local

import "github.com/kazdevl/twimec/domain"

type ChapterRepository struct {
	StorageLocation string
}

func NewChapterRepository(sl string) *ChapterRepository {
	return &ChapterRepository{
		StorageLocation: sl,
	}
}

func (c *ChapterRepository) GetList(contentID string) []domain.Chapter {
	// TODO impl
	return nil
}

func (c *ChapterRepository) Store(chapter domain.Chapter) error {
	// TODO impl
	return nil
}
