/**
  @author:Cofeesy
  @data:2022/6/24
  @note:
**/
package serializer

import "memorandumProject/model"

//序列化User容器
type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"user_name" form:"user_name" example:"xxx"`
	Status   string `json:"status" form:"status"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

/**
 * @Author Cofessy
 * @Description //对数据库层的User对象进行序列化返回给前端
 * @Date 0:24 2022/6/24
 * @Param user model.User ”数据库的User对象“
 * @return User ”对数据库User序列化后的新User，用于返回给前端“
 **/
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
