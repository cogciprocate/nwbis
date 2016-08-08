package main

import (
	"github.com/c0gent/unframed"
	"time"
	//"github.com/c0gent/unframed/log"
)

type inhouse struct {
	Id                     int
	OrganizerCharacterName string
	OrganizerCharacterId   int
	OrganizerUserName      string
	OrganizerUserId        int
	ScheduledAt            string
	AddedAt                time.Time
	TimeSince              string
}

type inhousesAdapter struct {
	unframed.DataAdapter
	Inhouses []*inhouse
}

func (da *inhousesAdapter) list() *inhousesAdapter {
	da.Query(da.newInhouse, db.Stmts["listInhouses"], 50)
	return da
}

func (da *inhousesAdapter) show(id int) *inhousesAdapter {
	da.Query(da.newInhouse, db.Stmts["showInhouse"], id)
	return da
}

func (da *inhousesAdapter) delete(id int) {
	da.Exec(db.Stmts["deleteInhouse"], id)
}

func (da *inhousesAdapter) save(inhouse *inhouse) (lastId int) {

	if inhouse.Id == 0 {
		lastId = da.Insert(db.Stmts["createInhouse"], inhouse.OrganizerCharacterName, inhouse.OrganizerUserName, inhouse.ScheduledAt)
	} else {
		da.Exec(db.Stmts["updateInhouse"], inhouse.Id, inhouse.OrganizerCharacterName, inhouse.OrganizerUserName, inhouse.ScheduledAt)
		lastId = inhouse.Id
	}
	return
}

func (da *inhousesAdapter) newInhouse() (inf interface{}) {

	inhouse := new(inhouse)
	da.Inhouses = append(da.Inhouses, inhouse)
	inf = inhouse

	return
}

func inhousesStmts() {
	d := unframed.Dbd.Pg

	db.AddStatement("createInhouse",
		d,
		`INSERT INTO inhouses (
			organizer_character_name, 
			organizer_user_name, 
			scheduled_at, 
			added_at
		) VALUES ($1, $2, $3, NOW()) 
		RETURNING id;`,
	)

	db.AddStatement("listInhouses",
		d,
		`
		SELECT 
			i.id,
			i.organizer_character_name,
			i.organizer_character_id,
			i.organizer_user_name,
			i.organizer_user_id,
			i.scheduled_at,
			i.added_at,
			(NOW() - i.added_at) as time_since
		FROM inhouses i
		ORDER BY i.id DESC
		LIMIT $1;
		`,
	)

	db.AddStatement("showInhouse",
		d,
		`
		SELECT 
			i.id,
			i.organizer_character_name,
			i.organizer_character_id,
			i.organizer_user_name,
			i.organizer_user_id,
			i.scheduled_at,
			i.added_at,
			(NOW() - i.added_at) as time_since
		FROM inhouses i
		WHERE i.id = $1;`,
	)

	db.AddStatement("updateInhouse",
		d,
		`UPDATE inhouses SET 
			organizer_character_name = $2, 
			organizer_user_name = $3, 
			scheduled_at = $4,
			added_at = NOW()
		WHERE id = $1;`,
	)

	db.AddStatement("deleteInhouse",
		d,
		`DELETE FROM inhouses
		WHERE id = $1;`,
	)

}
