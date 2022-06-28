package http

type responseSuccess struct {
	Result string	`json:"result"`
}

type responseErr struct {
	Err string `json:"error"`
}

func (s Server)