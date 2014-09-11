package main

import (
	_ "github.com/lib/pq"
	"github.com/nsan1129/unframed"
	"github.com/nsan1129/unframed/log"
)

var net *unframed.NetHandle
var db *unframed.DbHandle

var cfgFile string = "/etc/nwbis/config.json"

func main() {
	cfg := unframed.ReadConfig(cfgFile)

	db = unframed.NewDB(cfg.DbType, cfg.ConnStr)
	defer db.Close()
	log.Message(cfg.Wd)
	net = unframed.NewNet()

	unframed.DefaultPageTitle = "nwBiS"
	net.TemplateFiles(
		"tmpl/_base.html.tmpl",
		"tmpl/_side_table_1.html.tmpl",
	)
	//net.Get("/episodes/watch", episodesWatch)
	//net.Get("/episodes", episodesList)
	net.Dir("assets/")
	net.Dir("public/")

	registerControllers()

	db.PrepareStatements()
	net.LoadTemplates()

	log.Message("Serving nwBiS")
	net.Serve(cfg.ListenPort)

}

func registerControllers() {
	lfgsReg()
	lfmsReg()
	homeReg()
	codingChallengeReg()
}

