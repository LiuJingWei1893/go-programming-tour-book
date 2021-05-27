# 结构体

## 定义
```go
type persion struct {
    name string
    age int
}
```
如果结构体的成员数据类型一致，可以写成一行：
```go
type persion struct {
    name string
    age,index int
}
```

## 实例化

```go
var lw persion //返回结构体变量
lw := persion{} //返回结构体变量

lw := new{persion} //返回结构体指针
lw := &persion{} //返回结构体指针
```

## 值类型
struct是值类型，不是引用类型，不同的实例化在赋值时候的效果不一样，`var lw persion`和`lw := persion{}`是个结构体变量，是值，`lw := new{persion}`和`lw := &persion{}`是引用。
```go
var lw persion //返回结构体变量
lw := persion{} //返回结构体变量
用的时候是：
lw.name
lw.age

lw := new{persion} //返回结构体指针
lw := &persion{} //返回结构体指针
用的时候是：
(*lw).name
(*lw).age
```
但对引用类型的，go会做自解析，也就是实例化的时候是引用类型（返回结构体指针），用的时候也是：
```go
lw.age
lw.name

也就是go把lw.age自动解析成了(*lw).age
```
但用的时候还是要明确，到底实例化的时候是值，还是引用。值是copy，引用直接动实例化的struct

## 初始化属性

属性、
成员、
成员变量、
结构体成员、
结构体成员变量、
在这里说的是一个东西；

```go
lw := persion{
    name: tom,
    age: 17,
}
```
如果是初始化全部，可以不写键，只写值, 按照定义顺序写：
```go
lw := persion{
    tom,
    age,
    }
```

## 访问成员
```go
lw.name
lw.name = bob
```

## 匿名结构体
```go
lw := struct{
    name string
    age int
}{
    name: tom,
    age: 17,
}
```