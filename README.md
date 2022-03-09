# 整洁架构 ( Go )

## <a id="qianyan">前言</a>

主要围绕如下几点展开讨论，

1.  **分层规范**
    1.  服务分为哪几层
    2.  每一层的职责是什么
2.  **目录结构**
    1.  每一层拆分后，水平是否需要在进行分层
    2.  公共的功能放在什么地方
3.  **工程化规范**
    1.  errors 统一处理规范
    2.  API 接口的规范
4.  **其他**
    1.  如何保障测试
    2.  错误日志太多：规范日志的输出
    3.  告警太多，导致几乎没人再去关注告警

## 分层规范

从根目录，开始分为：

-   `/api`

    **职责：**

    -   定义接口协议 DTO 结构体的定义
    -   路由注册
    -   讲请求/响应序列化和反序列化

    **测试方法：**

    -   使用 [httptest](https://pkg.go.dev/net/http/httptest) 模拟 http server ， 然后请求接口。

    **测试关注点：**

    -   请求是否被转换成预定义的 请求结构体
    -   响应是否被转换成预期的 响应结构体

-   `/cmd`

    **职责**

    -   main.go 存放的目录。
    -   各种依赖的初始化。
    -   `func main` 函数中，做服务的初始化、依赖注入。

    **测试**

    不做测试。

    由于依赖比较多，所以会导致函数会很长。依赖注入的过程也比较繁琐，可以借助 [wire](https://pkg.go.dev/github.com/google/wire) 工具直接生成相应的代码。

-   `/internal`

    强制增加 `/internal` package，防止随意引用。

    -   可以避免循环引用的问题
    -   规范调用关系，如果不是我们自己服务的调用的话，那么就应该使用 rpc 的调用方式。

    <br />

    在这下面会存在如下目录，

    -   `/server` - http server, grpc server 的定义。

        里面依赖多个 service，每一个 service 算是未来的一个「微服务」能力。比如：UserService、FeedService 等等。

        这里的难点就是要如何定义好各种各样的 service 。

        **职责：**

        -   创建 http server，**管理 http server 的生命周期**。 ( 重点 )
        -   ( 类似于 grpc ) 使用 Register 的方式将 server 注入到 api 中，绑定 server 与 router 的关系。

        **测试：** 暂时不需要测试。

    -   `/service`

        **调用关系：** service —> usecase 中的 Usercase。

        **职责：**

        1.  **参数转换**，并做简单的参数校验；
        2.  这里面只做编排，不做任何业务逻辑。
        3.  做业务数据的渲染； ( 由于没有 BFF，所以将 BFF 的功能放到这一层做，但是会导致这一层的代码膨胀 )

    -   `/domain` - 保存 domain 级别的对象，其中包含：`domain object` 、 `value object` 、 `domain service` 。 domain object 里面包含各自负责领域的业务逻辑。

        这一层是业务的核心层级。

        **职责**

        -   **包含**
            -   具体的业务逻辑。
            -   这里面设计对象 ( Domain Object ) 的，可以将业务逻辑放到 Domain Object 中。参考文章： [链接 1](https://blog.csdn.net/abchywabc/article/details/79362975) ， [阿里技术专家详解 DDD 系列 第五讲：聊聊如何避免写流水账代码](https://zhuanlan.zhihu.com/p/366395817)
        -   **不包含**：UI 渲染；数据库或 RPC 框架的具体实现。
        -   这一层按照现在的分层模式，非常独立，不会向上依赖，也不会向下依赖。

        **测试：**

        -   repo 的依赖，由于是 interface 注入，所以直接 mock 的方式。 ( 后续会引入 Go 官方的 [gomock](https://pkg.go.dev/github.com/golang/mock/gomock) )
        -   测试重点为业务逻辑是否符合预期。

        **难点：**

        这一层的难点是，如何定义各种各样的业务用例 ( Usecase ) 。比如，

    -   `/usecase` - Use Cases，即 DDD 中的 `Application Service`，它主要的作用是对 domain 业务的编排。

    -   `/repo`

        **职责：** 各种数据依赖的具体实现方式。数据访问层，包括 DB、RPC、缓存等。这里面存放 PO 数据，这些数据就是 **简单的表映射**。

-   `/pkg`

    里面定义可以共享出去的工具。由于是可以直接让别人用，这里面的 package 当作基础依赖库使用。既然又是基础依赖库，它里面尽可能的不包含第三方依赖。

## 参考资料

-   [Microsoft - Design a DDD-oriented microservice](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/ddd-oriented-microservice)

-   https://github.com/bxcodec/go-clean-arch - 按照 Bob 大叔的整洁代码架构分层。

    > 简单的示例，看看就好，在复杂的业务场景下没有太多的参考价值。

-   [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Bob 大叔的 Clean 架构

    -   中文翻译：**[架构整洁之道](https://www.cnblogs.com/yjf512/archive/2012/09/10/2678313.html)**

-   https://github.com/alibaba/COLA - 阿里巴巴的 Clean Architecture 的示例。

    ![COLA 的架构](https://camo.githubusercontent.com/9541f7afd632322da151e2555d2529e254b18eadefadb86b9d743953a35298ce/68747470733a2f2f696d672d626c6f672e6373646e696d672e636e2f32303230313230393138323933343833382e706e67)

    这里的分层的数据讲的还是比较清晰。

-   https://github.com/manakuro/golang-clean-architecture - Golang 的一个 DDD 版本

    > 个人觉得不好，对于实际的场景来说，太复杂！

    推荐来源： [DDD 之代码架构](https://www.yasinshaw.com/articles/112) ，该文章总结了几个重点问题：

    -   聚合根的定义：永远的难题，很难定义。

-   [Domain Driven Design: Domain Service, Application Service](https://stackoverflow.com/questions/2268699/domain-driven-design-domain-service-application-service)

    在这里面介绍了什么是 Domain Service：

    -   集合了 domain object 的业务逻辑和业务规则，但是又无法将这些逻辑放到具体的对象里面。
    -   Domain Service 是不存在或最好不存在状态的。但是可以改变 Domain Object 中的状态。

-   [https://mp.weixin.qq.com/s/Xzlt_WcdcfLWhofafY3c6g](https://mp.weixin.qq.com/s/Xzlt_WcdcfLWhofafY3c6g)

    > 腾讯的整洁架构实践，个人觉得不是很好。最致命的一个原因是：直接使用了一个通用的 `model`，从前到后传递，这个就非常不合理。
    > 首先对于数据的职责定义还是需要做区分，比如需要分成： `DTO`、 `Domain Object` 、 `PO` 等数据类型，因为每个数据类型承担的职责起码就不一样。
    > 而文中直接定义的 `model` 就很类似于传统 `MVC` 的那一套，而 Bob 大叔在他的书里面就明确说了这是两套方法论。
    > 另外，Bob 大叔也说了，很多对象看着是类似或者一样，但是其实是两个东西。
