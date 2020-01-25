package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
)  

// CustomerDTO : data transfer object  (customer) table.
type CustomerDTO struct {
	CustomerID int32 `json:"customer_id"`
	StoreID int16 `json:"store_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	AddressID int16 `json:"address_id"`
	Activebool bool `json:"activebool"`
	CreateDate time.Time `json:"create_date"`
	LastUpdate time.Time `json:"last_update"`
	Active int32 `json:"active"`
	
}

// CustomerDTOToDAL : convert CustomerDTO to CustomerDAL
func (a *CustomerDTO) CustomerDTOToDAL() (*dal.CustomerDAL, error) { 
	customer := &dal.CustomerDAL{
		CustomerID:a.CustomerID,
		StoreID:a.StoreID,
		FirstName:a.FirstName,
		LastName:a.LastName,
		Email:a.Email,
		AddressID:a.AddressID,
		Activebool:a.Activebool,
		CreateDate:a.CreateDate,
		LastUpdate:a.LastUpdate,
		Active:a.Active,
		 
	}
	return customer, nil
}

// CustomerDALToDTO : convert CustomerDAL to CustomerDTO
func CustomerDALToDTO(a *dal.CustomerDAL) (*CustomerDTO, error) { 
	customer := &CustomerDTO{
		CustomerID:a.CustomerID,
		StoreID:a.StoreID,
		FirstName:a.FirstName,
		LastName:a.LastName,
		Email:a.Email,
		AddressID:a.AddressID,
		Activebool:a.Activebool,
		CreateDate:a.CreateDate,
		LastUpdate:a.LastUpdate,
		Active:a.Active,
		 
	}
	return customer, nil
}

// CustomerDALToDTOArr : convert Array of CustomerDAL to Array of CustomerDTO
func CustomerDALToDTOArr(customers []*dal.CustomerDAL) ([]*CustomerDTO, error) {
	var err error
	res := make([]*CustomerDTO, len(customers))
	for i, customer := range customers {
		res[i], err = CustomerDALToDTO(customer)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


