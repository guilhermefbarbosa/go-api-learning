package school_details

import (
	"github.com/google/uuid"
	"go-api-learning/src/structs/utils"
)

type NaturalPersonDetail struct {
	utils.WeakModel
	PersonID      uuid.UUID `json:"person_id"`
	SchoolGroupID uuid.UUID `json:"school_group_id"`
	ShortName     string    `json:"short_name"`
}
