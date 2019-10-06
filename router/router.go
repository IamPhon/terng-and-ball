package router

import (
	"net/http"
	"test/member/delivery/handler"

	"github.com/gorilla/mux"
)

type Router interface {
	Routes() *mux.Router
}

type routerImpl struct {
	mh handler.MemberHandler
}

func NewRouter(mh handler.MemberHandler) Router {
	return &routerImpl{mh: mh}
}

func (r *routerImpl) Routes() *mux.Router {
	rt := mux.NewRouter()
	rt.HandleFunc("/signup", r.mh.HandleSignup()).Methods(http.MethodPost)
	rt.HandleFunc("/signin", r.mh.HandleLogin()).Methods(http.MethodPost)
	rt.HandleFunc("/member", r.mh.HandleGetUser()).Methods(http.MethodGet)
	rt.Use(logging)
	return rt
}
