package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)  

// FilmDAL : data access layer  (film) table.
type FilmDAL struct {
	FilmID int32 `json:"film_id" gorm:"column:film_id;primary_key:true"`
	Title string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	ReleaseYear int32 `json:"release_year" gorm:"column:release_year"`
	LanguageID int16 `json:"language_id" gorm:"column:language_id"`
	RentalDuration int16 `json:"rental_duration" gorm:"column:rental_duration"`
	RentalRate interface {} `json:"rental_rate" gorm:"column:rental_rate"`
	Length int16 `json:"length" gorm:"column:length"`
	ReplacementCost interface {} `json:"replacement_cost" gorm:"column:replacement_cost"`
	Rating interface {} `json:"rating" gorm:"column:rating"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	SpecialFeatures interface {} `json:"special_features" gorm:"column:special_features"`
	Fulltext interface {} `json:"fulltext" gorm:"column:fulltext"`
	
}

// TableName sets the insert table name for this struct type
func (f *FilmDAL) TableName() string {
	return "film"
} 

// GetAllFilms : get all films.
func GetAllFilms() []*FilmDAL {
	films := []*FilmDAL{}
	db.DB().Find(&films)
	return films
}

// GetFilm : get one film by id.
func GetFilm(id int32) (*FilmDAL, error) {
	f := &FilmDAL{}
	result := db.DB().First(f, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return f, nil
}


// CreateFilm : create new film.
func CreateFilm(f *FilmDAL) (*FilmDAL, error) {
	result := db.DB().Create(f)
	if result.Error != nil {
		return nil, result.Error
	}
	return f, nil
}

// UpdateFilm : update exist film.
func UpdateFilm(f *FilmDAL) (*FilmDAL, error) {
	_, err := GetFilm(f.FilmID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(f)
	if result.Error != nil {
		return nil, result.Error
	}
	return f, nil
}

// DeleteFilm : delete film by id.
func DeleteFilm(id int32) error {
	f, err := GetFilm(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(f)
	return result.Error
}


