package service

import (
	"build-version/model/data"
	"build-version/model/dto"
	"build-version/model/request"
	"build-version/service/repository"
	_ "encoding/json"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

func StartSession(session *request.CreateSession) (*data.SessionData, error) {
	startTime := time.Now()
	if db, err := connectDb(); err != nil {
		return nil, err
	} else {
		if proj, err := repository.GetProjectByAccessToken(db, session.AccessToken); err != nil {
			log.Debugln("project: %v", err)
			return nil, err
		} else {
			sessionHistory := &data.SessionHistoryData{
				Id: uuid.New().String(),
				StartTs: startTime,
				AssociatedBranch:  session.Branch,
				Project:           proj.Id,
			}
			activeRecord := &data.ActiveSessionData{
				Id:                uuid.New().String(),
				SessionId:           sessionHistory.Id,
			}
			if err := repository.CreateSessionHistory(db, sessionHistory); err != nil {
				log.Debugln("session history: %v", err)
				return nil, errors.New("failed to create session history record")
			} else {
				if res, err := repository.CreateActiveSession(db, activeRecord); err != nil {
					log.Debugln("active session: %v", err)
					return nil, err
				} else {
					return res, nil
				}
			}
		}

	}
}

func EndSession(req *dto.UpdateSessionDTO) error {
	//TODO: Update Session History
	if db, err := connectDb(); err != nil {
		return err
	} else {
		if _, err := repository.GetProjectByAccessToken(db, req.AccessToken); err != nil {
			return errors.New("invalid access token")
		} else {
			record := &data.SessionHistoryData{
				Id:                req.SessionId,
				EndTs:             time.Now(),
				AssociatedVersion: "",
				AssociatedBranch:  "",
			}
			if err = repository.UpdateSessionHistory(db, record); err != nil {
				return err
			}  else {
				return repository.DeleteActiveSession(db, req.SessionId)
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

func CheckValidToken(accessToken string) bool {
	log.Debugf("Checking token...")
	if db, err := connectDb(); err  != nil {
		log.Debugln("Token Check Error", err)
		return false
	} else {
		if err = repository.QuickTokenCheck(db, accessToken); err != nil {
			return false
		}
		return true
	}
}

