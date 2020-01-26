package dto
import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dal"
)  

// FilmDTO : data transfer object  (film) table.
type FilmDTO struct {
	FilmID int32 `json:"film_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int32 `json:"release_year"`
	LanguageID int16 `json:"language_id"`
	RentalDuration int16 `json:"rental_duration"`
	RentalRate interface {} `json:"rental_rate"`
	Length int16 `json:"length"`
	ReplacementCost interface {} `json:"replacement_cost"`
	Rating interface {} `json:"rating"`
	LastUpdate time.Time `json:"last_update"`
	SpecialFeatures interface {} `json:"special_features"`
	Fulltext interface {} `json:"fulltext"`
	
}

// FilmDTOToDAL : convert FilmDTO to FilmDAL
func (a *FilmDTO) FilmDTOToDAL() (*dal.FilmDAL, error) { 
	film := &dal.FilmDAL{
		FilmID:a.FilmID,
		Title:a.Title,
		Description:a.Description,
		ReleaseYear:a.ReleaseYear,
		LanguageID:a.LanguageID,
		RentalDuration:a.RentalDuration,
		RentalRate:a.RentalRate,
		Length:a.Length,
		ReplacementCost:a.ReplacementCost,
		Rating:a.Rating,
		LastUpdate:a.LastUpdate,
		SpecialFeatures:a.SpecialFeatures,
		Fulltext:a.Fulltext,
		 
	}
	return film, nil
}

// FilmDALToDTO : convert FilmDAL to FilmDTO
func FilmDALToDTO(a *dal.FilmDAL) (*FilmDTO, error) { 
	film := &FilmDTO{
		FilmID:a.FilmID,
		Title:a.Title,
		Description:a.Description,
		ReleaseYear:a.ReleaseYear,
		LanguageID:a.LanguageID,
		RentalDuration:a.RentalDuration,
		RentalRate:a.RentalRate,
		Length:a.Length,
		ReplacementCost:a.ReplacementCost,
		Rating:a.Rating,
		LastUpdate:a.LastUpdate,
		SpecialFeatures:a.SpecialFeatures,
		Fulltext:a.Fulltext,
		 
	}
	return film, nil
}

// FilmDALToDTOArr : convert Array of FilmDAL to Array of FilmDTO
func FilmDALToDTOArr(films []*dal.FilmDAL) ([]*FilmDTO, error) {
	var err error
	res := make([]*FilmDTO, len(films))
	for i, film := range films {
		res[i], err = FilmDALToDTO(film)
		if err != nil { 
			return res, err
		}
	}
	return res, nil
}


