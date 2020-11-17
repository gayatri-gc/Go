package data

// db act as a dummy package level database.
var db map[string]bool

// init initialize a dummy db with some data
func init() {
	db = make(map[string]bool)
	db["gayatri61995@gmail.com"] = true
	db["gayatri"] = true
}

// UserExists check if the User is registered with the provided email.
func UserExists(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}