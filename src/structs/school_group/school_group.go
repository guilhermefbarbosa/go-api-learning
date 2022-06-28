package school_group

import (
	"go-api-learning/src/structs/utils"
)

type Person struct {
	utils.Model
	Name string `json:"name"`
}
