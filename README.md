## single u center service

### 已经完成 ：
1. 消息队列broker引入
2. gomicro自动日志。
3. 分布式日志
4. 项目私有化
5. xorm日志入jeager
6. prometheus监控
7. gin网关
8. jwt
9. 角色管理 Casbin 参考：single_ucenter\enlight_ucenter_client\tests\casbin
   1. 角色管理，已经测试，可以写入gateway也可以写入 ucenter
   2. middleware 参考 gateway中
   3. e.AddPermissionForUser("bob", "read") 不可用会导致报错

### 待完成：
 
 1. jwt可以写到common项目，再其他地方解析，生成（目前再gateway中生成和检测，符合基本逻辑）


### 特别注意

jaeger 不支持，gin.context 要用 c:= ctx.Request.Context()获取原生context