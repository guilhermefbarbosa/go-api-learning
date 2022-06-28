package school_domain

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"go-api-learning/src/structs/utils"
)

type SchoolDomain struct {
	utils.Model
	PersonID       uuid.UUID    `json:"person_id"`
	SchoolID       uuid.UUID    `json:"school_id"`
	Status         string       `json:"status"`
	RolesArray     []string     `json:"roles_array"`
	CustomMetadata pgtype.JSONB `json:"custom_metadata"`
}
