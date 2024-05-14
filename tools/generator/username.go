package generator

import (
	ung "github.com/sibiraj-s/unique-names-generator"
)

func NewUsername() string {
	return ung.New(ung.Options{
		Dictionaries: [][]string{
			ung.Adjectives,
			ung.Colors,
			ung.Animals,
		},
	})
}
