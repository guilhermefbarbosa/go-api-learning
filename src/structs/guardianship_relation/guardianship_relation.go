package guardianship_relation

import (
	"github.com/google/uuid"
	"go-api-learning/src/structs/utils"
)

type GuardianshipRelation struct {
	utils.WeakModel
	GuardianSchoolDomain uuid.UUID `json:"guardian_school_domain"`
	GuardianshipType     string    `json:"guardianship_type"`
	StudentSchoolDomain  uuid.UUID `json:"student_school_domain"`
	GuardianshipDegree   string    `json:"guardianship_degree"`
}
