package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GithubEvent struct {
	Type    string   `json:"type"`
	Repo    *Repo    `json:"repo"`
	Payload *Payload `json:"payload"`
}

type Repo struct {
	Name string `json:"name"`
}
type Payload struct {
	Size int `json:"size"`
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Введите имя пользователя например: MrGulo")
	}
	username := os.Args[1]

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			fmt.Println("Пользователь не найден")
		} else {
			fmt.Println("Ошибка API")
		}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения body")
	}

	var events []GithubEvent
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON")
		return
	}

	for _, event := range events {
		typeEvent := event.Type
		numberCommits := event.Payload.Size
		repoName := event.Repo.Name

		switch typeEvent {
		case "PushEvent":
			fmt.Printf("Pushed %d commits to %s/developer-roadmap\n", numberCommits, repoName)
		case "CreateEvent":
			fmt.Printf("Created branch/repository in %s\n", repoName)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", repoName)
		default:
			fmt.Printf("Event type %s in %s\n", typeEvent, repoName)

		}
	}
}
