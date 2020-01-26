package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)


// ConvertFilmID : covnert FilmID string to FilmID int32.
func ConvertFilmID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}


// GetAllFilms : get All films.
func GetAllFilms() ([]*dto.FilmDTO, error) {
	films := dal.GetAllFilms()
	return dto.FilmDALToDTOArr(films)
}

// GetFilm : get one film by id.
func GetFilm(id int32) (*dto.FilmDTO, error) {
	f, err := dal.GetFilm(id)
	if err != nil {
		return nil, err
	}
	return dto.FilmDALToDTO(f)
}


// CreateFilm : create new film.
func CreateFilm(f *dto.FilmDTO) (*dto.FilmDTO, error) {
	film, err := f.FilmDTOToDAL()
	if err != nil {
		return nil, err
	}
	newfilm, err := dal.CreateFilm(film)
	if err != nil {
		return nil, err
	}
	return dto.FilmDALToDTO(newfilm)
}

// UpdateFilm : update exist film.
func UpdateFilm(f *dto.FilmDTO) (*dto.FilmDTO, error) {
	film, err := f.FilmDTOToDAL()
	if err != nil {
		return nil, err
	}
	updatefilm, err := dal.UpdateFilm(film)
	if err != nil {
		return nil, err
	}
	return dto.FilmDALToDTO(updatefilm)
}

// DeleteFilm : delete film by id.
func DeleteFilm(id int32) error {
	return dal.DeleteFilm(id)
}


