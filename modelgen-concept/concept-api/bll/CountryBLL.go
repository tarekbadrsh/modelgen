package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)


// ConvertCountryID : covnert CountryID string to CountryID int32.
func ConvertCountryID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllCountries : get All countries.
func GetAllCountries() ([]*dto.CountryDTO, error) {
	countries := dal.GetAllCountries()
	return dto.CountryDALToDTOArr(countries)
}

// GetCountry : get one country by id.
func GetCountry(id int32) (*dto.CountryDTO, error) {
	c, err := dal.GetCountry(id)
	if err != nil {
		return nil, err
	}
	return dto.CountryDALToDTO(c)
}


// CreateCountry : create new country.
func CreateCountry(c *dto.CountryDTO) (*dto.CountryDTO, error) {
	country, err := c.CountryDTOToDAL()
	if err != nil {
		return nil, err
	}
	newcountry, err := dal.CreateCountry(country)
	if err != nil {
		return nil, err
	}
	return dto.CountryDALToDTO(newcountry)
}

// UpdateCountry : update exist country.
func UpdateCountry(c *dto.CountryDTO) (*dto.CountryDTO, error) {
	country, err := c.CountryDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatecountry, err := dal.UpdateCountry(country)
	if err != nil {
		return nil, err
	}
	return dto.CountryDALToDTO(updatecountry)
}

// DeleteCountry : delete country by id.
func DeleteCountry(id int32) error {
	return dal.DeleteCountry(id)
}


