package data

import "time"

type PlanData struct {
	Id string
	Name string
	RequestLimit int `db:"request_limit"`
	ConcurrentSession int `db:"concurrent_session"`
	Price float32
	Currency string
}

type OrganisationData struct {
	Id string
	Name string
	CreatedTs time.Time `db:"created_ts"`
	UpdatedTs time.Time `db:"updated_ts"`
	PlanType string `db:"plan_type"`
}

type ProjectData struct {
	Id string `db:"id"`
	Name string `db:"name"`
	AccessToken string `db:"access_code"`
	ExceededLimit bool `db:"exceeded_limit"`
	CreatedTs time.Time `db:"created_ts"`
	UpdatedTs time.Time `db:"updated_ts"`
	Organisation string `db:"organisation"`
}

type SessionHistoryData struct {
	Id string
	StartTs time.Time `db:"start_ts"`
	EndTs time.Time `db:"end_ts"`
	AssociatedVersion string `db:"associated_version"`
	AssociatedBranch string `db:"associated_branch"`
	Project string
}

type ActiveSessionData struct {
	Id string `db:"id"`
	SessionId string `db:"session"`
}

//NOTE: Join model
type SessionData struct {
	Id string
	ActiveSessionId string `db:"active_session_id"`
	StartTs time.Time `db:"start_ts"`
	EndTs time.Time `db:"end_ts"`
	AssociatedVersion string `db:"associated_version"`
	AssociatedBranch string `db:"associated_branch"`
	Project string
}






