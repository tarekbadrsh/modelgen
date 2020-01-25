package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/web-api/dal"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)


// ConvertCategoryID : covnert CategoryID string to CategoryID int32.
func ConvertCategoryID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllCategories : get All categories.
func GetAllCategories() ([]*dto.CategoryDTO, error) {
	categories := dal.GetAllCategories()
	return dto.CategoryDALToDTOArr(categories)
}

// GetCategory : get one category by id.
func GetCategory(id int32) (*dto.CategoryDTO, error) {
	c, err := dal.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return dto.CategoryDALToDTO(c)
}


// CreateCategory : create new category.
func CreateCategory(c *dto.CategoryDTO) (*dto.CategoryDTO, error) {
	category, err := c.CategoryDTOToDAL()
	if err != nil {
		return nil, err
	}
	newcategory, err := dal.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return dto.CategoryDALToDTO(newcategory)
}

// UpdateCategory : update exist category.
func UpdateCategory(c *dto.CategoryDTO) (*dto.CategoryDTO, error) {
	category, err := c.CategoryDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatecategory, err := dal.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return dto.CategoryDALToDTO(updatecategory)
}

// DeleteCategory : delete category by id.
func DeleteCategory(id int32) error {
	return dal.DeleteCategory(id)
}


