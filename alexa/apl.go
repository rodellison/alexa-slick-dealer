package alexa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	FetchAPL func() (*APLDocumentAndData, error)
	FileToRead string
)

func init() {
	FetchAPL = CreateAPLDocAndData
	FileToRead = os.Getenv("AplTemplate")
}

type APLDocumentAndData struct {
	APLDocument    APLDocument    `json:"document"`
	APLDataSources APLDataSources `json:"datasources"`
	APLSources     interface{}    `json:"sources,omitempty"`
}

//The types defined as []interface{} or just interface[] can stay as is, there's no need to break them down to
//struct level as they shouldn't be modified by code after loading
//They are broken to this level just to facilitate some file load and UNMarshal testing
type APLDocument struct {
	Type         string        `json:"type,omitempty"`
	Version      string        `json:"version,omitempty"`
	Theme        string        `json:"theme,omitempty"`
	Import       []interface{} `json:"import,omitempty"`
	Resources    []interface{} `json:"resources,omitempty"`
	Styles       interface{}   `json:"styles,omitempty"`
	Layouts      interface{}   `json:"layouts,omitempty"`
	MainTemplate interface{}   `json:"mainTemplate,omitempty"`
}

//The DataSources sections and types below needs to be adjusted to match the apl_template_export.json file as the properties and variables for the APL document
//are completely customized and unique for each skill.
//We need the properties to be accessible so as to dynamically populate them when responding to user requests
type APLDataSources struct {
	TemplateData TemplateData `json:"TemplateData,omitempty"`
}

type TemplateData struct {
	Type         string            `json:"type"`
	ObjectID     string            `json:"objectId"`
	Properties   APLDataProperties `json:"properties,omitempty"`
	Transformers []interface{}     `json:"transformers,omitempty"`
}

type APLDataProperties struct {
	Title                 string   `json:"Title,omitempty"`
	LayoutToUse           string   `json:"LayoutToUse,omitempty"`
	Locale                string   `json:"Locale,omitempty"`
	HeadingText           string   `json:"HeadingText,omitempty"`
	HintString            string   `json:"HintString,omitempty"`
	HomeImageUrlXSMALL    string   `json:"HomeImageUrlXSMALL,omitempty"`
	HomeImageUrlXLARGE    string   `json:"HomeImageUrlXLARGE,omitempty"`
	HelpImageUrlXSMALL    string   `json:"HelpImageUrlXSMALL,omitempty"`
	HelpImageUrlXLARGE    string   `json:"HelpImageUrlXLARGE,omitempty"`
	ItemsImageUrlXSMALL   string   `json:"ItemsImageUrlXSMALL,omitempty"`
	ItemsImageUrlXLARGE   string   `json:"ItemsImageUrlXLARGE,omitempty"`
	LogoUrl               string   `json:"LogoUrl,omitempty"`
	GeneralSquareImageUrl string   `json:"GeneralSquareImageUrl,omitempty"`
	ItemsText             []string `json:"ItemsText,omitempty"`
}

func CreateAPLDocAndData() (*APLDocumentAndData, error) {

	aplDoc := &APLDocumentAndData{
		APLDocument:    APLDocument{},
		APLDataSources: APLDataSources{},
	}

	//apl_template_export.json is the name given to the file downloaded from the Alexa Developer console (Display tool)
    fmt.Println("Opening APLTemplate JSON file: ", FileToRead)
	file, err := ioutil.ReadFile(FileToRead)
	if err != nil {
		fmt.Println("Error reading APL json file: ", err.Error())
		return nil, err
	} else {
		errUnMarshal := json.Unmarshal([]byte(file), &aplDoc)
		if errUnMarshal != nil {
			fmt.Println("Error Unmarshalling APL json file: ", err.Error())
			return nil, err
		}
	}

	return aplDoc, nil

}
