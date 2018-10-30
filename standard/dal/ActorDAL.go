package dal

import (
	"time"

	"github.com/tarekbadrshalaan/modelgen/standard/db"
)

// ActorDAL : data access layer (actor) table.
type ActorDAL struct {
	ActorID    int32     `json:"actor_id" gorm:"column:actor_id;primary_key:true"`
	FirstName  string    `json:"first_name" gorm:"column:first_name"`
	LastName   string    `json:"last_name" gorm:"column:last_name"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
}

// TableName sets the insert table name for this struct type.
func (a *ActorDAL) TableName() string {
	return "actor"
}

// GetAllActors : get all actors.
func GetAllActors() []*ActorDAL {
	actors := []*ActorDAL{}
	db.DB().Find(&actors)
	return actors
}

// GetActor : get one actor by id.
func GetActor(id int32) (*ActorDAL, error) {
	actor := &ActorDAL{}
	result := db.DB().First(actor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return actor, nil
}

// CreateActor : create new actor.
func CreateActor(actor *ActorDAL) (*ActorDAL, error) {
	result := db.DB().Create(actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return actor, nil
}

// UpdateActor : update exist actor.
func UpdateActor(actor *ActorDAL) (*ActorDAL, error) {
	_, err := GetActor(actor.ActorID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(actor)
	if result.Error != nil {
		return nil, result.Error
	}
	return actor, nil
}

// DeleteActor : delete actor by id.
func DeleteActor(id int32) error {
	actor, err := GetActor(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(actor)
	return result.Error
}
