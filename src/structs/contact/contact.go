package contact

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"go-api-learning/src/structs/utils"
)

type Contact struct {
	utils.Model
	PersonID        uuid.UUID    `json:"person_id"`
	SchoolDomainID  uuid.UUID    `json:"school_domain_id"`
	ContactType     string       `json:"contact_type"`
	ContactInfo     string       `json:"contact_info"`
	ContactPriority string       `json:"contact_priority"`
	CustomMetadata  pgtype.JSONB `json:"custom_metadata"`
}

func (this Contact) GetContact() string {
	return fmt.Sprintf(
		"%s",
		this.ContactInfo,
	)
}
