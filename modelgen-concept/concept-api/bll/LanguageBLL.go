package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)


// ConvertLanguageID : covnert LanguageID string to LanguageID int32.
func ConvertLanguageID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllLanguages : get All languages.
func GetAllLanguages() ([]*dto.LanguageDTO, error) {
	languages := dal.GetAllLanguages()
	return dto.LanguageDALToDTOArr(languages)
}

// GetLanguage : get one language by id.
func GetLanguage(id int32) (*dto.LanguageDTO, error) {
	l, err := dal.GetLanguage(id)
	if err != nil {
		return nil, err
	}
	return dto.LanguageDALToDTO(l)
}


// CreateLanguage : create new language.
func CreateLanguage(l *dto.LanguageDTO) (*dto.LanguageDTO, error) {
	language, err := l.LanguageDTOToDAL()
	if err != nil {
		return nil, err
	}
	newlanguage, err := dal.CreateLanguage(language)
	if err != nil {
		return nil, err
	}
	return dto.LanguageDALToDTO(newlanguage)
}

// UpdateLanguage : update exist language.
func UpdateLanguage(l *dto.LanguageDTO) (*dto.LanguageDTO, error) {
	language, err := l.LanguageDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatelanguage, err := dal.UpdateLanguage(language)
	if err != nil {
		return nil, err
	}
	return dto.LanguageDALToDTO(updatelanguage)
}

// DeleteLanguage : delete language by id.
func DeleteLanguage(id int32) error {
	return dal.DeleteLanguage(id)
}


