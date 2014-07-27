package main

import (
	"github.com/nsan1129/unframed"
)

func otherStmts() {
	d := unframed.Dbd.Pg

	
	db.AddStatement("listQueuePrefs",
		d,
		`
		SELECT
		*
		FROM queue_prefs;
		`,
	)
	db.AddStatement("listClasses",
		d,
		`
		SELECT
		*
		FROM classes;
		`,
	)
}
