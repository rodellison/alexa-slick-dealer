package alexa

// Modified version of Arien Malec's work
// https://github.com/arienmalec/alexa-go
// https://medium.com/@amalec/alexa-skills-with-go-54db0c21e758

const (
	HelpIntent     = "AMAZON.HelpIntent"
	CancelIntent   = "AMAZON.CancelIntent"
	StopIntent     = "AMAZON.StopIntent"
	FallbackIntent = "AMAZON.FallbackIntent"
	YesIntent      = "AMAZON.YesIntent"
	NoIntent       = "AMAZON.NoIntent"
)

const (
	LocaleItalian           = "it-IT"
	LocaleGerman            = "de-DE"
	LocaleAustralianEnglish = "en-AU"
	LocaleCanadianEnglish   = "en-CA"
	LocaleBritishEnglish    = "en-GB"
	LocaleIndianEnglish     = "en-IN"
	LocaleAmericanEnglish   = "en-US"
	LocaleSpanishUS         = "es-US"
	LocaleJapanese          = "ja-JP"
)

type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Body    ReqBody `json:"request"`
	Context Context `json:"context"`
}

type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

type ReqBody struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	Locale      string `json:"locale"`
	Intent      Intent `json:"intent,omitempty"`
	Reason      string `json:"reason,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
}

type Intent struct {
	Name  string          `json:"name"`
	Slots map[string]Slot `json:"slots"`
}

type Slot struct {
	Name        string      `json:"name"`
	Value       string      `json:"value"`
	Resolutions Resolutions `json:"resolutions"`
}

type Resolutions struct {
	ResolutionPerAuthority []struct {
		Values []struct {
			Value struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}

type Application struct {
	ApplicationID string `json:"applicationId,omitempty"`
}

type Context struct {
	System System `json:"system,omitempty"`
	//	Viewport Viewport  //to be defined later if needed
	//	Viewports Viewports   //to be defined if needed
}

//Example of what comes in the request context section..
//To see if user's device supports visual display of APL, check for presence of:
//context.System.device.supportedInterfaces.Alexa.Presentation.APL
/*
	"context": {
		"System": {
			"application": {
				"applicationId": "amzn1.ask.skill..."
			},
			"user": {
				"userId": "amzn1.ask.account.AG5RRFTTIOF44BY7IW..."
			},
			"device": {
				"deviceId": "amzn1.ask.device.AFXHAQJ2AOG...",
				"supportedInterfaces": {
					"Alexa.Presentation.APL": {
						"runtime": {
							"maxVersion": "1.3"
						}
					}
				}
			},
			"apiEndpoint": "https://api.amazonalexa.com",
			"apiAccessToken": "eyJ0eXAiOiJKV1QiLCJhbG...."
		},

*/

type System struct {
	Application    Application `json:"application,omitempty"`
	User           User        `json:"user,omitempty"`
	Device         Device      `json:"device,omitempty"`
	APIEndPoint    string      `json:"apiEndpoint,omitempty"`
	APIAccessToken string      `json:"apiAccessToken"`
}

type User struct {
	UserID string `json:"userId,omitempty"`
}

type Device struct {
	DeviceID            string              `json:"deviceId,omitempty"`
	SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces,omitempty"`
}

type SupportedInterfaces struct {
	APL APL `json:"Alexa.Presentation.APL,omitempty"`
}

type APL struct {
	Runtime Runtime `json:"runtime,omitempty"`
}

type Runtime struct {
	MaxVersion string `json:"maxVersion,omitempty"`
}

func IsEnglish(locale string) bool {
	switch locale {
	case LocaleAmericanEnglish:
		return true
	case LocaleIndianEnglish:
		return true
	case LocaleBritishEnglish:
		return true
	case LocaleCanadianEnglish:
		return true
	case LocaleAustralianEnglish:
		return true
	default:
		return false
	}
}

func SupportsAPL(request *Request) bool {
	if request.Context.System.Device.SupportedInterfaces.APL.Runtime.MaxVersion != "" {
		return true
	} else {
		return false
	}

}
