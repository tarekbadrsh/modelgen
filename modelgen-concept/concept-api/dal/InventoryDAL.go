package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/db"
)  

// InventoryDAL : data access layer  (inventory) table.
type InventoryDAL struct {
	InventoryID int32 `json:"inventory_id" gorm:"column:inventory_id;primary_key:true"`
	FilmID int16 `json:"film_id" gorm:"column:film_id"`
	StoreID int16 `json:"store_id" gorm:"column:store_id"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (i *InventoryDAL) TableName() string {
	return "inventory"
} 

// GetAllInventories : get all inventories.
func GetAllInventories() []*InventoryDAL {
	inventories := []*InventoryDAL{}
	db.DB().Find(&inventories)
	return inventories
}

// GetInventory : get one inventory by id.
func GetInventory(id int32) (*InventoryDAL, error) {
	i := &InventoryDAL{}
	result := db.DB().First(i, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}


// CreateInventory : create new inventory.
func CreateInventory(i *InventoryDAL) (*InventoryDAL, error) {
	result := db.DB().Create(i)
	if result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// UpdateInventory : update exist inventory.
func UpdateInventory(i *InventoryDAL) (*InventoryDAL, error) {
	_, err := GetInventory(i.InventoryID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(i)
	if result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// DeleteInventory : delete inventory by id.
func DeleteInventory(id int32) error {
	i, err := GetInventory(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(i)
	return result.Error
}


