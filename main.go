package main

import (
	"github.com/guths/bookxd/cmd"
	"github.com/guths/bookxd/db"
)

func main() {
	db.NewPostgresStore()

	cmd.Execute()
}
