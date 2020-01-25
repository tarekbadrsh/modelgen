package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/db"
)  

// RentalDAL : data access layer  (rental) table.
type RentalDAL struct {
	RentalID int32 `json:"rental_id" gorm:"column:rental_id;primary_key:true"`
	RentalDate time.Time `json:"rental_date" gorm:"column:rental_date"`
	InventoryID int32 `json:"inventory_id" gorm:"column:inventory_id"`
	CustomerID int16 `json:"customer_id" gorm:"column:customer_id"`
	ReturnDate time.Time `json:"return_date" gorm:"column:return_date"`
	StaffID int16 `json:"staff_id" gorm:"column:staff_id"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (r *RentalDAL) TableName() string {
	return "rental"
} 

// GetAllRentals : get all rentals.
func GetAllRentals() []*RentalDAL {
	rentals := []*RentalDAL{}
	db.DB().Find(&rentals)
	return rentals
}

// GetRental : get one rental by id.
func GetRental(id int32) (*RentalDAL, error) {
	r := &RentalDAL{}
	result := db.DB().First(r, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return r, nil
}


// CreateRental : create new rental.
func CreateRental(r *RentalDAL) (*RentalDAL, error) {
	result := db.DB().Create(r)
	if result.Error != nil {
		return nil, result.Error
	}
	return r, nil
}

// UpdateRental : update exist rental.
func UpdateRental(r *RentalDAL) (*RentalDAL, error) {
	_, err := GetRental(r.RentalID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(r)
	if result.Error != nil {
		return nil, result.Error
	}
	return r, nil
}

// DeleteRental : delete rental by id.
func DeleteRental(id int32) error {
	r, err := GetRental(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(r)
	return result.Error
}


