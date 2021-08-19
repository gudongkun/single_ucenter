package models

// XORM技术文档范例
type User struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
}
