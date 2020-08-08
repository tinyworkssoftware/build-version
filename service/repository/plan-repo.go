package repository

import (
	"build-version/config"
	"build-version/model/data"
	model "build-version/model/toml"
	"log"
)

var (
	conf *model.TomlConfig
)

func GetPlans() (*[]data.PlanData, error) {
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		defer db.Close()
		data := []data.PlanData{}
		err = db.Select(&data, "SELECT * FROM tbl_plan_type")
		log.Println(err)
		return &data, nil
	}
}