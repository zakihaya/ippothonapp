package controllers

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    "github.com/revel/revel" // あとで使います
    "ippothonapp/app/models" // revel new APP_NAME の APP_NAME
)

var (
    DbMap *gorp.DbMap // このデータベースマッパーからSQLを流す
)

func InitDB() {
    db, err := sql.Open("sqlite3", "./tmp/scrape.db")
    if err != nil {
        panic(err.Error())
    }
    DbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

    // ここで好きにテーブルを定義する
    t := DbMap.AddTable(models.Stock{}).SetKeys(true, "Id")
    t.ColMap("CompanyName").MaxSize = 200

    DbMap.CreateTables()
}

type GorpController struct {
    *revel.Controller
    Transaction *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
    txn, err := DbMap.Begin() // ここで開始したtransactionをCOMMITする
    if err != nil {
        panic(err)
    }
    c.Transaction = txn
    return nil
}

func (c *GorpController) Commit() revel.Result {
    if c.Transaction == nil {
        return nil
    }
    err := c.Transaction.Commit() // SQLによる変更をDBに反映
    if err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Transaction = nil // 正常終了した場合はROLLBACK処理に入らないようにする
    return nil
}

func (c *GorpController) Rollback() revel.Result {
    if c.Transaction == nil {
        return nil
    }
    err := c.Transaction.Rollback() // 問題があった場合変更前の状態に戻す
    if err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Transaction = nil
    return nil
}
