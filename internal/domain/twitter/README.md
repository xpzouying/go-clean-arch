# Twitter Domain Service

这里是一个 Domain Service，有一些跨 feed 和 user domain object 的业务逻辑需要放在这里面做。

比如：

刷 feed 时，我们需要做下列操作：

1. 从 feed 获取 feed 信息。
2. 根据 feed 中的 uid，获取 user 信息。

