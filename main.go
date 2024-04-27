package main

func main() {
	server := NewAPIServer(":443", "./certs/cert.pem", "./certs/key.pem")
	server.Run()
}
