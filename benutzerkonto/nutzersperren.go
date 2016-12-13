package benutzerkonto

import (
	"github.austin.utexas.edu/kriegrj/go-github/github"
	"log"
)

// Sperren suspends a user.
func Sperren(user *string, client *github.Client) (*github.Response, error) {
	u, _, err := client.Users.Get(*user)
	if err != nil {
		return nil, err
	}
	log.Printf("Revoke %v's account %v", *u.Name, *u.Login)

	resp, err := client.Users.Suspend(*u.Login)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Freigeben unsuspends a user.
func Freigeben(user *string, client *github.Client) (*github.Response, error) {
	u, _, err := client.Users.Get(*user)
	if err != nil {
		return nil, err
	}
	log.Printf("Unsuspend %v's account %v", *u.Name, *u.Login)

	resp, err := client.Users.Unsuspend(*u.Login)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
