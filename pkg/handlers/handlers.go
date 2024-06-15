package handlers

import (
	"net/http"

	"github.com/motoshi-suzuki-hp/go-cource/pkg/config"
	"github.com/motoshi-suzuki-hp/go-cource/pkg/models"
	"github.com/motoshi-suzuki-hp/go-cource/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Reposiory

// Repository is the repository type
type Reposiory struct {
	App *config.AppConfig
}

// NewRepo coreates a new repository
func NewRepo(a *config.AppConfig) *Reposiory {
	return &Reposiory{
		App: a,
	}
}

// NewHAndlers sets the repository for the handlers
func NewHandlers(r *Reposiory) {
	Repo = r
}


// Home is the home page handler
func (m *Reposiory) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Reposiory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"


	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}


