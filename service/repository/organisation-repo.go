package repository

import (
	"build-version/config"
	"build-version/model/data"
)

func getOrganisationById(organisationId string) (*data.OrganisationData, error){
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		ret := data.OrganisationData{}
		row := db.QueryRow("SELECT * FROM tbl_organisation WHERE id = ?", organisationId)
		if err = row.Scan(&ret); err != nil {
			return nil, err
		} else {
			return &ret, nil
		}
	}
}

func getOrganisationByName(name string) (*data.OrganisationData, error) {
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		ret := data.OrganisationData{}
		row := db.QueryRow("SELECT * FROM tbl_organisation WHERE name = ?", name)
		if err = row.Scan(&ret); err != nil {
			return nil, err
		} else {
			return &ret, nil
		}
	}
}

func createOrganisation(data *data.OrganisationData) error {
	query := `
		INSERT INTO tbl_organisation(id, name, plan_type)
		VALUES(?,?,?);
	`
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return err
	} else {
		if _, err := db.Exec(query, data.Id, data.Name, data.PlanType); err != nil {
			return err
		} else {
			return nil
		}
	}
}
