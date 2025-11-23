package role_menu_def

type AuthMenu struct {
	MenuId          int   `json:"menu_id"`
	OperationIdList []int `json:"operation_id_list"`
}

type RolesAuthed struct {
	SystemId int   `json:"system_id"`
	RoleIds  []int `json:"role_ids"`

	Authed []AuthMenu `json:"authed"`
}
