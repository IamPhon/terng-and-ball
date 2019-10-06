package handler

import "net/http"

// MemberHandler is for router
type MemberHandler interface {
	HandleSignup() http.HandlerFunc
	HandleLogin() http.HandlerFunc
	HandleGetUser() http.HandlerFunc
}

func NewMemberHandler() MemberHandler {
	return new(memberHandlerImpl)
}

type memberHandlerImpl struct{}

func (m *memberHandlerImpl) HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func (m *memberHandlerImpl) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (m *memberHandlerImpl) HandleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
