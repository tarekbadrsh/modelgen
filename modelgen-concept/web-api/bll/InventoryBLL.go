package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertInventoryID : covnert InventoryID string to InventoryID int32.
func ConvertInventoryID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllInventories : get All inventories.
func GetAllInventories() ([]*dto.InventoryDTO, error) {
	inventories := dal.GetAllInventories()
	return dto.InventoryDALToDTOArr(inventories)
}

// GetInventory : get one inventory by id.
func GetInventory(id int32) (*dto.InventoryDTO, error) {
	i, err := dal.GetInventory(id)
	if err != nil {
		return nil, err
	}
	return dto.InventoryDALToDTO(i)
}


// CreateInventory : create new inventory.
func CreateInventory(i *dto.InventoryDTO) (*dto.InventoryDTO, error) {
	inventory, err := i.InventoryDTOToDAL()
	if err != nil {
		return nil, err
	}
	newinventory, err := dal.CreateInventory(inventory)
	if err != nil {
		return nil, err
	}
	return dto.InventoryDALToDTO(newinventory)
}

// UpdateInventory : update exist inventory.
func UpdateInventory(i *dto.InventoryDTO) (*dto.InventoryDTO, error) {
	inventory, err := i.InventoryDTOToDAL()
	if err != nil {
		return nil, err
	}
	updateinventory, err := dal.UpdateInventory(inventory)
	if err != nil {
		return nil, err
	}
	return dto.InventoryDALToDTO(updateinventory)
}

// DeleteInventory : delete inventory by id.
func DeleteInventory(id int32) error {
	return dal.DeleteInventory(id)
}


