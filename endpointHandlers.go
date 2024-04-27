package main

import "net/http"

func userHandler(resWriter http.ResponseWriter, req *http.Request) {
	userId := req.PathValue("userId")
	token := req.PathValue("token")

	resWriter.Write([]byte("The user id is : " + userId + " token: " + token))
}
