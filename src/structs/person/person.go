package person

import (
	"github.com/jackc/pgtype"
	"go-api-learning/src/structs/utils"
)

type Person struct {
	utils.Model
	TaxId              string       `json:"tax_id"`
	TaxIdType          string       `json:"tax_id_type"`
	PersonType         string       `json:"person_type"`
	Name               string       `json:"name"`
	NationalityIsoCode string       `json:"nationality_iso_code"`
	ExternalIds        pgtype.JSONB `json:"external_ids"`
	CustomMetadata     pgtype.JSONB `json:"custom_metadata"`
}
