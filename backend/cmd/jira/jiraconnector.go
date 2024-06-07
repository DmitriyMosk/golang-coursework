package jira

import (
	"encoding/json"
	"fmt"
	"golang-coursework/cmd/config"
	"golang-coursework/database"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type PageInfo struct {
	PageCount     int `json:"pageCount"`
	CurrentPage   int `json:"currentPage"`
	ProjectsCount int `json:"projectsCount"`
}

func UpdateProjectIssues(projectKey string) error {
	url := fmt.Sprintf("%s/rest/api/2/project/%s", config.GConfig.ProgramSettings.JiraURL, projectKey)
	logrus.Infof("Requesting issues for project: %s", url) // Логирование URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Errorf("Failed to create request: %v", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Failed to send request: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Failed to get issues: %v", resp.Status)
		return fmt.Errorf("failed to get issues: %v", resp.Status)
	}

	var responseBody struct {
		Issues []database.Issue `json:"issues"`
	}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		logrus.Errorf("Failed to decode response body: %v", err)
		return err
	}

	logrus.Infof("Number of issues found: %d", len(responseBody.Issues)) // Логирование количества задач

	if len(responseBody.Issues) == 0 {
		logrus.Warn("No issues found for the given project key")
	}

	// Multithreaded upload to DB
	var wg sync.WaitGroup
	for _, issue := range responseBody.Issues {
		wg.Add(1)
		go func(issue database.Issue) {
			defer wg.Done()
			// Save issue to DB
			if err := database.DB.Create(&issue).Error; err != nil {
				logrus.Errorf("Failed to save issue to DB: %v", err)
			} else {
				logrus.Infof("Successfully saved issue to DB: %s", issue.ID)
			}
		}(issue)
	}
	wg.Wait()

	return nil
}

func GetProjects(limit, page int, search string) ([]database.Project, PageInfo, error) {
	url := fmt.Sprintf("%s/rest/api/2/project", config.GConfig.ProgramSettings.JiraURL)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("1")
		return nil, PageInfo{}, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("2")
		fmt.Println("failed to send request: %v", err)
		return nil, PageInfo{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("3")
		return nil, PageInfo{}, fmt.Errorf("unexpected status code: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("4")

		return nil, PageInfo{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var projects []database.Project
	if err := json.Unmarshal(body, &projects); err != nil {
		fmt.Println("5")

		return nil, PageInfo{}, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Filter projects by search query
	var filteredProjects []database.Project
	for _, project := range projects {
		if search == "" || contains(project.Name, search) || contains(project.Key, search) {
			filteredProjects = append(filteredProjects, project)
		}
	}

	// Paginate
	start := (page - 1) * limit
	end := start + limit
	if start > len(filteredProjects) {
		start = len(filteredProjects)
	}
	if end > len(filteredProjects) {
		end = len(filteredProjects)
	}

	pageInfo := PageInfo{
		PageCount:     (len(filteredProjects) + limit - 1) / limit,
		CurrentPage:   page,
		ProjectsCount: len(filteredProjects),
	}

	return filteredProjects[start:end], pageInfo, nil
}

func contains(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

/*

fmt.Println("--------------------------------------------------------")
fmt.Println(issues)
fmt.Println("--------------------------------------------------------")

*/
