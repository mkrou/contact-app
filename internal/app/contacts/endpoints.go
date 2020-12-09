package contacts

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints contains all endpoints of contacts module
type Endpoints struct {
	GetContacts   endpoint.Endpoint
	AddContact    endpoint.Endpoint
	EditContact   endpoint.Endpoint
	DeleteContact endpoint.Endpoint
}

//NewContactEndpoints will initialize contact endpoints
func NewContactEndpoints(s ContactServiceModule) Endpoints {
	return Endpoints{
		GetContacts:   makeGetContacts(s),
		AddContact:    makeAddContact(s),
		EditContact:   makeEditContact(s),
		DeleteContact: makeDeleteContact(s),
	}
}

func makeGetContacts(s ContactServiceModule) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {

		contacts, err := s.GetContacts(ctx)
		if err != nil {
			return nil, fmt.Errorf("get contacts endpoint: %w", err)
		}

		return contacts, nil
	}
}

type AddContactRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

func makeAddContact(s ContactServiceModule) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddContactRequest)

		contact, err := s.AddContact(ctx, Contact{
			Name:    req.Name,
			Surname: req.Surname,
			Phone:   req.Phone,
		})
		if err != nil {
			return nil, fmt.Errorf("add contact endpoint: %w", err)
		}

		return contact, nil
	}
}

type EditContactRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}

func makeEditContact(s ContactServiceModule) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EditContactRequest)

		contact, err := s.EditContact(ctx, Contact{
			ID:      req.ID,
			Name:    req.Name,
			Surname: req.Surname,
			Phone:   req.Phone,
		})
		if err != nil {
			return nil, fmt.Errorf("edit contact endpoint: %w", err)
		}

		return contact, nil
	}
}

type DeleteContactRequest struct {
	ID int `json:"id"`
}

func makeDeleteContact(s ContactServiceModule) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteContactRequest)

		err := s.DeleteContact(ctx, req.ID)
		if err != nil {
			return nil, fmt.Errorf("add contact endpoint: %w", err)
		}

		return struct {}{}, nil
	}
}
