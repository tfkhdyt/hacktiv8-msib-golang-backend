package main

var students = []*Student{}

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Grade int32  `json:"grade"`
}

func init() {
	students = append(students, &Student{
		ID:    "s001",
		Name:  "Taufik",
		Grade: 6,
	}, &Student{
		ID:    "s002",
		Name:  "Fauzi",
		Grade: 4,
	})
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(ID string) *Student {
	for _, student := range students {
		if student.ID == ID {
			return student
		}
	}

	return nil
}
