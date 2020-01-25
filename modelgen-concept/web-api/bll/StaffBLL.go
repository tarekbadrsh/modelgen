package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertStaffID : covnert StaffID string to StaffID int32.
func ConvertStaffID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllStaffs : get All staffs.
func GetAllStaffs() ([]*dto.StaffDTO, error) {
	staffs := dal.GetAllStaffs()
	return dto.StaffDALToDTOArr(staffs)
}

// GetStaff : get one staff by id.
func GetStaff(id int32) (*dto.StaffDTO, error) {
	s, err := dal.GetStaff(id)
	if err != nil {
		return nil, err
	}
	return dto.StaffDALToDTO(s)
}


// CreateStaff : create new staff.
func CreateStaff(s *dto.StaffDTO) (*dto.StaffDTO, error) {
	staff, err := s.StaffDTOToDAL()
	if err != nil {
		return nil, err
	}
	newstaff, err := dal.CreateStaff(staff)
	if err != nil {
		return nil, err
	}
	return dto.StaffDALToDTO(newstaff)
}

// UpdateStaff : update exist staff.
func UpdateStaff(s *dto.StaffDTO) (*dto.StaffDTO, error) {
	staff, err := s.StaffDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatestaff, err := dal.UpdateStaff(staff)
	if err != nil {
		return nil, err
	}
	return dto.StaffDALToDTO(updatestaff)
}

// DeleteStaff : delete staff by id.
func DeleteStaff(id int32) error {
	return dal.DeleteStaff(id)
}


