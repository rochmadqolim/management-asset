package main

import (
	"go_inven_ctrl/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Exec()
}
