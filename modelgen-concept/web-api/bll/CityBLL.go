package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertCityID : covnert CityID string to CityID int32.
func ConvertCityID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllCities : get All cities.
func GetAllCities() ([]*dto.CityDTO, error) {
	cities := dal.GetAllCities()
	return dto.CityDALToDTOArr(cities)
}

// GetCity : get one city by id.
func GetCity(id int32) (*dto.CityDTO, error) {
	c, err := dal.GetCity(id)
	if err != nil {
		return nil, err
	}
	return dto.CityDALToDTO(c)
}


// CreateCity : create new city.
func CreateCity(c *dto.CityDTO) (*dto.CityDTO, error) {
	city, err := c.CityDTOToDAL()
	if err != nil {
		return nil, err
	}
	newcity, err := dal.CreateCity(city)
	if err != nil {
		return nil, err
	}
	return dto.CityDALToDTO(newcity)
}

// UpdateCity : update exist city.
func UpdateCity(c *dto.CityDTO) (*dto.CityDTO, error) {
	city, err := c.CityDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatecity, err := dal.UpdateCity(city)
	if err != nil {
		return nil, err
	}
	return dto.CityDALToDTO(updatecity)
}

// DeleteCity : delete city by id.
func DeleteCity(id int32) error {
	return dal.DeleteCity(id)
}


