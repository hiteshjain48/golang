package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5242248788342-5273380460768-uKEbPuUhx2qqW87Fj1EJtmxc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A057PG5MAQZ-5242321498006-a944f7bc624cfc262d7da94205042c2b7e273f8b747a66d6ed7f15a0ac094030")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	
}