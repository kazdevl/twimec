package local

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kazdevl/twimec/domain"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigGet(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample"
	repo := NewConfigRepository(storagePath)
	config := repo.Get(contentID)

	latestTime, err := time.Parse(time.RFC3339, "2021-11-21T14:24:58.000Z")
	assert.NoError(t, err)
	want := domain.ConfigContentAcquisition{
		ContentID:     contentID,
		AuthorName:    "sample",
		Keyword:       "sample",
		LatestChapter: 2,
		LatestTime:    latestTime,
	}
	t.Log(latestTime)
	assert.Equal(t, want, *config)
}

func Test_ConfigAll(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	repo := NewConfigRepository(storagePath)
	configs := repo.All()

	latestTime, err := time.Parse(time.RFC3339, "2021-11-21T14:24:58.000Z")
	assert.NoError(t, err)
	want := []domain.ConfigContentAcquisition{
		{
			ContentID:     "sample",
			AuthorName:    "sample",
			Keyword:       "sample",
			LatestChapter: 2,
			LatestTime:    latestTime,
		},
		{
			ContentID:     "sample1",
			AuthorName:    "sample1",
			Keyword:       "sample1",
			LatestChapter: 2,
			LatestTime:    latestTime,
		},
	}
	assert.Equal(t, want, configs)
}

func Test_ConfigStore(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample3"
	latestTime, err := time.Parse(time.RFC3339, "2021-11-21T14:24:58.000Z")
	assert.NoError(t, err)
	config := domain.ConfigContentAcquisition{
		ContentID:     contentID,
		AuthorName:    "sample3",
		Keyword:       "sample3",
		LatestChapter: 2,
		LatestTime:    latestTime,
	}

	repo := NewConfigRepository(storagePath)
	err = repo.Store(config)
	assert.NoError(t, err)

	filename := fmt.Sprintf("%s/%s/acquisition.json", storagePath, contentID)
	wantByte, err := os.ReadFile(filename)
	t.Log(string(wantByte))
	var want domain.ConfigContentAcquisition
	err = json.Unmarshal(wantByte, &want)
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, want, config)

	err = os.RemoveAll(filepath.Join(storagePath, contentID))
	assert.NoError(t, err)
}

func Test_ConfigUpdate(t *testing.T) {
	storagePath, _ := filepath.Abs("./testdata/config/contents")
	contentID := "sample3"
	contentDir := filepath.Join(storagePath, contentID)
	err := os.Mkdir(contentDir, 0777)
	assert.NoError(t, err)
	latestTime, err := time.Parse(time.RFC3339, "2021-11-21T14:24:58.000Z")
	assert.NoError(t, err)
	config := domain.ConfigContentAcquisition{
		ContentID:     contentID,
		AuthorName:    "sample3",
		Keyword:       "sample3",
		LatestChapter: 0,
		LatestTime:    latestTime,
	}
	f, err := os.Create(fmt.Sprintf("%s/acquisition.json", contentDir))
	assert.NoError(t, err)

	err = json.NewEncoder(f).Encode(config)
	f.Close()
	assert.NoError(t, err)

	repo := NewConfigRepository(storagePath)
	want := domain.ConfigContentAcquisition{
		ContentID:     contentID,
		AuthorName:    "sample4",
		Keyword:       "sample4",
		LatestChapter: 10,
		LatestTime:    latestTime,
	}
	err = repo.Update(want)
	assert.NoError(t, err)

	f, err = os.Open(fmt.Sprintf("%s/acquisition.json", contentDir))
	defer f.Close()
	assert.NoError(t, err)
	var result domain.ConfigContentAcquisition
	err = json.NewDecoder(f).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, want, result)

	err = os.RemoveAll(contentDir)
	assert.NoError(t, err)
}
