package proxy

import "fmt"

type daoProxier interface {
	Get(id int) string
	Update(id int, data string)
	Delete(id int)
	Add(data string)
}

type daoer interface {
	get(id int) string
	update(id int, data string)
	delete(id int)
	add(data string)
}

// DAO works directly with the database
type dao struct {
}

// DAOProxy declaring connection to DAO
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

// get receives data by id
func (d *dao) get(id int) string {
	if d == nil {
		return "Connection doesn't exist\n"
	} else {
		return "Doing inner Get logic\n"
	}
}

// update updates data by id
func (d *dao) update(id int, data string) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Update logic\n")
	}
}

// delete deletes data by id
func (d *dao) delete(id int) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Delete logic\n")
	}
}

// add adds data
func (d *dao) add(data string) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Add logic\n")
	}
}

// createConnection check if a connection exist
// Creates connection if not
func (d *daoProxy) createConnection() {
	if d.connection == nil {
		d.connection = &dao{}
	}
}

func NewDAOProxy() *daoProxy {
	return &daoProxy{}
}
