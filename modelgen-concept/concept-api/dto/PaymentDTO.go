package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
)  

// PaymentDTO : data transfer object  (payment) table.
type PaymentDTO struct {
	PaymentID int32 `json:"payment_id"`
	CustomerID int16 `json:"customer_id"`
	StaffID int16 `json:"staff_id"`
	RentalID int32 `json:"rental_id"`
	Amount interface {} `json:"amount"`
	PaymentDate time.Time `json:"payment_date"`
	
}

// PaymentDTOToDAL : convert PaymentDTO to PaymentDAL
func (a *PaymentDTO) PaymentDTOToDAL() (*dal.PaymentDAL, error) { 
	payment := &dal.PaymentDAL{
		PaymentID:a.PaymentID,
		CustomerID:a.CustomerID,
		StaffID:a.StaffID,
		RentalID:a.RentalID,
		Amount:a.Amount,
		PaymentDate:a.PaymentDate,
		 
	}
	return payment, nil
}

// PaymentDALToDTO : convert PaymentDAL to PaymentDTO
func PaymentDALToDTO(a *dal.PaymentDAL) (*PaymentDTO, error) { 
	payment := &PaymentDTO{
		PaymentID:a.PaymentID,
		CustomerID:a.CustomerID,
		StaffID:a.StaffID,
		RentalID:a.RentalID,
		Amount:a.Amount,
		PaymentDate:a.PaymentDate,
		 
	}
	return payment, nil
}

// PaymentDALToDTOArr : convert Array of PaymentDAL to Array of PaymentDTO
func PaymentDALToDTOArr(payments []*dal.PaymentDAL) ([]*PaymentDTO, error) {
	var err error
	res := make([]*PaymentDTO, len(payments))
	for i, payment := range payments {
		res[i], err = PaymentDALToDTO(payment)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


