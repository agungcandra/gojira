package config

import "github.com/agungcandra/gojira/entity"

func NewCredential(username, password string) *entity.Credential {
	return &entity.Credential{
		Username: username,
		Password: password,
	}
}
