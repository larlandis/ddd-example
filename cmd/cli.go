package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/larlandis/ddd-example/pkg/contact"
	"github.com/olekukonko/tablewriter"
)

//CLI defines command line interface
type CLI interface {
	Run()
}

type cli struct {
	contact contact.Service
}

func (cli *cli) Run() {
	reader := bufio.NewReader(os.Stdin)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColumnSeparator("|")
	table.SetRowLine(true)
	table.SetHeader([]string{"UserID", "Nickname", "Email", "Address"})
	for {
		fmt.Print("\nIngrese un ID de usuario: (vacio para terminar) ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// get user id from prompt
		userID := strings.TrimSuffix(cmdString, "\n")
		if userID == "" {
			return
		}
		// get contact info
		contactInfo, err := cli.contact.GetContact(userID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		address := "---"
		if len(contactInfo.Addresses) > 0 {
			address = contactInfo.Addresses[0]
		}
		table.Append([]string{
			contactInfo.UserID,
			contactInfo.UserName,
			contactInfo.Email,
			address,
		})
		table.Render()
	}
}

// NewCLI returns CLI implementation
func NewCLI(cServ contact.Service) CLI {
	return &cli{
		contact: cServ,
	}
}
