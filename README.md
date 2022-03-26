# Somnus_ORM
> 实现一个对象关系映射，通过使用描述对象和数据库之间映射的元数据，将面向对象语言程序中的对象子对持久化到关系数据库中。

## 1.database/sql基础
1.实现一个简单的log库：能够日志分级（Info、Error、Disabled 三级）
2.不同层级日志显示时使用不同的颜色区分
3.显示打印日志代码对应的文件名和行号
```
1.创建两个日志实例分别用于打印info和err，并且绑定输出Printf/Println方法
2.设置日志分级：使用三个便利，通过控制Output来控制日志是否打印
  在设置了Errlevel和infolevel满足条件的时候会被丢弃
```
4.Sessioin和db交互：封装执行语句和变量的写入和各种操作的执行

## 2.对象表结构映射
<em>难度加大了！！</em>

1.Dialect:将Go语言的类型映射为数据库中的类型
```
dialect.go整体注册、管理抽象数据库，抽象
sqlit3.go将Go类型进行了具体SQLite数据映射、判断表是否存在、init()加载注册
```
2.Schema：对象-表的转换，包括域获取、对象解析
3.与数据库交互的session中需要增加dialect和reftable(持久化解析结果，防止连续多次解析相同数据)
4.创建 Engine 实例时，获取 driver 对应的 dialect