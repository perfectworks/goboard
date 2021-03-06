package storage

import "gopkg.in/gorp.v1"

const (
	// ResourceProject means project resource
	ResourceProject = "PROJECT"
)

const (
	// OperationGet allow user to get resources
	OperationGet = "GET"

	// OperationPut allow user to create a new resource
	OperationPut = "PUT"

	// OperationPost allow user to update an existed resource
	OperationPost = "POST"

	// OperationDelete allow user to delete an existed resource
	OperationDelete = "DELETE"
)

// RolePrivilege means allow specified role to execute Operation on Resource
type RolePrivilege struct {
	Resource  string `db:"resource"`
	Operation string `db:"operation"`
	RoleID    int    `db:"role_id"`
}

func initRolePrivilegeTable(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(RolePrivilege{}, "role_privileges")
}

// QueryRolePrivilegesByRoleID will fetch all RolePrivilege by a specified RoleID
func QueryRolePrivilegesByRoleID(roleID int, dbmap *gorp.DbMap) (rolePrivileges []RolePrivilege, err error) {
	_, err = dbmap.Select(&rolePrivileges, "select * from role_privileges where role_id = ?", roleID)
	return
}

// Save will insert a RolePrivilege record to database
func (rp *RolePrivilege) Save(dbmap *gorp.DbMap) (err error) {
	return dbmap.Insert(rp)
}

// AuthProject check user privilege to project
func AuthProject(userID int, resource string, operation string, projectID int, dbmap *gorp.DbMap) (authorized bool, err error) {
	count, err := dbmap.SelectInt("select count(*) as count from user_roles ur"+
		" join roles r on r.id = ur.role_id"+
		" join role_privileges rp on rp.role_id = r.id"+
		" where ur.user_id = ? and (ur.project_id = ? or r.scope = ?) and rp.resource = ? and rp.operation = ?", userID, projectID, RoleGlobal, resource, operation)

	authorized = count > 0

	return
}
