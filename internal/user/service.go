package user

import (
	"restapi/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
}

//func (s *Storage) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
//	// TODO for next one
//	return
//}
