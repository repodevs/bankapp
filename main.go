package main

import (
	"github.com/repodevs/bankapp/api"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	// migrations.Migrate()
	api.StartAPI()
}
