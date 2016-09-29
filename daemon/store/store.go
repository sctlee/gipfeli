package store

// ReadTx ...
type ReadTx interface {
	lookup(table, index, id string) Object
	get(table, id string) Object
	find(table string, appendResult func(Object), bys ...By) error
}

// Tx is a read/write transaction. Note that transaction does not imply
// any internal batching. The purpose of this transaction is to give the
// user a guarantee that its changes won't be visible to other transactions
// until the transaction is over.
type Tx interface {
	ReadTx
	create(table string, o Object) error
	update(table string, o Object) error
	delete(table, id string) error
}
