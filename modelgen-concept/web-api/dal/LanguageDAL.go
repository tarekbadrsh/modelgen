package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)  

// LanguageDAL : data access layer  (language) table.
type LanguageDAL struct {
	LanguageID int32 `json:"language_id" gorm:"column:language_id;primary_key:true"`
	Name interface {} `json:"name" gorm:"column:name"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
func (l *LanguageDAL) TableName() string {
	return "language"
} 

// GetAllLanguages : get all languages.
func GetAllLanguages() []*LanguageDAL {
	languages := []*LanguageDAL{}
	db.DB().Find(&languages)
	return languages
}

// GetLanguage : get one language by id.
func GetLanguage(id int32) (*LanguageDAL, error) {
	l := &LanguageDAL{}
	result := db.DB().First(l, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return l, nil
}


// CreateLanguage : create new language.
func CreateLanguage(l *LanguageDAL) (*LanguageDAL, error) {
	result := db.DB().Create(l)
	if result.Error != nil {
		return nil, result.Error
	}
	return l, nil
}

// UpdateLanguage : update exist language.
func UpdateLanguage(l *LanguageDAL) (*LanguageDAL, error) {
	_, err := GetLanguage(l.LanguageID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(l)
	if result.Error != nil {
		return nil, result.Error
	}
	return l, nil
}

// DeleteLanguage : delete language by id.
func DeleteLanguage(id int32) error {
	l, err := GetLanguage(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(l)
	return result.Error
}


