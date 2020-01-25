package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertAddressID : covnert AddressID string to AddressID int32.
func ConvertAddressID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllAddresses : get All addresses.
func GetAllAddresses() ([]*dto.AddressDTO, error) {
	addresses := dal.GetAllAddresses()
	return dto.AddressDALToDTOArr(addresses)
}

// GetAddress : get one address by id.
func GetAddress(id int32) (*dto.AddressDTO, error) {
	a, err := dal.GetAddress(id)
	if err != nil {
		return nil, err
	}
	return dto.AddressDALToDTO(a)
}


// CreateAddress : create new address.
func CreateAddress(a *dto.AddressDTO) (*dto.AddressDTO, error) {
	address, err := a.AddressDTOToDAL()
	if err != nil {
		return nil, err
	}
	newaddress, err := dal.CreateAddress(address)
	if err != nil {
		return nil, err
	}
	return dto.AddressDALToDTO(newaddress)
}

// UpdateAddress : update exist address.
func UpdateAddress(a *dto.AddressDTO) (*dto.AddressDTO, error) {
	address, err := a.AddressDTOToDAL()
	if err != nil {
		return nil, err
	}
	updateaddress, err := dal.UpdateAddress(address)
	if err != nil {
		return nil, err
	}
	return dto.AddressDALToDTO(updateaddress)
}

// DeleteAddress : delete address by id.
func DeleteAddress(id int32) error {
	return dal.DeleteAddress(id)
}


