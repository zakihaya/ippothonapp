package controllers

import (
    "github.com/revel/revel"
)

func init() {
    revel.OnAppStart(InitDB) // DBやテーブルの作成
    revel.InterceptMethod((*GorpController).Begin, revel.BEFORE) // transaction開始
    revel.InterceptMethod((*GorpController).Commit, revel.AFTER) // 変更反映
    revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY) // 異常時処理
}
