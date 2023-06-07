package actors

import "time"

type Actors struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	Verified  string    `json:"verified"`
	Active    string    `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type RegisterApproval struct {
	Id           int    `json:"id"`
	AdminId      int    `json:"admin_id"`
	SuperAdminId int    `json:"super_admin_id"`
	Status       string `json:"status"`
}
