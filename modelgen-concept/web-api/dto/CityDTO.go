package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
)  

// CityDTO : data transfer object  (city) table.
type CityDTO struct {
	CityID int32 `json:"city_id"`
	City string `json:"city"`
	CountryID int16 `json:"country_id"`
	LastUpdate time.Time `json:"last_update"`
	
}

// CityDTOToDAL : convert CityDTO to CityDAL
func (a *CityDTO) CityDTOToDAL() (*dal.CityDAL, error) { 
	city := &dal.CityDAL{
		CityID:a.CityID,
		City:a.City,
		CountryID:a.CountryID,
		LastUpdate:a.LastUpdate,
		 
	}
	return city, nil
}

// CityDALToDTO : convert CityDAL to CityDTO
func CityDALToDTO(a *dal.CityDAL) (*CityDTO, error) { 
	city := &CityDTO{
		CityID:a.CityID,
		City:a.City,
		CountryID:a.CountryID,
		LastUpdate:a.LastUpdate,
		 
	}
	return city, nil
}

// CityDALToDTOArr : convert Array of CityDAL to Array of CityDTO
func CityDALToDTOArr(cities []*dal.CityDAL) ([]*CityDTO, error) {
	var err error
	res := make([]*CityDTO, len(cities))
	for i, city := range cities {
		res[i], err = CityDALToDTO(city)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


