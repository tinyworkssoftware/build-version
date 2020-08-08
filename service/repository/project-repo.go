package repository

import (
	"build-version/config"
	"build-version/model/data"
)

func getProjectById(projectId string) (*data.ProjectData, error){
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		defer db.Close()
		ret := data.ProjectData{}
		row := db.QueryRow("SELECT * FROM tbl_project WHERE id = ?", projectId)
		if err = row.Scan(&ret); err != nil {
			return nil, err
		} else {
			return &ret, nil
		}
	}
}

func getProjectByName(name string) (*data.ProjectData, error) {
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		defer db.Close()
		ret := data.ProjectData{}
		row := db.QueryRow("SELECT * FROM tbl_project WHERE name = ?", name)
		if err = row.Scan(&ret); err != nil {
			return nil, err
		} else {
			return &ret, nil
		}
	}
}

func GetAllProjects() (*[]data.ProjectData, error) {
	query := `SELECT * FROM tbl_project;`
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return nil, err
	} else {
		defer db.Close()
		var records []data.ProjectData
		if err = db.Select(&records, query); err != nil {
			return nil, err
		} else {
			return &records, nil
		}
	}
}

func createProject(data *data.ProjectData) error {
	query := `
		INSERT INTO tbl_project(id, name, organisation)
		VALUES(?,?,?);
	`
	conf = config.GetAppConfig()
	if db, err := ConnectDb(conf); err != nil {
		return err
	} else {
		defer db.Close()
		if _, err := db.Exec(query, data.Id, data.Name, data.Organisation); err != nil {
			return err
		} else {
			return nil
		}
	}
}