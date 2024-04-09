package routes

import (
	"net/http"
	"strconv"

	"github.com/fanchann/grpc_math/helpers"
	"github.com/fanchann/grpc_math/proto"
	"github.com/fanchann/grpc_math/web"
)

type IRoutes interface {
	Add(w http.ResponseWriter, r *http.Request)
	Multiply(w http.ResponseWriter, r *http.Request)
	Reduce(w http.ResponseWriter, r *http.Request)
	Devide(w http.ResponseWriter, r *http.Request)
}

type Routes struct {
	proto.MathServiceClient
}

func NewRoutes(p proto.MathServiceClient) IRoutes {
	return &Routes{MathServiceClient: p}
}

func (ro *Routes) Add(w http.ResponseWriter, r *http.Request) {
	result, err := onAction(ro.MathServiceClient, "add", r)
	if err != nil {
		helpers.WriteToResponseBody(w, http.StatusBadRequest, toWebResponse(false, err.Error()))
		return
	}
	helpers.WriteToResponseBody(w, http.StatusOK, toWebResponse(true, result))
}

func (ro *Routes) Multiply(w http.ResponseWriter, r *http.Request) {
	result, err := onAction(ro.MathServiceClient, "multiply", r)
	if err != nil {
		helpers.WriteToResponseBody(w, http.StatusBadRequest, toWebResponse(false, err.Error()))
		return
	}
	helpers.WriteToResponseBody(w, http.StatusOK, toWebResponse(true, result))
}

func (ro *Routes) Reduce(w http.ResponseWriter, r *http.Request) {
	result, err := onAction(ro.MathServiceClient, "reduce", r)
	if err != nil {
		helpers.WriteToResponseBody(w, http.StatusBadRequest, toWebResponse(false, err.Error()))
		return
	}
	helpers.WriteToResponseBody(w, http.StatusOK, toWebResponse(true, result))
}

func (ro *Routes) Devide(w http.ResponseWriter, r *http.Request) {
	result, err := onAction(ro.MathServiceClient, "divide", r)
	if err != nil {
		helpers.WriteToResponseBody(w, http.StatusBadRequest, toWebResponse(false, err.Error()))
		return
	}
	helpers.WriteToResponseBody(w, http.StatusOK, toWebResponse(true, result))
}

func toWebResponse[X any](status bool, data X) *web.Responses {
	return &web.Responses{
		Success: status,
		Data:    data,
	}
}

func onAction(m proto.MathServiceClient, action string, r *http.Request) (*proto.Response, error) {
	switch action {
	case "add":
		a := r.PathValue("a")
		b := r.PathValue("b")

		aInt, err := strconv.Atoi(a)
		bInt, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}

		result, err := m.Add(r.Context(), &proto.Request{A: int64(aInt), B: int64(bInt)})
		if err != nil {
			return nil, err
		}
		return result, nil
	case "multiply":
		a := r.PathValue("a")
		b := r.PathValue("b")

		aInt, err := strconv.Atoi(a)
		bInt, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}

		result, err := m.Multiply(r.Context(), &proto.Request{A: int64(aInt), B: int64(bInt)})
		if err != nil {
			return nil, err
		}
		return result, nil
	case "divide":
		a := r.PathValue("a")
		b := r.PathValue("b")

		aInt, err := strconv.Atoi(a)
		bInt, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}

		result, err := m.Devide(r.Context(), &proto.Request{A: int64(aInt), B: int64(bInt)})
		if err != nil {
			return nil, err
		}
		return result, nil
	case "reduce":
		a := r.PathValue("a")
		b := r.PathValue("b")

		aInt, err := strconv.Atoi(a)
		bInt, err := strconv.Atoi(b)
		if err != nil {
			return nil, err
		}

		result, err := m.Reduce(r.Context(), &proto.Request{A: int64(aInt), B: int64(bInt)})
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, nil
	}
}
