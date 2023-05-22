package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5242248788342-5323673192992-ASFINe23BefEhF3rgLlQHy8e")
	os.Setenv("CHANNEL_ID","C05747ARBF0")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"upload.txt"}

	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)

		if err!=nil {
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n",file.Name,file.URL)
	}
}