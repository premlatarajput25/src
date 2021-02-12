package dataAccess

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/premlatarajput25/StudentProject/beans"
)

type StudentAccess struct {
	Db *sql.DB
}

func (studentAccess StudentAccess) GetAllStudents() (students []beans.Student, err error) {

	rows, err := studentAccess.Db.Query("select * from students")
	if err != nil {
		log.Println("111111  " + err.Error())
		return nil, err
	}
	for rows.Next() {
		var id int64
		var name string
		var age int
		var class int
		err2 := rows.Scan(&id, &name, &age, &class)
		if err2 != nil {
			log.Println("22222  " + err.Error())
			return nil, err2
		}
		student := beans.Student{id, name, age, class}
		students = append(students, student)

	}
	return students, err
}
func (studentAccess StudentAccess) UpdateStudent(student *beans.Student) (int64, error) {
	result, err := studentAccess.Db.Exec("update students set name = ?,age =?,class =? where id =?", student.Name, student.Age, student.Class, student.Id)
	if err != nil {
		log.Println("111111  " + err.Error())
		return 0, err
	}
	rowsAffected, _ := result.RowsAffected()

	return rowsAffected, nil

}
func (studentAccess StudentAccess) InsertStudent(student *beans.Student) (err error) {
	// log("Adding  Student ")
	result, err := studentAccess.Db.Exec("insert into students(age,class,name) values(?,?,?)", student.Age, student.Class, student.Name)
	if err != nil {
		log.Println("111111  " + err.Error())
		return err
	} else {

		id, err1 := result.LastInsertId()
		student.Id = id
		fmt.Printf("id of student %d \n", uint64(id))
		if err1 != nil {
			return err1
		}

	}

	return err
}
func (studentAccess StudentAccess) DeleteStudent(id int64) (int64, error) {
	// log("Removing  Student ")
	result, err := studentAccess.Db.Exec("delete   from students where id = ?", id)
	if err != nil {
		log.Println("111111  " + err.Error())
		return 0, err
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, err
}

// func (studentAccess StudentAccess) getStudentWithId(id int64) (students beans.Student, err error) {

// 	row, err := studentAccess.Db.Query("select * from students where id =?", id)
// 	if err != nil {
// 		log.Println("111111  " + err.Error())
// 		return nil, err
// 	}
// 	// var student beans.Student
// 	for row.Next() {
// 		var id int64
// 		var name string
// 		var age int
// 		var class int
// 		err2 := row.Scan(&id, &name, &age, &class)
// 		if err2 != nil {
// 			log.Println("22222  " + err.Error())
// 			return nil, err2
// 		}
// 		fmt.Println("yes  : ", id, name, age, class)
// 		student := beans.Student{id, name, age, class}
// 		students = append(students, student)

// 	}
// 	return &students, err
// }
