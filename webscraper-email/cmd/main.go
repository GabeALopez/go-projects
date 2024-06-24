package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: add logic to enter into application. Like a login
func main() {
	fmt.Println("Welcome to the Anime Update Email List!")
	fmt.Println("Get updated when your favorite anime is being released within an hour.")

	anime := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addAnime(anime)
		case "2":
			viewAnime(anime)
		case "3":
			deleteAnime(anime)
		case "4":
			fmt.Println("Exiting the application.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

	}

	//TODO: add logic under here to loop
	/*anime := map[string]int{
		"Demon Slayer":          166240,
		"Jobless Reincarnation": 166873,
		"Spice and Wolf":        145728,
	}

	url := "http://google.com"

	data, err := scrapper.ScrapeWebsite(url)

	if err != nil {
		log.Fatalf("Error scraping website: %w", err)
	}

	fmt.Println(data)
	*/
}

func displayMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Add Anime")
	fmt.Println("2. View Your Current Anime")
	fmt.Println("3. Delete Anime")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
}

func addAnime(animeMap map[string]int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter anime title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	//TODO: Add logic to input title and do call to get id
	animeMap[title] = id
	//TODO: add logic to add to a file
	fmt.Println("Anime added to file!")
}

func viewAnime(animeMap map[string]int) {
	fmt.Println("\nYour Tasks:")
	for title := range animeMap {
		fmt.Printf("Title: %s\n", title)
	}
}

func deleteAnime(animeMap map[string]int) {
	fmt.Print("Enter task number to delete: ")
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')

	//TODO: add deleting anime from file logic
	if _, exists := animeMap[title]; exists {
		delete(animeMap, title)
		fmt.Println("Anime was deleted!")
		return
	}

	fmt.Println("Anime was not deleted :(")
	return

}
