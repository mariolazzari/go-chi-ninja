package main

import (
	"context"
	"fmt"

	"github.com/mariolazzari/go-chi-ninja/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
