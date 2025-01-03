package todogo

// The data types that will be used globally should be defined the the root of the project
type User struct {
	Id       int    `json:"-"`    // Skip this field in json encoded string
	Name     string `json:"name"` // rename the field to name in json encoding
	Username string `json:"username"`
	Password string `json:"password"`
}
