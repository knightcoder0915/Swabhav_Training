package authentication

type Role struct {
	RoleType  string
	BibaLevel int
	BellLevel int
}

func NewRole(role string, level1, level2 int) *Role {
	return &Role{RoleType: role, BibaLevel: level1, BellLevel: level2}
}
