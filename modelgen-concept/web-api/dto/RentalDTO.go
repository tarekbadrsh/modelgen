package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
)  

// RentalDTO : data transfer object  (rental) table.
type RentalDTO struct {
	RentalID int32 `json:"rental_id"`
	RentalDate time.Time `json:"rental_date"`
	InventoryID int32 `json:"inventory_id"`
	CustomerID int16 `json:"customer_id"`
	ReturnDate time.Time `json:"return_date"`
	StaffID int16 `json:"staff_id"`
	LastUpdate time.Time `json:"last_update"`
	
}

// RentalDTOToDAL : convert RentalDTO to RentalDAL
func (a *RentalDTO) RentalDTOToDAL() (*dal.RentalDAL, error) { 
	rental := &dal.RentalDAL{
		RentalID:a.RentalID,
		RentalDate:a.RentalDate,
		InventoryID:a.InventoryID,
		CustomerID:a.CustomerID,
		ReturnDate:a.ReturnDate,
		StaffID:a.StaffID,
		LastUpdate:a.LastUpdate,
		 
	}
	return rental, nil
}

// RentalDALToDTO : convert RentalDAL to RentalDTO
func RentalDALToDTO(a *dal.RentalDAL) (*RentalDTO, error) { 
	rental := &RentalDTO{
		RentalID:a.RentalID,
		RentalDate:a.RentalDate,
		InventoryID:a.InventoryID,
		CustomerID:a.CustomerID,
		ReturnDate:a.ReturnDate,
		StaffID:a.StaffID,
		LastUpdate:a.LastUpdate,
		 
	}
	return rental, nil
}

// RentalDALToDTOArr : convert Array of RentalDAL to Array of RentalDTO
func RentalDALToDTOArr(rentals []*dal.RentalDAL) ([]*RentalDTO, error) {
	var err error
	res := make([]*RentalDTO, len(rentals))
	for i, rental := range rentals {
		res[i], err = RentalDALToDTO(rental)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


