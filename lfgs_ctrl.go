package main

import (
	"net/http"
	//"github.com/nsan1129/unframed"
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
	sr.Get("/delete/{Id}", lfgsDelete)
	sr.Get("/oust/{Id}", lfgsOust)
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

	da := new(lfgsAdapter)
	ca := new(classesAdapter).list()
	qa := new(queuePrefsAdapter).list()
	ra := new(rankingPagesAdapter).list()

	id := 0
	net.SetSession(r)
	if val,ok := net.Session.Values["lfg_id"]; ok {
	    id = val.(int)
		//log.Message("Loading LFg, Id:", id)
	}

	if id == 0 {
		_ = da.newLfg()
	} else {
		da.show(id)
	}
	net.ExeTmpl(w, "lfgsForm", da, ca, qa, ra)
}

func lfgsSave(w http.ResponseWriter, r *http.Request) {

	dm := new(lfg)

	net.DecodeForm(dm, r)

	lastId := new(lfgsAdapter).save(dm)

	net.SetSession(r)
	net.Session.Values["lfg_id"] = lastId
	net.Session.AddFlash("new LFG Saved: ")
	net.Session.AddFlash(net.Session.Values["lfg_id"].(int))
	net.Session.Save(r,w)

	http.Redirect(w, r, "/lfgs/list", http.StatusFound)

}

func lfgsDelete(w http.ResponseWriter, r *http.Request) {

	/* id := net.QueryUrl("Id", r) */
	net.SetSession(r)
	if val,ok := net.Session.Values["lfg_id"]; ok {
	    id := val.(int)
	    da := new(lfgsAdapter)
		da.delete(id)
		net.Session.AddFlash("LFG Listing Deleted: ")
		net.Session.AddFlash(net.Session.Values["lfg_id"].(int))
	} else {
		net.Session.AddFlash("Could Not Delete LFG Listing: ")
		net.Session.AddFlash(net.Session.Values["lfg_id"].(int))
	}
	
	net.Session.Values["lfg_id"] = 0
	net.Session.Save(r,w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func lfgsOust(w http.ResponseWriter, r *http.Request) {
	net.SetSession(r)
	id := net.QueryUrl("Id", r)

	if val,ok := net.Session.Values["ousts"]; ok {
		net.Session.Values["ousts"] = net.StrAppendInt(val.(string), id, ",")
	} else {
		net.Session.Values["ousts"] = net.StrAppendInt("", id, "")
	}

	net.Session.Save(r,w)
	da := new(lfgsAdapter)
	da.oust(id)

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetLfgsList() *lfgsAdapter {
	return new(lfgsAdapter).list()
}
