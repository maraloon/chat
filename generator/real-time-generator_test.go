package main

import (
	infrastructure "chat/infrastucture"
	"testing"
)

var db = infrastructure.NewDatabase()

// WARN: запускай внутри контейнера

func TestAsyncMessages(t *testing.T) {
	// create100Chats(db)
	create100Users(db)
	// createMessages(db)
}
