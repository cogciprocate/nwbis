package main

import (
	"net/http"
	//"github.com/c0gent/unframed"
	"github.com/c0gent/unframed/log"
)

func homeReg() {
	homeTemplates()
	homeRoutes()

	otherStmts()

}

func homeTemplates() {
	net.TemplateFiles(
		"tmpl/home.html.tmpl",
		"tmpl/help.html.tmpl",
	)
}

func homeRoutes() {
	net.Get("/", home)
	net.Get("/help", help)
}

func home(w http.ResponseWriter, r *http.Request) {

	lfgsA := new(lfgsAdapter).list()
	lfmsA := new(lfmsAdapter).list()
	ra := new(rankingPagesAdapter).list()

	net.SetSession(r)

	if flashes := net.Session.Flashes(); len(flashes) > 0 {
        // Just print the flash values.
        log.Message(flashes)
    }

    lfm_id := 0
	if val,ok := net.Session.Values["lfm_id"]; ok {
		lfm_id = val.(int)
	}

	lfg_id := 0
	if val,ok := net.Session.Values["lfg_id"]; ok {
		lfg_id = val.(int)
	}

	ou := ""
	if val,ok := net.Session.Values["ousts"]; ok {
		ou = val.(string)
		//log.Message("ousts session value found and sent to tmpl")
	}	

    net.Session.Save(r,w)
	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "home", lfgsA, lfmsA, ra, lfg_id, lfm_id, ou)
}

func help(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "help", nil)
}