# opentracing

 试想一个简单的web网站。当用户访问你的首页时，
 web服务器发起两个HTTP调用，其中每个调用又访问了数据库。
 这个过程是否简单直白，我们可以不费什么力气就能发现请求缓慢的原因。
 如果你考虑到调用延迟，你可以为每个调用分布式唯一的ID，并通过HTTP头进行传递。
 如果请求耗时过长，你通过使用唯一ID来grep日志文件，发现问题出在哪里。
 现在，想想一下，你的web网站变得流行起来，你开始使用分布式架构，你的应用需要跨越多个机器，
 多个服务来工作。随着机器和服务数量的增长，日志文件能明确解决问题的机会越来越少。
 确定问题发生的原因将越来越困难。这时，你发现投入调用流程追踪能力是非常有价值的。
 
 
 一个span可以和一个或者多个span间存在因果关系。
 OpenTracing定义了两种关系：ChildOf 和 FollowsFrom。
 `这两种引用类型代表了子节点和父节点间的直接因果关系`。
 未来，OpenTracing将支持非因果关系的span引用关系。
 （例如：多个span被批量处理，span在同一个队列中，等等）
 
 `ChildOf 引用`: 一个span可能是一个父级span的孩子，即"ChildOf"关系。
 在"ChildOf"引用关系下，父级span某种程度上取决于子span。
 下面这些情况会构成"ChildOf"关系：
* 一个RPC调用的服务端的span，和RPC服务客户端的span构成ChildOf关系
* 一个sql insert操作的span，和ORM的save方法的span构成ChildOf关系
* 很多span可以并行工作（或者分布式工作）都可能是一个父级的span的子项，
 他会合并所有子span的执行结果，并在指定期限内返回
 
 ```cassandraql
    [-Parent Span---------]
         [-Child Span----]

    [-Parent Span--------------]
         [-Child Span A----]
          [-Child Span B----]
        [-Child Span C----]
         [-Child Span D---------------]
         [-Child Span E----]
```
 
 `FollowsFrom 引用`: 一些父级节点不以任何方式依然他们子节点的执行结果，这种情况下，我们说这些子span和父span之间是"FollowsFrom"的因果关系。
 "FollowsFrom"关系可以被分为很多不同的子类型，未来版本的OpenTracing中将正式的区分这些类型
 
 下面都是合理的表述一个"FollowFrom"关系的父子节点关系的时序图。
 
 ```cassandraql
    [-Parent Span-]  [-Child Span-]


    [-Parent Span--]
     [-Child Span-]


    [-Parent Span-]
                [-Child Span-]
```


## Logs

每个span可以进行多次Logs操作，每一次Logs操作，都需要一个带时间戳的时间名称，以及可选的任意大小的存储结构。
标准中定义了一些日志（logging）操作的一些常见用例和相关的log事件的键值，可参考Data Conventions Guidelines 数据约定指南。


### Data Conventions 数据约定

#### Span Naming, Span命名

Span 可以包含很多的tags、logs和baggage，但是始终需要一个高度概括的operation name。这些应该是一个简单的字符串，代表span中进行的工作类型。这个字符串应该是工作类型的逻辑名称，例如代表一个RPC或者一次HTTP的调用的端点，亦或对于代表SQL的span，使用SELECT or INSERT作为逻辑名，等等。

其次，span可以存在一个可选的tag，叫做component，他的值可以典型的代表一个进程、框架、类库或者模块名称。这个tag会很有价值。