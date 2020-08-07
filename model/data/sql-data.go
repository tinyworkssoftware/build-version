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
	Id string
	Name string
	CreatedTs time.Time `db:"created_ts"`
	UpdatedTs time.Time `db:"updated_ts"`
	Organisation string `db:"organisation"`
}





