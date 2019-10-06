package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test/member/models"
	"test/member/usecase"
)

// MemberHandler is for router
type MemberHandler interface {
	HandleSignup() http.HandlerFunc
	HandleLogin() http.HandlerFunc
	HandleGetUser() http.HandlerFunc
}

func NewMemberHandler(mu usecase.MemberUsecase) MemberHandler {
	return &memberHandlerImpl{usecase: mu}
}

type memberHandlerImpl struct {
	usecase usecase.MemberUsecase
}

func handleResponse(w http.ResponseWriter, result interface{}) {
	b, err := json.Marshal(result)
	if err != nil {
		w.Write([]byte(`{"error":"internal server error","code":500}`))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (mh *memberHandlerImpl) HandleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		signupReq := models.SignupRequest{}
		json.Unmarshal(b, &signupReq)
		fmt.Println("----------")
		fmt.Println(signupReq)
		// validation
		handleResponse(w, mh.usecase.Signup(signupReq))
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
