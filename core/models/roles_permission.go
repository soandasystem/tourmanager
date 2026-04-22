package models

type RolesPermissions struct {
	ID         string `json:"_id,omitempty"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

// Resp  response struct
type RolesPermissionsResp struct {
	ID         string `json:"id"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

func (RolesPermissionsResp) TableName() string {
	return "roles_permissions" // Nombre de la tabla en la base de datos
}

type RolesPermissionsListResponse struct {
	Items      []RolesPermissionsResp `json:"items"`
	TotalCount int64                  `json:"totalCount"`
}

// Create---Req  request struct
type CreateRolesPermissionsReq struct {
	ID         string `gorm:"primaryKey;autoIncrement"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

func (CreateRolesPermissionsReq) TableName() string {
	return "roles_permissions" // Nombre de la tabla en la base de datos
}

type UpdateRolesPermissionsReq struct {
	ID         string  `json:"-"`
	RolesId    *int64  `json:"roles_id"`
	Permission *string `json:"permission"`
	Actions    *string `json:"actions"`
	CompanyId  *int64  `json:"company_id"`
}

func (UpdateRolesPermissionsReq) TableName() string {
	return "roles_permissions" // Nombre de la tabla en la base de datos
}
