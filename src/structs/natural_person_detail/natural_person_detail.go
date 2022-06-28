package natural_person_detail

import (
	"github.com/google/uuid"
	"go-api-learning/src/structs/utils"
)

type NaturalPersonDetail struct {
	utils.WeakModel
	PersonID      uuid.UUID `json:"person_id"`
	PreferredName string    `json:"trading_name"`
	BirthDate     string    `json:"birth_date"`
	MothersName   string    `json:"mothers_name"`
}
