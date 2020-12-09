package contacts

import (
	"context"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
)

type ContactStorageRepository interface {
	GetContacts(ctx context.Context) ([]Contact, error)
	AddContact(ctx context.Context, contact Contact) (Contact, error)
	EditContact(ctx context.Context, contact Contact) (Contact, error)
	DeleteContact(ctx context.Context, id int) error
}

type ContactServiceModule interface {
	GetContacts(ctx context.Context) ([]Contact, error)
	AddContact(ctx context.Context, contact Contact) (Contact, error)
	EditContact(ctx context.Context, contact Contact) (Contact, error)
	DeleteContact(ctx context.Context, id int) error
}

func NewContactService(repository ContactStorageRepository) *ContactService {
	return &ContactService{
		repository: repository,
	}
}

type ContactService struct {
	repository ContactStorageRepository
}

func (p ContactService) GetContacts(ctx context.Context) ([]Contact, error) {
	return p.repository.GetContacts(ctx)
}

func (p ContactService) AddContact(ctx context.Context, contact Contact) (Contact, error) {
	err := validation.ValidateStruct(&contact,
		validation.Field(&contact.Name, validation.Required, validation.Length(3, 30)),
		validation.Field(&contact.Surname, validation.Required, validation.Length(3, 30)),
		validation.Field(&contact.Phone, validation.Required, validation.Length(11, 11)),
	)
	if err != nil {
		return Contact{}, fmt.Errorf("add contact: validation: %w", err)
	}

	return p.repository.AddContact(ctx, contact)
}

func (p ContactService) EditContact(ctx context.Context, contact Contact) (Contact, error) {
	err := validation.ValidateStruct(&contact,
		validation.Field(&contact.Name, validation.Required, validation.Length(3, 30)),
		validation.Field(&contact.Surname, validation.Required, validation.Length(3, 30)),
		validation.Field(&contact.Phone, validation.Required, validation.Length(11, 11)),
	)
	if err != nil {
		return Contact{}, fmt.Errorf("add contact: validation: %w", err)
	}

	return p.repository.EditContact(ctx, contact)
}

func (p ContactService) DeleteContact(ctx context.Context, id int) error {
	return p.repository.DeleteContact(ctx, id)
}
