package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    
}

func (s *Student) Name(a string) {
    fmt.Println("name:" + a)
}

func (s *Student) Age(a int) {
    fmt.Println(a + 7)
}

func main() {
    
    var ts interface{}
    
    fmt.Println(reflect.ValueOf(ts).IsValid())
    
    var n string = "shuchao"
    s := &Student{}
    
    t := reflect.ValueOf(s)
    
    // todo 添加isValid判断
    for i := 0; i < t.NumMethod(); i++ {
        //取得所有函数名
        fmt.Println(reflect.TypeOf(s).Method(i).Name)
    }
    
    fn := reflect.ValueOf(s).MethodByName("Name")
    if fn.IsValid() {
        fn.Call([]reflect.Value{reflect.ValueOf(n)})
    }
}
