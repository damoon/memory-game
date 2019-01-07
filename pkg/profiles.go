package memory

import (
	"fmt"
	"time"
)

type userID = string
type passwordResetToken = string

type model struct {
	profiles       map[userID]profile
	passwordResets passwordResets
}

func initModel() model {
	return model{
		profiles:       initProfiles(),
		passwordResets: initPasswordResets(),
	}
}

type profile struct {
	id, nick, email, passwordHash string
}

func initProfiles() map[userID]profile {
	return map[userID]profile{}
}

type passwordResets struct {
	requested map[userID]bool
	tokens    map[passwordResetToken]passwordReset
}

type passwordReset struct {
	validUntil time.Time
	userID     userID
}

func initPasswordResets() passwordResets {
	return passwordResets{
		requested: map[userID]bool{},
		tokens:    map[passwordResetToken]passwordReset{},
	}
}

func signUp(m model, msg UserSignedUp) (model, error) {

	id := msg.GetUserID()
	_, found := m.profiles[id]
	if found {
		return m, fmt.Errorf("a profile with id %s already exists", id)
	}

	m.profiles[id] = profile{
		id:           msg.GetUserID(),
		nick:         msg.GetNick(),
		email:        msg.GetEmail(),
		passwordHash: msg.GetPasswordHash(),
	}

	return m, nil
}
