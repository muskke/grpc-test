# grpc-test

## 测试grpc各版本的服务兼容性

---


#### v1:初始协议
```shell
2022-04-13 23:48:03  file=v1/client.go:32 level=info DoTestV1 req: [name:"Test" age:100 others:"height: 180cm,weight: 65.0kg"]
2022-04-13 23:48:03  file=v1/client.go:39 level=info DoTestV1 rsp: [version:20 code:200 data:{key:"name" value:"Test"} data:{key:"others" value:"height: 180cm,weight: 65.0kg"}]
```
> 结果：数据正常

#### v2:增加、弃用、插入字段
```shell
2022-04-13 23:47:28  file=v2/client.go:34 level=info DoTestV1 req: [name:"Test" age:100 height:180.5 others:"height: 180cm,weight: 65.0kg" weight:65]
2022-04-13 23:47:28  file=v2/client.go:41 level=info DoTestV1 rsp: [version:VERSION_2 code:200 data:{key:"name" value:"Test"} data:{key:"others" value:"height: 180cm,weight: 65.0kg"} height:180.5 weight:65]
```
> 结果：数据正常

#### v3:交换、修改序号、移除字段
* v1:
```shell
#server：
Handler.Test - req: [name:"Test"  age:100  100:"height: 180cm,weight: 65.0kg"]
Handler.Test - rsp: [version:VERSION_2  code:200  data:{key:"name"  value:"Test"}]
Handler.Test - req: [name:"Test"  age:100  100:"height: 180cm,weight: 65.0kg"]
Handler.Test - rsp: [version:VERSION_2  code:200  data:{key:"name"  value:"Test"}]
#client：
DoTestV1 req: [name:"Test" age:100 others:"height: 180cm,weight: 65.0kg"]
DoTestV1 rsp: [version:20 code:200 data:{key:"name" value:"Test"} 4:0x3ff199999999999a]
GrpcClient.Test rsp: [version:20 code:200 data:{key:"name" value:"Test"} 4:0x3ff199999999999a]
```
> 结果：不正常，version的值变成了20，正确应为VERSION_2。结尾多了未能解析的乱码。

* v2:
```shell
#server:
Handler.Test - req: [name:"Test"  age:100  height:180.5  5:0x42820000]
Handler.Test - rsp: [version:VERSION_2  code:200  height:180.5  data:{key:"name"  value:"Test"}]
Handler.Test - req: [name:"Test"  age:100  height:180.5  5:0x42820000]
Handler.Test - rsp: [version:VERSION_2  code:200  height:180.5  data:{key:"name"  value:"Test"}]
#client:
DoTestV2 req: [name:"Test"  age:100  height:180.5  weight:65]
DoTestV2 rsp: [version:VERSION_2  code:200  data:{key:"name"  value:"Test"}  4:0x3ff199999999999a]
GrpcClient.Test rsp: [version:VERSION_2  code:200  data:{key:"name"  value:"Test"} 4:0x3ff199999999999a]
```
> 结果：不正常：server端缺失weight属性，client缺失了height和weight属性，两端都多了未解析的乱码

* v3:
```shell
#server:
Handler.Test - req: [name:"Test"  age:100  height:180.5  weight:65]
Handler.Test - rsp: [version:VERSION_2  code:200  height:180.5  weight:65  data:{key:"name"  value:"Test"}]
Handler.Test - req: [name:"Test"  age:100  height:180.5  weight:65]
Handler.Test - rsp: [version:VERSION_2  code:200  height:180.5  weight:65  data:{key:"name"  value:"Test"}]
client:
DoTestV3 req: [name:"Test"  age:100  height:180.5  weight:65]
DoTestV3 rsp: [version:VERSION_2  code:200  height:180.5  weight:65  data:{key:"name"  value:"Test"}]
GrpcClient.Test rsp: [version:VERSION_2  code:200  height:180.5  weight:65  data:{key:"name"  value:"Test"}]
```
> 结果：正常

#### 结论
* 修改字段序号，会导致协议无法正确解析
* 移除字段，服务端会忽略该字段
* 交换序号，将导致两个字段都无法解析
* 只要字段和序号能够对应上，就没问题