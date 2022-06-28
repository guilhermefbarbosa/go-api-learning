package juridical_person_detail

import (
	"github.com/google/uuid"
	"go-api-learning/src/structs/utils"
)

type JuridicalPersonDetail struct {
	utils.WeakModel
	PersonID     uuid.UUID `json:"person_id"`
	TradingName  string    `json:"trading_name"`
	BusinessType string    `json:"business_type"`
}
