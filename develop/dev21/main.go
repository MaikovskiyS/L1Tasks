package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/
func main() {
	var db Database
	adapter := New(db)
	var id int
	adapter.Get(id)
}

//Target
type DbClient interface {
	Get(id int)
}
type DBClient struct{}

func (c *DBClient) Get(id int) {
}

//adapter
type Adapter struct {
	db Database
}

func New(db Database) DbClient {
	return &Adapter{
		db: db,
	}
}
func (a *Adapter) Get(id int) {
	dbId := Dbid(id)
	a.db.GetFromDb(dbId)
}

//Adaptee
type Database interface {
	GetFromDb(id Dbid)
}
type Dbid int
