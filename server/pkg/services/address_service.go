package services

import "jatin/pkg/schemas"

type AddressServiceTemplate interface {
	CreateAddress() (*schemas.Address, error)
	UpdateAddress() (*schemas.Address, error)
	DeleteAddress() (*schemas.Address, error)
	GetAddressesByUserId() (*[]schemas.Address, error)
	GetAddressByAddressId() (*schemas.Address, error)
}
