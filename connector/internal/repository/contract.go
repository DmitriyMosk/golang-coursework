package repository

import "golang-coursework/connector/internal/models"

type IConnectorRepository interface {
	PushIssues(issues []models.TransformedIssue) error
	CheckProjectExists(title string) (bool, error)
}
