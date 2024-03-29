package main

import (
	infrastructure "chat/infrastucture"
	"testing"
)

func BenchmarkCreateMessage(b *testing.B) {
	db := infrastructure.NewDatabase()

	for i := 0; i < 100; i++ {
		createMessage(db)
	}
}
