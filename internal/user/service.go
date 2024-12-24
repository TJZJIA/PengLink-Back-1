package user

import (
	"PengLink-Back-1/models"
)

func RegisterUser(username, password, role string) error {
	// 创建新用户
	user := models.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	// 保存用户到数据库
	// 假设这里有保存到数据库的逻辑
	// db.Create(&user)
	if err := saveUserToDB(&user); err != nil {
		return err
	}
	return nil
}

func LoginUser(username, password string) (*models.User, error) {
	// 假设这里有从数据库中查找用户的逻辑
	// var user models.User
	// db.Where("username = ? AND password = ?", username, password).First(&user)
	// 如果用户不存在或密码不匹配，返回错误
	// if user.ID == 0 {
	// 	return nil, errors.New("invalid username or password")
	// }
	// 这里为了示例，假设用户存在且密码匹配
	user := models.User{
		Username: username,
		Password: password,
		Role:     "user", // 示例角色
	}
	return &user, nil
}

// 假设这是一个保存用户到数据库的函数
func saveUserToDB(user *models.User) error {
	// 这里添加实际的数据库保存逻辑
	// 例如：db.Create(user)
	return nil
}
