package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// TODO: Create general request
func generalRequest(animeTitles []string) []map[string]int {

}

// TODO: Create request processing function
func processAddAnimeRequest(animeTitles []string) map[string]int {

	// This does not matter as title takes all in graphql
	// R̶e̶a̶d̶ ̶f̶o̶r̶ ̶J̶a̶p̶a̶n̶e̶s̶e̶ ̶a̶n̶d̶ ̶e̶n̶g̶l̶i̶s̶h̶ ̶t̶i̶t̶l̶e̶s̶

	/*

		# Summary so far:

		- Worked on processing the anime that come in as slice of maps
			1. [ ] Plan to work on editing each individual map in slice
			2. [ ] Return the map so it can be added to file
			3. [ ] Create the API request to site

	*/
	for {
		fmt.Print("Enter anime titles seperated by a comma (ex: <title>, <title>): ")
		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		userTitles := strings.Split(line, ",")

		// 1. Do API request and grab the titles and IDs
		animeMapFromAPI := generalRequest(userTitles)
		// 2. Confirm data to user
		for title, id := range animeMapFromAPI {
			fmt.Println("%s : %s", title, id)

		}
		//TODO: Allow to redo a single anime title
		fmt.Println("Are these the anime you are looking to add? (y/n): ")
		//Grab input from user
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "y":
			break
		case "n":
			continue
		default:
			fmt.Println("Not a valid choice")
		}

	}

	// 3. Return the map with title and ID

}

func readAnimeFile() map[string]int {

	for {
		file, err := os.Open("cmd/anime.txt")
		if err != nil {
			//TODO: log these errors in a file
			//TODO: After 5 consecutive fails email me
			fmt.Println("Error opening file:", err)
			fmt.Println("Retrying in 1 minute...")
			time.Sleep(time.Minute)
			continue
		}

		fileInfo, err := file.Stat()
		if err != nil {
			//TODO: log these errors in a file
			fmt.Println("Error getting file information:", err)
			fmt.Println("Retrying in 5 seconds")
			time.Sleep(time.Second * 5)
			file.Close()
			continue
		}

		if fileInfo.Size() > 0 {
			fmt.Println("File has anime :)")
		} else {
			fmt.Println("File does not have anime :(")
		}

		file.Close()
		break
	}

	animeMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			title := parts[0]
			id := parts[1]

			var idInt int

			_, err := fmt.Sscanf(id, "%d", &idInt)
			if err != nil {
				//TODO: log these errors in a file
				fmt.Printf("Error converting id '%s' to integer: %v\n", id, err)
				fmt.Printf("Skipping %s:%s", title, id)
				continue
			}
			animeMap[title] = idInt
		} else {
			//TODO: log these errors in a file
			fmt.Printf("Invalid line format: '%s'\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		//TODO: log these errors in a file
		fmt.Println("Error scanning file:", err)
		return nil
	}

	return animeMap

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

	var animeTitles []string

	for titles := range animeMap {
		animeTitles = append(animeTitles, titles)
	}

	processAddAnimeRequest(animeTitles)
	//TODO: add logic to add to a file
	fmt.Println("Anime added to file!")
}

func viewAnime(animeMap map[string]int) {
	fmt.Println("\nHere are your current anime:")
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

// TODO: add logic to enter into application. Like a login
func main() {
	fmt.Println("Welcome to the Anime Update Email List!")
	fmt.Println("Get updated when your favorite anime is being released within an hour.")

	var animeMap map[string]int

	for {
		animeMap := readAnimeFile()
		if animeMap == nil {
			fmt.Println("Retrying in a minute...")
			time.Sleep(time.Minute)
			continue
		}
		break
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		displayMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addAnime(animeMap)
		case "2":
			viewAnime(animeMap)
		case "3":
			deleteAnime(animeMap)
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
