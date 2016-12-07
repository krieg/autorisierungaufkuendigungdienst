package benutzerkonto

import (
	"fmt"
	"github.austin.utexas.edu/kriegrj/go-github/github"
)

// Sperren suspends a user.
func Sperren(user *string, client *github.Client) {
	fmt.Println("User account to revoke:", *user)
	u, _, err := client.Users.Get(*user)
	if err != nil {
		fmt.Println("Users.Get %v returned error: %v", *user, err)
	}
	fmt.Println(*u.Login, *u.Name)

	resp, err := client.Users.Suspend(*u.Login)
	if err != nil {
		fmt.Println("Users.Suspend returned error:", err)
	} else {
		fmt.Println("Response following user suspension:", *resp)
	}
}
