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
		return HandleLaunchIntent(request)
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

func HandleLaunchIntent(request alexa.Request) alexa.Response {

	var primarybuilder alexa.SSMLBuilder
	var repromptbuilder alexa.SSMLBuilder
	var response alexa.Response

	primarybuilder.Say("Hello!, and welcome to Slick Dealer. You can ask a question like, What are the frontpage deals, or What are the popular deals.")
	primarybuilder.Pause("1000")
	primarybuilder.Say("For instructions on what you can say, please say help me.")
	repromptbuilder.Say("Please ask a question like, What are the frontpage deals, or What are the popular deals.")


	if alexa.SupportsAPL(&request)	{
		response = alexa.NewAPLAskResponse("Welcome to Slick Dealer",
			primarybuilder.Build(),
			repromptbuilder.Build(),
			"You can ask a question like, What are the front page deals, or What are the popular deals.",
			false,
			nil,
			"Home")
	} else {
		response = alexa.NewSimpleAskResponse("Welcome to Slick Dealer",
			primarybuilder.Build(),
			repromptbuilder.Build(),
			"You can ask a question like, What are the front page deals, or What are the popular deals.",
			false,
			nil)
	}

	return response

}

func HandleDealIntent(request alexa.Request, dealType string, resumingPrior bool) alexa.Response {

	var feedResponse FeedResponse
	var primarybuilder alexa.SSMLBuilder
	var repromptbuilder alexa.SSMLBuilder
	var sessAttrData map[string]interface{}

	if resumingPrior {

		incomingSessionAttrs := request.Session.Attributes
		incomingData, _ := json.Marshal(incomingSessionAttrs["dataToSave"])
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

		sessAttrData = make(map[string]interface{})
		sessAttrData["dataToSave"] = feedResponse

	} else {
		for _, item := range feedResponse.Channel.Item {
			primarybuilder.Say(item.Title)
			primarybuilder.Pause("1000")
		}
		sessAttrData = nil
		primarybuilder.Say("There are no additional deals. Please ask another question like, What are the popular deals or What are the frontpage deals. Say Cancel to exit. ")
		primarybuilder.Pause("1000")
	}

	if alexa.SupportsAPL(&request) {
		return alexa.NewAPLAskResponse(dealType + " Deals",
			primarybuilder.Build(),
			repromptbuilder.Build(),
			"Here are your deals",
			false,
			sessAttrData,
			"ItemsList")
	} else {
		return alexa.NewSimpleAskResponse(dealType + " Deals",
			primarybuilder.Build(),
			repromptbuilder.Build(),
			"Here are your deals",
			false,
			sessAttrData)
	}

}

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var primaryBuilder alexa.SSMLBuilder
	var repromptBuilder alexa.SSMLBuilder
	var response alexa.Response

	primaryBuilder.Say("OK, Here are some of the things you can ask:")
	primaryBuilder.Pause("1000")
	primaryBuilder.Say("What are the frontpage deals or,")
	primaryBuilder.Pause("500")
	primaryBuilder.Say("What are the popular deals.")

	repromptBuilder.Say("Please ask a question like, What are the frontpage deals, or " +
		"What are the popular deals. Say Cancel if you'd like to quit.")
	repromptBuilder.Pause("500")

	if alexa.SupportsAPL(&request) {
		response = alexa.NewAPLAskResponse("Slick Dealer Help",
			primaryBuilder.Build(),
			repromptBuilder.Build(),
			"Please ask a question like, What are the frontpage deals, or What are the popular deals. You can also say Cancel to exit.",
			false,
			nil,
			"Help")
	} else {
		response = alexa.NewSimpleAskResponse("Slick Dealer Help",
			primaryBuilder.Build(),
			repromptBuilder.Build(),
			"Please ask a question like, What are the frontpage deals, or What are the popular deals. You can also say Cancel to exit.",
			false,
			nil)
	}

	return response
}

func HandleStopIntent(request alexa.Request) alexa.Response {

	var primarybuilder alexa.SSMLBuilder
	var response alexa.Response

	primarybuilder.Say("Thanks and have a great day!, Goodbye.")
	response = alexa.NewSimpleTellResponse(os.Getenv("SkillTitle"),
		primarybuilder.Build(),
		"Thanks and have a great day!, Goodbye.",
		true,
		nil)

	return response
}

func Handler(request alexa.Request) (alexa.Response, error) {

	//Ensure this lambda function/code is invoked through the associated Alexa Skill, and not called directly
	if request.Session.Application.ApplicationID != os.Getenv("AppARN") {
		var primarybuilder alexa.SSMLBuilder
		primarybuilder.Say("Sorry, not authorized. Please enable and use this skill through an approved Alexa device.")
		return alexa.NewSimpleTellResponse("Not authorized",
			primarybuilder.Build(),
			"Not authorized, Please enable and use this skill through an approved Alexa device.",
			true,
			nil), nil
		} else {
		return IntentDispatcher(request), nil
	}

}

func main() {
	lambda.Start(Handler)
}
