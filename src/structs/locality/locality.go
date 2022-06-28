package locality

import (
	"fmt"
	"go-api-learning/src/structs/utils"
)

type Locality struct {
	utils.Model
	IbgeCode       string `json:"ibge_code"`
	City           string `json:"city"`
	State          string `json:"state"`
	CountryName    string `json:"country_name"`
	CountryIsoCode string `json:"country_iso_code"`
}

func (this Locality) GetLocality() string {
	return fmt.Sprintf(
		"%s - %s - %s",
		this.City,
		this.State,
		this.CountryName)
}
