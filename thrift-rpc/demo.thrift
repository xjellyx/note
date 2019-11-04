namespace py example

// 一个data测试
struct Data {
    1: string text
    2:list<string> stringArr
}

// 学生信息
struct Student {
    1: string name
    2: string sex
    3: string className
    4: i32   age
    5: string uid
}

// 输入的form
struct FormStudent {
    1: string name
    2: string sex
    3: string className
    4: i32   age
}

// 基本方法
service BaseService {
    Data getData(1:Data data)
    Student getStudentByUID(1:string uid)
    Student modifyStudent(1:string uid,2:FormStudent form)
}


