package main

import (
	infrastructure "chat/infrastucture"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"gorm.io/gorm"
)

// var user = &infrastructure.User{}
// var chat = &infrastructure.Chat{}
// var message = &infrastructure.Message{}

// TODO: i don't know am I need it,
// but name is so cool, so I'll try to use it
type RealTimeGenerator struct{}

func main() {
	db := infrastructure.NewDatabase()

	// create100Chats(db)
	// create100Users(db)
	// connectChatsWithUsers() // TODO: later
	createMillionMessages(db)
	// createMillionMessagesAsync(db)
}

func create100Chats(db *gorm.DB) {
	var count int64
	db.Model(&infrastructure.Chat{}).Count(&count)

	if count == 0 {
		for i := 0; i < 100; i++ {
			db.Create(&infrastructure.Chat{Name: generateRandomChatName()})
		}
	}
}

func create100Users(db *gorm.DB) {
	var count int64
	db.Model(&infrastructure.User{}).Count(&count)

	if count == 0 {
		for _, name := range pokemonNames {
			db.Create(&infrastructure.User{Nickname: name})
		}
	}

}

func createMillionMessagesAsync(db *gorm.DB) {
	var wg sync.WaitGroup

	// Define a function to create a message
	createMessage := func() {
		defer wg.Done() // Notify WaitGroup that this goroutine is done

		// Create message in the database
		createMessage(db)
	}

	// Launch 10 goroutines
	for i := 0; i < 100000; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go createMessage()
	}

	// Wait for all goroutines to finish
	wg.Wait()

}
func createMillionMessages(db *gorm.DB) {
	// var count int64
	// db.Model(&infrastructure.Message{}).Count(&count)

	// if count == 0 {
		// TODO: вот тут поставить 1млн,
		// banchmark времени исполнения
		// сделать 10 потоков, снова посмотреть время
		for i := 0; i < 100000; i++ {
			createMessage(db)
		}
	// }
}

func createMessage(db *gorm.DB) {
	var msg = chatMessages[rand.Intn(len(chatMessages))]
	db.Create(&infrastructure.Message{
		// TODO: hardcoded ids
		ChatId: uint(rand.Intn(99) + 1),
		UserId: uint(rand.Intn(99) + 1),
		Text:   msg,
	})
}

func generateRandomChatName() string {
	rand.Seed(time.Now().UnixNano())
	adjective := adjectives[rand.Intn(len(adjectives))]
	animal := animals[rand.Intn(len(animals))]
	return fmt.Sprintf("%s %s", adjective, animal)
}

var adjectives = []string{
	"Cool", "Happy", "Brave", "Clever", "Friendly",
	"Playful", "Cheerful", "Lively", "Charming", "Smart",
	"Funny", "Kind", "Radiant", "Daring", "Gentle",
	"Joyful", "Silly", "Unique", "Witty", "Zesty",
}

var animals = []string{
	"Pandas", "Tigers", "Lions", "Elephants", "Dolphins",
	"Penguins", "Giraffes", "Koalas", "Otters", "Foxes",
	"Zebras", "Bears", "Owls", "Monkeys", "Snakes",
	"Horses", "Kangaroos", "Cheetahs", "Wolves", "Ducks",
}

var chatMessages = []string{
	"Hello!",
	"How are you?",
	"Nice weather today!",
	"What's up?",
	"How's your day going?",
	"Hey there!",
	"What are you up to?",
	"Did you see that?",
	"I'm here!",
	"Howdy!",
}

var pokemonNames = []string{
	"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon",
	"Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie",
	"Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill",
	"Pidgey", "Pidgeotto", "Pidgeot", "Rattata", "Raticate",
	"Spearow", "Fearow", "Ekans", "Arbok", "Pikachu",
	"Raichu", "Sandshrew", "Sandslash", "Nidoran♀", "Nidorina",
	"Nidoqueen", "Nidoran♂", "Nidorino", "Nidoking", "Clefairy",
	"Clefable", "Vulpix", "Ninetales", "Jigglypuff", "Wigglytuff",
	"Zubat", "Golbat", "Oddish", "Gloom", "Vileplume",
	"Paras", "Parasect", "Venonat", "Venomoth", "Diglett",
	"Dugtrio", "Meowth", "Persian", "Psyduck", "Golduck",
	"Mankey", "Primeape", "Growlithe", "Arcanine", "Poliwag",
	"Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam",
	"Machop", "Machoke", "Machamp", "Bellsprout", "Weepinbell",
	"Victreebel", "Tentacool", "Tentacruel", "Geodude", "Graveler",
	"Golem", "Ponyta", "Rapidash", "Slowpoke", "Slowbro",
	"Magnemite", "Magneton", "Farfetch'd", "Doduo", "Dodrio",
	"Seel", "Dewgong", "Grimer", "Muk", "Shellder",
	"Cloyster", "Gastly", "Haunter", "Gengar", "Onix",
	"Drowzee", "Hypno", "Krabby", "Kingler", "Voltorb",
}
