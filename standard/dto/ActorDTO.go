package dto

import (
	"time"

	"github.com/pkg/errors"
	"github.com/tarekbadrshalaan/modelgen/standard/dal"
)

// ActorDTO : data transfer object of (actor) table.
type ActorDTO struct {
	ActorID    int32     `json:"actor_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	LastUpdate time.Time `json:"last_update"`
}

// ActorDTOToDAL : convert ActorDTO to ActorDAL
func (a *ActorDTO) ActorDTOToDAL() (*dal.ActorDAL, error) {
	actor := &dal.ActorDAL{
		ActorID:    a.ActorID,
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		LastUpdate: a.LastUpdate,
	}
	return actor, nil
}

// ActorDALToDTO : convert ActorDAL to ActorDTO
func ActorDALToDTO(a *dal.ActorDAL) (*ActorDTO, error) {
	actor := &ActorDTO{
		ActorID:    a.ActorID,
		FirstName:  a.FirstName,
		LastName:   a.LastName,
		LastUpdate: a.LastUpdate,
	}
	return actor, nil
}

// ActorDALToDTOArr : convert Array of ActorDAL to Array of ActorDTO
func ActorDALToDTOArr(actors []*dal.ActorDAL) ([]*ActorDTO, error) {
	var err error
	res := make([]*ActorDTO, len(actors), len(actors))
	for i, actor := range actors {
		res[i], err = ActorDALToDTO(actor)
		if err != nil {
			return res, errors.Wrapf(err, "convert at %d:%v", i, actor)
		}
	}
	return res, nil
}
