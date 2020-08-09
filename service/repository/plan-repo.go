package repository

import (
	"build-version/model/data"
	model "build-version/model/toml"
	"github.com/jmoiron/sqlx"
)

var (
	conf *model.TomlConfig
)

func GetPlans(db *sqlx.DB,) (*[]data.PlanData, error) {
	var data = make([]data.PlanData, 0)
	if err := db.Select(&data, "SELECT * FROM tbl_plan_type"); err != nil {
		return nil, err
	} else {
		return &data, nil
	}
}