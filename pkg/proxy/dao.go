package proxy

import "fmt"

// dao works directly with the database
type dao struct {
}

// update updates data by id
func (d *dao) update(id int, data string) {
	if d == nil {
		fmt.Printf("Connection doesn't exist\n")
	} else {
		fmt.Printf("Doing inner Update logic\n")
	}
}

// get receives data by id
func (d *dao) get(id int) string {
	if d == nil {
		return "Connection doesn't exist\n"
	} else {
		return "Doing inner Get logic\n"
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
