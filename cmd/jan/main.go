package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sig1nt/JAN/internal/discord"
)

var id = flag.String("webhook", "", "the discord webhook to use")
var token = flag.String("token", "", "the discord webhook token")

func main() {
	flag.Parse()
	if *id == "" || *token == "" {
		fmt.Fprintln(os.Stderr, "webhook and token must both be specified")
		os.Exit(1)
	}

	webhook := &discord.Webhook{
		ID:    *id,
		Token: *token,
	}
	if err := webhook.Execute("hello from Go!"); err != nil {
		fmt.Fprintf(os.Stderr, "execute failed: %v", err)
		os.Exit(2)
	}
}
