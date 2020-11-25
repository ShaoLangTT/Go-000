学习笔记


[文章书籍问题点记录](https://shimo.im/docs/GYvDrQT8qW8RgkGY)<br />  
---------------------------------------------------------------

KnowledgePoint
---------------------
**1.Nginx -> API GetWay -> BFF -> SERVER**

Nginx:进行反向代理

API GetWay:负责的是下游服务，统一收敛的公共能力，比如你提到的鉴权，防水墙(安全问题),缓存，
          频控等等。公共能力如果每次变更，下游所有服务都需要在发布一般，所以才有了 gw 这一层。 
          并且这一层可以各个接口的监控，大盘的整体健康度等。所以是很有必要的
          
BFF:数据组装层（这里可以有多个BFF，每个BFF可以做不同的事情）

SERVER：单个服务应用（比如：用户服务，订单服务）

现在一个订单请求进来，先经过Nginx，再通过API GetWay进行相关鉴权，再到相关的BFF层,BFF层会调用用户服务和订单服务的
相关数据接口，进行聚合返回给客户端


### https主要用于传输层的数据安全，签名用于应用层防止数据篡改


###### 发送时要保守，接收时要开放：是指系统和外部系统做交互时的一个指导性原则，和具体技术无关。这里可以理解为，对外部系统应该尽可能的宽容(开放)，对内应当尽可能的严苛(保守)，是从人类社会交往的"严于律己宽于律人"原则类比而来


**Documents**<br /> 
-------------------
1.[Kubernetes教程](https://www.kuboard.cn/learning/)<br />  
2.[gRPC 官方文档中文版](http://doc.oschina.net/grpc)<br />  
3.[grpc-go](https://github.com/grpc/grpc-go)<br />  
4.[Golang gRPC实践 连载](https://segmentfault.com/a/1190000007880647)<br />  
5.[如何通过gRPC实现高效远程过程调用？](https://time.geekbang.org/column/article/247812)<br />  
6.[《跟煎鱼学 Go》](https://eddycjy.com/go-categories/)<br />  
7.[四层和七层负载均衡的区别](https://kb.cnblogs.com/page/188170/)<br />  



