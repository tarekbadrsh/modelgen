package dal

import ( 
	"time"
	
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/db"
)  

// ActorDAL : data access layer  (actor) table.
type ActorDAL struct {
	ActorID int32 `json:"actor_id" gorm:"column:actor_id;primary_key:true"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName string `json:"last_name" gorm:"column:last_name"`
	LastUpdate time.Time `json:"last_update" gorm:"column:last_update"`
	
}

// TableName sets the insert table name for this struct type
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
	a := &ActorDAL{}
	result := db.DB().First(a, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}


// CreateActor : create new actor.
func CreateActor(a *ActorDAL) (*ActorDAL, error) {
	result := db.DB().Create(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// UpdateActor : update exist actor.
func UpdateActor(a *ActorDAL) (*ActorDAL, error) {
	_, err := GetActor(a.ActorID)
	if err != nil {
		return nil, err
	}
	result := db.DB().Save(a)
	if result.Error != nil {
		return nil, result.Error
	}
	return a, nil
}

// DeleteActor : delete actor by id.
func DeleteActor(id int32) error {
	a, err := GetActor(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete(a)
	return result.Error
}


