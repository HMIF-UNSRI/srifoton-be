package user

func (u *User) GetUserRoleString() string {
	return string(u.Role)
}

func (u *User) SetUserRoleString(r string) {
	switch r {
	case string(Base):
		u.Role = Base
	default:
		u.Role = Admin
	}
}
