package json

import (
	"encoding/json"

	"github.com/mmanjoura/nomad/backend"
	appUser "github.com/mmanjoura/nomad/backend/user"
)

type User struct{}
type UserUpdate struct{}

func (r *User) DecodeUser(input []byte) (*appUser.User, error) {
	u := &appUser.User{}
	if err := json.Unmarshal(input, u); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.User.DecodeUser")
	}
	return u, nil
}

func (r *User) EncodeUser(input *appUser.User) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.User.EncodeUser")
	}
	return rawMsg, nil
}

func (r *UserUpdate) DecodeUserUpdate(input []byte) (*appUser.UserUpdate, error) {
	upd := &appUser.UserUpdate{}
	if err := json.Unmarshal(input, upd); err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.UpdateUser.DecodeUserUpdate")
	}
	return upd, nil
}

func (r *UserUpdate) EncodeUserUpdate(input *appUser.UserUpdate) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, backend.Errorf(backend.EINTERNAL, "serializer.UpdateUser.EncodeUserUpdate")
	}
	return rawMsg, nil
}
