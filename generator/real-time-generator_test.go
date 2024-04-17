package main

import (
	infrastructure "chat/infrastucture"
	"testing"
)

var db = infrastructure.NewDatabase()

// WARN: запускай внутри контейнера

func TestAsyncMessages(t *testing.T) {
    createMillionMessagesAsync(db)
}

// func TestMessages(t *testing.T) {
//     createMillionMessages(db)
// }

// func BenchmarkCreateMillionMessagesAsync(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		createMillionMessagesAsync(db)
// 	}
// }

// func BenchmarkCreateMillionMessages(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		createMillionMessages(db)
// 	}
// }
