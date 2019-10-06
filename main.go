package main

import (
	"fmt"
	"net/http"
	"test/member/delivery/handler"
	"test/member/usecase"
	"test/router"
)

func main() {
	mu := usecase.NewMemberUsecase()
	mh := handler.NewMemberHandler(mu)
	r := router.NewRouter(mh)
	fmt.Println("Server starting at port 8080...")
	http.ListenAndServe(":8080", r.Routes())
}
