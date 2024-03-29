package main

import (
	infrastructure "chat/infrastucture"
	"testing"
)

func BenchmarkCreateMillionMessagesAsync(b *testing.B) {
	db := infrastructure.NewDatabase()
    createMillionMessagesAsync(db)
}

func BenchmarkCreateMillionMessages(b *testing.B) {
	db := infrastructure.NewDatabase()
    createMillionMessages(db)
}
