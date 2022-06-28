package http

type successResp struct {
	Result string `json:"result"`
}

type errResp struct {
	Err string `json:"error"`
}
