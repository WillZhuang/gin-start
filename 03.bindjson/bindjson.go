package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"pwd"`
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}

func main() {

	engine := gin.Default()

	userGroup := engine.Group("/user")
	userGroup.POST("/register", POST_register)

	// http://localhost:8080/hello?name=davie&classes=软件工程
	engine.GET("/hello", GET_hello)
	engine.POST("/register", POST_register)
	engine.POST("/addstudent", POST_addstudent)

	engine.Run()
}

func GET_hello(context *gin.Context) {

	fmt.Println(context.FullPath())

	var student Student
	err := context.ShouldBindQuery(&student)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(student.Name)
	fmt.Println(student.Classes)
	context.Writer.Write([]byte("hello," + student.Name))

}

func POST_register(context *gin.Context) {
	fmt.Println(context.FullPath())
	var register Register
	if err := context.ShouldBind(&register); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(register.UserName)
	fmt.Println(register.Phone)
	context.Writer.Write([]byte(register.UserName + " Register "))

}

func POST_addstudent(context *gin.Context) {
	fmt.Println(context.FullPath())
	var person Person
	if err := context.BindJSON(&person); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("姓名：" + person.Name)
	fmt.Println("年龄：", person.Age)
	context.Writer.Write([]byte(" 添加记录：" + person.Name))
}