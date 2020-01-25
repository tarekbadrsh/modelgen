package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dto"
)


// ConvertRentalID : covnert RentalID string to RentalID int32.
func ConvertRentalID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllRentals : get All rentals.
func GetAllRentals() ([]*dto.RentalDTO, error) {
	rentals := dal.GetAllRentals()
	return dto.RentalDALToDTOArr(rentals)
}

// GetRental : get one rental by id.
func GetRental(id int32) (*dto.RentalDTO, error) {
	r, err := dal.GetRental(id)
	if err != nil {
		return nil, err
	}
	return dto.RentalDALToDTO(r)
}


// CreateRental : create new rental.
func CreateRental(r *dto.RentalDTO) (*dto.RentalDTO, error) {
	rental, err := r.RentalDTOToDAL()
	if err != nil {
		return nil, err
	}
	newrental, err := dal.CreateRental(rental)
	if err != nil {
		return nil, err
	}
	return dto.RentalDALToDTO(newrental)
}

// UpdateRental : update exist rental.
func UpdateRental(r *dto.RentalDTO) (*dto.RentalDTO, error) {
	rental, err := r.RentalDTOToDAL()
	if err != nil {
		return nil, err
	}
	updaterental, err := dal.UpdateRental(rental)
	if err != nil {
		return nil, err
	}
	return dto.RentalDALToDTO(updaterental)
}

// DeleteRental : delete rental by id.
func DeleteRental(id int32) error {
	return dal.DeleteRental(id)
}


