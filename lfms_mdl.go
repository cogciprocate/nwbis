package main

import (
	"github.com/c0gent/unframed"
	"time"
	//"github.com/c0gent/unframed/log"
)

type lfm struct {
	Id            int
	CharacterName string
	CharacterId   int
	UserName      string
	UserId        int
	Rainbow       string
	AddedAt       time.Time
	TimeSince     string
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

func (da *lfmsAdapter) save(lfm *lfm) (lastId int) {

	if lfm.Id == 0 {
		lastId = da.Insert(db.Stmts["createLfm"], lfm.CharacterName, lfm.UserName, lfm.Rainbow)
	} else {
		da.Exec(db.Stmts["updateLfm"], lfm.Id, lfm.CharacterName, lfm.UserName, lfm.Rainbow)
		lastId = lfm.Id
	}
	return
}

func (da *lfmsAdapter) newLfm() (inf interface{}) {

	lfm := new(lfm)
	da.Lfms = append(da.Lfms, lfm)
	inf = lfm

	return
}

func lfmsStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("createLfm",
		d,
		`INSERT INTO lfms (
			character_name, 
			user_name, 
			rainbow, 
			added_at
		) VALUES ($1, $2, $3, NOW()) 
		RETURNING id;`,
	)

	db.AddStatement("listLfms",
		d,
		`
		SELECT 
			l.id,
			l.character_name,
			l.character_id,
			l.user_name,
			l.user_id,
			l.rainbow,
			l.added_at,
			(NOW() - l.added_at) as time_since
		FROM lfms l
		WHERE (NOW() - l.added_at) < (interval '480 minutes')
		ORDER BY l.id DESC
		LIMIT $1;
		`,
	)

	db.AddStatement("showLfm",
		d,
		`
		SELECT 
			l.id,
			l.character_name,
			l.character_id,
			l.user_name,
			l.user_id,
			l.rainbow,
			l.added_at,
			(NOW() - l.added_at) as time_since

		FROM lfms l
		WHERE l.id = $1;`,
	)

	db.AddStatement("updateLfm",
		d,
		`UPDATE lfms SET 
			character_name = $2, 
			user_name = $3, 
			rainbow = $4,
			added_at = NOW()
		WHERE id = $1;`,
	)

	db.AddStatement("deleteLfm",
		d,
		`DELETE FROM lfms
		WHERE id = $1;`,
	)

}
