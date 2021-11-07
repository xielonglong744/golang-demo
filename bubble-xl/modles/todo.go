package modles

import (
	"bubble-xl/dao"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func TodoCreate(todo *Todo) error {
	err := dao.DB.Create(todo).Error
	return err
}

func TodoDelete(id string) error {
	err := dao.DB.Where("id= ?", id).Delete(Todo{}).Error
	return err
}

func TodoUpdate(todo *Todo) error  {
	err := dao.DB.Save(todo).Error
	return err
}

func TodoAllList()  (todo []*Todo,err error){
	err = dao.DB.Find(&todo).Error
	if err != nil {
		return nil, err
	}
	return
}

func TodoOneList(id string) (todo *Todo,err error) {
	todo = new(Todo)
	err = dao.DB.Where("id=?",id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}