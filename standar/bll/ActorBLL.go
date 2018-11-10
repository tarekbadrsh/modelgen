package bll

import (
	"strconv"

	"github.com/tarekbadrshalaan/modelgen/standar/dal"
	"github.com/tarekbadrshalaan/modelgen/standar/dto"
)

// ConvertActorID : covnert ActorID string to ActorID int32.
func ConvertActorID(str string) (int32, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := int32(pram)
	return id, nil
}

// GetAllActors : get All actors.
func GetAllActors() ([]*dto.ActorDTO, error) {
	actors := dal.GetAllActors()
	return dto.ActorDALToDTOArr(actors)
}

// GetActor : get one actor by id.
func GetActor(id int32) (*dto.ActorDTO, error) {
	actor, err := dal.GetActor(id)
	if err != nil {
		return nil, err
	}
	return dto.ActorDALToDTO(actor)
}

// CreateActor : create new actor.
func CreateActor(a *dto.ActorDTO) (*dto.ActorDTO, error) {
	actor, err := a.ActorDTOToDAL()
	if err != nil {
		return nil, err
	}
	newactor, err := dal.CreateActor(actor)
	if err != nil {
		return nil, err
	}
	return dto.ActorDALToDTO(newactor)
}

// UpdateActor : update exist actor.
func UpdateActor(a *dto.ActorDTO) (*dto.ActorDTO, error) {
	actor, err := a.ActorDTOToDAL()
	if err != nil {
		return nil, err
	}
	updateactor, err := dal.UpdateActor(actor)
	if err != nil {
		return nil, err
	}
	return dto.ActorDALToDTO(updateactor)
}

// DeleteActor : delete actor by id.
func DeleteActor(id int32) error {
	return dal.DeleteActor(id)
}
