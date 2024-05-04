package api

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserPermissions struct {
	UserID      string       `json:"user_id"`
	Permissions []Permission `json:"permissions"`
}
type Permission struct {
	Resources ResourcesSelector `json:"resources"`
	Allows    []ResourceAction  `json:"allows,omitempty"`
	Denies    []ResourceAction  `json:"denies,omitempty"`
}
type ResourcesSelector string
type ResourceAction string
