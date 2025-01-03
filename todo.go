package todogo

// todo_list
type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// relation user_list
type UserList struct {
	Id     int
	UserId int
	ListId int
}

// todo_item
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// list_item
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
