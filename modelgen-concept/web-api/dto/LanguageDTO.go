package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dal"
)  

// LanguageDTO : data transfer object  (language) table.
type LanguageDTO struct {
	LanguageID int32 `json:"language_id"`
	Name interface {} `json:"name"`
	LastUpdate time.Time `json:"last_update"`
	
}

// LanguageDTOToDAL : convert LanguageDTO to LanguageDAL
func (a *LanguageDTO) LanguageDTOToDAL() (*dal.LanguageDAL, error) { 
	language := &dal.LanguageDAL{
		LanguageID:a.LanguageID,
		Name:a.Name,
		LastUpdate:a.LastUpdate,
		 
	}
	return language, nil
}

// LanguageDALToDTO : convert LanguageDAL to LanguageDTO
func LanguageDALToDTO(a *dal.LanguageDAL) (*LanguageDTO, error) { 
	language := &LanguageDTO{
		LanguageID:a.LanguageID,
		Name:a.Name,
		LastUpdate:a.LastUpdate,
		 
	}
	return language, nil
}

// LanguageDALToDTOArr : convert Array of LanguageDAL to Array of LanguageDTO
func LanguageDALToDTOArr(languages []*dal.LanguageDAL) ([]*LanguageDTO, error) {
	var err error
	res := make([]*LanguageDTO, len(languages))
	for i, language := range languages {
		res[i], err = LanguageDALToDTO(language)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


