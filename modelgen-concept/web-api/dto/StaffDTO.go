package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
)  

// StaffDTO : data transfer object  (staff) table.
type StaffDTO struct {
	StaffID int32 `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	AddressID int16 `json:"address_id"`
	Email string `json:"email"`
	StoreID int16 `json:"store_id"`
	Active bool `json:"active"`
	Username string `json:"username"`
	Password string `json:"password"`
	LastUpdate time.Time `json:"last_update"`
	Picture []uint8 `json:"picture"`
	
}

// StaffDTOToDAL : convert StaffDTO to StaffDAL
func (a *StaffDTO) StaffDTOToDAL() (*dal.StaffDAL, error) { 
	staff := &dal.StaffDAL{
		StaffID:a.StaffID,
		FirstName:a.FirstName,
		LastName:a.LastName,
		AddressID:a.AddressID,
		Email:a.Email,
		StoreID:a.StoreID,
		Active:a.Active,
		Username:a.Username,
		Password:a.Password,
		LastUpdate:a.LastUpdate,
		Picture:a.Picture,
		 
	}
	return staff, nil
}

// StaffDALToDTO : convert StaffDAL to StaffDTO
func StaffDALToDTO(a *dal.StaffDAL) (*StaffDTO, error) { 
	staff := &StaffDTO{
		StaffID:a.StaffID,
		FirstName:a.FirstName,
		LastName:a.LastName,
		AddressID:a.AddressID,
		Email:a.Email,
		StoreID:a.StoreID,
		Active:a.Active,
		Username:a.Username,
		Password:a.Password,
		LastUpdate:a.LastUpdate,
		Picture:a.Picture,
		 
	}
	return staff, nil
}

// StaffDALToDTOArr : convert Array of StaffDAL to Array of StaffDTO
func StaffDALToDTOArr(staffs []*dal.StaffDAL) ([]*StaffDTO, error) {
	var err error
	res := make([]*StaffDTO, len(staffs))
	for i, staff := range staffs {
		res[i], err = StaffDALToDTO(staff)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


