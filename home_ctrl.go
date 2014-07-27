package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
)

func homeReg() {
	homeTemplates()
	homeRoutes()

	otherStmts()

	net.RegType(&Session{})
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

	//log.Message(len(dataModel.Lfgs))
	net.ExeTmpl(w, "home", lfgsA, lfmsA, ra)
}

func help(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "help", nil)
}