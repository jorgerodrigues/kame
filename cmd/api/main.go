package main

import (
	"fmt"
	"os"

	"github.com/jorgerodrigues/upkame/internal/server"
)

func main() {

	server := server.NewServer()
	fmt.Printf("Server running on port %v 🚀 \n", os.Getenv("PORT"))
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
