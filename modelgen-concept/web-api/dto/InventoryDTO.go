package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
)  

// InventoryDTO : data transfer object  (inventory) table.
type InventoryDTO struct {
	InventoryID int32 `json:"inventory_id"`
	FilmID int16 `json:"film_id"`
	StoreID int16 `json:"store_id"`
	LastUpdate time.Time `json:"last_update"`
	
}

// InventoryDTOToDAL : convert InventoryDTO to InventoryDAL
func (a *InventoryDTO) InventoryDTOToDAL() (*dal.InventoryDAL, error) { 
	inventory := &dal.InventoryDAL{
		InventoryID:a.InventoryID,
		FilmID:a.FilmID,
		StoreID:a.StoreID,
		LastUpdate:a.LastUpdate,
		 
	}
	return inventory, nil
}

// InventoryDALToDTO : convert InventoryDAL to InventoryDTO
func InventoryDALToDTO(a *dal.InventoryDAL) (*InventoryDTO, error) { 
	inventory := &InventoryDTO{
		InventoryID:a.InventoryID,
		FilmID:a.FilmID,
		StoreID:a.StoreID,
		LastUpdate:a.LastUpdate,
		 
	}
	return inventory, nil
}

// InventoryDALToDTOArr : convert Array of InventoryDAL to Array of InventoryDTO
func InventoryDALToDTOArr(inventories []*dal.InventoryDAL) ([]*InventoryDTO, error) {
	var err error
	res := make([]*InventoryDTO, len(inventories))
	for i, inventory := range inventories {
		res[i], err = InventoryDALToDTO(inventory)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


