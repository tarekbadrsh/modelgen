package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertStoreID : covnert StoreID string to StoreID int32.
func ConvertStoreID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllStores : get All stores.
func GetAllStores() ([]*dto.StoreDTO, error) {
	stores := dal.GetAllStores()
	return dto.StoreDALToDTOArr(stores)
}

// GetStore : get one store by id.
func GetStore(id int32) (*dto.StoreDTO, error) {
	s, err := dal.GetStore(id)
	if err != nil {
		return nil, err
	}
	return dto.StoreDALToDTO(s)
}


// CreateStore : create new store.
func CreateStore(s *dto.StoreDTO) (*dto.StoreDTO, error) {
	store, err := s.StoreDTOToDAL()
	if err != nil {
		return nil, err
	}
	newstore, err := dal.CreateStore(store)
	if err != nil {
		return nil, err
	}
	return dto.StoreDALToDTO(newstore)
}

// UpdateStore : update exist store.
func UpdateStore(s *dto.StoreDTO) (*dto.StoreDTO, error) {
	store, err := s.StoreDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatestore, err := dal.UpdateStore(store)
	if err != nil {
		return nil, err
	}
	return dto.StoreDALToDTO(updatestore)
}

// DeleteStore : delete store by id.
func DeleteStore(id int32) error {
	return dal.DeleteStore(id)
}


