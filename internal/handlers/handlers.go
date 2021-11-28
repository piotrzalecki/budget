package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
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
	tDataAll, err := m.DB.AllTransactionsData()
	if err != nil {
		log.Println("Error retriving all transactions data", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["tdata"] = tDataAll

	render.Template(w, r, "tdata.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsDataDetails(w http.ResponseWriter, r *http.Request) {
	tLid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error retriving id from url", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	td, err := m.DB.GetTransactionDataById(tLid)
	if err != nil {
		log.Println("Error retriving all transactions data", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["tdata"] = td
	render.Template(w, r, "tdata_details.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsDataNew(w http.ResponseWriter, r *http.Request) {

	tt, err := m.DB.AllTransactionTypes()
	if err != nil {
		log.Println("Error retriving all transactions types", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tc, err := m.DB.AllTransactionCategories()
	if err != nil {
		log.Println("Error retriving all transactions categories", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tr, err := m.DB.AllRecurentTransactions()
	if err != nil {
		log.Println("Error retriving all transactions recurences", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["ttypes"] = tt
	data["tcats"] = tc
	data["trec"] = tr

	render.Template(w, r, "tdata_new.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsDataNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("td_name")
	desc := r.Form.Get("td_desc")
	quote := r.Form.Get("td_quote")
	date := r.Form.Get("td_date")
	repeatUntil := r.Form.Get("td_repeat")
	category := r.Form.Get("tc_id")
	ttype := r.Form.Get("tt_id")
	rcur := r.Form.Get("tr_id")

	//TODO: Implement better form validation
	if name == "" || desc == "" || quote == "" || date == "" {
		log.Println("can't post for new transaction data, one of required fields is empty")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	quoteFloat, err := strconv.ParseFloat(quote, 32)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't parse date ")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	parsedRepeatUntill, err := time.Parse(layout, repeatUntil)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't parse repeat until ")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing category")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	typeId, err := strconv.Atoi(ttype)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing type")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	recurId, err := strconv.Atoi(rcur)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing recurence")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	transactionData := models.TransactionData{
		Name:                 name,
		Description:          desc,
		RepeatUntil:          parsedRepeatUntill,
		TransactionDate:      parsedDate,
		TransactionQuote:     float32(quoteFloat),
		TransactionCategory:  models.TransactionCategory{Id: categoryId},
		TransactionType:      models.TransactionType{Id: typeId},
		TransactionRecurence: models.TransactionRecurence{Id: recurId},
	}

	_, err = m.DB.CreateTransactionData(transactionData)
	if err != nil {
		log.Println(err)
		log.Println(fmt.Printf("creating new transaction data failed name: %s FAILED", name))
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)

}

func (m *Repository) TransactionDataDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tdId := r.Form.Get("id")
	tdIdParsed, err := strconv.Atoi(tdId)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't get transaction data id from uri")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}
	err = m.DB.DeleteTransactionData(tdIdParsed)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't delete rtransaction data")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
}

func (m *Repository) TransactionDataUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve transaction data id")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	tdata, err := m.DB.GetTransactionDataById(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't get transaction data database")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	tt, err := m.DB.AllTransactionTypes()
	if err != nil {
		log.Println("Error retriving all transactions types", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tc, err := m.DB.AllTransactionCategories()
	if err != nil {
		log.Println("Error retriving all transactions categories", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tr, err := m.DB.AllRecurentTransactions()
	if err != nil {
		log.Println("Error retriving all transactions recurences", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["tdata"] = tdata
	data["ttypes"] = tt
	data["tcats"] = tc
	data["trec"] = tr

	render.Template(w, r, "tdata_update.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionDataUpdatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("td_id")
	name := r.Form.Get("td_name")
	desc := r.Form.Get("td_desc")
	quote := r.Form.Get("td_quote")
	date := r.Form.Get("td_date")
	repeatUntil := r.Form.Get("td_repeat")
	category := r.Form.Get("tc_id")
	ttype := r.Form.Get("tt_id")
	rcur := r.Form.Get("tr_id")

	//TODO: Implement better form validation
	if name == "" || desc == "" || quote == "" || date == "" {
		log.Println("can't post for new transaction data, one of required fields is empty")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	quoteFloat, err := strconv.ParseFloat(quote, 32)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't parse date ")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}
	parsedRepeatUntill, err := time.Parse(layout, repeatUntil)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't parse repeat until ")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing category")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	typeId, err := strconv.Atoi(ttype)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing type")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	recurId, err := strconv.Atoi(rcur)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing recurence")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	tranId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("error parsing id")
		http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
		return
	}

	transactionData := models.TransactionData{
		Id:                   tranId,
		Name:                 name,
		Description:          desc,
		RepeatUntil:          parsedRepeatUntill,
		TransactionDate:      parsedDate,
		TransactionQuote:     float32(quoteFloat),
		TransactionCategory:  models.TransactionCategory{Id: categoryId},
		TransactionType:      models.TransactionType{Id: typeId},
		TransactionRecurence: models.TransactionRecurence{Id: recurId},
	}

	err = m.DB.UpdateTransactionsData(transactionData)
	if err != nil {
		log.Println(err)
		log.Println(fmt.Printf("updating transaction data failed name: %s FAILED", name))
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)

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

func (m *Repository) FlowBoard(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	recTrans, err := m.DB.GetAllActiveRecurentTransactions(lastOfMonth)
	if err != nil {
		log.Println("Error retriving all transactions data", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tDataAll := models.TransactionsData{}

	for _, rt := range recTrans {
		startDate := rt.TransactionDate
		addTime := rt.TransactionRecurence.AddTime
		yearAdd, err := strconv.Atoi(strings.Split(addTime, "-")[0])
		if err != nil {
			log.Println("Error rconverting str to int", err)
			http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
			return
		}
		yearMonth, err := strconv.Atoi(strings.Split(addTime, "-")[1])
		if err != nil {
			log.Println("Error rconverting str to int", err)
			http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
			return
		}
		yearDays, err := strconv.Atoi(strings.Split(addTime, "-")[2])
		if err != nil {
			log.Println("Error rconverting str to int", err)
			http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
			return
		}

		for {
			if startDate.After(firstOfMonth) && startDate.Before(lastOfMonth) {
				rt.TransactionDate = startDate
				tDataAll = append(tDataAll, rt)
			}
			if startDate.After(lastOfMonth) {
				break
			}
			startDate = startDate.AddDate(yearAdd, yearMonth, yearDays)
		}
	}

	singleTransactions, err := m.DB.GetSingleTransactionsForDates(firstOfMonth, lastOfMonth)
	if err != nil {
		log.Println("Error retreiving single transactions", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}
	for _, st := range singleTransactions {
		tDataAll = append(tDataAll, st)
	}

	// //remove logged tranasctions
	tranLogs, err := m.DB.AllTransactionsLogsForDates(firstOfMonth, lastOfMonth)
	if err != nil {
		log.Println("Error retriving transaction logs", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	// calculate expected occurnces for each transaction in transaction log
	expectedTO := make(map[int]int) // [transaction id]number of cocrences

	for _, tl := range tranLogs {
		td, err := m.DB.GetTransactionDataById(tl.TransactionData.Id)
		if err != nil {
			log.Println("Error single transaction", err)
			http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
			return
		}
		if td.TransactionType.Name == "RT" {
			transactionStartDate := td.TransactionDate
			frequency := td.TransactionRecurence.AddTime
			yearAdd, err := strconv.Atoi(strings.Split(frequency, "-")[0])
			if err != nil {
				log.Println("Error rconverting str to int", err)
				http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
				return
			}
			yearMonth, err := strconv.Atoi(strings.Split(frequency, "-")[1])
			if err != nil {
				log.Println("Error rconverting str to int", err)
				http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
				return
			}
			yearDays, err := strconv.Atoi(strings.Split(frequency, "-")[2])
			if err != nil {
				log.Println("Error rconverting str to int", err)
				http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
				return
			}
			counter := 0
			for {
				if transactionStartDate.After(firstOfMonth.AddDate(0, 0, -1)) && transactionStartDate.Before(lastOfMonth) {
					counter += 1
				}
				transactionStartDate = transactionStartDate.AddDate(yearAdd, yearMonth, yearDays)
				if transactionStartDate.After(lastOfMonth) {
					break
				}
			}
			expectedTO[td.Id] = counter
		}
	}

	// cleaning logged transactions
	for _, stl := range tranLogs {
		fmt.Println(expectedTO)
		remove := false
		for i, sd := range tDataAll {
			if stl.TransactionData.Id == sd.Id {
				// remove logged single transactions
				fmt.Println("----> matched", sd.Id, stl.Id)
				if sd.TransactionType.Name == "ST" {
					remove = true
				}
				// remove logged recurent trnsactions
				if sd.TransactionType.Name == "RT" {
					if expectedTO[sd.Id] > 0 {
						fmt.Println("Decreasing one for", sd.Name)
						expectedTO[sd.Id] -= 1
						remove = true
					}
				}
				if remove {
					copy(tDataAll[i:], tDataAll[i+1:])
					tDataAll[len(tDataAll)-1] = models.TransactionData{}
					tDataAll = tDataAll[:len(tDataAll)-1]
				}
				break
			}

		}
	}

	balance, err := m.DB.GetLatestBalanceQuote()
	if err != nil {
		log.Println("Error retriving latest balance", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	for _, tr := range tDataAll {
		balance += tr.TransactionQuote
	}
	sort.Sort(tDataAll)
	data := make(map[string]interface{})
	data["tdata"] = tDataAll
	data["balance"] = balance

	render.Template(w, r, "flowboard.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLog(w http.ResponseWriter, r *http.Request) {
	// var tDataFormAll []models.TransactionData
	tLogAll, err := m.DB.AllTransactionsLogs()
	if err != nil {
		log.Println("Error retriving all transactions data", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["tlog"] = tLogAll

	render.Template(w, r, "tlog.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLogDetails(w http.ResponseWriter, r *http.Request) {
	tLid, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Error retriving id from url", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	tl, err := m.DB.GetTransactionLogById(tLid)
	if err != nil {
		log.Println("Error retriving transaacion log by id", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	data := make(map[string]interface{})
	data["tlog"] = tl
	render.Template(w, r, "tlog_details.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLogNew(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve recurent transaction id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	tData, err := m.DB.GetTransactionDataById(id)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve recurent transaction id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}
	data := make(map[string]interface{})
	data["tdata"] = tData

	render.Template(w, r, "tlog_new.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLogNewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tid := r.Form.Get("tid")
	tquote := r.Form.Get("tquote")
	tdate := r.Form.Get("tdate")
	uId := r.Form.Get("user_id")

	if tid == "" || tquote == "" || tdate == "" || uId == "" {
		log.Println("can't post for new transaction log, one of required fields is empty")
		fmt.Println(tid)
		fmt.Println(tquote)
		fmt.Println(tdate)
		fmt.Println(uId)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	transactionId, err := strconv.Atoi(tid)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve categroy id from uri")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}

	userId, err := strconv.Atoi(uId)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve categroy id from uri")
		http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
		return
	}

	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, tdate)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't parse date ")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	quoteFloat, err := strconv.ParseFloat(tquote, 32)
	if err != nil {
		fmt.Println(err)
		log.Fatal("can't retrieve category id")
		http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
		return
	}

	newlog := models.TransactionLog{
		TransactionData:  models.TransactionData{Id: transactionId},
		TransactionQuote: float32(quoteFloat),
		TransactionDate:  parsedDate,
		CreatedBy:        models.User{Id: userId},
		UpdateBy:         models.User{Id: userId},
	}

	NewTLId, err := m.DB.CreateTransactionLog(newlog)

	if err != nil {
		log.Println(fmt.Printf("creating log for transaction id: %s, quote: %s FAILED", tid, tquote))
		http.Redirect(w, r, "/dashboard/tcats/new", http.StatusSeeOther)

	}

	balance, err := m.DB.GetLatestBalanceQuote()
	if err != nil {
		log.Println("Error retriving latest balance", err)
		http.Redirect(w, r, "/dashboard", http.StatusInternalServerError)
		return
	}

	newBalanceQuote := balance + float32(quoteFloat)

	nab := models.AccountBalance{
		Balance:            newBalanceQuote,
		BalanceTransaction: models.TransactionLog{Id: NewTLId},
	}

	_, err = m.DB.CreateAccountBalance(nab)
	if err != nil {
		log.Println("Error updateing balance", err)
		http.Redirect(w, r, "/flowboard", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/flowboard", http.StatusSeeOther)

}
