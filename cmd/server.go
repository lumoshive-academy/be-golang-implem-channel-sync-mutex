package cmd

import (
	"fmt"
	"log"
	"net/http"
	"session-24/internal/wire"
)

func APiserver(app *wire.App) {
	fmt.Println("Server running on port 8080")

	if err := http.ListenAndServe(":8080", app.Route); err != nil {
		log.Fatal("can't run service")
	}

}
