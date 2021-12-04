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
	handleError(w, r, err, "can't retrieve all transaction categories from database", "/dashboard/tcats")

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
	handleError(w, r, err, fmt.Sprintf("creating transaction categoryof name %s failed", newcat.Name), "/dashboard/tcats/new")
	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)

}

func (m *Repository) TransactionCategoryDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	catId := r.Form.Get("id")
	catIdint, err := strconv.Atoi(catId)
	handleError(w, r, err, fmt.Sprintf("can't retrieve category of id %s from uri", catId), "/dashboard/tcats")

	err = m.DB.DeleteTransactionCategory(catIdint)
	handleError(w, r, err, fmt.Sprintf("can't delete category of id %d", catIdint), "/dashboard/tcats")

	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)
}

func (m *Repository) TransactionCategoryUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to intger", idPara), "/dashboard/tcats")

	category, err := m.DB.GetTransactionCategoryById(id)
	handleError(w, r, err, fmt.Sprintf("can't retriece category by id %d", id), "/dashboard/tcats")

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
	handleError(w, r, err, fmt.Sprintf("catn't convert %s to integer", id), "/dashboard/tcats")

	var newcat models.TransactionCategory

	newcat.Id = idInt
	newcat.Name = name
	newcat.Description = desc

	err = m.DB.UpdateTransactionCategory(newcat)
	handleError(w, r, err, fmt.Sprintf("updating category name: %s, description: %s failed", name, desc), "/dashboard/tcats")
	http.Redirect(w, r, "/dashboard/tcats", http.StatusSeeOther)

}

func (m *Repository) TransactionsData(w http.ResponseWriter, r *http.Request) {
	tDataAll, err := m.DB.AllTransactionsData()
	handleError(w, r, err, "error retreiving all transactions data", "/dashboard")
	data := make(map[string]interface{})
	data["tdata"] = tDataAll

	render.Template(w, r, "tdata.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsDataDetails(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tLid, err := strconv.Atoi(id)
	handleError(w, r, err, fmt.Sprintf("can't convrt %s to integer", id), "/dashboard/tdata")

	td, err := m.DB.GetTransactionDataById(tLid)
	handleError(w, r, err, fmt.Sprintf("error retriving tranasction data with id %d", tLid), "/dashboard/tdata")

	data := make(map[string]interface{})
	data["tdata"] = td
	render.Template(w, r, "tdata_details.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsDataNew(w http.ResponseWriter, r *http.Request) {

	tt, err := m.DB.AllTransactionTypes()
	handleError(w, r, err, "error retriving all transactions types", "/dashboard/tdata")

	tc, err := m.DB.AllTransactionCategories()
	handleError(w, r, err, "error retriving all transactions categories", "/dashboard/tdata")

	tr, err := m.DB.AllRecurentTransactions()
	handleError(w, r, err, "error retriving all transactions recurences", "/dashboard/tdata")

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
	handleError(w, r, err, fmt.Sprintf("can't convert %s to float32", quote), "/dashboard/tdata")

	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, date)
	handleError(w, r, err, fmt.Sprintf("can't parse date %s for layout %s", date, layout), "/dashboard/tdata")

	parsedRepeatUntill, err := time.Parse(layout, repeatUntil)
	handleError(w, r, err, fmt.Sprintf("can't parse date %s for layout %s", repeatUntil, layout), "/dashboard/tdata")

	categoryId, err := strconv.Atoi(category)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", category), "/dashboard/tdata")

	typeId, err := strconv.Atoi(ttype)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", ttype), "/dashboard/tdata")

	recurId, err := strconv.Atoi(rcur)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", rcur), "/dashboard/tdata")

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
	handleError(w, r, err, fmt.Sprintf("error reating transaction data %v", transactionData), "/dashboard/tdata")
	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)

}

func (m *Repository) TransactionDataDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tdId := r.Form.Get("id")
	tdIdParsed, err := strconv.Atoi(tdId)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", tdId), "/dashboard/tdata")

	err = m.DB.DeleteTransactionData(tdIdParsed)
	handleError(w, r, err, fmt.Sprintf("can't delete transaction data with id %d", tdIdParsed), "/dashboard/tdata")
	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
}

func (m *Repository) TransactionDataUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", idPara), "/dashboard/tdata")

	tdata, err := m.DB.GetTransactionDataById(id)
	handleError(w, r, err, fmt.Sprintf("can't get transaction data by id %d", id), "/dashboard/tdata")

	tt, err := m.DB.AllTransactionTypes()
	handleError(w, r, err, "can't get all transactions types", "/dashboard/tdata")

	tc, err := m.DB.AllTransactionCategories()
	handleError(w, r, err, "can't get all transactions categories", "/dashboard/tdata")

	tr, err := m.DB.AllRecurentTransactions()
	handleError(w, r, err, "can't get all recurnet transactions", "/dashboard/tdata")

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
	handleError(w, r, err, fmt.Sprintf("can't convert %s to float32", quote), "/dashboard/tdata")

	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	handleError(w, r, err, fmt.Sprintf("can't parse %s for layout %s", date, layout), "/dashboard/tdata")

	parsedRepeatUntill, err := time.Parse(layout, repeatUntil)
	handleError(w, r, err, fmt.Sprintf("can't parse %s for layout %s", repeatUntil, layout), "/dashboard/tdata")

	categoryId, err := strconv.Atoi(category)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", category), "/dashboard/tdata")

	typeId, err := strconv.Atoi(ttype)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", ttype), "/dashboard/tdata")

	recurId, err := strconv.Atoi(rcur)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", rcur), "/dashboard/tdata")

	tranId, err := strconv.Atoi(id)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", id), "/dashboard/tdata")

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
	handleError(w, r, err, fmt.Sprintf("updating transaction data failde for %v", transactionData), "/dashboard/tdata")

	http.Redirect(w, r, "/dashboard/tdata", http.StatusSeeOther)
}

// Transaction Recurence
func (m *Repository) TransactionRecurence(w http.ResponseWriter, r *http.Request) {

	tr, err := m.DB.AllRecurentTransactions()
	handleError(w, r, err, "can't retrieve all recurent transactions", "/dashboard/")

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
	handleError(w, r, err, fmt.Sprintf("can't create recurent transaction %v", recurt), "/dashboard/trecurence")

	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)

}

func (m *Repository) TransactionRecurenceDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rId := r.Form.Get("id")
	recId, err := strconv.Atoi(rId)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", rId), "/dashboard/trecurence")

	err = m.DB.DeleteRecurentTransaction(recId)
	handleError(w, r, err, fmt.Sprintf("can't delete recurent transaction with id %d", recId), "/dashboard/trecurence")

	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)
}

func (m *Repository) TransactionRecurenceUpdateGet(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", idPara), "/dashboard/trecurence")

	recurentt, err := m.DB.GetRecurentTransactionById(id)
	handleError(w, r, err, fmt.Sprintf("can't retrieve recurent transaction with id %d", id), "/dashboard/trecurence")

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
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", id), "/dashboard/trecurence")

	var recurt models.TransactionRecurence

	recurt.Id = idInt
	recurt.Name = name
	recurt.Description = desc
	recurt.AddTime = addTime

	err = m.DB.UpdateRecurentTransaction(recurt)
	handleError(w, r, err, fmt.Sprintf("can't update recurent transaction %v", recurt), "/dashboard/trecurence")

	http.Redirect(w, r, "/dashboard/trecurence", http.StatusSeeOther)

}

func (m *Repository) TransactionTypes(w http.ResponseWriter, r *http.Request) {

	tt, err := m.DB.AllTransactionTypes()
	handleError(w, r, err, "can't retreive all transaction types", "/dashboard/trecurence")

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
	handleError(w, r, err, "can't retireve all active recurent tranasactions", "/dashboard")

	tDataAll := models.TransactionsData{}

	for _, rt := range recTrans {
		startDate := rt.TransactionDate
		addTime := rt.TransactionRecurence.AddTime
		addTimeArr := strings.Split(addTime, "-")

		yearAdd, err := strconv.Atoi(addTimeArr[0])
		handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", addTimeArr[0]), "/dashboard")

		yearMonth, err := strconv.Atoi(addTimeArr[1])
		handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", addTimeArr[1]), "/dashboard")

		yearDays, err := strconv.Atoi(addTimeArr[2])
		handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", addTimeArr[2]), "/dashboard")

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
	handleError(w, r, err, fmt.Sprintf("can't get single transactions for dates %v and %v", firstOfMonth, lastOfMonth), "/dashboard")

	for _, st := range singleTransactions {
		tDataAll = append(tDataAll, st)
	}

	// //remove logged tranasctions
	tranLogs, err := m.DB.AllTransactionsLogsForDates(firstOfMonth, lastOfMonth)
	handleError(w, r, err, fmt.Sprintf("can't get all transactions logs for dates %v and %v", firstOfMonth, lastOfMonth), "/dashboard")

	// calculate expected occurnces for each transaction in transaction log
	expectedTO := make(map[int]int) // [transaction id]number of cocrences

	for _, tl := range tranLogs {
		td, err := m.DB.GetTransactionDataById(tl.TransactionData.Id)
		handleError(w, r, err, fmt.Sprintf("can't get tranasctions data by id %d", tl.TransactionData.Id), "/dashboard")

		if td.TransactionType.Name == "RT" {
			transactionStartDate := td.TransactionDate
			frequency := td.TransactionRecurence.AddTime
			frequencyArr := strings.Split(frequency, "-")

			yearAdd, err := strconv.Atoi(frequencyArr[0])
			handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", frequencyArr[0]), "/dashboard")

			yearMonth, err := strconv.Atoi(frequencyArr[1])
			handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", frequencyArr[1]), "/dashboard")

			yearDays, err := strconv.Atoi(frequencyArr[2])
			handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", frequencyArr[2]), "/dashboard")

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
	handleError(w, r, err, "can't retrieve latest balance", "/dashboard")

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
	handleError(w, r, err, "can't retrieve all transactions data", "/dashboard")

	data := make(map[string]interface{})
	data["tlog"] = tLogAll

	render.Template(w, r, "tlog.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLogDetails(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tLid, err := strconv.Atoi(id)

	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", id), "/dashboard/tlog")

	tl, err := m.DB.GetTransactionLogById(tLid)
	handleError(w, r, err, fmt.Sprintf("can't get transaction log by id %d", tLid), "/dashboard/tlog")

	data := make(map[string]interface{})
	data["tlog"] = tl
	render.Template(w, r, "tlog_details.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) TransactionsLogNew(w http.ResponseWriter, r *http.Request) {
	idPara := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idPara)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", idPara), "/dashboard/tlog")

	tData, err := m.DB.GetTransactionDataById(id)
	handleError(w, r, err, fmt.Sprintf("can't get transaction data by id %d", id), "/dashboard/tlog")

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
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", tid), "/dashboard/tlog")

	userId, err := strconv.Atoi(uId)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to integer", uId), "/dashboard/tlog")

	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, tdate)
	handleError(w, r, err, fmt.Sprintf("can't parse date %s for layout %s", tdate, layout), "/dashboard/tdata")

	quoteFloat, err := strconv.ParseFloat(tquote, 32)
	handleError(w, r, err, fmt.Sprintf("can't convert %s to float32", tquote), "/dashboard/tlog")

	newlog := models.TransactionLog{
		TransactionData:  models.TransactionData{Id: transactionId},
		TransactionQuote: float32(quoteFloat),
		TransactionDate:  parsedDate,
		CreatedBy:        models.User{Id: userId},
		UpdateBy:         models.User{Id: userId},
	}

	newTLId, err := m.DB.CreateTransactionLog(newlog)
	handleError(w, r, err, fmt.Sprintf("can't create transaction log %v", newlog), "/dashboard/tlog")

	balance, err := m.DB.GetLatestBalanceQuote()
	handleError(w, r, err, "can't retreive latest balance quote", "/dashboard/tlog")

	newBalanceQuote := balance + float32(quoteFloat)

	nab := models.AccountBalance{
		Balance:            newBalanceQuote,
		BalanceTransaction: models.TransactionLog{Id: newTLId},
	}

	_, err = m.DB.CreateAccountBalance(nab)
	handleError(w, r, err, fmt.Sprintf("can't create account balance %v", nab), "/dashboard/tlog")

	http.Redirect(w, r, "/flowboard", http.StatusSeeOther)

}
