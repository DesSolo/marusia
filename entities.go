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
	} `json:"meta"`
	Request struct {
		Command           string `json:"command"`
		OriginalUtterance string `json:"original_utterance"`
		Type              string `json:"type"`
		NLU               struct {
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

// IsNewSession ...
func (r *Request) IsNewSession() bool {
	return r.Session.New
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
		Text       string `json:"text"`
		TTS        string `json:"tts"`
		EndSession bool   `json:"end_session"`
	} `json:"response"`
	Session struct {
		SessionID string `json:"session_id"`
		MessageID int    `json:"message_id"`
		UserID    string `json:"user_id"`
	} `json:"session"`
	Version string `json:"version"`
}

// LoadSession prepare respons from request
func (resp *Response) LoadSession(req *Request) {
	resp.Session.SessionID = req.Session.SessionID
	resp.Session.MessageID = req.Session.MessageID
	resp.Session.UserID = req.Session.UserID
	resp.Version = req.Version
}

// Text ...
func (resp *Response) Text(s string) {
	resp.Response.Text = s
}

// TTS text to speech
func (resp *Response) TTS(s string) {
	resp.Response.TTS = s
}

// EndSession mark session as ended
func (resp *Response) EndSession() {
	resp.Response.EndSession = true
}
