package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/web-api/db"
)  

// CountryDAL : data access layer  (country) table.
type CountryDAL struct {
	CountryID int32 `json:"country_id" gorm:"column:country_id;primary_key:true"`
	Country string `json:"country" gorm:"column:country"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (c *CountryDAL) TableName() string {
	return "country"
} 

// GetAllCountries : get all countries.
func GetAllCountries() []*CountryDAL {
	countries := []*CountryDAL{}
	db.DB().Find(&countries)
	return countries
}

// GetCountry : get one country by id.
func GetCountry(id int32) (*CountryDAL, error) {
	c := &CountryDAL{}
	result := db.DB().First(c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}


// CreateCountry : create new country.
func CreateCountry(c *CountryDAL) (*CountryDAL, error) {
	result := db.DB().Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// UpdateCountry : update exist country.
func UpdateCountry(c *CountryDAL) (*CountryDAL, error) {
	_, err := GetCountry(c.CountryID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// DeleteCountry : delete country by id.
func DeleteCountry(id int32) error {
	c, err := GetCountry(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(c)
	return result.Error
}


