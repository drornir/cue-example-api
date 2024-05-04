package api

Examples: {
	AUser: #User & {
		id:    "abcd"
		name:  "dn"
		email: "me@drornir.dev"
	}
	APermissions: #UserPermissions & {
		user_id: AUser.id
		permissions: [
			{
				resources: "my-selector"
				allows: ["read", "write"]
			},
			{
				resources: "my-selector-2"
				allows: []
				denies: ["write"]
			},
		]
	}
}
