package proxy

import "fmt"

// daoProxy declaring connection to DAO
type daoProxy struct {
	connection *dao
}

// Get creates connection if not exist
// And returns data by id
func (d *daoProxy) Get(id int) string {
	d.createConnection()
	fmt.Printf("Getting..\n")
	return d.connection.get(id)
}

// Update creates connection if not exist
// And updates data by id
func (d *daoProxy) Update(id int, data string) {
	d.createConnection()
	fmt.Printf("Updating..\n")
	d.connection.update(id, data)
	fmt.Printf("Now user#%v data is '%v'\n", id, data)
}

// Delete creates connection if not exist
// And deletes data by id
func (d *daoProxy) Delete(id int) {
	d.createConnection()
	fmt.Printf("Deleting..\n")
	d.connection.delete(id)
	fmt.Printf("Deleted user with id = %v\n", id)
}

// Add creates connection if not exist
// And adds data
func (d *daoProxy) Add(data string) {
	d.createConnection()
	fmt.Printf("Adding..\n")
	d.connection.add(data)
	fmt.Printf("Added new user with data '%v'\n", data)
}

// createConnection check if a connection exist
// Creates connection if not
func (d *daoProxy) createConnection() {
	if d.connection == nil {
		d.connection = &dao{}
	}
}

// NewDAOProxy returns new daoProxy
func NewDAOProxy() *daoProxy {
	return &daoProxy{}
}
