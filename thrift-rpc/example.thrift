namespace py example

struct Student {
    1: string name
    2: string sex
    3: string className
    4: i32   age
    5: string uid
}

struct FormStudent {
    1: string name
    2: string sex
    3: string className
    4: i32   age
}

service BaseService {
    Student getStudentByUID(1:string uid)
    Student modifyStudent(1:string uid,2:FormStudent form)
}