package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/db"
)  

// AddressDAL : data access layer  (address) table.
type AddressDAL struct {
	AddressID int32 `json:"address_id" gorm:"column:address_id;primary_key:true"`
	Address string `json:"address" gorm:"column:address"`
	Address2 string `json:"address2" gorm:"column:address2"`
	District string `json:"district" gorm:"column:district"`
	CityID int16 `json:"city_id" gorm:"column:city_id"`
	PostalCode string `json:"postal_code" gorm:"column:postal_code"`
	Phone string `json:"phone" gorm:"column:phone"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (a *AddressDAL) TableName() string {
	return "address"
} 

// GetAllAddresses : get all addresses.
func GetAllAddresses() []*AddressDAL {
	addresses := []*AddressDAL{}
	db.DB().Find(&addresses)
	return addresses
}

// GetAddress : get one address by id.
func GetAddress(id int32) (*AddressDAL, error) {
	a := &AddressDAL{}
	result := db.DB().First(a, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}


// CreateAddress : create new address.
func CreateAddress(a *AddressDAL) (*AddressDAL, error) {
	result := db.DB().Create(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// UpdateAddress : update exist address.
func UpdateAddress(a *AddressDAL) (*AddressDAL, error) {
	_, err := GetAddress(a.AddressID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// DeleteAddress : delete address by id.
func DeleteAddress(id int32) error {
	a, err := GetAddress(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(a)
	return result.Error
}


