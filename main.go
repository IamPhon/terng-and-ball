package main

import (
	"fmt"
	"net/http"
	"test/member/handler"
	"test/router"
)

func main() {
	mh := handler.NewMemberHandler()
	r := router.NewRouter(mh)
	fmt.Println("Server starting at port 8080...")
	http.ListenAndServe(":8080", r.Routes())
}
