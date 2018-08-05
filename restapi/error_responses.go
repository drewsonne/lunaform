package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"net/http"
	"github.com/go-openapi/runtime"
)

type CommonServerErrorResponder struct {
	Payload *models.ServerError
	code    int
}

func NewServerError(code int, errorString string) (r *CommonServerErrorResponder) {
	return &CommonServerErrorResponder{
		Payload: &models.ServerError{
			Message: swag.String(errorString),
			Status: swag.String(http.StatusText(
				int(code),
			)),
			StatusCode: swag.Int32(int32(code)),
		},
		code: int(code),
	}
}

// WriteResponse to the client
func (cser *CommonServerErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(cser.code)
	if cser.Payload != nil {
		payload := cser.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (cser *CommonServerErrorResponder) Error() {

}
