package v1_model

type User struct { // https://gorm.io/docs/models.html
	Id   int    `gorm:"column:id;type:int;comment:标识;primaryKey" json:"id"`
	User string `gorm:"column:user;type:varchar;comment:用户名" json:"user"`
	Pswd string `gorm:"column:pswd;type:varchar;comment:密码" json:"omitempty"`
	V    bool   `gorm:"column:v;type:bool;default:true;comment:有效性" json:"v"`
}
