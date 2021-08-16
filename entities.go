package marusia

// RequestType ...
type RequestType string

const (
	// SimpleUtterance ...
	SimpleUtterance RequestType = "SimpleUtterance"
	// ButtonPressed ...
	ButtonPressed RequestType = "ButtonPressed"
)

// Request structure of the incoming message
type Request struct {
	Meta struct {
		ClientID string `json:"client_id"`
		Locale   string `json:"locale"`
		Timezone string `json:"timezone"`
		//Interfaces struct{
		//	Screen string
		//}
	} `json:"meta"`
	Request struct {
		Command           string `json:"command"`
		OriginalUtterance string `json:"original_utterance"`
		Type              string `json:"type"`
		//Payload string
		NLU struct {
			Tokens   []string
			Entities []string
		} `json:"nlu"`
	} `json:"request"`
	Session struct {
		SessionID string `json:"session_id"`
		UserID    string `json:"user_id"`
		SkillID   string `json:"skill_id"`
		New       bool   `json:"new"`
		MessageID int    `json:"message_id"`
	} `json:"session"`
	Version string
}

// OriginalUtterance message text
func (r *Request) OriginalUtterance() string {
	return r.Request.OriginalUtterance

}

// CardType ...
type CardType string

const (
	// BigImage ...
	BigImage CardType = "BigImage"
	// ItemsList ...
	ItemsList CardType = "ItemsList"
)

// Response ...
type Response struct {
	Response struct {
		Text string `json:"text"`
		TTS  string `json:"tts"`
		// Card struct {
		// 	Type    CardType `json:"type"`
		// 	ImageID int      `json:"image_id"`
		// } `json:"card,omitempty"`
		// Buttons []struct {
		// 	Title string `json:"title"`
		// 	//Payload struct{}
		// 	URL string `json:"url"`
		// } `json:"buttons,omitempty"`
		EndSession bool `json:"end_session"`
	} `json:"response"`
	Session struct {
		SessionID string `json:"session_id"`
		MessageID int    `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"session"`
	Version string `json:"version"`
}

// LoadSession prepare respons from request
func (resp *Response) LoadSession(req *Request) *Response {
	resp.Session.SessionID = req.Session.SessionID
	resp.Session.MessageID = req.Session.MessageID
	resp.Session.UserID = req.Session.UserID
	resp.Version = req.Version
	return resp
}

// Text ...
func (resp *Response) Text(s string) *Response {
	resp.Response.Text = s
	return resp
}

// TTS text to speech
func (resp *Response) TTS(s string) *Response {
	resp.Response.TTS = s
	return resp
}

// EndSession mark session as ended
func (resp *Response) EndSession() *Response {
	resp.Response.EndSession = true
	return resp
}
