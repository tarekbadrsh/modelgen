package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
)  

// StoreDTO : data transfer object  (store) table.
type StoreDTO struct {
	StoreID int32 `json:"store_id"`
	ManagerStaffID int16 `json:"manager_staff_id"`
	AddressID int16 `json:"address_id"`
	LastUpdate time.Time `json:"last_update"`
	
}

// StoreDTOToDAL : convert StoreDTO to StoreDAL
func (a *StoreDTO) StoreDTOToDAL() (*dal.StoreDAL, error) { 
	store := &dal.StoreDAL{
		StoreID:a.StoreID,
		ManagerStaffID:a.ManagerStaffID,
		AddressID:a.AddressID,
		LastUpdate:a.LastUpdate,
		 
	}
	return store, nil
}

// StoreDALToDTO : convert StoreDAL to StoreDTO
func StoreDALToDTO(a *dal.StoreDAL) (*StoreDTO, error) { 
	store := &StoreDTO{
		StoreID:a.StoreID,
		ManagerStaffID:a.ManagerStaffID,
		AddressID:a.AddressID,
		LastUpdate:a.LastUpdate,
		 
	}
	return store, nil
}

// StoreDALToDTOArr : convert Array of StoreDAL to Array of StoreDTO
func StoreDALToDTOArr(stores []*dal.StoreDAL) ([]*StoreDTO, error) {
	var err error
	res := make([]*StoreDTO, len(stores))
	for i, store := range stores {
		res[i], err = StoreDALToDTO(store)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


