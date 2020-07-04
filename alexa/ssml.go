package alexa

type SSML struct {
	text  string
	pause string
}

type SSMLBuilder struct {
	SSML []SSML
}

func (builder *SSMLBuilder) Say(text string) {
	//NOTE!! ParseString result is a Lowercase string
	text = ParseString(text)
	builder.SSML = append(builder.SSML, SSML{text: text})
}

func (builder *SSMLBuilder) Pause(pause string) {
	builder.SSML = append(builder.SSML, SSML{pause: pause})
}

func (builder *SSMLBuilder) Build() string {
	var response string
	for index, ssml := range builder.SSML {
		if ssml.text != "" {
			response += ssml.text + " "
		} else if ssml.pause != "" && index != len(builder.SSML)-1 {
			response += "<break time='" + ssml.pause + "ms'/> "
		}
	}
	return "<speak>" + response + "</speak>"
}
