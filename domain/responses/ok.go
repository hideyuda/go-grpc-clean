package responses

type OK struct {
	OK bool `json:"ok"`
}

func NewOK(ok bool) OK {
	return OK{ok}
}
