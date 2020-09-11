package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string    `json:"menuId" gorm:"comment:'菜单ID'"`
	AuthorityId string    `json:"-" gorm:"comment:'角色ID'"`
	// self-referential has many
	Children    []SysMenu `json:"children" gorm:"foreignKey:MenuId"`
}
