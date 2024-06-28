package main

import (
	infrastructure "chat/infrastucture"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"gorm.io/gorm"
)

// TODO: i don't know am I need it,
// but name is so cool, so I'll try to use it
type RealTimeGenerator struct{}

const (
	goroutinesNum = 64
	chatsNum      = 100
	usersNum      = 100
	messagesNum   = 100000
)

func main() {
	db := infrastructure.NewDatabase()

	createChats(db)
	createUsers(db)
	// connectChatsWithUsers() // TODO: later
	createMessages(db)
}

func createChats(db *gorm.DB) {
	// var count int64
	// db.Model(&infrastructure.Chat{}).Count(&count)

	workers(chatsNum, func() {
		db.Create(&infrastructure.Chat{Name: generateRandomChatName()})
	})
}

func createUsers(db *gorm.DB) {
	// var count int64
	// db.Model(&infrastructure.User{}).Count(&count)

	var i = 0
	var pokemonName = func() string {
		name := pokemonNames[i]
		i++
		return name
	}

	workers(usersNum, func() {
		db.Create(&infrastructure.User{Nickname: pokemonName()})
	})
}

func createMessages(db *gorm.DB) {
	workers(messagesNum, func() { createMessage(db) })
}

func workers(jobsNum int, closure func()) {
	var wg sync.WaitGroup
	jobs := make(chan int, jobsNum)

	for w := 0; w < goroutinesNum; w++ {
		go func() {
			for range jobs {
				closure()
				wg.Done()
			}
		}()
	}

	wg.Add(jobsNum)
	for j := 1; j <= jobsNum; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
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
