package repository

import (
	"build-version/model/data"
	"github.com/jmoiron/sqlx"
)

func GetProjectById(db *sqlx.DB, projectId string) (*data.ProjectData, error){
	record := data.ProjectData{}
	if err := db.QueryRowx("SELECT * FROM tbl_project WHERE id = ?", projectId).StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetProjectByName(db *sqlx.DB, name string) (*data.ProjectData, error) {
	record := data.ProjectData{}
	if err := db.QueryRowx("SELECT * FROM tbl_project WHERE name = ?", name).StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetProjectByAccessToken(db *sqlx.DB, accessToken string) (*data.ProjectData, error) {
	record := data.ProjectData{}
	if err := db.QueryRowx("SELECT * FROM tbl_project WHERE access_code = ?", accessToken).StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetAllProjects(db *sqlx.DB) (*[]data.ProjectData, error) {
	query := `SELECT * FROM tbl_project;`
	var records = make([]data.ProjectData, 0)
	if err := db.Select(&records, query); err != nil {
		return nil, err
	} else {
		return &records, nil
	}
}

func CreateProject(db *sqlx.DB, data *data.ProjectData) error {
	query := `
		INSERT INTO tbl_project(id, name, organisation, access_code)
		VALUES(?,?,?,?);
	`
	if _, err := db.Exec(query, data.Id, data.Name, data.Organisation, data.AccessToken); err != nil {
			return err
		} else {
			return nil
		}
}

func UpdateProject(db *sqlx.DB, data *data.ProjectData) error {
	query := `
		UPDATE tbl_project 
		SET name = ?, access_code = ?, updated_ts = CURRENT_TIMESTAMP
		WHERE id = ?
	`
	if _, err := db.Exec(query, data.Name, data.AccessToken, data.Id); err != nil {
		return err
	} else {
		return nil
	}
}