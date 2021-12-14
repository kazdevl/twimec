package set

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kazdevl/twimec/domain"
	"github.com/kazdevl/twimec/repository"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type contentArgs struct {
	Title      string
	AuthorName string
	Keyword    string
}

func NewCmd(cRepo repository.ConfigRepository, ciRepo repository.ContentInfoRepository) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set -t -a -k",
		Short: "set title, auhtor name, keyword",
		Long:  `set is for setting title, auhtor name, keyword`,
		Run: func(cmd *cobra.Command, args []string) {
			if !validateFlagValues(cmd.Flags()) {
				return
			}
			var c contentArgs
			c.Title, c.AuthorName, c.Keyword = getFlagValues(cmd.Flags())
			err := proccess(c, cRepo, ciRepo)
			log.Println(err)
		},
	}
	cmd.LocalFlags().StringP("title", "t", "", "set title")
	cmd.LocalFlags().StringP("author", "a", "", "set author_name")
	cmd.LocalFlags().StringP("keyword", "k", "", "set keyword")
	return cmd
}

func validateFlagValues(fset *pflag.FlagSet) bool {
	if validateFlagValue(fset, "title") && validateFlagValue(fset, "author") && validateFlagValue(fset, "keyword") {
		return true
	}
	return false
}

func validateFlagValue(fSet *pflag.FlagSet, flag string) bool {
	t, err := fSet.GetString(flag)
	if err != nil {
		return false
	}
	if len(t) == 0 {
		return false
	}
	return true
}

func getFlagValues(fSet *pflag.FlagSet) (t, a, k string) {
	t, _ = fSet.GetString("title")
	a, _ = fSet.GetString("author")
	k, _ = fSet.GetString("keyword")
	return
}

func proccess(c contentArgs, cRepo repository.ConfigRepository, ciRepo repository.ContentInfoRepository) error {
	homeDir, _ := os.UserHomeDir()
	guid := xid.New().String()
	if err := os.Mkdir(filepath.Join(homeDir, "twimec", "storage", "config", "contents", guid), 0777); err != nil {
		return err
	}
	contentinfo := domain.ContentInfo{
		ID:         guid,
		AuthorName: c.AuthorName,
		Title:      c.Title,
	}
	config := domain.ConfigContentAcquisition{
		ContentID:  guid,
		AuthorName: c.AuthorName,
		Keyword:    c.Keyword,
		LatestTime: time.Now().In(time.UTC),
	}
	if err := ciRepo.Store(contentinfo); err != nil {
		return err
	}
	if err := cRepo.Store(config); err != nil {
		return err
	}
	return nil
}
