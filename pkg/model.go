package memory

import "github.com/google/uuid"

type memory struct {
	profiles             map[uuid.UUID]profile
	profileVerifications map[string]uuid.UUID
	emailChanges         map[string]uuid.UUID
}

type profile struct {
	uuid                      uuid.UUID
	nick, email, passwordHash string
}
