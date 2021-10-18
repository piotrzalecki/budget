package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
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

	if name == "" || desc == "" {
		log.Println("can't post for new transaction category, one of required fields is empty")
		http.Redirect(w, r, "/dashboard/tcats/new", http.StatusSeeOther)
		return
	}

	var newcat models.TransactionCategory

	newcat.Name = name
	newcat.Description = desc

	_, err := m.DB.CreateTransactionCategory(newcat)

	if err != nil {
		log.Println(fmt.Printf("creating category name: %s, description: %s FAILED", name, desc))
		http.Redirect(w, r, "/dashboard/tcats/new", http.StatusSeeOther)

	}

	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)

}

func (m *Repository) TransactionCategoryDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	catId := r.Form.Get("id")
	catIdint, err := strconv.Atoi(catId)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve categroy id from uri")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}
	err = m.DB.DeleteTransactionCategory(catIdint)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't delete category")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
}

func (m *Repository) TransactionCategoryUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}

	category, err := m.DB.GetTransactionCategoryById(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't get category from database")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["category"] = category

	render.Template(w, r, "tcat_update.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionCategoryUpdatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	name := r.Form.Get("name")
	desc := r.Form.Get("desc")
	if name == "" || desc == "" || id == "" {
		log.Println("can't post for new transaction category, one of required fields is empty")
		http.Redirect(w, r, "/dashboard/tcats/new", http.StatusSeeOther)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}

	var newcat models.TransactionCategory

	newcat.Id = idInt
	newcat.Name = name
	newcat.Description = desc

	err = m.DB.UpdateTransactionCategory(newcat)

	if err != nil {
		log.Println(fmt.Printf("updating category name: %s, description: %s FAILED", name, desc))
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)

}

func (m *Repository) TransactionsData(w http.ResponseWriter, r *http.Request) {
	// var tDataFormAll []models.TransactionData
	// tDataAll, err := m.DB.AllTransactionsData()
	// if err != nil {
	// 	log.Println("Error retriving all transactions data", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }

	// for _, td := range tDataAll {
	// 	tt, err := m.DB.GetTransactionTypeById(td.Type)
	// 	if err != nil {
	// 		log.Println("Error retriving transaction type by id", err)
	// 		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	ttf := models.TransactionsTypesForm{
	// 		Id:          tt.Id,
	// 		Name:        tt.Name,
	// 		Description: tt.Description,
	// 		Recurence:   models.RecurentTransactions{},
	// 		CreatedAt:   tt.CreatedAt,
	// 		UpdatedAt:   tt.UpdatedAt,
	// 	}

	// 	tc, err := m.DB.GetTransactionCategoryById(td.Category)
	// 	if err != nil {
	// 		log.Println("Error retriving transaction category by id", err)
	// 		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	tDataForm := models.TransactionsDataForm{
	// 		Id:            td.Id,
	// 		Name:          td.Name,
	// 		Description:   td.Description,
	// 		ExpectedQuote: td.ExpectedQuote,
	// 		Type:          ttf,
	// 		Category:      tc,
	// 		CreatedAt:     td.CreatedAt,
	// 		UpdatedAt:     td.UpdatedAt,
	// 	}

	// 	tDataFormAll = append(tDataFormAll, tDataForm)
	// }

	// data := make(map[string]interface{})
	// data["tdata"] = tDataFormAll

	// render.Template(w, r, "tdata.page.tmpl", &models.TemplateData{
	// 	Data: data,
	// })

}

func (m *Repository) TransactionsDataDetails(w http.ResponseWriter, r *http.Request) {
	// tid, err := strconv.Atoi(chi.URLParam(r, "id"))
	// if err != nil {
	// 	log.Println("Error retriving id from url", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }

	// td, err := m.DB.GetTransactionDataById(tid)
	// if err != nil {
	// 	log.Println("Error retriving all transactions data", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }

	// tt, err := m.DB.GetTransactionTypeById(td.Type)
	// if err != nil {
	// 	log.Println("Error retriving transaction type by id", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }

	// rt, err := m.DB.GetRecurentTransactionById(tt.Recurence)
	// if err != nil {
	// 	log.Println("Error retriving recurent transactionby id", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }

	// ttf := models.TransactionsTypesForm{
	// 	Id:          tt.Id,
	// 	Name:        tt.Name,
	// 	Description: tt.Description,
	// 	Recurence:   rt,
	// 	CreatedAt:   tt.CreatedAt,
	// 	UpdatedAt:   tt.UpdatedAt,
	// }

	// tc, err := m.DB.GetTransactionCategoryById(td.Category)
	// if err != nil {
	// 	log.Println("Error retriving transaction category by id", err)
	// 	http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
	// 	return
	// }
	// tDataForm := models.TransactionsDataForm{
	// 	Id:            td.Id,
	// 	Name:          td.Name,
	// 	Description:   td.Description,
	// 	ExpectedQuote: td.ExpectedQuote,
	// 	Type:          ttf,
	// 	Category:      tc,
	// 	CreatedAt:     td.CreatedAt,
	// 	UpdatedAt:     td.UpdatedAt,
	// }

	// data := make(map[string]interface{})
	// data["tdata"] = tDataForm

	// render.Template(w, r, "tdata_details.page.tmpl", &models.TemplateData{
	// 	Data: data,
	// })

}

// Transaction Recurence
func (m *Repository) TransactionRecurence(w http.ResponseWriter, r *http.Request) {

	tr, err := m.DB.AllRecurentTransactions()
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve transaction categories from database")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["trecurence"] = tr
	render.Template(w, r, "trecurence.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionRecurenceNew(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "trecurence_new.page.tmpl", &models.TemplateData{})

}

func (m *Repository) PostTransactionRecurenceNew(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	desc := r.Form.Get("description")
	addTime := r.Form.Get("add_time")

	if name == "" || desc == "" || addTime == "" {
		log.Println("can't post for new recuring transaction, one of required fields is empty")
		http.Redirect(w, r, "/dashboard/trecurence/new", http.StatusSeeOther)
		return
	}

	var recurt models.TransactionRecurence

	recurt.Name = name
	recurt.Description = desc
	recurt.AddTime = addTime

	_, err := m.DB.CreateRecurentTransaction(recurt)

	if err != nil {
		log.Println(fmt.Printf("creating recurent transaction name: %s FAILED\n %v", name, err))
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)

	}

	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)

}

func (m *Repository) TransactionRecurenceDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rId := r.Form.Get("id")
	recId, err := strconv.Atoi(rId)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't recurent transaction id from uri")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}
	err = m.DB.DeleteRecurentTransaction(recId)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't delete recurent transaction")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
}

func (m *Repository) TransactionRecurenceUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve recurent transaction id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	recurentt, err := m.DB.GetRecurentTransactionById(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't get category from database")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["trecurence"] = recurentt

	render.Template(w, r, "trecurence_update.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionRecurenceUpdatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	name := r.Form.Get("name")
	desc := r.Form.Get("description")
	addTime := r.Form.Get("add_time")

	if name == "" || desc == "" || addTime == "" {
		log.Println("can't post for new recuring transaction, one of required fields is empty")
		http.Redirect(w, r, "/dashboard/trecurence/new", http.StatusSeeOther)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	var recurt models.TransactionRecurence

	recurt.Id = idInt
	recurt.Name = name
	recurt.Description = desc
	recurt.AddTime = addTime

	err = m.DB.UpdateRecurentTransaction(recurt)

	if err != nil {
		log.Println(fmt.Printf("updating recurent transaction name: %s FAILED", name))
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)

}

func (m *Repository) TransactionTypes(w http.ResponseWriter, r *http.Request) {

	tt, err := m.DB.AllTransactionTypes()
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve transaction types from database")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["ttypes"] = tt
	render.Template(w, r, "ttypes.page.tmpl", &models.TemplateData{
		Data: data,
	})

}
