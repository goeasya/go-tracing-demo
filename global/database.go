package global

import "go-tracing-demo/pkg/database"

var GDBManager database.Manager

type DBConfig struct {
	Type string
	ConnInfo string
}

func MustInitDBManager(cfg *DBConfig) (err error){
	if cfg.Type == "gorm" {
		GDBManager, err = database.NewGormDB(cfg.ConnInfo)
	}
	return err
}

func GetDBManager() database.Manager {
	return GDBManager
}
