package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
)  

// AddressDTO : data transfer object  (address) table.
type AddressDTO struct {
	AddressID int32 `json:"address_id"`
	Address string `json:"address"`
	Address2 string `json:"address2"`
	District string `json:"district"`
	CityID int16 `json:"city_id"`
	PostalCode string `json:"postal_code"`
	Phone string `json:"phone"`
	LastUpdate time.Time `json:"last_update"`
	
}

// AddressDTOToDAL : convert AddressDTO to AddressDAL
func (a *AddressDTO) AddressDTOToDAL() (*dal.AddressDAL, error) { 
	address := &dal.AddressDAL{
		AddressID:a.AddressID,
		Address:a.Address,
		Address2:a.Address2,
		District:a.District,
		CityID:a.CityID,
		PostalCode:a.PostalCode,
		Phone:a.Phone,
		LastUpdate:a.LastUpdate,
		 
	}
	return address, nil
}

// AddressDALToDTO : convert AddressDAL to AddressDTO
func AddressDALToDTO(a *dal.AddressDAL) (*AddressDTO, error) { 
	address := &AddressDTO{
		AddressID:a.AddressID,
		Address:a.Address,
		Address2:a.Address2,
		District:a.District,
		CityID:a.CityID,
		PostalCode:a.PostalCode,
		Phone:a.Phone,
		LastUpdate:a.LastUpdate,
		 
	}
	return address, nil
}

// AddressDALToDTOArr : convert Array of AddressDAL to Array of AddressDTO
func AddressDALToDTOArr(addresses []*dal.AddressDAL) ([]*AddressDTO, error) {
	var err error
	res := make([]*AddressDTO, len(addresses))
	for i, address := range addresses {
		res[i], err = AddressDALToDTO(address)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


