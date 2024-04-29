package api

type User struct {
	ID    string
	Name  string
	Email string
}

type (
	UserPermissions struct {
		UserID      string
		Permissions []Permission
	}
	Permission struct {
		Resources ResourcesSelector
		Allows    []ResourceAction
		Denies    []ResourceAction
	}
	ResourcesSelector string
	ResourceAction    string
)
