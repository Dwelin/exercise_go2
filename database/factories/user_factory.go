// Package factories 存放工厂方法
package factories

import (
	"gohub/app/models/user"
	"gohub/pkg/helpers"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$joNqg8Qvwn5go7bCNUB4guDQTp8wu2vJecf7PAnWKujjlF6Wyk7S.",
		}
		objs = append(objs, model)
	}

	return objs
}
