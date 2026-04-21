package utils_auth

// RoleHierarchy defines the permission level for each role
var RoleHierarchy = map[string]int{
	"member": 1,
	"admin":  2,
}

// HasPermission checks if the user's role meets or exceeds the required role level
func HasPermission(userRole, requiredRole string) bool {
	userLevel, userExists := RoleHierarchy[userRole]
	requiredLevel, requiredExists := RoleHierarchy[requiredRole]

	if !userExists || !requiredExists {
		return false
	}

	return userLevel >= requiredLevel
}
