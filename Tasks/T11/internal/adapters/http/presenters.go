package http

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

//type successResp struct {
//	Result string `json:"result"`
//}
//
//type errResp struct {
//	Err string `json:"error"`
//}

func (a AdapterHTTP) respondSuccess(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	if _, err := fmt.Fprintf(w, "{\"result\":\"%s\"}", msg); err != nil {
		a.logger.Warnw("error writing response", zap.Error(err))
	}
}

func (a AdapterHTTP) respondError(w http.ResponseWriter, msg string, status int, err error) {
	a.logger.Info(err)
	http.Error(w, fmt.Sprintf("{\"error\":\"%s\"}", msg), status)
}
