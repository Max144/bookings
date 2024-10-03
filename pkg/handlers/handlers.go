package handlers

import (
	"github.com/Max144/bookings/pkg/config"
	"github.com/Max144/bookings/pkg/models"
	"github.com/Max144/bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = m.App.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationMake(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) StandardRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "standard-room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) DeluxeRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "deluxe-room.page.tmpl", &models.TemplateData{})
}

func (m *Repository) FamilyRoom(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "family-room.page.tmpl", &models.TemplateData{})
}
