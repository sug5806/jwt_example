package db_model

type User struct {
	Id    uint64
	Uuid  string
	Email string
	Pass  string
}

func (u User) TableName() string {
	return "user"
}
