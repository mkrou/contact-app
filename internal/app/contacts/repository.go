package contacts

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func NewContactPostgresRepository(db *pg.DB) *ContactPostgresRepository {
	return &ContactPostgresRepository{
		db: db,
	}
}

type ContactPostgresRepository struct {
	db *pg.DB
}

//GetContacts will return all contact records
func (cpr *ContactPostgresRepository) GetContacts(ctx context.Context) ([]Contact, error) {
	var contacts = []Contact{}

	err := cpr.db.ModelContext(ctx, &contacts).Order("id asc").Select()
	if err != nil {
		return nil, fmt.Errorf("postgres repository: get contacts: %w", err)
	}

	return contacts, nil
}

//AddContact will create new contact record
func (cpr *ContactPostgresRepository) AddContact(ctx context.Context, contact Contact) (Contact, error) {
	if _, err := cpr.db.ModelContext(ctx, &contact).Insert(); err != nil {
		return Contact{}, fmt.Errorf("postgres repository: create contact: %w", err)
	}

	return contact, nil
}

//EditContact will edit a contact record by ID
func (cpr *ContactPostgresRepository) EditContact(ctx context.Context, contact Contact) (Contact, error) {
	result, err := cpr.db.ModelContext(ctx, &contact).WherePK().Update()
	if err != nil {
		return Contact{}, fmt.Errorf("postgres repository: edit contact: %w", err)
	}

	if result.RowsAffected() == 0 {
		return Contact{}, nil
	}

	return contact, nil
}

//DeleteContact will delete a contact record by ID
func (cpr *ContactPostgresRepository) DeleteContact(ctx context.Context, id int) error {
	contact := Contact{
		ID: id,
	}

	if _, err := cpr.db.ModelContext(ctx, &contact).WherePK().Delete(); err != nil {
		return fmt.Errorf("postgres repository: delete contact: %w", err)
	}

	return nil
}
