package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Dbm *gorp.DbMap
	Txn *gorp.Transaction
)

func InitDB() {
	//db.Init()
	db, err := sql.Open("mysql", "issue_manager:Example_passwd123_@/issue_management")
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	Dbm.TraceOn("[gorp]", log.New(os.Stdout, "ginsample:", log.Lmicroseconds))
	/*
		Dbm.AddTableWithName(ServiceData{}, "service").SetKeys(true, "ID")
		Dbm.AddTableWithName(ServiceIssueData{}, "service_issue").SetKeys(true, "ID")
		Dbm.AddTableWithName(UserData{}, "user").SetKeys(true, "ID")
	*/
}

type GorpController struct {
}

func (c *GorpController) Begin() error {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	Txn = txn
	return nil
}

func (c *GorpController) Commit() error {
	if Txn == nil {
		return nil
	}
	if err := Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return nil
}

func (c *GorpController) Rollback() error {
	if Txn == nil {
		return nil
	}
	if err := Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	Txn = nil
	return nil
}
