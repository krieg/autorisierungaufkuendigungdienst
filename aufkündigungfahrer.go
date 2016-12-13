package main

import (
	"flag"
	"github.austin.utexas.edu/kriegrj/autorisierungaufkuendigungdienst/authentifizierung"
	"github.austin.utexas.edu/kriegrj/autorisierungaufkuendigungdienst/benutzerkonto"
	"github.austin.utexas.edu/kriegrj/go-github/github"
	"log"
	"os"
	"os/user"
)

func main() {

	// Flag represents command-line arguments
	type Flag struct {
		uPtr *string
		oPtr *string
		err  error
	}
	usr, e := user.Current()
	if e != nil {
		log.Println("Error retrieving current user: %v", e)
	}
	token := os.Getenv("GHE_TOKEN")
	if token == "" {
		log.Fatal("No OAuth token !\n\n")
	} else {
		log.Println("Successfully loaded OAuth token for", usr.Username)
	}

	client := github.NewClient(authentifizierung.Authn(token))
	f := new(Flag)
	f.uPtr = flag.String("u", "", "user account")
	f.oPtr = flag.String("o", "", "operation")
	flag.Parse()
	if *f.oPtr == "suspend" {
		_, f.err = benutzerkonto.Sperren(f.uPtr, client)
	} else if *f.oPtr == "unsuspend" {
		_, f.err = benutzerkonto.Freigeben(f.uPtr, client)
	}
	if f.err != nil {
		log.Printf("%v failed due to %v", *f.oPtr, f.err)
	} else {
		log.Printf("%v successful for %v", *f.oPtr, *f.uPtr)
	}
}
