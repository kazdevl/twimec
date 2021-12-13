package usecase

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/kazdevl/twimec/domain"
	"github.com/kazdevl/twimec/repository"
)

type Cronjob struct {
	tclient     TwitterClient
	chapterRepo repository.ChapterRepository
	configRepo  repository.ConfigRepository
}

func NewCronjob(tc TwitterClient, crr repository.ChapterRepository, cgr repository.ConfigRepository) *Cronjob {
	return &Cronjob{
		tclient:     tc,
		chapterRepo: crr,
		configRepo:  cgr,
	}
}

func (c *Cronjob) FetchContents() {
	configs := c.configRepo.All()
	authors := make([]string, 0, len(configs))
	links := make([]string, 0, len(configs))
	for _, config := range configs {
		pagesList, latestTime, err := c.tclient.FetchContent(config.AuthorName, config.Keyword, config.LatestTime)
		if err != nil {
			log.Println(err)
			continue
		}
		for j, pages := range pagesList {
			if err := c.chapterRepo.Store(domain.Chapter{
				ContentID: config.ContentID,
				Index:     config.LatestChapter + j + 1,
				Icon:      pages[0],
				Pages:     strings.Join(pages, ","),
			}); err != nil {
				log.Println(err)
				continue
			}
		}

		if len(pagesList) != 0 {
			links = append(links, fmt.Sprintf("http://localhost:6666/%s/%d", config.ContentID, config.LatestChapter+1))
			authors = append(authors, config.AuthorName)
		}
		config.LatestTime = latestTime
		config.LatestChapter = config.LatestChapter + len(pagesList)
		if err := c.configRepo.Store(config); err != nil {
			log.Println(err)
			continue
		}
	}
	c.Notify(authors)
	c.OpenNewContents(links)
}

func (c *Cronjob) Notify(authros []string) error {
	if len(authros) == 0 {
		return errors.New("zero length")
	}
	body := fmt.Sprintf("You can read content by the following authors\n%s", strings.Join(authros, "\n"))
	err := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "New Contents Avairable!!" subtitle "Check Your Safari" sound name "Hero"`, body)).Start()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cronjob) OpenNewContents(links []string) error {
	if len(links) == 0 {
		return errors.New("zero length")
	}
	command := `
	tell application "Safari"
		open location "%s"
		activate
	end tell
	`
	for _, link := range links {
		err := exec.Command("osascript", "-e", fmt.Sprintf(command, link)).Start()
		if err != nil {
			return err
		}
	}
	return nil
}
