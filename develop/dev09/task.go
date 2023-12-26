package main

import (
	"os"
	"wget/wget"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком
*/

func main() {
	os.Exit(wget.CLI(os.Args[1:]))
}
