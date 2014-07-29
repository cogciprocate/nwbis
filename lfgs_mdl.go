package main

import (
	"github.com/nsan1129/unframed"
	"time"
	//"github.com/nsan1129/unframed/log"
)

type lfg struct {
	Id        int
	CharacterName string
	CharacterId   int
	UserName       string
	UserId    int
	ClassId	int
	RankingPage   int
	QueuePrefId int
	AddedAt time.Time
	ClassAbbr  string
	ClassName	string
	QueuePrefName string
	TimeSince string
}

type lfgsAdapter struct {
	unframed.DataAdapter
	Lfgs []*lfg
}

func (da *lfgsAdapter) list() *lfgsAdapter {
	da.Query(da.newLfg, db.Stmts["listLfgs"], 50)
	return da
}

func (da *lfgsAdapter) show(id int) *lfgsAdapter {
	da.Query(da.newLfg, db.Stmts["showLfg"], id)
	//log.Message(da.Lfgs[0].AddedAt)
	return da
}

func (da *lfgsAdapter) delete(id int) {
	da.Exec(db.Stmts["deleteLfg"], id)
}

func (da *lfgsAdapter) save(lfg *lfg) (lastId int) {
	
	if lfg.Id == 0 {
		lastId = da.Insert(db.Stmts["createLfg"], lfg.CharacterName, lfg.UserName, lfg.ClassId, lfg.RankingPage, lfg.QueuePrefId)
	} else {
		da.Exec(db.Stmts["updateLfg"], lfg.Id, lfg.CharacterName, lfg.UserName, lfg.ClassId, lfg.RankingPage, lfg.QueuePrefId)
		lastId = lfg.Id
	}
	return
}

func (da *lfgsAdapter) newLfg() (inf interface{}) {
	
	lfg := new(lfg)
	da.Lfgs = append(da.Lfgs, lfg)
	inf = lfg
	
	return
}
