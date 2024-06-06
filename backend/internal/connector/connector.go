package connector

import (
	"golang-coursework/pkg/config"
	"golang-coursework/pkg/jira"
	"golang-coursework/pkg/logging"
	"golang-coursework/pkg/storage"
	"time"

	log "github.com/sirupsen/logrus"
)

func Run() error { //TODO add goroutine
	cfg, err := config.LoadConfig("configs/config.yml")
	if err != nil {
		return err
	}

	db, err := storage.NewDatabase(*cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	logger := logging.NewLogger()

	client := jira.NewClient(cfg.Connector.JiraURL, logger)

	retryWait := cfg.Connector.InitialRetryWait
	for {
		issues, err := client.FetchIssues(cfg.Connector.IssuesPerRequest)
		if err != nil {
			log.Println("Error fetching issues, retrying...")
			time.Sleep(time.Duration(retryWait) * time.Second)
			retryWait *= 2
			if retryWait > cfg.Connector.MaxRetryWait {
				return err
			}
			continue
		}

		err = db.SaveIssues(issues)
		if err != nil {
			return err
		}

		retryWait = cfg.Connector.InitialRetryWait
	}
}
