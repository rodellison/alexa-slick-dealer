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
				"applicationId": "amzn1.ask.skill.42628c60-eb3d-4200-be0d-8b45c7ca1c6e"
			},
			"user": {
				"userId": "amzn1.ask.account.AG5RRFTTIOF44BY7IWFB3P2EARKJ4DLJKBYFSW6HAMAF4E6IOCHQZDGSCXLK5KQK2IPTF5UF5J77WZJ66GLI6J3F26EZTRAZCED65UMLQJDDD3FRTP3IYA5DA46J6MT4XX6N23TWWAYJ2LHTHDLQG23X5FAKKSN254VVLNCTBDE64DLKDXO6CVVTNHKYQOOUBTL4STBZ6BOBUKQ"
			},
			"device": {
				"deviceId": "amzn1.ask.device.AFXHAQJ2AOGKW4AWMTAQKSJITPZDNTQKFWFQ773LTRAZVJO5HQMQ3WKBQQTUV74556RJJO5ZCPS4XUNQF2CCWD5JY6YA4LIY7J54WEV2SZ7CV5FXZFQ6RDQ6BIQTSBS64643SQJSOMGYORW5PH4FXQ5O3S4Q",
				"supportedInterfaces": {
					"Alexa.Presentation.APL": {
						"runtime": {
							"maxVersion": "1.3"
						}
					}
				}
			},
			"apiEndpoint": "https://api.amazonalexa.com",
			"apiAccessToken": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjEifQ.eyJhdWQiOiJodHRwczovL2FwaS5hbWF6b25hbGV4YS5jb20iLCJpc3MiOiJBbGV4YVNraWxsS2l0Iiwic3ViIjoiYW16bjEuYXNrLnNraWxsLjQyNjI4YzYwLWViM2QtNDIwMC1iZTBkLThiNDVjN2NhMWM2ZSIsImV4cCI6MTU5MzcwNDM2OCwiaWF0IjoxNTkzNzA0MDY4LCJuYmYiOjE1OTM3MDQwNjgsInByaXZhdGVDbGFpbXMiOnsiY29udGV4dCI6IkFBQUFBQUFBQUFDNWxlNEpLUWVEZ0tRSXlPaGIrekFnSkFFQUFBQUFBQUFUcGwyU0l1azJGaEpDRVg5dkphQi95d3Q3bW8wTmZVb3pUOGZseVJZbWtGL0hDSmk2VS9ic05hVE1oelV5aHhtclBRWVRaalJWYXdCOUVPMFBKREgrdmpNNGJUdUk2M01tU25tWDRZdFRNQUVOVnQwYThuRk1oNjRVWFExS1BZVHNXOS9rdFl5d0t3Q1k1ZG4zNEdXWDBGR2NYNzY5TXlWdUs4UmFJakc0RDhVU1dQMzhFZFZhd0k5aXdjdk02eW42RjBFVXI3bW01T0Y4M3V1RWdYOWV4bDhCYllRMDZPUDlQV0Z4MTQ0T0YzVGlQSUdXeVcxais5aEhqSHBuZVNzYjI4bUFXekZ5T0R2SEtHSHdyVU1ydHJSVWJyUjBIK2t6eUtGNG90dTFVbmxvUkZFeEFtZU54Z3pGalV4bkxCS1lTeTI2WU5mOGhmTWc1TUU0THpvZnlNbzh4N0RqcE5weUt5SzROdUhBZTVVV3F4L3YvWkhFdkhyUVM3UjczUmFBZXNzUiIsImNvbnNlbnRUb2tlbiI6bnVsbCwiZGV2aWNlSWQiOiJhbXpuMS5hc2suZGV2aWNlLkFGWEhBUUoyQU9HS1c0QVdNVEFRS1NKSVRQWkROVFFLRldGUTc3M0xUUkFaVkpPNUhRTVEzV0tCUVFUVVY3NDU1NlJKSk81WkNQUzRYVU5RRjJDQ1dENUpZNllBNExJWTdKNTRXRVYyU1o3Q1Y1RlhaRlE2UkRRNkJJUVRTQlM2NDY0M1NRSlNPTUdZT1JXNVBINEZYUTVPM1M0USIsInVzZXJJZCI6ImFtem4xLmFzay5hY2NvdW50LkFHNVJSRlRUSU9GNDRCWTdJV0ZCM1AyRUFSS0o0RExKS0JZRlNXNkhBTUFGNEU2SU9DSFFaREdTQ1hMSzVLUUsySVBURjVVRjVKNzdXWko2NkdMSTZKM0YyNkVaVFJBWkNFRDY1VU1MUUpEREQzRlJUUDNJWUE1REE0Nko2TVQ0WFg2TjIzVFdXQVlKMkxIVEhETFFHMjNYNUZBS0tTTjI1NFZWTE5DVEJERTY0RExLRFhPNkNWVlROSEtZUU9PVUJUTDRTVEJaNkJPQlVLUSJ9fQ.AhSCps877GQRMExWybtosHaB-kurouoOVrOiqO4EwZ8j_aC__-5s2Yrov0OGB-e8Eib_M6-xPI8KcJ_o4A9hvKL45sUeMYD7oI5xWc6EGGmW3NA7Y9ocClqko0RrccXiUE148TqU_mMwX27q4VhCSli7QD7AQa0YC5GGbrefUAR_GQk_cGYZstMReejrP5ingZOVE0b-dytR2m3jJqV--ilez8aI7p5oRFzO6u7wpBXjh9BxkjSHgF1Sw30biF8w7Pnpt-zc6YdMTH5q4zalZTRro1ODxnuXuej4xfxmQ3zwpKh8nqIx6r2YnWC0gaKj9SACTHDxMXlzWDz4oPPBxw"
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
