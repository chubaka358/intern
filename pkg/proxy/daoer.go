package proxy

// DAOer provides dao interface
type DAOer interface {
	get(id int) string
	update(id int, data string)
	delete(id int)
	add(data string)
}
