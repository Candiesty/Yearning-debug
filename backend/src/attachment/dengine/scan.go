package dengine

import (
	"Yearning-go/src/attachment/dmessage"

	"gorm.io/gorm"
)

const MAX_COL = 300

var cols [MAX_COL]string

func FetchTables(db *gorm.DB) []string {
	var res []string

	rows, _ := db.Raw("select tablename from pg_tables where schemaname='public'").Rows()
	col, _ := rows.Columns()
	dmessage.PrintV("Fetch Tables:", col, len(col))
	_tmp := UnifiedLabel(col)
	for rows.Next() {
		rows.Scan(_tmp...)
		for i := 0; i < len(col); i++ {
			j := *_tmp[i].(*string)
			res = append(res, j)
		}
	}
	dmessage.PrintV(res)
	return res
}

func FetchField(t string, db *gorm.DB) []string {
	var res []string

	rows, _ := db.Raw("SELECT column_name FROM information_schema.columns WHERE table_name = ?;", t).Rows()
	col, _ := rows.Columns()
	dmessage.PrintV("Fetch Table:", col, len(col))
	_tmp := UnifiedLabel(col)
	for rows.Next() {
		rows.Scan(_tmp...)
		for i := 0; i < len(col); i++ {
			j := *_tmp[i].(*string)
			res = append(res, j)
		}
	}
	dmessage.PrintV(res)
	return res
}

func FetchTable(t string, db *gorm.DB) ([]string, int) {
	var res []string

	rows, _ := db.Raw("SELECT column_name,data_type,is_nullable,column_default FROM information_schema.columns WHERE table_name = ?;", t).Rows()
	col, _ := rows.Columns()
	dmessage.PrintV("Fetch Table:", col, len(col))
	_tmp := UnifiedLabel(col)
	for rows.Next() {
		rows.Scan(_tmp...)
		for i := 0; i < len(col); i++ {
			j := *_tmp[i].(*string)
			res = append(res, j)
		}
	}
	dmessage.PrintV(res)
	return res, len(col)
}
