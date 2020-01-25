package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dal"
)  

// CountryDTO : data transfer object  (country) table.
type CountryDTO struct {
	CountryID int32 `json:"country_id"`
	Country string `json:"country"`
	LastUpdate time.Time `json:"last_update"`
	
}

// CountryDTOToDAL : convert CountryDTO to CountryDAL
func (a *CountryDTO) CountryDTOToDAL() (*dal.CountryDAL, error) { 
	country := &dal.CountryDAL{
		CountryID:a.CountryID,
		Country:a.Country,
		LastUpdate:a.LastUpdate,
		 
	}
	return country, nil
}

// CountryDALToDTO : convert CountryDAL to CountryDTO
func CountryDALToDTO(a *dal.CountryDAL) (*CountryDTO, error) { 
	country := &CountryDTO{
		CountryID:a.CountryID,
		Country:a.Country,
		LastUpdate:a.LastUpdate,
		 
	}
	return country, nil
}

// CountryDALToDTOArr : convert Array of CountryDAL to Array of CountryDTO
func CountryDALToDTOArr(countries []*dal.CountryDAL) ([]*CountryDTO, error) {
	var err error
	res := make([]*CountryDTO, len(countries))
	for i, country := range countries {
		res[i], err = CountryDALToDTO(country)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


