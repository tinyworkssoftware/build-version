package repository

import (
	"build-version/model/data"
	"github.com/jmoiron/sqlx"
)

func getProjectById(db *sqlx.DB, projectId string) (*data.ProjectData, error){
	ret := data.ProjectData{}
	row := db.QueryRow("SELECT * FROM tbl_project WHERE id = ?", projectId)
	if err := row.Scan(&ret); err != nil {
		return nil, err
	} else {
		return &ret, nil
	}
}

func getProjectByName(db *sqlx.DB, name string) (*data.ProjectData, error) {
	ret := data.ProjectData{}
	row := db.QueryRow("SELECT * FROM tbl_project WHERE name = ?", name)
	if err := row.Scan(&ret); err != nil {
			return nil, err
		} else {
			return &ret, nil
		}
}

func GetAllProjects(db *sqlx.DB) (*[]data.ProjectData, error) {
	query := `SELECT * FROM tbl_project;`
	var records []data.ProjectData
	if err := db.Select(&records, query); err != nil {
		return nil, err
	} else {
		return &records, nil
	}
}

func createProject(db *sqlx.DB, data *data.ProjectData) error {
	query := `
		INSERT INTO tbl_project(id, name, organisation)
		VALUES(?,?,?);
	`
	if _, err := db.Exec(query, data.Id, data.Name, data.Organisation); err != nil {
			return err
		} else {
			return nil
		}
}