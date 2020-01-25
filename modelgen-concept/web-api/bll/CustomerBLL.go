package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertCustomerID : covnert CustomerID string to CustomerID int32.
func ConvertCustomerID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllCustomers : get All customers.
func GetAllCustomers() ([]*dto.CustomerDTO, error) {
	customers := dal.GetAllCustomers()
	return dto.CustomerDALToDTOArr(customers)
}

// GetCustomer : get one customer by id.
func GetCustomer(id int32) (*dto.CustomerDTO, error) {
	c, err := dal.GetCustomer(id)
	if err != nil {
		return nil, err
	}
	return dto.CustomerDALToDTO(c)
}


// CreateCustomer : create new customer.
func CreateCustomer(c *dto.CustomerDTO) (*dto.CustomerDTO, error) {
	customer, err := c.CustomerDTOToDAL()
	if err != nil {
		return nil, err
	}
	newcustomer, err := dal.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return dto.CustomerDALToDTO(newcustomer)
}

// UpdateCustomer : update exist customer.
func UpdateCustomer(c *dto.CustomerDTO) (*dto.CustomerDTO, error) {
	customer, err := c.CustomerDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatecustomer, err := dal.UpdateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return dto.CustomerDALToDTO(updatecustomer)
}

// DeleteCustomer : delete customer by id.
func DeleteCustomer(id int32) error {
	return dal.DeleteCustomer(id)
}


