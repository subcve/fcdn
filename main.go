package fcdn

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/projectdiscovery/cdncheck"
)

func main() {
	// uses projectdiscovery endpoint with cached data to avoid ip ban
	// Use cdncheck.New() if you want to scrape each endpoint (don't do it too often or your ip can be blocked)
	scanner := bufio.NewScanner(os.Stdin)
	client, err := cdncheck.NewWithCache()
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		if scanner.Text() != "" {
			if found, result, err := client.Check(net.ParseIP(scanner.Text())); err == nil && !found {
				fmt.Println(scanner.Text(), result)
			}
		}
	}
}
