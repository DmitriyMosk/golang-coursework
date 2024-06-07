package jira

import (
	"encoding/json"
	"fmt"
	"golang-coursework/database"
	"io/ioutil"
	"net/http"
	"time"
)

type JiraClient struct {
	BaseURL string
	Client  *http.Client
}

func NewJiraClient(baseURL string) *JiraClient {
	return &JiraClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

func (jc *JiraClient) GetIssues(projectKey string, startAt int, maxResults int) ([]database.Issue, error) {
	url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&startAt=%d&maxResults=%d", jc.BaseURL, projectKey, startAt, maxResults)

	resp, err := jc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get issues: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result database.SearchResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.Issues, nil
}
