package data

// Remove removes the user with the given ID
// dataset is locked for writing during this process
func Remove(id Key) {
	dataset := <-lock
	delete(dataset, id)
	lock <- dataset
}
