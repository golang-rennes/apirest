package data

import "sort"

// GetAll returns all users (sorted by ID)
func GetAll() []*User {
	users := make([]*User, 0, len(dataset))
	for _, u := range dataset {
		users = append(users, u)
	}
	sort.Sort(userSlice(users))
	return users
}

// GetByFirstname returns users with matching firstname
func GetByFirstname(firstname string) []*User {
	users := make([]*User, 0, 10)
	for _, u := range dataset {
		if u.Firstname == firstname {
			users = append(users, u)
		}
	}
	sort.Sort(userSlice(users))
	return users
}

// Get returns the user with the given ID (nil if none)
func Get(id Key) *User {
	return dataset[id]
}
