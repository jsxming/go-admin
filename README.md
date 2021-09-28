<!--
 * @Description: 
 * @Autor: 小明～
 * @Date: 2021-09-15 12:03:29
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-18 09:14:46
-->
## gin react后台开发系统
***
### 包含功能
1. 使用viper读取配置文件
2. JWT token验证获取
3. 使用gin作为web开发框架
4. 使用gorm2x 作为数据库请求库

### MVC 项目结构分层
> routers 负责处理路由、参数验证和请求转发
> Service 业务逻辑的入口，从这里开始请求参数一定是合法的，业务逻辑和业务流程都在这里处理
> Model 数据模型层负责和数据存储打交道
