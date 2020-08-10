package response

import "time"

type Error struct {
	ErrorMessage string `json:"error_message"`
	CorrelationId string `json:"correlation_id"`
	TransactionTs time.Time `json:"transaction_ts"`
}

type OrganisationComposite struct {
	OrganisationId string
	OrganisationName string
}
