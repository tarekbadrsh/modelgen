package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)  

// CategoryDAL : data access layer  (category) table.
type CategoryDAL struct {
	CategoryID int32 `json:"category_id" gorm:"column:category_id;primary_key:true"`
	Name string `json:"name" gorm:"column:name"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (c *CategoryDAL) TableName() string {
	return "category"
} 

// GetAllCategories : get all categories.
func GetAllCategories() []*CategoryDAL {
	categories := []*CategoryDAL{}
	db.DB().Find(&categories)
	return categories
}

// GetCategory : get one category by id.
func GetCategory(id int32) (*CategoryDAL, error) {
	c := &CategoryDAL{}
	result := db.DB().First(c, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}


// CreateCategory : create new category.
func CreateCategory(c *CategoryDAL) (*CategoryDAL, error) {
	result := db.DB().Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// UpdateCategory : update exist category.
func UpdateCategory(c *CategoryDAL) (*CategoryDAL, error) {
	_, err := GetCategory(c.CategoryID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

// DeleteCategory : delete category by id.
func DeleteCategory(id int32) error {
	c, err := GetCategory(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(c)
	return result.Error
}


