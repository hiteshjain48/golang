package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5242248788342-5273380460768-uKEbPuUhx2qqW87Fj1EJtmxc")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A057PG5MAQZ-5242321498006-a944f7bc624cfc262d7da94205042c2b7e273f8b747a66d6ed7f15a0ac094030")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob claculator",
		Examples: []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob	
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}