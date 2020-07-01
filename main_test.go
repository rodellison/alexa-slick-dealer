package main

import (
	"github.com/rodellison/alexa-slick-dealer/alexa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleLaunchIntent(t *testing.T) {

	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "LaunchRequest",
		},
		Context: alexa.Context{},
	}

	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}


func TestHandleFrontpageDealIntent(t *testing.T) {
	
	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "FrontpageDealIntent",
				Slots: nil,
			},

		},
		Context: alexa.Context{},
	}
	
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandlePopularDealIntent(t *testing.T) {
	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "PopularDealIntent",
				Slots: nil,
			},

		},
		Context: alexa.Context{},
	}

	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty")
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleHelpIntent(t *testing.T) {
	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.HelpIntent",
				Slots: nil,
			},

		},
		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleDealResumeDetails(t *testing.T) {
	sessionAttrData := make(map[string]interface{})
	sessionAttrData["dataToSave"] = "some data"

	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{
			Attributes: sessionAttrData,
		},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.YesIntent",
				Slots: nil,
			},

		},		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleNoIntent(t *testing.T) {

	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.NoIntent",
				Slots: nil,
			},

		},		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleStopIntent(t *testing.T) {
	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.StopIntent",
				Slots: nil,
			},

		},		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}
func TestHandleCancelIntent(t *testing.T) {

	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.CancelIntent",
				Slots: nil,
			},

		},		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}

func TestHandleFallback(t *testing.T) {

	theRequest := &alexa.Request{
		Version: "1.0",
		Session: alexa.Session{},
		Body:    alexa.ReqBody{
			Type: "IntentRequest",
			Intent: alexa.Intent{
				Name:  "AMAZON.FallbackIntent",
				Slots: nil,
			},

		},		Context: alexa.Context{},
	}
	response, _ := Handler(*theRequest)
	assert.NotEmpty(t, response, "The response should not be empty", false)
	assert.NotEmpty(t, response.Body.OutputSpeech, "There should be output speech")
}