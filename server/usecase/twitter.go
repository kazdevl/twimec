package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kazdevl/twimec/domain"
	"github.com/sivchari/gotwtr"
)

type TwitterClient interface {
	FetchContent(name, keyword string, previous time.Time) ([]domain.Pages, time.Time, error)
}

type TClient struct {
	c         *gotwtr.Client
	allowLike bool
}

func NewTClient(allowLike bool, token string) *TClient {
	return &TClient{
		c:         gotwtr.New(token),
		allowLike: allowLike,
	}
}

func (t *TClient) FetchContent(name, keyword string, previous time.Time) ([]domain.Pages, time.Time, error) {
	query := fmt.Sprintf("from:%s -is:retweet \"%s\"", name, keyword)
	res, err := t.c.SearchRecentTweets(context.Background(), query, &gotwtr.SearchTweetsOption{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionAttachmentsMediaKeys},
		MediaFields: []gotwtr.MediaField{gotwtr.MediaFieldMediaKey, gotwtr.MediaFieldURL},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldAttachments, gotwtr.TweetFieldCreatedAt},
		MaxResults:  100,
		StartTime:   previous,
	})
	if err != nil {
		return nil, previous, err
	}
	if len(res.Tweets) == 0 {
		return nil, previous, errors.New("no tweets")
	}
	pagesList := make([]domain.Pages, len(res.Tweets))
	for index, tweet := range res.Tweets {
		if t.allowLike {
			t.like(tweet.AuthorID, tweet.ID)
		}
		pagesList[index] = t.getTweetImageLinks(tweet, res.Includes.Media)
	}
	latestTime, _ := time.Parse(time.RFC3339, res.Tweets[len(res.Tweets)-1].CreatedAt)
	return pagesList, latestTime, nil
}

func (t *TClient) like(userID, tweetID string) {
	// TODO impl
}

func (t *TClient) getTweetImageLinks(tweet *gotwtr.Tweet, media []*gotwtr.Media) domain.Pages {
	links := make(domain.Pages, 0, 4)
	for _, key := range tweet.Attachments.MediaKeys {
		for _, m := range media {
			if key == m.MediaKey {
				links = append(links, m.URL)
			}
		}
	}
	return links
}
