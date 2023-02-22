package repository

type SQLExecuter interface {
	Get(name string, dest interface{}, query string, args ...interface{}) error
	// GetByJsonTag(name string, dest interface{}, query string, args ...interface{}) error
	Select(name string, dest interface{}, query string, args ...interface{}) error
	Exec(name string, query string, args ...interface{}) (int64, error)
}

type DB interface {
	Begin() (Tx, error)
}

type Tx interface {
	Rollback() error
	Commit() error
}
