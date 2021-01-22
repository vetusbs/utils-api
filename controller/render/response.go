package render

type Response struct {
	Values []interface{} `json:"values"`
}

func NewResponse(values []interface{}) Response {
	return Response{Values: values}
}
