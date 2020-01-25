package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dal"
)  

// CategoryDTO : data transfer object  (category) table.
type CategoryDTO struct {
	CategoryID int32 `json:"category_id"`
	Name string `json:"name"`
	LastUpdate time.Time `json:"last_update"`
	
}

// CategoryDTOToDAL : convert CategoryDTO to CategoryDAL
func (a *CategoryDTO) CategoryDTOToDAL() (*dal.CategoryDAL, error) { 
	category := &dal.CategoryDAL{
		CategoryID:a.CategoryID,
		Name:a.Name,
		LastUpdate:a.LastUpdate,
		 
	}
	return category, nil
}

// CategoryDALToDTO : convert CategoryDAL to CategoryDTO
func CategoryDALToDTO(a *dal.CategoryDAL) (*CategoryDTO, error) { 
	category := &CategoryDTO{
		CategoryID:a.CategoryID,
		Name:a.Name,
		LastUpdate:a.LastUpdate,
		 
	}
	return category, nil
}

// CategoryDALToDTOArr : convert Array of CategoryDAL to Array of CategoryDTO
func CategoryDALToDTOArr(categories []*dal.CategoryDAL) ([]*CategoryDTO, error) {
	var err error
	res := make([]*CategoryDTO, len(categories))
	for i, category := range categories {
		res[i], err = CategoryDALToDTO(category)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


