package stuctpack

//for Storing Student Details
type Student struct{
	Id int 
	name string 
	department string 
	age uint 
	isAvailable bool 
}

var students []Student

func Add(id int, name string, department string, age uint, isAvailable bool){
	var student1 Student
	student1.Id = id
	student1.name = name
	student1.department = department
	student1.age = age
	student1.isAvailable = isAvailable
	students = append(students, student1)
}

func Get() []Student{
	return students
}


// func main(){
// 	// ans := B.Palindrome("malayala")
// 	// fmt.Println(ans)

// 	Add(1,"Sakthi","Computer Science",19,true)
// 	Add(2,"Ram"," Science",20,false)
// 	Add(3,"Thiru","Mechanical",21,false)
// 	Add(4,"Siva","Civil",19,true)

// 	Student := Get();
// 	fmt.Printf("%v",Student)
// }