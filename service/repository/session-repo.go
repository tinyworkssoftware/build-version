package repository

import (
	"build-version/model/data"
	"github.com/jmoiron/sqlx"
)

func CreateSessionHistory(db *sqlx.DB, request *data.SessionHistoryData) (*data.SessionHistoryData, error) {
	query := `
		INSERT INTO tbl_session_history(id, associated_version, associated_branch, project, session)
		VALUES (:id, :associated_version, :associated_branch, :project , :session)
	`
	if _, err := db.NamedExec(query, request); err != nil {
		return nil, err
	}
	return GetSessionHistoryById(db, request.Id)
}

func UpdateSessionHistory(db *sqlx.DB) {

}

func CreateActiveSession(db *sqlx.DB, request *data.ActiveSessionData) (*data.ActiveSessionData, error) {
	query := `
		INSERT INTO tbl_active_session(id, associated_version, associated_branch, project, session)
		VALUES (:id, :associated_version, :associated_branch, :project , :session)
	`
	if _, err := db.NamedExec(query, request); err != nil {
		return nil, err
	}
	return GetActiveSessionById(db, request.Id)
}

func UpdateActiveSession(db *sqlx.DB) {

}

func GetSessionHistoryById(db *sqlx.DB, id string) (*data.SessionHistoryData, error) {
	var record data.SessionHistoryData
	if err := db.QueryRowx("SELECT * FROM tbl_session_history WHERE id = ?", id).
		StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetActiveSessionById(db *sqlx.DB, id string) (*data.ActiveSessionData, error) {
	var record data.ActiveSessionData
	if err := db.QueryRowx("SELECT * FROM tbl_active_session WHERE id = ?", id).
		StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}

func GetActiveSessions(db *sqlx.DB) (*[]data.ActiveSessionData, error) {
	var record = make([]data.ActiveSessionData, 0)
	if err := db.Select(&record, "SELECT * FROM tbl_active_session"); err != nil {
		return nil, err
	} else {
		return &record, nil
	}
}