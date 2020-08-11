package repository

import (
	"build-version/model/data"
	"errors"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func CreateSessionHistory(db *sqlx.DB, request *data.SessionHistoryData) error {
	query := `
		INSERT INTO tbl_session_history(id, associated_version, associated_branch, project)
		VALUES (:id, :associated_version, :associated_branch, :project)
	`
	if _, err := db.NamedExec(query, request); err != nil {
		return err
	}

	return nil
}

func UpdateSessionHistory(db *sqlx.DB, historyData *data.SessionHistoryData ) error {
	query := `
		UPDATE tbl_session_history
		SET associated_branch = :associated_branch, associated_version = :associated_version, end_ts = :end_ts
		WHERE id = :id;
	`
	if _, err := db.NamedExec(query, historyData); err != nil {
		return err
	} else {
		return err
	}
}

func DeleteActiveSession(db *sqlx.DB, activeSessionId string) error {
	var query = `
		DELETE FROM tbl_active_session 
		WHERE session = ?
	`
	if _, err := db.Exec(query, activeSessionId); err != nil {
		return err
	}

	return nil
}

func CreateActiveSession(db *sqlx.DB, request *data.ActiveSessionData) (*data.SessionData, error) {
	query := `
		INSERT INTO tbl_active_session(id, session)
		VALUES (:id, :session)
	`
	if _, err := db.NamedExec(query, request); err != nil {
		return nil, err
	}
	return GetSessionById(db, request.SessionId)
}

func UpdateActiveSession(db *sqlx.DB) {

}

func GetSessionById(db *sqlx.DB, id string) (*data.SessionData, error) {
	var query = `
		SELECT s.id AS 'active_session_id', h.* FROM tbl_active_session s
		INNER JOIN tbl_session_history h ON s.session = h.id
		WHERE h.id = ?;
	`
	var record data.SessionData
	if err := db.QueryRowx(query, id).
		StructScan(&record); err != nil {
		return nil, err
	} else {
		return &record, nil
	}

}

func QuickTokenCheck(db *sqlx.DB, accessToken string) error {
	var query = `
		SELECT id
		FROM tbl_project 
		WHERE access_code = ?
		LIMIT 1
	`
	if row, err := db.Queryx(query, accessToken); err != nil {
		log.Debugln(err)
		return err
	} else {
		if row.Next() {
			return nil
		} else {
			return errors.New("not found")
		}
	}

}

//func GetSessionHistoryById(db *sqlx.DB, id string) (*data.SessionHistoryData, error) {
//
//}

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