package pullreqsum

import (
	"fmt"
	"strings"
)

type ConsolePrinter struct {
	MessageSender     string
	MessageRecipients []string
	MessageSubject    string
}

func (cp *ConsolePrinter) Update(s Subject) {
	fmt.Printf("Email Sender: %s\n", cp.MessageSender)
	fmt.Printf("Email Recipients: %v\n", strings.Join(cp.MessageRecipients, ", "))
	fmt.Printf("Email Subject: %s\n", cp.MessageSubject)
	fmt.Println("-----------------")
	for _, pr := range s.GetPullRequests() {
		fmt.Printf("%s: %s\n#%d created on %v\n\n", *pr.State, *pr.Title, *pr.Number, *pr.CreatedAt)
	}
}
