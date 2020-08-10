package repository

import (
	"build-version/model/data"
	"github.com/jmoiron/sqlx"
)

func GetOrganisationById(db *sqlx.DB, organisationId string) (*data.OrganisationData, error){
	var record data.OrganisationData
	if err := db.QueryRowx("SELECT * FROM tbl_organisation WHERE id = ?", organisationId).StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetOrganisationByName(db *sqlx.DB, name string) (*data.OrganisationData, error) {
	ret := data.OrganisationData{}
	row := db.QueryRow("SELECT * FROM tbl_organisation WHERE name = ?", name)
	if err := row.Scan(&ret); err != nil {
		return nil, err
	} else {
		return &ret, nil
	}
}

func GetAllOrganisation(db *sqlx.DB, ) (*[]data.OrganisationData, error) {
	query := `SELECT * FROM tbl_organisation;`
	var records = make([]data.OrganisationData, 0)
	if err := db.Select(&records, query); err != nil {
		return nil, err
	} else {
			return &records, nil
	}
}

func CreateOrganisation(db *sqlx.DB, data *data.OrganisationData) error {
	query := `
		INSERT INTO tbl_organisation(id, name, plan_type)
		VALUES(?,?,?);
	`
	if _, err := db.Exec(query, data.Id, data.Name, data.PlanType); err != nil {
		return err
	} else {
		return nil
	}
}

