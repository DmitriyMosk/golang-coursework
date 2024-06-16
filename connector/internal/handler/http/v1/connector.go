package v1

import (
	"encoding/json"
	"golang-coursework/connector/config"
	"golang-coursework/connector/internal/models"
	"golang-coursework/connector/internal/repository"
	"golang-coursework/connector/internal/service"
	"golang-coursework/connector/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ConnectorHandler struct {
	connectorSvc   service.Connector
	transformerSvc service.Transformer
	connectorRep   repository.IConnectorRepository
	log            *logger.Logger
	cfg            *config.Config
}

func NewConnectorHandler(services *service.Services, repositories *repository.Repositories, log *logger.Logger, cfg *config.Config) *ConnectorHandler {
	return &ConnectorHandler{
		log:            log,
		connectorSvc:   services.Connector,
		transformerSvc: services.Transformer,
		connectorRep:   repositories.ConnectorRepository,
		cfg:            cfg,
	}
}

func (handler *ConnectorHandler) GetConnectorHandler(router *mux.Router) {
	router.HandleFunc("/updateProject",
		handler.UpdateProject).Methods(http.MethodPost)
	router.HandleFunc("/projects",
		handler.GetProjects).Methods(http.MethodGet)
}

func (handler *ConnectorHandler) UpdateProject(writer http.ResponseWriter, request *http.Request) {
	projectName := request.URL.Query()["project"]
	if len(projectName) == 0 {
		errorWriter(writer, handler.log, "error: no projects in request.", http.StatusBadRequest)
		return
	}

	issues, err := handler.connectorSvc.GetProjectIssues(projectName[0])
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.MarshalIndent(issues, "", "\t")
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	transformedIssues := handler.transformerSvc.TransformData(issues)
	err = handler.connectorRep.PushIssues(transformedIssues)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
}

func (handler *ConnectorHandler) GetProjects(writer http.ResponseWriter, request *http.Request) {

	limit, page, search, err := getProjectParametersFromRequest(request)

	if err != nil {
		errorWriter(writer, handler.log, "error: Incorrect getProjects parameter.", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	projects, pages, err := handler.connectorSvc.GetProjects(limit, page, search)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var issueResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: "/api/v1/connector/projects"},
		},
		Info:    projects,
		Message: "Connector get projects",
		Name:    "",
		PageInfo: models.Page{
			TotalPageCount:     pages.TotalPageCount,
			CurrentPageNumber:  pages.CurrentPageNumber,
			TotalProjectsCount: pages.TotalProjectsCount,
		},
		Status: true,
	}

	response, errResp := json.MarshalIndent(issueResponse, "", "\t")
	_, errWrite := writer.Write(response)

	if errResp != nil || errWrite != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
}

func getProjectParametersFromRequest(request *http.Request) (int, int, string, error) {
	defaultLimit := 20
	defaultPage := 1
	defaultSearch := ""

	var err error
	urlQuery := request.URL.Query()

	if limit, ok := urlQuery["limit"]; ok {
		defaultLimit, err = strconv.Atoi(limit[0])
		if err != nil {
			return defaultLimit, defaultPage, defaultSearch, err
		}
	}

	if page, ok := urlQuery["page"]; ok {
		defaultPage, err = strconv.Atoi(page[0])
		if err != nil {
			return defaultLimit, defaultPage, defaultSearch, err
		}
	}

	if search, ok := urlQuery["search"]; ok {
		defaultSearch = search[0]
	}

	return defaultLimit, defaultPage, defaultSearch, nil
}
