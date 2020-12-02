package dao

import (
	"geek_example/common/db"
	"geek_example/week_02/model"
	"github.com/pkg/errors"
)

var dbConnector = db.Db

//新增
func InsertRecord(person *model.Person) error {

	dbConnector.AutoMigrate(person)
	create := dbConnector.Create(person)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

//根据条件查询
func SelectByExample(person *model.Person) (*model.Person, error) {
	record := new(model.Person)
	find := dbConnector.Where(person).Find(record)
	if find.Error != nil {
		return nil, errors.Wrap(find.Error, "SelectByExample is error")
	}
	defer find.Close()

	return record, nil
}

//更新
func UpdateById(person model.Person) error {

	update := dbConnector.Model(&model.Person{}).Update(person)

	if update.Error != nil {
		return errors.Wrap(update.Error, "UpdateById is error")
	}
	defer update.Close()

	return nil
}

//根据id删除
func DeleteById(id int) error {

	delete := dbConnector.Where("id = " + string(id)).Delete(&model.Person{})
	if delete.Error != nil {
		return errors.Wrap(delete.Error, "DeleteById is error")
	}
	defer delete.Close()

	return nil

}
