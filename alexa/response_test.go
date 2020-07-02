package alexa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSimpleResponse(t *testing.T) {

	response := NewSimpleResponse("TextTitle1", "TextValue1")
	assert.Contains(t, response.Body.OutputSpeech.Text, "TextValue1")

}

func TestNewSimpleResponseWithCard(t *testing.T) {

	response := NewSimpleResponseWithCard("TextTitle1", "TextValue1")
	assert.Contains(t, response.Body.OutputSpeech.Text, "TextValue1")
	assert.Contains(t, response.Body.Card.Title, "TextTitle1")

}

func TestNewRepromptResponse(t *testing.T) {

	response := NewRepromptResponse("TextValue1", "RepromptsTextValue1")
	assert.Contains(t, response.Body.Reprompt.OutputSpeech.Text, "RepromptsTextValue1")

}
func TestNewSSMLResponseWithEndSession(t *testing.T) {


	response := NewSSMLResponse("TextTitle1", "<speak>TextValue1</speak>", "RepromptTextValue1", true, nil)
	assert.Contains(t, response.Body.OutputSpeech.SSML, "TextValue1")
	assert.True(t, response.Body.ShouldEndSession)

}
func TestNewSSMLResponseWithContinueSession(t *testing.T) {

	response := NewSSMLResponse("TextTitle1", "<speak>TextValue1</speak>", "RepromptTextValue1", false, map[string]interface{}{})
	assert.Contains(t, response.Body.OutputSpeech.SSML, "TextValue1")
	assert.False(t, response.Body.ShouldEndSession)

}