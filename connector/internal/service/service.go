package service

import (
	"golang-coursework/connector/config"
	"golang-coursework/connector/internal/models"
	"golang-coursework/connector/internal/repository"
	"golang-coursework/connector/pkg/logger"
)

type Connector interface {
	GetProjectIssues(projectName string) ([]models.Issue, error)
	GetProjects(limit int, page int, search string) ([]models.Project, models.Page, error)
}

type Transformer interface {
	TransformData(issues []models.Issue) []models.TransformedIssue
}

type Services struct {
	Connector   Connector
	Transformer Transformer
}

type ServicesDependencies struct {
	JiraRepositoryUrl string
}

func NewServices(repositories *repository.Repositories, deps ServicesDependencies, log *logger.Logger, cfg *config.Config) *Services {
	return &Services{
		Connector:   NewConnectorService(repositories.ConnectorRepository, deps.JiraRepositoryUrl, log, cfg),
		Transformer: NewTransformerService(),
	}
}
