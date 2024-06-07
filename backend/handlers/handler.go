package handlers

import (
	"encoding/json"
	"fmt"
	"golang-coursework/cmd/config"
	"golang-coursework/cmd/etl"
	"golang-coursework/cmd/jira"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UpdateProject(db *gorm.DB, jiraClient *jira.JiraClient, config *config.Config, projectKey string) error {
	return etl.LoadIssues(db, jiraClient, projectKey, config.ProgramSettings.ThreadCount, config.ProgramSettings.IssueInOneRequest, config.ProgramSettings.MaxTimeSleep, config.ProgramSettings.MinTimeSleep)
}

func UpdateProjectV2(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр строки запроса "project"
	projectKey := r.URL.Query().Get("project")
	fmt.Println("--------------------------------------------------------")
	fmt.Println(projectKey)
	fmt.Println("--------------------------------------------------------")

	if projectKey == "" {
		logrus.Error("Missing project key")
		http.Error(w, "Missing project key", http.StatusBadRequest)
		return
	}

	// Обновляем задачи проекта
	if err := jira.UpdateProjectIssues(projectKey); err != nil {
		logrus.Errorf("Failed to update project issues: %v", err)
		http.Error(w, "Failed to update project issues", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Project issues updated successfully"))
}

func GetProjectsV2(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 20
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}
	search := r.URL.Query().Get("search")

	projects, pageInfo, err := jira.GetProjects(limit, page, search)
	if err != nil {
		http.Error(w, "Failed to get projects_1", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"projects": projects,
		"pageInfo": map[string]int{
			"currentPage":   page,
			"totalPages":    (len(projects) + limit - 1) / limit,
			"projectsCount": pageInfo.ProjectsCount,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
