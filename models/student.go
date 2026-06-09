package models

// 合并提交2

type Student struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int    `json:"add_time"` // 驼峰命名
}

// 数据库表名称
func (Student) TableName() string {
	return "student"
}
