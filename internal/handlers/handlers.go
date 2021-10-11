package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piotrzalecki/budget/internal/config"
	"github.com/piotrzalecki/budget/internal/driver"
	"github.com/piotrzalecki/budget/internal/models"
	"github.com/piotrzalecki/budget/internal/render"
	"github.com/piotrzalecki/budget/internal/repository"
	"github.com/piotrzalecki/budget/internal/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}

}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (m *Repository) LoginPost(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

// TRANSACTION CATEGORY HANDLERS
func (m *Repository) TransactionCategory(w http.ResponseWriter, r *http.Request) {

	tcats, err := m.DB.AllTransactionCategories()
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve transaction categories from database")
		http.Redirect(w, r, "/tcats", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["tcats"] = tcats

	render.Template(w, r, "tcat.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) PostTransactionCategory(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (m *Repository) UpdateTransactionCategory(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (m *Repository) DeleteTransactionCategory(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

}

func (m *Repository) TransactionCategoryNew(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "tcat_new.page.tmpl", &models.TemplateData{})

}

func (m *Repository) PostTransactionCategoryNew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	desc := r.Form.Get("desc")
	fmt.Println(name)
	fmt.Println(desc)

	if name == "" || desc == "" {
		log.Println("can't post for new transaction category, one of required fields is empty")
		http.Redirect(w, r, "/tcats/new", http.StatusSeeOther)
		return
	}

	var newcat models.TransactionCategory

	newcat.Name = name
	newcat.Description = desc

	err := m.DB.CreateTransactionCategory(newcat)

	if err != nil {
		log.Println(fmt.Printf("creating category name: %s, description: %s FAILED", name, desc))
		http.Redirect(w, r, "/tcats/new", http.StatusSeeOther)

	}

	http.Redirect(w, r, "/tcats", http.StatusSeeOther)

}
