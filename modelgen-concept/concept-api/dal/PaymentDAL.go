package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/db"
)  

// PaymentDAL : data access layer  (payment) table.
type PaymentDAL struct {
	PaymentID int32 `json:"payment_id" gorm:"column:payment_id;primary_key:true"`
	CustomerID int16 `json:"customer_id" gorm:"column:customer_id"`
	StaffID int16 `json:"staff_id" gorm:"column:staff_id"`
	RentalID int32 `json:"rental_id" gorm:"column:rental_id"`
	Amount interface {} `json:"amount" gorm:"column:amount"`
	PaymentDate time.Time `json:"payment_date" gorm:"column:payment_date"`
	
}

// TableName sets the insert table name for this struct type
func (p *PaymentDAL) TableName() string {
	return "payment"
} 

// GetAllPayments : get all payments.
func GetAllPayments() []*PaymentDAL {
	payments := []*PaymentDAL{}
	db.DB().Find(&payments)
	return payments
}

// GetPayment : get one payment by id.
func GetPayment(id int32) (*PaymentDAL, error) {
	p := &PaymentDAL{}
	result := db.DB().First(p, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}


// CreatePayment : create new payment.
func CreatePayment(p *PaymentDAL) (*PaymentDAL, error) {
	result := db.DB().Create(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

// UpdatePayment : update exist payment.
func UpdatePayment(p *PaymentDAL) (*PaymentDAL, error) {
	_, err := GetPayment(p.PaymentID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

// DeletePayment : delete payment by id.
func DeletePayment(id int32) error {
	p, err := GetPayment(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(p)
	return result.Error
}


