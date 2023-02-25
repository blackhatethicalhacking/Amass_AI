// made with <3 by Chris 'SaintDruG' Abou-Chabke - PoC Concept for fetching subdomains, asn and more using ChatGPT - project purpose - try to integrate it with Amass
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get user's name
	fmt.Println("Hi! What's your name?")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hi %s! How can I help you today with your recon?\n", name)

	// Set the API Key
	apiKey := "<API_CHAT_GPT_HERE"

	// Set the OpenAI model to use
	model := "text-davinci-003"

	// Set the max number of tokens to generate
	maxTokens := 4000

	// Set the temperature for the model
	temperature := 1.0

	// Main loop to continuously ask for input and generate responses
	for {
		fmt.Println()
		fmt.Printf("%s: ", name)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "quit" {
			fmt.Println("Exiting... Have a 1337 day! from Black Hat Ethical Hacking")
			break
		}

		// Generate response from OpenAI API
		response := generateResponse(input, apiKey, model, maxTokens, temperature)

		// Print the response
		fmt.Println()
		fmt.Printf("ChatGPT: %s\n", strings.TrimSpace(response))
	}
}

// Function to generate a response from OpenAI API without logic
func generateResponse(input string, apiKey string, model string, maxTokens int, temperature float64) string {
	url := "https://api.openai.com/v1/completions"
	payload := strings.NewReader(fmt.Sprintf("{\"model\":\"%s\",\"prompt\":\"%s\",\"max_tokens\":%d,\"temperature\":%f}", model, input, maxTokens, temperature))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(res.Body).Decode(&data)

	return data["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
}
