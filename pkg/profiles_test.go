package memory

import (
	"testing"
)

func Test_signUp(t *testing.T) {

	m := initModel()

	m1 := UserSignedUp{
		UserID: "abc",
	}

	m, err := signUp(m, m1)
	if err != nil {
		t.Errorf("error = %v", err)
	}
}

func Test_missingUserID(t *testing.T) {

	m := initModel()

	m1 := UserSignedUp{}

	m, err := signUp(m, m1)
	if err != nil {
		return
	}
	t.Errorf("sign up should error in case of missing user id")
}
