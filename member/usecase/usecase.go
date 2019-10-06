package usecase

import "test/member/models"

type MemberUsecase interface {
	Signup(models.SignupRequest) bool
}

func NewMemberUsecase() MemberUsecase {
	return &memberUsecaseImpl{}
}

type memberUsecaseImpl struct{}

func (u *memberUsecaseImpl) Signup(r models.SignupRequest) bool {
	return true
}
