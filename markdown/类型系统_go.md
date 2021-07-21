
[toc]
## 语言逻辑层
### 接口
> 接口在go语言中的用途

#### 非侵入式
接口是一组方法的集合，结构体只要实现了某个接口声明的全部函数，则表示该结构体实现了该接口，无需显式的继承该接口，在文法上解耦了结构体和接口

#### 接口类型
##### 鸭子类型
一种代码风格化，并不是实际的数据类型。
在鸭子类型中，关注的不是对象的类型本身，而是它是如何使用的。例如，在不使用鸭子类型的语言中，我们可以编写一个函数，它接受一个类型为鸭的对象，并调用它的走和叫方法。在使用鸭子类型的语言中，这样的一个函数可以接受一个任意类型的对象，并调用它的走和叫方法。如果这些需要被调用的方法不存在，那么将引发一个运行时错误。任何拥有这样的正确的走和叫方法的对象都可被函数接受的这种行为引出了以上表述，这种决定类型的方式因此得名。鸭子类型通常得益于不测试方法和函数中参数的类型，而是依赖文档、清晰的代码和测试来确保正确使用。**从静态类型语言转向动态类型语言的用户通常试图添加一些静态的（在运行之前的）类型检查，从而影响了鸭子类型的益处和可伸缩性，并约束了语言的动态特性。**
###### 样例
目的是排序
1. 常规思路是函数判断传入的参数类型为字符串数组或整形数组或者其他类型的数组，再分别进行排序操作，但增加新类型时，需要在函数内部再次变更，扩展性较差。
2. go语言中进行了解耦操作，定义多种结构体分别实现了各自的Sorter接口，在执行sort操作时，形参为Sorter接口类型，使外部sort函数无需变动，若增加新类型，只需定义新的结构体实现Sorter接口即可。
```go
// 定义Sorter接口为通用方法集体
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// 定义两个结构体
type Xi []int
type Xs []string

func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(x Sorter) {
    for i := 0; i < x.Len()-1; i++ {
        for j := i + 1; j < x.Len(); j++ {
            if x.Less(i, j) {
                x.Swap(i, j)
            }
        }
    }
}

```

###### 好处
1. Golang接口是协议、是虚的，有隔离的作用；
2. 能够实现高内聚低耦合高复用，可以防止出现面条式程序；
3. 更容易划分模块和多人开发；
4. 很容易实现各种设计模式；
5. 减少开发量，提高通用性。
##### 空接口
* 可作为所有类型的泛化，可理解为any类型
```go
  // 此时func可以传入任意参数
  func (args ...interface {}) {

  }
```

### 接口赋值
* 场景
  * 将对象赋值给接口
  * 将一个接口赋值给另一个接口
* 要点
  * **谁实现了接口，谁就可以赋值给接口**
### 接口查询
```go
var file1 Writer := ...
if demo, ok := file.(isWriter);ok{
}
```


## 底层
### 类型元数据
* 全局唯一
* 用于描述类型本身
#### 存储内容
* 类型名称
* 类型大小
* 对齐边界
* 是否自定义（区分内置类型和自定义类型）
* ..etc
##### 结构体
<!-- * _type + 其他描述信息 + uncommontype(若为自定义类型) -->
* 内置类型为 _type + 其他描述信息 如下demo1
* 自定义类型为 _type + 其他描述信息 +  ncommontype
```go
// runtime._type
type _type struct {
  size uintptr //
  ptrdata uintptr
  hash uint32
  tflag tflag
  align uint8
  fieldalign uint8
  kind uint8
  ...
}
type uncommontype struct {
  pkgpath nameOff //记录类型所在包路径
  mcount uint16   // 记录该类型关联的方法数目
  ...    uint16   
  moff   uint16   // 方法元数据组成的数组在内存中相对偏移量，用于定位类型绑定的方法
  ...    uint16
}


// 方法类型元数据
type method struct {
  name nameOff
  mtyp typeOff
  ifn textOff
  tfn textOff
}



//demo1
type slicetype struct {
  typ _type
  elem *_type // 指向slice内部存储元素类型的元数据
}
```