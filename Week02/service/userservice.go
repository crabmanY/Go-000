package service

import (
	"geek_example/week_02/dao"
	"geek_example/week_02/model"
	"github.com/pkg/errors"
	"log"
)

//单纯的获取个人信息,不做什么操作
//如果有错误的话需要抛给上层处理
func GetUserInfoWithoutSomeThing(id int) (*model.Person, error) {
	person := new(model.Person)
	person.UserId = id
	return dao.SelectByExample(person)
}

//获取个人信息之后会有额外的逻辑操作
func GetUserInfoWithSomeThing(id int, reduceNum int) (*model.Person, error) {
	person := new(model.Person)
	person.UserId = id
	record, err := dao.SelectByExample(person)

	//dao层出现错误
	if err != nil {
		return record, err
	}

	//进行后续操作,后续操作可能会出错
	num, reduceErr := reduceInNum(0)

	if reduceErr != nil {
		return nil, errors.Wrap(reduceErr, "reduceNum is error")
	}

	log.Printf("reduceInNum is %d", num)

	return record, err

}

func reduceInNum(num int) (*int, error) {
	switch num {
	case 0:
		return nil, errors.New("num is must not zero")
	default:
		return new(int), nil
	}
}
