package proxy

// DAOProxier provides daoProxy interface
type DAOProxier interface {
	Get(id int) string
	Update(id int, data string)
	Delete(id int)
	Add(data string)
}
