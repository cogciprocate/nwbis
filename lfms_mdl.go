package main

import (
	"github.com/nsan1129/unframed"
	"time"
	//"github.com/nsan1129/unframed/log"
)

type lfm struct {
	Id        int
	CharacterName string
	CharacterId   int
	UserName       string
	UserId    int
	Rainbow	bool
	AddedAt time.Time
	TimeSince string
}

type lfmsAdapter struct {
	unframed.DataAdapter
	Lfms []*lfm
}

func (da *lfmsAdapter) list() *lfmsAdapter {
	da.Query(da.newLfm, db.Stmts["listLfms"], 50)
	return da
}

func (da *lfmsAdapter) show(id int) *lfmsAdapter {
	da.Query(da.newLfm, db.Stmts["showLfm"], id)
	return da
}

func (da *lfmsAdapter) delete(id int) {
	da.Exec(db.Stmts["deleteLfm"], id)
}

func (da *lfmsAdapter) save(lfm *lfm) {
	
	if lfm.Id == 0 {
		da.Exec(db.Stmts["createLfm"], lfm.CharacterName, lfm.UserName, lfm.Rainbow)
	} else {
		da.Exec(db.Stmts["updateLfm"], lfm.Id, lfm.CharacterName, lfm.UserName, lfm.Rainbow)
	}
	
}

func (da *lfmsAdapter) newLfm() (inf interface{}) {
	
	lfm := new(lfm)
	da.Lfms = append(da.Lfms, lfm)
	inf = lfm
	
	
	return
}
