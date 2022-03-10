package instagram

// Err err
type Err struct {
	Code      int    `json:"code"`
	FBTrackID string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

func (e *Err) Error() string {
	return e.Message
}
