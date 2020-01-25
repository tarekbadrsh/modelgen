package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)  

// StaffDAL : data access layer  (staff) table.
type StaffDAL struct {
	StaffID int32 `json:"staff_id" gorm:"column:staff_id;primary_key:true"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName string `json:"last_name" gorm:"column:last_name"`
	AddressID int16 `json:"address_id" gorm:"column:address_id"`
	Email string `json:"email" gorm:"column:email"`
	StoreID int16 `json:"store_id" gorm:"column:store_id"`
	Active bool `json:"active" gorm:"column:active"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	Picture []uint8 `json:"picture" gorm:"column:picture"`
	
}

// TableName sets the insert table name for this struct type
func (s *StaffDAL) TableName() string {
	return "staff"
} 

// GetAllStaffs : get all staffs.
func GetAllStaffs() []*StaffDAL {
	staffs := []*StaffDAL{}
	db.DB().Find(&staffs)
	return staffs
}

// GetStaff : get one staff by id.
func GetStaff(id int32) (*StaffDAL, error) {
	s := &StaffDAL{}
	result := db.DB().First(s, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}


// CreateStaff : create new staff.
func CreateStaff(s *StaffDAL) (*StaffDAL, error) {
	result := db.DB().Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

// UpdateStaff : update exist staff.
func UpdateStaff(s *StaffDAL) (*StaffDAL, error) {
	_, err := GetStaff(s.StaffID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

// DeleteStaff : delete staff by id.
func DeleteStaff(id int32) error {
	s, err := GetStaff(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(s)
	return result.Error
}


