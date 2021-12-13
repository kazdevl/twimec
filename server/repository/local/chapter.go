package local

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kazdevl/twimec/domain"
)

type ChapterRepository struct {
	storagePath string
}

func NewChapterRepository(sp string) *ChapterRepository {
	return &ChapterRepository{
		storagePath: sp,
	}
}

func (c *ChapterRepository) GetList(contentID string) []domain.Chapter {
	contentDir := filepath.Join(c.storagePath, contentID)
	des, err := os.ReadDir(contentDir)
	if err != nil {
		return nil
	}
	result := make([]domain.Chapter, len(des))
	for index, de := range des {
		pagesByte, err := os.ReadFile(filepath.Join(contentDir, de.Name()))
		if err != nil {
			return nil
		}
		pages := string(pagesByte)
		pagesSlice := strings.Split(pages, ",")
		result[index] = domain.Chapter{
			ContentID: contentID,
			Index:     index,
			Icon:      pagesSlice[0],
			Pages:     string(pages),
		}
	}
	return result
}

func (c *ChapterRepository) Store(chapter domain.Chapter) error {
	contentDir := filepath.Join(c.storagePath, chapter.ContentID)
	des, err := os.ReadDir(contentDir)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s/%d.txt", contentDir, len(des)))
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(chapter.Pages)
	return nil
}
