package slack

import  ( 
	"fmt"
    "os"
    "time"
    "internal/logger/logger"
	"github.com/slack-go/slack"
)


const (
  CHANNEL_ID = "INSERT_YOUR CHANNEL_ID"
  TOKEN_ID = "INSERT_YOUR TOKEN_ID  "
)


func InitSlackClient(tokenid string) *slack.Client {
    client := slack.New(TOKEN_ID)
    return client
}


func getSlackChannelHistory(){
  api := InitSlackClient()
  historyParams := slack.HistoryParameters{Latest: "", Oldest: "0", Count: 2, Inclusive: false, Unreads:false,}
  history, err := api.GetChannelHistory(CHANNEL_ID, historyParams)
  if err != nil {
 		logger.Error("Cannot get Slack History Reasons : "+err)
  		return 
  }

  return history.Messages

}

func SlackPostMessage(msg string) error {
  api := InitSlackClient()
  channelID, timestamp, err := api.PostMessage(CHANNEL_ID,slack.MsgOptionText(msg, false))  
  if err != nil {
    return err
  }
  return nil
}