> goroutine: 不用共享内存来通信，而是以通信来共享内存
### 协程(goroutine)
语言级别(用户级别)的执行体，由用户来参与协程的调度，运行在线程之中，可谓轻量级线程
### channel
### 并发机制
#### goroutine模型
* M / machine
  * 即一个内核线程，goroutine运行在其中
* P / processor
  * goroutine的上下文空间，即其所需要的资源。P的数量也代表了真正的并发数
* G / goroutine
  * goroutine的运行需要前两者的配合
* 三者依赖
  * 
#### 工作流
