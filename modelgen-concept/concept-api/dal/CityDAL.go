package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/db"
)  

// CityDAL : data access layer  (city) table.
type CityDAL struct {
	CityID int32 `json:"city_id" gorm:"column:city_id;primary_key:true"`
	City string `json:"city" gorm:"column:city"`
	CountryID int16 `json:"country_id" gorm:"column:country_id"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (c *CityDAL) TableName() string {
	return "city"
} 

// GetAllCities : get all cities.
func GetAllCities() []*CityDAL {
	cities := []*CityDAL{}
	db.DB().Find(&cities)
	return cities
}

// GetCity : get one city by id.
func GetCity(id int32) (*CityDAL, error) {
	c := &CityDAL{}
	result := db.DB().First(c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}


// CreateCity : create new city.
func CreateCity(c *CityDAL) (*CityDAL, error) {
	result := db.DB().Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// UpdateCity : update exist city.
func UpdateCity(c *CityDAL) (*CityDAL, error) {
	_, err := GetCity(c.CityID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// DeleteCity : delete city by id.
func DeleteCity(id int32) error {
	c, err := GetCity(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(c)
	return result.Error
}


