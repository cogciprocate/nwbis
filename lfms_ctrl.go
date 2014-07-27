package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
)

func lfmsReg() {
	lfmsTemplates()
	lfmsRoutes()
	lfmsStmts()

	net.RegType(&lfm{})
}

func lfmsTemplates() {
	net.TemplateFiles(
		"tmpl/_lfms_list.html.tmpl",
		"tmpl/_lfms_form.html.tmpl",
		"tmpl/lfms_page.html.tmpl",
	)
	//log.Message("lfmsTemplates run")
}

func lfmsRoutes() {
	sr := net.Subrouter("/lfms")
	sr.Get("/list", home)
	sr.Get("/form/{Id}", lfmsForm)
	sr.Post("/save", lfmsSave)
	sr.Post("/delete", lfmsDelete)
}

func lfmsPage(w http.ResponseWriter, r *http.Request) {
	da := new(lfmsAdapter).list()
	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "lfmsPage", da)
}

func lfmsList(w http.ResponseWriter, r *http.Request) {
	da := new(lfmsAdapter).list()
	//log.Message(len(dataModel.Lfms))
	net.ExeTmpl(w, "lfmsList", da)
}


func lfmsForm(w http.ResponseWriter, r *http.Request) {

	id := net.QueryUrl("Id", r)

	da := new(lfmsAdapter)

	if id == 0 {
		_ = da.newLfm()
	} else {
		da.show(id)
	}
	net.ExeTmpl(w, "lfmsForm", da)
}

func lfmsSave(w http.ResponseWriter, r *http.Request) {

	da := new(lfm)

	net.DecodeForm(da, r)

	new(lfmsAdapter).save(da)

	http.Redirect(w, r, "/lfms/list", http.StatusFound)

}

func lfmsDelete(w http.ResponseWriter, r *http.Request) {

	var da struct{ Id int }

	net.DecodeForm(&da, r)

	ss := new(lfmsAdapter)
	ss.delete(da.Id)
}

func GetLfmsList() *lfmsAdapter {
	return new(lfmsAdapter).list()
}