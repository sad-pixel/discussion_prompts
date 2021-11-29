package main

import (
	"encoding/json"
	"log"
	"os"

	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		ShowHelpMenu()
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sheetDbEndpoint := os.Getenv("SHEETDB_ENDPOINT")
	discordEndpoint := os.Getenv("DISCORD_ENDPOINT")
	discordRoleId := os.Getenv("DISCORD_ROLE_ID")

	log.Println("SheetDb: ", sheetDbEndpoint)
	log.Println("Webhook: ", discordEndpoint)
	log.Println("Role: ", discordRoleId)

	switch os.Args[1] {
	case "help":
		ShowHelpMenu()
	case "post":
		PostRandomPrompt(discordRoleId, discordEndpoint, sheetDbEndpoint)
	default:
		fmt.Println("incorrect subcommand")
		os.Exit(1)
	}

	// SendWebhook("1", "e", discordEndpoint)
}

func ShowHelpMenu() {
	fmt.Println("Usage:")
	fmt.Println("$ discussion-prompts [subcommand]")
	fmt.Println()
	fmt.Println("Subcommands:")
	fmt.Println("help - Display the help menu")
	fmt.Println("post - Post a discussion prompt")
}

func PostRandomPrompt(roleId string, discordEndpoint string, sheetDbEndpoint string) {
	prompt := FetchRandomPrompt(sheetDbEndpoint)
	SendWebhook(roleId, prompt.Prompt, discordEndpoint)
	MarkPromptPosted(sheetDbEndpoint, prompt.Id)
}

type WebHookPayload struct {
	Content string `json:"content"`
}

func SendWebhook(roleId string, prompt string, discordEndpoint string) {
	url := "https://discord.com/api/webhooks/914741202392326165/D4TQGAf7FdiLCzcwRMOA30e8Dn9zuavqZ_7KrTaBcQycRCS9JVexJofU3Coe0_nJBx6H"

	jsonPayload, _ := json.Marshal(&WebHookPayload{
		Content: "Hello everyone <@&" + roleId + ">! \n\n" +
			"Today's Discussion Question is: " + prompt + "\n\n" +
			"If you have any question suggestions feel free to DM a staff member or open a ticket. Happy discussion :)",
	})

	payload := strings.NewReader(string(jsonPayload))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println("Webhook sent with prompt: " + prompt)
	log.Println("Discord response: ", body)
}

type DiscussionPrompt struct {
	Id       string `json:"id"`
	Prompt   string `json:"prompt"`
	IsPosted string `json:"is_posted"`
}

func FetchRandomPrompt(sheetDbEndpoint string) DiscussionPrompt {
	url := sheetDbEndpoint + "search?sort_order=random&limit=1&is_posted=N&single_object=true"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prompt DiscussionPrompt
	json.Unmarshal(body, &prompt)

	log.Println("Fetched Prompt ID: "+prompt.Id, " - ", prompt.Prompt)
	return prompt
}

func MarkPromptPosted(sheetDbEndpoint string, promptId string) {
	url := sheetDbEndpoint + "id/" + promptId

	payload := strings.NewReader("{\n\t\"data\": [\n\t\t{\n\t\t\t\"is_posted\": \"Y\"\n\t\t}\n\t]\n}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println("Prompt " + promptId + " marked as posted.")
	log.Println("SheetDB response: ", body)
}
