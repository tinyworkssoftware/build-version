package request

type CreateOrg struct {
	Name     string
	PlanType string `json:"plan_type"`
}