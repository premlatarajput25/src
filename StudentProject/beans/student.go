package beans

import "fmt"

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class int    `json:"class"`
	//subject Subject
}

func (student Student) toString() string {
	return fmt.Sprintf("Id : %d, Name : %s, Age : %d, Class : %d",
		student.Id,
		student.Name,
		student.Age,
		student.Class)
}
