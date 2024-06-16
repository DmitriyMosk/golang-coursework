package v1

import (
	"encoding/json"
	"fmt"
	"golang-coursework/backend/resource/config"
	"golang-coursework/backend/resource/internal/models"
	repository2 "golang-coursework/backend/resource/internal/repository"
	"golang-coursework/backend/resource/pkg/logger"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ResourceHandler struct {
	resourceRep repository2.IResourceRepository
	log         *logger.Logger
	cfg         *config.Config
}

func NewResourceHandler(repositories *repository2.Repositories, log *logger.Logger, cfg *config.Config) *ResourceHandler {
	return &ResourceHandler{
		log:         log,
		resourceRep: repositories.ResourceRepository,
		cfg:         cfg,
	}
}

func (handler *ResourceHandler) GetResourceHandler(router *mux.Router) {
	router.HandleFunc("/issues/{id:[0-9]+}",
		handler.getIssue).Methods(http.MethodGet)
	router.HandleFunc("/project",
		handler.getProject).Methods(http.MethodGet)
	router.HandleFunc("/projects",
		handler.getProjects).Methods(http.MethodGet)

	router.HandleFunc("/issues/",
		handler.postIssue).Methods(http.MethodPost)
	router.HandleFunc("/projects/",
		handler.postProject).Methods(http.MethodPost)

	router.HandleFunc("/project",
		handler.deleteProject).Methods(http.MethodPost)

}

func (handler *ResourceHandler) getIssue(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	issue, err := handler.resourceRep.GetIssueInfo(id)

	fmt.Println(issue)
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var issueResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: fmt.Sprintf("/api/v1/issues/%d", id)},
		},
		Info:    issue,
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(issueResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

}

func (handler *ResourceHandler) getProject(writer http.ResponseWriter, request *http.Request) {
	projectName := request.URL.Query()["project"]
	if len(projectName) == 0 {
		errorWriter(writer, handler.log, "error: no projects in request.", http.StatusBadRequest)
		return
	}

	project, err := handler.resourceRep.GetProjectInfo(projectName[0])

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var projectResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: fmt.Sprintf("/api/v1/resource/project?project=%s", projectName[0])},
		},
		Info:    project,
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(projectResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
}

func (handler *ResourceHandler) getProjects(writer http.ResponseWriter, request *http.Request) {
	projects, err := handler.resourceRep.GetProjects()

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var projectResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: "/api/v1/resource/projects"},
		},
		Info:    projects,
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(projectResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
}

func (handler *ResourceHandler) postIssue(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var issueInfo models.IssueInfo
	err = json.Unmarshal(body, &issueInfo)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = handler.resourceRep.InsertIssue(issueInfo)
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var issuesResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: "/api/v1/resource/issues/"},
		},
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(issuesResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (handler *ResourceHandler) postProject(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var projectInfo models.ProjectInfo
	err = json.Unmarshal(body, &projectInfo)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := handler.resourceRep.InsertProject(projectInfo)
	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var projectResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: fmt.Sprintf("/api/v1/resource/project?project=%d", id)},
		},
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(projectResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *ResourceHandler) deleteProject(writer http.ResponseWriter, request *http.Request) {
	projectName := request.URL.Query()["project"]
	if len(projectName) == 0 {
		errorWriter(writer, handler.log, "error: no projects in request.", http.StatusBadRequest)
		return
	}
	err := handler.resourceRep.DeleteProject(projectName[0])

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	var projectResponse = models.ResponseStruct{
		Links: models.ListOfReferences{
			Issues:   models.Link{Href: "/api/v1/resource/issues"},
			Projects: models.Link{Href: "/api/v1/resource/project"},
			Graphs:   models.Link{Href: "/api/v1/graph"},
			Self:     models.Link{Href: fmt.Sprintf("/api/v1/resource/project?project=%s", projectName[0])},
		},
		Message: "",
		Name:    "",
		Status:  true,
	}

	response, err := json.MarshalIndent(projectResponse, "", "\t")

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = writer.Write(response)

	if err != nil {
		errorWriter(writer, handler.log, err.Error(), http.StatusBadRequest)
		return
	}
}
