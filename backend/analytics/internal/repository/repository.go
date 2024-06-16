package repository

import (
	"database/sql"
	postrege_sql "golang-coursework/backend/analytics/internal/repository/postrege-sql"
)

type Repositories struct {
	AnalyticsRepository IAnalyticsRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		AnalyticsRepository: postrege_sql.NewAnalyticsRepository(db),
	}
}
