package api

import "strings"

User: #User
#User: {
	name: strings.MinRunes(2)
}

UserPermissions: #UserPermissions
Permission: #Permission
ResourcesSelector: #ResourcesSelector
ResourceAction: #ResourceAction
