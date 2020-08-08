package service

import (
	"build-version/model/data"
	"build-version/service/repository"
)

func GetAllPlans() (*[]data.PlanData, error) {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if results, err := repository.GetPlans(db); err != nil {
			return nil, err
		} else {
			return results, nil
		}
	}
}
