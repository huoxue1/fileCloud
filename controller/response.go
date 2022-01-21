package controller

type Response struct {
	Ok    bool        `json:"ok"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func ok(msg string, data interface{}) Response {
	return Response{
		Ok:    true,
		Msg:   msg,
		Data:  data,
		Error: nil,
	}
}

func fail(msg string, err interface{}) Response {
	return Response{
		Ok:    false,
		Msg:   msg,
		Data:  nil,
		Error: err,
	}
}
