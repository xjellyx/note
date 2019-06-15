namespace py example

struct Data {
    1: string text
    2:list<string> stringArr
}

service BaseService {
    Data doFormat(1:Data data)
}