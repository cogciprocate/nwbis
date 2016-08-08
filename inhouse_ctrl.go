package main

import (
	"github.com/c0gent/unframed/log"
	"net/http"
	//"github.com/c0gent/unframed"
)

func inhousesReg() {
	inhousesTemplates()
	inhousesRoutes()
	inhousesStmts()

	net.RegType(&inhouse{})
}

func inhousesTemplates() {
	net.TemplateFiles(
		"tmpl/_inhouses.html.tmpl",
	)
	//log.Message("inhousesTemplates run")
}

func inhousesRoutes() {
	sr := net.Subrouter("/inhouses")
	sr.Get("/list", inhousesList)
	sr.Get("/form/{Id}", inhousesForm)
	sr.Post("/save", inhousesSave)
	sr.Get("/delete/{Id}", inhousesDelete)
}

/*
func inhousesPage(w http.ResponseWriter, r *http.Request) {
	da := new(inhousesAdapter).list()
	net.SetSession(r)
	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "inhousesPage", da)
}
*/

func inhousesList(w http.ResponseWriter, r *http.Request) {
	da := new(inhousesAdapter).list()
	net.SetSession(r)
	//log.Message(len(dataModel.Inhouses))
	if flashes := net.Session.Flashes(); len(flashes) > 0 {
		// Just print the flash values.
		//log.Message(flashes)
	} else {
		net.Session.AddFlash("Hello, flash messages world!")
		//log.Message("No flashes found.")
	}
	net.Session.Save(r, w)
	net.ExeTmpl(w, "inhousesList", da)
}

func inhousesForm(w http.ResponseWriter, r *http.Request) {

	da := new(inhousesAdapter)

	id := 0
	net.SetSession(r)
	if val, ok := net.Session.Values["inhouse_id"]; ok {
		id = val.(int)
		log.Message("Loading Inhouse, Id:", id)
	}

	if id == 0 {
		_ = da.newInhouse()
	} else {
		da.show(id)
	}
	net.ExeTmpl(w, "inhousesForm", da)
}

func inhousesSave(w http.ResponseWriter, r *http.Request) {

	dm := new(inhouse)

	net.DecodeForm(dm, r)

	lastId := new(inhousesAdapter).save(dm)

	net.SetSession(r)
	net.Session.Values["inhouse_id"] = lastId
	net.Session.AddFlash("new Inhouse Saved: ")
	net.Session.AddFlash(net.Session.Values["inhouse_id"].(int))
	net.Session.Save(r, w)

	http.Redirect(w, r, "/inhouses/list", http.StatusFound)

}

func inhousesDelete(w http.ResponseWriter, r *http.Request) {

	net.SetSession(r)
	if val, ok := net.Session.Values["inhouse_id"]; ok {
		id := val.(int)
		ss := new(inhousesAdapter)
		ss.delete(id)
		net.Session.AddFlash("Inhouse Listing Deleted: ")
		net.Session.AddFlash(net.Session.Values["inhouse_id"].(int))
	} else {
		net.Session.AddFlash("Could Not Delete Inhouse Listing: ")
		net.Session.AddFlash(net.Session.Values["inhouse_id"].(int))
	}

	net.Session.Values["inhouse_id"] = 0
	net.Session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetInhousesList() *inhousesAdapter {
	return new(inhousesAdapter).list()
}
