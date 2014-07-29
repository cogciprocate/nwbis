package main

import (
	"github.com/nsan1129/unframed"
)

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
			(NOW() - l.added_at) as time_since

		FROM lfgs l, classes c, queue_prefs q
		WHERE c.id = l.class_id 
			AND q.id = l.queue_pref_id 
			AND (NOW() - l.added_at) < (interval '240 minutes')
		ORDER BY l.id DESC
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
			(NOW() - l.added_at) as time_since
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
			added_at = NOW()
		WHERE id = $1;`,
	)
	
	db.AddStatement("deleteLfg",
		d,
		`DELETE FROM lfgs
		WHERE id = $1;`,
	)
}
