package data

// Key is a unique ID type
type Key uint

// User represents a user resource, with unique ID
type User struct {
	ID             Key      `json:"id"`
	Firstname      string   `json:"firstname"`
	MiddleInitials []string `json:"middle_initials,omitempty"`
	Lastname       string   `json:"lastname"`
}

// userSlice implements sort.Sortable interface for User slices
type userSlice []*User

func (u userSlice) Len() int           { return len(u) }
func (u userSlice) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }
func (u userSlice) Less(i, j int) bool { return u[i].ID < u[j].ID }

// dataset is hardcoded…
// source: https://fr.wikipedia.org/wiki/Liste_des_pr%C3%A9sidents_des_%C3%89tats-Unis
var dataset = map[Key]*User{
	1789: &User{ID: 1789, Firstname: "George", Lastname: "Washington"},
	1797: &User{ID: 1797, Firstname: "John", Lastname: "Adams"},
	1801: &User{ID: 1801, Firstname: "Thomas", Lastname: "Jefferson"},
	1809: &User{ID: 1809, Firstname: "James", Lastname: "Madison"},
	1817: &User{ID: 1817, Firstname: "James", Lastname: "Monroe"},
	1825: &User{ID: 1825, Firstname: "John", MiddleInitials: []string{"Q"}, Lastname: "Adams"},
	1829: &User{ID: 1829, Firstname: "Andrew", Lastname: "Jackson"},
	1837: &User{ID: 1837, Firstname: "Martin", Lastname: "Van Buren"},
	1841: &User{ID: 1841, Firstname: "William", MiddleInitials: []string{"H"}, Lastname: "Harrison"},
	1842: &User{ID: 1842, Firstname: "John", Lastname: "Tyler"},
	1845: &User{ID: 1845, Firstname: "James", MiddleInitials: []string{"K"}, Lastname: "Polk"},
	1849: &User{ID: 1849, Firstname: "Zachary", Lastname: "Taylor"},
	1850: &User{ID: 1850, Firstname: "Millard", Lastname: "Fillmore"},
	1853: &User{ID: 1853, Firstname: "Franklin", Lastname: "Pierce"},
	1857: &User{ID: 1857, Firstname: "James", Lastname: "Buchanan"},
	1861: &User{ID: 1861, Firstname: "Abraham", Lastname: "Lincoln"},
	1865: &User{ID: 1865, Firstname: "Andrew", Lastname: "Johnson"},
	1869: &User{ID: 1869, Firstname: "Ulysses", MiddleInitials: []string{"S"}, Lastname: "Grant"},
	1877: &User{ID: 1877, Firstname: "Rutherford", MiddleInitials: []string{"B"}, Lastname: "Hayes"},
	1881: &User{ID: 1881, Firstname: "James", MiddleInitials: []string{"A"}, Lastname: "Garfield"},
	1882: &User{ID: 1882, Firstname: "Chester", MiddleInitials: []string{"A"}, Lastname: "Arthur"},
	1885: &User{ID: 1885, Firstname: "Stephen Grover", Lastname: "Cleveland"},
	1889: &User{ID: 1889, Firstname: "Benjamin", Lastname: "Harrison"},
	1897: &User{ID: 1897, Firstname: "William", Lastname: "McKinley"},
	1901: &User{ID: 1901, Firstname: "Theodore", Lastname: "Roosevelt"},
	1909: &User{ID: 1909, Firstname: "William", MiddleInitials: []string{"H"}, Lastname: "Taft"},
	1913: &User{ID: 1913, Firstname: "Thomas Woodrow", Lastname: "Wilson"},
	1921: &User{ID: 1921, Firstname: "Warren", MiddleInitials: []string{"G"}, Lastname: "Harding"},
	1923: &User{ID: 1923, Firstname: "John Calvin", Lastname: "Coolidge"},
	1929: &User{ID: 1929, Firstname: "Herbert", MiddleInitials: []string{"C"}, Lastname: "Hoover"},
	1933: &User{ID: 1933, Firstname: "Franklin", MiddleInitials: []string{"D"}, Lastname: "Roosevelt"},
	1945: &User{ID: 1945, Firstname: "Harry", MiddleInitials: []string{"S"}, Lastname: "Truman"},
	1953: &User{ID: 1953, Firstname: "Dwight", MiddleInitials: []string{"D"}, Lastname: "Eisenhower"},
	1961: &User{ID: 1961, Firstname: "John", MiddleInitials: []string{"F"}, Lastname: "Kennedy"},
	1963: &User{ID: 1963, Firstname: "Lyndon", MiddleInitials: []string{"B"}, Lastname: "Johnson"},
	1969: &User{ID: 1969, Firstname: "Richard", MiddleInitials: []string{"M"}, Lastname: "Nixon"},
	1974: &User{ID: 1974, Firstname: "Gerald", MiddleInitials: []string{"R"}, Lastname: "Ford"},
	1977: &User{ID: 1977, Firstname: "Jimmy", Lastname: "Carter"},
	1981: &User{ID: 1981, Firstname: "Ronald", MiddleInitials: []string{"W"}, Lastname: "Reagan"},
	1989: &User{ID: 1989, Firstname: "George", MiddleInitials: []string{"H", "W"}, Lastname: "Bush"},
	1993: &User{ID: 1993, Firstname: "Bill", Lastname: "Clinton"},
	2001: &User{ID: 2001, Firstname: "George", MiddleInitials: []string{"W"}, Lastname: "Bush"},
	2009: &User{ID: 2009, Firstname: "Barack", Lastname: "Obama"},
	2017: &User{ID: 2017, Firstname: "Donald", Lastname: "Trump"},
}

var lastID Key = 2017

var lock = make(chan map[Key]*User, 1)

func init() {
	lock <- dataset
}
