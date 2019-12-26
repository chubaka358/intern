package proxy

import "fmt"

type daoConnecter interface {
	Get(id int) string
	Update(id int, data string)
	Delete(id int)
	Add(data string)
}

// DAO works directly with the database
type dao struct {
}

// DAOProxy declaring connection to DAO
type DAOProxy struct {
	connection *dao
}

// Get receives data by id
func (d *dao) Get(id int) string {
	if d == nil {
		return "Connection doesn't exist\n"
	} else {
		return "Doing inner Get logic\n"
	}
}

// Update updates data by id
func (d *dao) Update(id int, data string) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Update logic\n")
	}
}

// Delete deletes data by id
func (d *dao) Delete(id int) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Delete logic\n")
	}
}

// Add adds data
func (d *dao) Add(data string) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Add logic\n")
	}
}

// Get creates connection if not exist
// And returns data by id
func (d *DAOProxy) Get(id int) string {
	d.createConnection()
	fmt.Printf("Getting..\n")
	return d.connection.Get(id)
}

// Update creates connection if not exist
// And updates data by id
func (d *DAOProxy) Update(id int, data string) {
	d.createConnection()
	fmt.Printf("Updating..\n")
	d.connection.Update(id, data)
	fmt.Printf("Now user#%v data is '%v'\n", id, data)
}

// Delete creates connection if not exist
// And deletes data by id
func (d *DAOProxy) Delete(id int) {
	d.createConnection()
	fmt.Printf("Deleting..\n")
	d.connection.Delete(id)
	fmt.Printf("Deleted user with id = %v\n", id)
}

// Add creates connection if not exist
// And adds data
func (d *DAOProxy) Add(data string) {
	d.createConnection()
	fmt.Printf("Adding..\n")
	d.connection.Add(data)
	fmt.Printf("Added new user with data '%v'\n", data)
}

// createConnection check if a connection exist
// Creates connection if not
func (d *DAOProxy) createConnection() {
	if d.connection == nil {
		d.connection = &dao{}
	}
}
