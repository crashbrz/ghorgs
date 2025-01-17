package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Organization struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
	URL   string `json:"url"`
}

func listGitHubOrgs(token string, baseURL string, minimal bool, outputFile string, limit int) {
	url := fmt.Sprintf("%s/organizations", baseURL)
	client := &http.Client{}
	var output []string
	count := 0

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating the request:", err)
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Set("Accept", "application/vnd.github.v3+json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making the request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
			return
		}

		var orgs []Organization
		if err := json.NewDecoder(resp.Body).Decode(&orgs); err != nil {
			fmt.Println("Error decoding the response:", err)
			return
		}

		for _, org := range orgs {
			if limit > 0 && count >= limit {
				break
			}

			if minimal {
				output = append(output, org.Login)
				fmt.Println(org.Login)
			} else {
				info := fmt.Sprintf("- %s (ID: %d, URL: %s)", org.Login, org.ID, org.URL)
				output = append(output, info)
				fmt.Println(info)
			}
			count++
		}

		if limit > 0 && count >= limit {
			break
		}

		links := resp.Header.Get("Link")
		if links == "" {
			break
		}

		nextURL := parseNextURL(links)
		if nextURL == "" {
			break
		}

		url = nextURL
	}

	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer file.Close()

		for _, line := range output {
			file.WriteString(line + "\n")
		}
		fmt.Printf("Results saved to %s\n", outputFile)
	}
}

func parseNextURL(linkHeader string) string {
	links := strings.Split(linkHeader, ",")
	for _, link := range links {
		parts := strings.Split(link, ";")
		if len(parts) == 2 && strings.TrimSpace(parts[1]) == "rel=\"next\"" {
			return strings.Trim(parts[0], "<> ")
		}
	}
	return ""
}

func main() {
	token := flag.String("t", "", "GitHub Access Token")
	baseURL := flag.String("u", "https://api.github.com", "Base URL of the GitHub API")
	listOrgs := flag.Bool("O", false, "List GitHub organizations")
	minimal := flag.Bool("m", false, "Output minimal information (login only)")
	outputFile := flag.String("o", "", "File to save the output")
	limit := flag.Int("l", 0, "Limit the number of results fetched. Omitting -l, will fetch all orgs")
	flag.Parse()

	if *token == "" {
		fmt.Println("Usage: -t <GitHub Access Token> -u <GitHub API Base URL> [-O for organizations] [-m for minimal output] [-o for output file] [-l for limit]")
		os.Exit(1)
	}

	if !*listOrgs {
		fmt.Println("Please specify -O flag to list organizations.")
		return
	}

	fmt.Println("==== Listing Organizations ====")
	listGitHubOrgs(*token, *baseURL, *minimal, *outputFile, *limit)
}
