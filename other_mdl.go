package main

import (
	"github.com/nsan1129/unframed"
	//"time"
	//"github.com/nsan1129/unframed/log"
)

type queuePref struct {
	Id        int
	Name string
}

type queuePrefsAdapter struct {
	unframed.DataAdapter
	QueuePrefs []*queuePref
}

	func (da *queuePrefsAdapter) list() *queuePrefsAdapter {
		da.Query(da.newQueuePref, db.Stmts["listQueuePrefs"])
		return da
	}

	func (da *queuePrefsAdapter) newQueuePref() (inf interface{}) {
		
		dr := new(queuePref)
		da.QueuePrefs = append(da.QueuePrefs, dr)
		inf = dr
		
		return
	}


type class struct {
	Id        int
	Name string
	Abbr string
}

type classesAdapter struct {
	unframed.DataAdapter
	Classes []*class
}

	func (da *classesAdapter) list() *classesAdapter {
		da.Query(da.newClass, db.Stmts["listClasses"])
		return da
	}

	func (da *classesAdapter) newClass() (inf interface{}) {
		
		dr := new(class)
		da.Classes = append(da.Classes, dr)

		inf = dr
		return
	}

type rankingPage struct {
	Id int
	Name string
}

type rankingPagesAdapter struct {
	unframed.DataAdapter
	RankingPages []*rankingPage
}
	func (da *rankingPagesAdapter) list() *rankingPagesAdapter {
		da.RankingPages = append(da.RankingPages, &rankingPage{1, "1-19"})
		da.RankingPages = append(da.RankingPages, &rankingPage{2, "20-99"})
		da.RankingPages = append(da.RankingPages, &rankingPage{3, "100+"})
		return da
	}

	func (da *rankingPagesAdapter) AsText(x int) (page string) {
		if (x <= len(da.RankingPages)) {
			page = da.RankingPages[x-1].Name
		} else {
			page = "?"
		}
		return

	}