package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rodellison/alexa-slick-dealer/alexa"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type FeedResponse struct {
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func RequestFeed(mode string) (FeedResponse, error) {
	endpoint, _ := url.Parse("https://slickdeals.net/newsearch.php")
	queryParams := endpoint.Query()
	queryParams.Set("mode", mode)
	queryParams.Set("searcharea", "deals")
	queryParams.Set("searchin", "first")
	queryParams.Set("rss", "1")
	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	if err != nil {
		return FeedResponse{}, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var feedResponse FeedResponse
		xml.Unmarshal(data, &feedResponse)
		return feedResponse, nil
	}
}

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response

	if request.Body.Type == "LaunchRequest" {
		return alexa.NewRepromptResponse("Hello!, and welcome to Slick Dealer. You can ask a question like, get me the front page deals, or give me the popular deals.", "For instructions on what you can say, please say help me.")
	}

	switch request.Body.Intent.Name {
	case "FrontpageDealIntent":
		response = HandleDealIntent(request, "Frontpage", false)
	case "PopularDealIntent":
		response = HandleDealIntent(request, "Popular", false)
	case alexa.YesIntent:
		response = HandleDealIntent(request, "", true)
	case alexa.NoIntent:
		response = HandleHelpIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case alexa.StopIntent:
		response = HandleStopIntent(request)
	case alexa.CancelIntent:
		response = HandleStopIntent(request)
	case alexa.FallbackIntent:
		response = HandleHelpIntent(request)
	}
	return response
}

func HandleDealIntent(request alexa.Request, dealType string, resumingPrior bool) alexa.Response {

	var feedResponse FeedResponse
	var primarybuilder alexa.SSMLBuilder
	var repromptbuilder alexa.SSMLBuilder

	if resumingPrior {

		incomingSessionAttrs := request.Session.Attributes
		incomingData, _ :=  json.Marshal(incomingSessionAttrs["dataToSave"])
		json.Unmarshal(incomingData, &feedResponse)

	} else {

		if dealType == "Frontpage" {
			feedResponse, _ = RequestFeed("frontpage")
			primarybuilder.Say("Here are the current Front page deals:")
		} else {
			feedResponse, _ = RequestFeed("popdeals")
			primarybuilder.Say("Here are the current popular deals:")
		}
		primarybuilder.Pause("1000")
	}

	if len(feedResponse.Channel.Item) > 3 {
		for j := 0; j < 3; j++ {
			thisItem := feedResponse.Channel.Item[j]
			primarybuilder.Say(thisItem.Title)
			primarybuilder.Pause("1000")
		}

		repromptString := "Would you like to hear more deals"
		primarybuilder.Say(repromptString)
		primarybuilder.Pause("1000")

		repromptbuilder.Say(repromptString)
		repromptbuilder.Pause("1000")

		//Save session attributes data for reentry, should the user answer yes to 'more' details..
		feedResponse.Channel.Item = feedResponse.Channel.Item[3:]

		sessAttrData := make(map[string]interface{})
		sessAttrData["dataToSave"] = feedResponse

		return alexa.NewSSMLResponse(dealType + " Deals", primarybuilder.Build(), repromptbuilder.Build(), false, sessAttrData)

	} else {
		for _, item := range feedResponse.Channel.Item {
			primarybuilder.Say(item.Title)
			primarybuilder.Pause("1000")
		}
		primarybuilder.Say("There are no additional deals")
		primarybuilder.Pause("1000")
		return alexa.NewSSMLResponse(dealType + " Deals", primarybuilder.Build(), "", true, nil)
	}

}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	var repromptBuilder alexa.SSMLBuilder

	builder.Say("OK, Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("What are the frontpage deals.")
	builder.Pause("1000")
	builder.Say("What are the popular deals.")

	repromptBuilder.Say("Please say, What are the frontpage deals, or What are the popular deals")
	repromptBuilder.Pause("500")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build(), repromptBuilder.Build(), false, nil)
}

func HandleStopIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Cancel and Stop", "Thanks and have a great day!, Goodbye.")
}

func Handler(request alexa.Request) (alexa.Response, error) {

	//Ensure this lambda function/code is invoked through the associated Alexa Skill, and not called directly
	if request.Session.Application.ApplicationID != os.Getenv("AppARN") {
		return alexa.NewSimpleResponse("Not authorized", "Please enable and use this skill through an approved Alexa device."), nil
	} else {
		return IntentDispatcher(request), nil
	}

}

func main() {
	lambda.Start(Handler)
}
