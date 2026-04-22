package models

import "time"

type Users struct {
	ID               string    `json:"_id,omitempty"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	RolesId          int64     `json:"roles_id"`
	Active           int       `json:"active"`
	Author           string    `json:"author"`
	Company_id       int64     `json:"company_id"`
	ResetToken       string    `json:"reset_token"`
	ResetTokenExpira time.Time `json:"reset_token_expira"`
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	UpdatedDate      time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type UsersResp struct {
	ID               string    `json:"id"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	RolesId          int64     `json:"roles_id"`
	Active           int       `json:"active"`
	Author           string    `json:"author"`
	Company_id       int64     `json:"company_id"`
	ResetToken       string    `json:"reset_token"`
	ResetTokenExpira time.Time `json:"reset_token_expira"`
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	UpdatedDate      time.Time `gorm:"autoUpdateTime"`
}

func (UsersResp) TableName() string {
	return "users" // Nombre de la tabla en la base de datos
}

// Create---Req  request struct
type CreateUsersReq struct {
	ID               string    `gorm:"primaryKey;autoIncrement"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	RolesId          int64     `json:"roles_id"`
	Active           int       `json:"active"`
	Author           string    `json:"author"`
	Company_id       int64     `json:"company_id"`
	ResetToken       string    `json:"reset_token"`
	ResetTokenExpira time.Time `json:"reset_token_expira"`
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	UpdatedDate      time.Time `gorm:"autoUpdateTime"`
}

func (CreateUsersReq) TableName() string {
	return "users" // Nombre de la tabla en la base de datos
}

type UpdateUsersReq struct {
	ID               string     `json:"-"`
	Username         *string    `json:"username"`
	Name             *string    `json:"name"`
	Password         *string    `json:"password"`
	Email            *string    `json:"email"`
	Phone            *string    `json:"phone"`
	RolesId          *int64     `json:"roles_id"`
	Active           *int       `json:"active"`
	Author           *string    `json:"author"`
	Company_id       *int64     `json:"company_id"`
	ResetToken       *string    `json:"reset_token"`
	ResetTokenExpira *time.Time `json:"reset_token_expira"`
	CreatedDate      *time.Time `gorm:"autoCreateTime"`
	UpdatedDate      *time.Time `gorm:"autoUpdateTime"`
}

func (UpdateUsersReq) TableName() string {
	return "users" // Nombre de la tabla en la base de datos
}

type UsersInf struct {
	ID          string      `json:"id"`
	Username    string      `json:"username"`
	Name        string      `json:"name"`
	Password    string      `json:"password,omitempty"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	RolesId     int64       `json:"roles_id"`
	Rol         RolesReport `json:"rol" gorm:"foreignKey:RolesId;references:ID"`
	Active      int         `json:"active"`
	Author      string      `json:"author"`
	CompanyId   int64       `json:"company_id"`
	CreatedDate time.Time   `gorm:"autoCreateTime"`
	UpdatedDate time.Time   `gorm:"autoUpdateTime"`
}

func (UsersInf) TableName() string {
	return "users" // Nombre de la tabla en la base de datos
}

type UsersListResponse struct {
	Items      []UsersInf `json:"items"`
	TotalCount int64      `json:"totalCount"`
}

type UsersReport struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	CompanyId int64  `json:"company_id"`
}

func (UsersReport) TableName() string {
	return "users" // Nombre de la tabla en la base de datos
}
