package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)  

// CustomerDAL : data access layer  (customer) table.
type CustomerDAL struct {
	CustomerID int32 `json:"customer_id" gorm:"column:customer_id;primary_key:true"`
	StoreID int16 `json:"store_id" gorm:"column:store_id"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName string `json:"last_name" gorm:"column:last_name"`
	Email string `json:"email" gorm:"column:email"`
	AddressID int16 `json:"address_id" gorm:"column:address_id"`
	Activebool bool `json:"activebool" gorm:"column:activebool"`
	CreateDate time.Time `json:"create_date" gorm:"column:create_date"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	Active int32 `json:"active" gorm:"column:active"`
	
}

// TableName sets the insert table name for this struct type
func (c *CustomerDAL) TableName() string {
	return "customer"
} 

// GetAllCustomers : get all customers.
func GetAllCustomers() []*CustomerDAL {
	customers := []*CustomerDAL{}
	db.DB().Find(&customers)
	return customers
}

// GetCustomer : get one customer by id.
func GetCustomer(id int32) (*CustomerDAL, error) {
	c := &CustomerDAL{}
	result := db.DB().First(c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}


// CreateCustomer : create new customer.
func CreateCustomer(c *CustomerDAL) (*CustomerDAL, error) {
	result := db.DB().Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// UpdateCustomer : update exist customer.
func UpdateCustomer(c *CustomerDAL) (*CustomerDAL, error) {
	_, err := GetCustomer(c.CustomerID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// DeleteCustomer : delete customer by id.
func DeleteCustomer(id int32) error {
	c, err := GetCustomer(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(c)
	return result.Error
}


