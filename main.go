package main

func main() {
	var server = NewAPIServer(":8000")
	server.Run()
}
