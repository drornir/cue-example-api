package api

import "strings"

#User: {
	ID: =~#"^[A-Za-z]{4,}$"
	Name: strings.MinRunes(2)
}
