package data

// Set adds a new user with given ID
// dataset is locked for writing during this process
func Set(id Key, user *User) {
	dataset := <-lock
	user.ID = id
	dataset[id] = user
	if id > lastID {
		lastID = id
	}
	lock <- dataset
}

// Add adds a new user with a uniquely generated ID
// dataset is locked for writing during this process
func Add(user *User) {
	dataset := <-lock
	lastID++
	user.ID = lastID
	dataset[lastID] = user
	lock <- dataset
}
