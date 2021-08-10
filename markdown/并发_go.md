> go：支持管道，信号，socket
### 进程
#### IPC ( Linux ) [ 进程见通信 ]
* 基于通信的IPC方法
  * 数据传送
    * 管道 pipe ----> 字节流
    * 消息队列 mq ---> 结构化数据
  * 共享内存
    * 共享内存区
* 基于信号的IPC方法
  * 操作系统的信号机制 signal ----> 异步
* 基于同步的IPC方法
  * 信号量 semaphore

> goroutine: 不用共享内存来通信，而是以通信来共享内存
### 协程(goroutine)
语言级别(用户级别)的执行体，由用户来参与协程的调度，运行在线程之中，可谓轻量级线程
####  channel
#### 并发机制
#### goroutine模型
* M / machine
  * 即一个内核线程，goroutine运行在其中
* P / processor
  * goroutine的上下文空间，即其所需要的资源。P的数量也代表了真正的并发数
* G / goroutine --- 用户态线程
  * goroutine的运行需要前两者的配合
* 三者依赖
  * 
#####  工作流

#### goroutine调度器的设计策略
##### 复用线程
###### work stealing
###### hand off

##### 利用并行
##### 抢占
##### 全局G队列
