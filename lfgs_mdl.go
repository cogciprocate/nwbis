package main

import (
	"github.com/c0gent/unframed"
	"time"
	//"github.com/c0gent/unframed/log"
)

type lfg struct {
	Id            int
	CharacterName string
	CharacterId   int
	UserName      string
	UserId        int
	ClassId       int
	RankingPage   int
	QueuePrefId   int
	AddedAt       time.Time
	ClassAbbr     string
	ClassName     string
	QueuePrefName string
	TimeSince     string
	Ousts         int
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

func (da *lfgsAdapter) oust(id int) {
	da.Exec(db.Stmts["oustLfg"], id)
}

func (da *lfgsAdapter) IsOusted(lfgId int, oustList string, ousts int) string {
	if ousts > 2 {
		return `class="strikethrough"`
	}
	if net.IntInStr(oustList, lfgId, ",") {
		return `class="strikethrough"`
	}
	return ``
}

func lfgsStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("createLfg",
		d,
		`INSERT INTO lfgs (
			character_name,
			user_name, 
			class_id, 
			ranking_page, 
			queue_pref_id,
			added_at
		) VALUES ($1, $2, $3, $4, $5, NOW()) 
		RETURNING id;`,
	)

	db.AddStatement("listLfgs",
		d,
		`SELECT 
			l.id,
			l.character_name,
			l.character_id,
			l.user_name,
			l.user_id,
			l.class_id,
			l.ranking_page,
			l.queue_pref_id,
			l.added_at,
			c.abbr,
			c.name,
			q.name,
			(NOW() - l.added_at) as time_since,
			l.ousts

		FROM lfgs l, classes c, queue_prefs q
		WHERE c.id = l.class_id 
			AND q.id = l.queue_pref_id 
			AND (NOW() - l.added_at) < (interval '60 minutes')
			AND l.ousts < 3
		ORDER BY l.added_at DESC
		LIMIT $1;
		`,
	)

	db.AddStatement("showLfg",
		d,
		`
		SELECT 
			l.id,
			l.character_name,
			l.character_id,
			l.user_name,
			l.user_id,
			l.class_id,
			l.ranking_page,
			l.queue_pref_id,
			l.added_at,
			c.abbr,
			c.name,
			q.name,
			(NOW() - l.added_at) as time_since,
			l.ousts
		FROM lfgs l, classes c, queue_prefs q
		WHERE c.id = l.class_id 
			AND q.id = l.queue_pref_id
			AND l.id = $1
			;
		`,
	)

	//lfg.id, lfg.CharacterName, lfg.UserName, lfg.ClassId, lfg.CurrentPage, lfg.QueuePrefId
	db.AddStatement("updateLfg",
		d,
		`UPDATE lfgs SET 
			character_name = $2, 
			user_name = $3, 
			class_id = $4, 
			ranking_page = $5, 
			queue_pref_id = $6,
			added_at = NOW(),
			ousts = 0
		WHERE id = $1;`,
	)

	db.AddStatement("deleteLfg",
		d,
		`UPDATE lfgs SET 
			ousts = ousts + 100
		WHERE id = $1;`,
	)

	db.AddStatement("oustLfg",
		d,
		`UPDATE lfgs SET 
			ousts = ousts + 1
		WHERE id = $1;`,
	)
	/**/
}
