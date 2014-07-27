package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
)

func lfgsReg() {
	lfgsTemplates()
	lfgsRoutes()
	lfgsStmts()

	net.RegType(&lfg{})
}

func lfgsTemplates() {
	net.TemplateFiles(
		"tmpl/_lfgs_list.html.tmpl",
		"tmpl/_lfgs_form.html.tmpl",
		"tmpl/lfgs_page.html.tmpl",
	)
	//log.Message("lfgsTemplates run")
}

func lfgsRoutes() {
	sr := net.Subrouter("/lfgs")
	sr.Get("/list", home)
	sr.Get("/form/{Id}", lfgsForm)
	sr.Post("/save", lfgsSave)
	sr.Post("/delete", lfgsDelete)
}

/*
func lfgsPage(w http.ResponseWriter, r *http.Request) {
	da := new(lfgsAdapter).list()
	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "lfgsPage", da)
}


func lfgsList(w http.ResponseWriter, r *http.Request) {
	da := new(lfgsAdapter).list()
	ra := new(rankingPagesAdapter).list()
	//log.Message(len(ra.RankingPages))
	net.ExeTmpl(w, "lfgsList", da, ra)
}
*/

func lfgsForm(w http.ResponseWriter, r *http.Request) {

	id := net.QueryUrl("Id", r)

	da := new(lfgsAdapter)
	ca := new(classesAdapter).list()
	qa := new(queuePrefsAdapter).list()
	ra := new(rankingPagesAdapter).list()

	if id == 0 {
		_ = da.newLfg()
	} else {
		da.show(id)
	}
	net.ExeTmpl(w, "lfgsForm", da, ca, qa, ra)
}

func lfgsSave(w http.ResponseWriter, r *http.Request) {

	da := new(lfg)

	net.DecodeForm(da, r)

	new(lfgsAdapter).save(da)

	http.Redirect(w, r, "/lfgs/list", http.StatusFound)

}

func lfgsDelete(w http.ResponseWriter, r *http.Request) {

	var da struct{ Id int }

	net.DecodeForm(&da, r)

	ss := new(lfgsAdapter)
	ss.delete(da.Id)
}

func GetLfgsList() *lfgsAdapter {
	return new(lfgsAdapter).list()
}