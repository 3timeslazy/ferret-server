package http

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func InternalError() middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		rw.WriteHeader(http.StatusInternalServerError)

		if err := producer.Produce(rw, "Internal error"); err != nil {
			panic(err)
		}
	})
}

func NotFound() middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		rw.WriteHeader(http.StatusNotFound)

		if err := producer.Produce(rw, "Not found"); err != nil {
			panic(err)
		}
	})
}

func Bad(message string) middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		rw.WriteHeader(http.StatusBadRequest)

		if message == "" {
			message = "Bad request"
		}

		if err := producer.Produce(rw, message); err != nil {
			panic(err)
		}
	})
}

func Unauthorized() middleware.Responder {
	return middleware.ResponderFunc(func(rw http.ResponseWriter, producer runtime.Producer) {
		rw.WriteHeader(http.StatusUnauthorized)

		if err := producer.Produce(rw, "Unauthorized"); err != nil {
			panic(err)
		}
	})
}
