package request

type CreateOrg struct {
	Name     string
	PlanType string `json:"plan_type"`
}

type CreateProject struct {
	Name string `json:"name"`
	Organisation string `json:"organisation"`
}

type CreateSession struct {
	AccessToken string
	Branch string
	StartTs string
}