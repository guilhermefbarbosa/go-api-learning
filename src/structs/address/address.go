package address

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"go-api-learning/src/structs/utils"
)

type Address struct {
	utils.Model
	PersonID       uuid.UUID    `json:"person_id"`
	SchoolDomainID uuid.UUID    `json:"school_domain_id"`
	Street         string       `json:"street"`
	StreetNumber   int          `json:"street_number"`
	Neighborhood   string       `json:"neighborhood"`
	LocalityID     uuid.UUID    `json:"locality_id"`
	ZipCode        string       `json:"zip_code"`
	CustomMetadata pgtype.JSONB `json:"custom_metadata"`
}

func (this Address) GetAddress() string {
	return fmt.Sprintf(
		"%s, %d, Bairro %s, %s",
		this.Street,
		this.StreetNumber,
		this.Neighborhood,
		this.ZipCode)
}
