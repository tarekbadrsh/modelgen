package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dto"
)


// ConvertPaymentID : covnert PaymentID string to PaymentID int32.
func ConvertPaymentID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllPayments : get All payments.
func GetAllPayments() ([]*dto.PaymentDTO, error) {
	payments := dal.GetAllPayments()
	return dto.PaymentDALToDTOArr(payments)
}

// GetPayment : get one payment by id.
func GetPayment(id int32) (*dto.PaymentDTO, error) {
	p, err := dal.GetPayment(id)
	if err != nil {
		return nil, err
	}
	return dto.PaymentDALToDTO(p)
}


// CreatePayment : create new payment.
func CreatePayment(p *dto.PaymentDTO) (*dto.PaymentDTO, error) {
	payment, err := p.PaymentDTOToDAL()
	if err != nil {
		return nil, err
	}
	newpayment, err := dal.CreatePayment(payment)
	if err != nil {
		return nil, err
	}
	return dto.PaymentDALToDTO(newpayment)
}

// UpdatePayment : update exist payment.
func UpdatePayment(p *dto.PaymentDTO) (*dto.PaymentDTO, error) {
	payment, err := p.PaymentDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatepayment, err := dal.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}
	return dto.PaymentDALToDTO(updatepayment)
}

// DeletePayment : delete payment by id.
func DeletePayment(id int32) error {
	return dal.DeletePayment(id)
}


