package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/db"
)  

// StoreDAL : data access layer  (store) table.
type StoreDAL struct {
	StoreID int32 `json:"store_id" gorm:"column:store_id;primary_key:true"`
	ManagerStaffID int16 `json:"manager_staff_id" gorm:"column:manager_staff_id"`
	AddressID int16 `json:"address_id" gorm:"column:address_id"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (s *StoreDAL) TableName() string {
	return "store"
} 

// GetAllStores : get all stores.
func GetAllStores() []*StoreDAL {
	stores := []*StoreDAL{}
	db.DB().Find(&stores)
	return stores
}

// GetStore : get one store by id.
func GetStore(id int32) (*StoreDAL, error) {
	s := &StoreDAL{}
	result := db.DB().First(s, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}


// CreateStore : create new store.
func CreateStore(s *StoreDAL) (*StoreDAL, error) {
	result := db.DB().Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

// UpdateStore : update exist store.
func UpdateStore(s *StoreDAL) (*StoreDAL, error) {
	_, err := GetStore(s.StoreID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

// DeleteStore : delete store by id.
func DeleteStore(id int32) error {
	s, err := GetStore(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(s)
	return result.Error
}


