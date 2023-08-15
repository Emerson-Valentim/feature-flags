package repository

type FlagEntity struct {
	Id    string
	Name  string
	State bool
}

type InsertFun func(flag FlagEntity) (*FlagEntity, error)
type FindFun func(ids []string) ([]FlagEntity, error)
type DeleteFun func(id string) error
type UpdateFun func(flag FlagEntity) (*FlagEntity, error)

type Repository struct {
	Insert InsertFun
	Find   FindFun
	Delete DeleteFun
	Update UpdateFun
}
