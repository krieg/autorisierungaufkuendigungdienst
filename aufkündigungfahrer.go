package main

import (
	"flag"
	"fmt"
	"github.austin.utexas.edu/kriegrj/autorisierungaufkuendigungdienst/authentifizierung"
	"github.austin.utexas.edu/kriegrj/autorisierungaufkuendigungdienst/benutzerkonto"
	"github.austin.utexas.edu/kriegrj/go-github/github"
	"log"
	"os"
	"os/user"
)

func main() {

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error retrieving current user: %v", err)
	}
	token := os.Getenv("GHE_TOKEN")
	if token == "" {
		log.Fatal("No OAuth token !\n\n")
	} else {
		fmt.Println("Successfully loaded OAuth token for", usr.Username)
	}

	client := github.NewClient(authentifizierung.Authn(token))
	uPtr := flag.String("u", "", "user account to revoke")
	flag.Parse()
	resp, err := benutzerkonto.Sperren(uPtr, client)

}
