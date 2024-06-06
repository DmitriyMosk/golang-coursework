package jira

import (
	log "github.com/sirupsen/logrus"
)

type Client struct {
	BaseURL string
	Logger  *log.Logger
}

func NewClient(baseURL string, logger *log.Logger) *Client {
	return &Client{
		BaseURL: baseURL,
		Logger:  logger,
	}
}

func (c *Client) FetchIssues(count int) ([]Issue, error) {
	// реализация для получения задач из JIRA
	return nil, nil
}

type Issue struct {
	// поля задачи
}
