package service

import (
	"build-version/model/data"
	"build-version/model/request"
	"build-version/service/repository"
	_ "encoding/json"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

func StartSession(session *request.CreateSession) (*data.SessionHistoryData, error) {
	startTime := time.Now()
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if proj, err := repository.GetProjectByAccessToken(db, session.AccessToken); err != nil {
			return nil, err
		} else {
			if proj == nil {
				return nil, errors.New("Invalid AccessToken")
			} else {
				activeRecord := &data.ActiveSessionData{
					Id:                uuid.New().String(),
					StartTs:           startTime,
					AssociatedBranch:  session.Branch,
					Session:           uuid.New().String(),
					Project:           proj.Id,
				}
				sessionHistory := &data.SessionHistoryData{
					Id: activeRecord.Id,
					StartTs: startTime,
					AssociatedBranch:  session.Branch,
					Session:           activeRecord.Session,
					Project:           activeRecord.Project,

				}
				go repository.CreateActiveSession(db, activeRecord)
				if res, err := repository.CreateSessionHistory(db, sessionHistory); err != nil {
					//TODO: rollback
					return nil, err
				} else {
					return res, nil
				}
			}
		}
	}
}

func GetActiveSessions() (*[]data.ActiveSessionData, error) {
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if record, err := repository.GetActiveSessions(db); err != nil {
			return nil, err
		} else {
			log.Info(record)
			return record, nil
		}
	}
}

func EndSession() {

}

