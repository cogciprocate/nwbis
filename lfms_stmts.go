package main

import (
	"github.com/nsan1129/unframed"
)

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
