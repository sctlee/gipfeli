package store

// Object ...
type Object interface {
	ID() string
	JSONData() (string, error)
	SetContent(jsonData string) error
}

// ObjectConfig ...
type ObjectConfig struct {
	Name     string
	Table    *TableSchema
	Objecter func() Object
}

// TableSchema ...
type TableSchema struct {
	Name    string
	Indexes map[string]*IndexSchema
}

// IndexSchema contains the schema for an index
type IndexSchema struct {
	Name         string
	AllowMissing bool
	Unique       bool
}
