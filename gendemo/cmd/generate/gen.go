package main

import (
	"github.com/ag9920/db-demo/gendemo/dal/model"
	"gorm.io/gen"
)

/*
 * 代码生成逻辑
 */
func main() {

	// 生成代码
	g := gen.NewGenerator(gen.Config{
		// 代码生成目录
		OutPath: "../../dal/query",

		// Mode 为 true 时，会在 OutPath
		// 目录下生成一个 model 目录，用于存放生成的 model 文件
		Mode: gen.WithDefaultQuery,
	})

	// 生成 model
	g.ApplyBasic(model.Passport{}, model.User{})

	// 生成 query
	g.ApplyInterface(func(model.Method) {}, model.User{})

	// 生成 query 实现
	g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	// Execute 执行
	g.Execute()
}
