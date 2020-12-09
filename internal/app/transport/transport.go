package transport

import (
	"context"
	"encoding/json"
	"errors"
	"contact/internal/app/contacts"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//NewTransport will create http handler with all api methods
func NewTransport(contactsEndpoint contacts.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/contacts").Handler(kithttp.NewServer(
		contactsEndpoint.GetContacts,
		decodeNoRequest,
		encodeResponse(http.StatusOK),
	))

	r.Methods(http.MethodPost).Path("/contacts").Handler(kithttp.NewServer(
		contactsEndpoint.AddContact,
		decodeAddContactRequest,
		encodeResponse(http.StatusCreated),
	))

	r.Methods(http.MethodPut).Path("/contacts/{id:[0-9]+}").Handler(kithttp.NewServer(
		contactsEndpoint.EditContact,
		decodeEditContactRequest,
		encodeResponse(http.StatusOK),
	))

	r.Methods(http.MethodDelete).Path("/contacts/{id:[0-9]+}").Handler(kithttp.NewServer(
		contactsEndpoint.DeleteContact,
		decodeDeleteContactRequest,
		encodeResponse(http.StatusOK),
	))

	return r
}

//encodeResponse will response encoded to json object with provided status code
func encodeResponse(status int) func(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return func(_ context.Context, w http.ResponseWriter, response interface{}) error {
		w.WriteHeader(status)

		return json.NewEncoder(w).Encode(response)
	}
}

//decodeNoRequest will do nothing special, but needed for empty requests
func decodeNoRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return nil, nil
}

//decodeAddContactRequest decode contacts.AddContactRequest from json
func decodeAddContactRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req contacts.AddContactRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

//decodeEditContactRequest decode contacts.EditContactRequest from json
func decodeEditContactRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req contacts.EditContactRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errors.New("transport error: id has wrong format")
	}

	req.ID = id

	return req, nil
}

//decodeDeleteContactRequest decode contacts.DeleteContactRequest from json
func decodeDeleteContactRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req contacts.DeleteContactRequest

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return nil, errors.New("transport error: id has wrong format")
	}

	req.ID = id

	return req, nil
}
