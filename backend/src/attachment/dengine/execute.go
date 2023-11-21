package dengine

import (
	"Yearning-go/src/attachment/dmessage"
	"Yearning-go/src/lib"
	"Yearning-go/src/model"
	"errors"
	"fmt"
	"time"

	drive "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func UnifiedLabel(col []string) []any {
	var s []any
	for i := 0; i < len(col); i++ {
		var t string
		s = append(s, &t)
	}
	return s
}

// 连接到数据库，传入的CoreDataSource应被解密，针对需要
func ConnPG(u *model.CoreDataSource) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", u.IP, u.Username, u.Password, u.DataBase, u.Port)
	dmessage.PrintV(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		dmessage.PrintV("failed to connect postgresql")
		return nil, err
	}
	dmessage.PrintV("connect postgresql")
	return db, err
}

func ConnMYSQL(u *model.CoreDataSource) (*gorm.DB, error) {
	dsn, err := model.InitDSN(model.DSN{
		Username: u.Username,
		Password: u.Password,
		Host:     u.IP,
		Port:     u.Port,
		DBName:   u.DataBase,
		CA:       u.CAFile,
		Cert:     u.Cert,
		Key:      u.KeyFile,
		DBType:   u.DBType,
	})
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(drive.New(drive.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		dmessage.PrintV("failed to connect mysql")
		return nil, err
	}
	dmessage.PrintV("connect mysql")
	return db, err
}

func SuperConnDB(u *model.CoreDataSource) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch u.DBType {
	case 0: // Mysql
		db, err = ConnMYSQL(u)
	case 1: // PG
		db, err = ConnPG(u)
	default:
		dmessage.PrintV("failed connect")
		err = errors.New("failed connect")
	}
	return db, err
}

func DEC(u *model.CoreDataSource) {
	if u.Password != "" && lib.Decrypt(model.JWT, u.Password) != "" {
		u.Password = lib.Decrypt(model.JWT, u.Password)
	}
}

func ExecuteSQL(order *model.CoreSqlOrder, u *model.CoreDataSource) error {
	var err error
	DEC(u)
	if u.DBType == 0 {
		u.DataBase = order.DataBase
	}
	db, err := SuperConnDB(u)
	if err != nil {
		return err
	}

	result := db.Exec(order.SQL)
	dmessage.PrintV(result)
	err = result.Error
	if err != nil {
		//执行失败 4
		model.DB().Model(&model.CoreSqlOrder{}).Where("work_id =?", order.WorkId).Updates(map[string]interface{}{"status": 4})
		model.DB().Create(&model.CoreSqlRecord{
			ID:        order.ID,
			WorkId:    order.WorkId,
			SQL:       order.SQL,
			State:     "失败",
			Affectrow: uint(result.RowsAffected),
			Time:      time.Now().Format("2006-01-02 15:04"),
			Error:     err.Error(),
		})
		return err
	}
	//执行成功 1
	model.DB().Model(&model.CoreSqlOrder{}).Where("work_id =?", order.WorkId).Updates(map[string]interface{}{"status": 1})
	model.DB().Create(&model.CoreSqlRecord{
		ID:        order.ID,
		WorkId:    order.WorkId,
		SQL:       order.SQL,
		State:     "成功",
		Affectrow: uint(result.RowsAffected),
		Time:      time.Now().Format("2006-01-02 15:04"),
		Error:     "无",
	})
	return nil
}
