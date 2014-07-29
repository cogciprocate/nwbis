package main

import (
	"net/http"
	"github.com/nsan1129/unframed/log"
	//"github.com/nsan1129/unframed"
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
	sr.Get("/delete/{Id}", lfmsDelete)
}

/*
func lfmsPage(w http.ResponseWriter, r *http.Request) {
	da := new(lfmsAdapter).list()
	net.SetSession(r)
	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "lfmsPage", da)
}

func lfmsList(w http.ResponseWriter, r *http.Request) {
	da := new(lfmsAdapter).list()
	net.SetSession(r)
	//log.Message(len(dataModel.Lfms))
	if flashes := net.Session.Flashes(); len(flashes) > 0 {
        // Just print the flash values.
        log.Message(flashes)
    } else {
    	net.Session.AddFlash("Hello, flash messages world!")
        log.Message("No flashes found.")
    }
    net.Session.Save(r,w)
	net.ExeTmpl(w, "lfmsList", da)
}
*/

func lfmsForm(w http.ResponseWriter, r *http.Request) {

	da := new(lfmsAdapter)

	id := 0
	net.SetSession(r)
	if val,ok := net.Session.Values["lfm_id"]; ok {
		id = val.(int)
		log.Message("Loading LFM, Id:", id) 
	}

	if id == 0 {
		_ = da.newLfm()
	} else {
		da.show(id)
	}
	net.ExeTmpl(w, "lfmsForm", da)
}

func lfmsSave(w http.ResponseWriter, r *http.Request) {

	dm := new(lfm)

	net.DecodeForm(dm, r)

	lastId := new(lfmsAdapter).save(dm)

	net.SetSession(r)
	net.Session.Values["lfm_id"] = lastId
	net.Session.AddFlash("new LFM Saved: ")
	net.Session.AddFlash(net.Session.Values["lfm_id"].(int))
	net.Session.Save(r,w)

	http.Redirect(w, r, "/lfms/list", http.StatusFound)

}

func lfmsDelete(w http.ResponseWriter, r *http.Request) {

	net.SetSession(r)
	if val,ok := net.Session.Values["lfm_id"]; ok {
	    id := val.(int)
	    ss := new(lfmsAdapter)
		ss.delete(id)
		net.Session.AddFlash("LFM Listing Deleted: ")
		net.Session.AddFlash(net.Session.Values["lfm_id"].(int))
	} else {
		net.Session.AddFlash("Could Not Delete LFM Listing: ")
		net.Session.AddFlash(net.Session.Values["lfm_id"].(int))
	}
	
	net.Session.Values["lfm_id"] = 0
	net.Session.Save(r,w)

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetLfmsList() *lfmsAdapter {
	return new(lfmsAdapter).list()
}