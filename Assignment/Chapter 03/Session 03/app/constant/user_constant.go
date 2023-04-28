package constant

type UserRoleType string

var UserRole = struct {
	HEAD        UserRoleType
	ADMIN_CAR   UserRoleType
	ADMIN_MOTOR UserRoleType
	ADMIN_CYCLE UserRoleType
}{
	HEAD:        "HEAD",
	ADMIN_CAR:   "ADMIN_MOBIL",
	ADMIN_MOTOR: "ADMIN_MOTOR",
	ADMIN_CYCLE: "ADMIN_SEPEDA",
}

type UserRoleTypeId int

var UserRoleId = struct {
	HEAD        UserRoleTypeId
	ADMIN_CAR   UserRoleTypeId
	ADMIN_MOTOR UserRoleTypeId
	ADMIN_CYCLE UserRoleTypeId
}{
	HEAD:        1,
	ADMIN_CAR:   2,
	ADMIN_MOTOR: 3,
	ADMIN_CYCLE: 4,
}
