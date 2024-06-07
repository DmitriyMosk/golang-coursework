package etl

import (
	"log"
	"sync"
	"time"

	"golang-coursework/cmd/jira"
	"golang-coursework/database"

	"gorm.io/gorm"
)

func LoadIssues(db *gorm.DB, jiraClient *jira.JiraClient, projectKey string, threadCount int, issueInOneRequest int, maxTimeSleep int, minTimeSleep int) error {
	var wg sync.WaitGroup

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func(startAt int) {
			defer wg.Done()

			for {
				issues, err := jiraClient.GetIssues(projectKey, startAt, issueInOneRequest)
				if err != nil {
					log.Printf("failed to get issues: %v", err)
					time.Sleep(time.Duration(minTimeSleep) * time.Millisecond)
					continue
				}

				if len(issues) == 0 {
					break
				}

				for _, Issue := range issues {
					dbIssue := database.Issue{
						ID:          Issue.ID,
						Key:         Issue.Key,
						Summary:     Issue.Summary,
						Description: Issue.Description,
					}

					if err := db.Create(&dbIssue).Error; err != nil {
						log.Printf("failed to save issue: %v", err)
					}
				}

				startAt += len(issues)
				time.Sleep(time.Duration(minTimeSleep) * time.Millisecond)
			}
		}(i * issueInOneRequest)
	}

	wg.Wait()
	return nil
}

/*
dbIssue := database.Issue{
						ID:          Issue.id,
						Key:         issue.Key,
						Summary:     issue.Fields.Summary,
						Description: issue.Fields.Description,
					}

*/
