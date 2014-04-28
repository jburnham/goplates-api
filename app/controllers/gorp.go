package controllers

import (
	"database/sql"
	"goplates-api/app/models"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	//setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
	//for col, size := range colSizes {
	//t.ColMap(col).MaxSize = size
	//}
	//}

	Dbm.AddTable(models.Employee{}).SetKeys(true, "Id")
	//setColumnSizes(t, map[string]int{
	//"Name":  100,
	//"Email": 255,
	//})
	Dbm.AddTable(models.Plate{}).SetKeys(false, "Plate")
	//setColumnSizes(t, map[string]int{
	//"Plate": 20,
	//})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	employees := []*models.Employee{
		&models.Employee{Name: "John Doe"},
	}
	plates := []*models.Plate{
		&models.Plate{"ABC123", 1},
	}
	for _, employee := range employees {
		if err := Dbm.Insert(employee); err != nil {
			panic(err)
		}
	}
	for _, plate := range plates {
		if err := Dbm.Insert(plate); err != nil {
			panic(err)
		}
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
