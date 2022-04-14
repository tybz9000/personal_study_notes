# Spring

对spring的理解

框架，容器，基石

- AOP 面向切面编程
- IOC 控制反转

默认单例

- Bean来管理对象

### SPEL

- Spring 表达式语言（简称“SpEL”）是一种强大的表达式语言，支持在运行时查询和操作对象
- 虽然 SpEL 是 Spring 产品组合中表达式评估的基础，但它不直接与 Spring 绑定，可以独立使用
- 基本使用方法

```
SpelExpressionParser parser = new SpelExpressionParser();
Expression exp = parser.parseExpression("'Hello World'"); 
String value = (String) exp.getValue();
```

- 加入些运算

```
Expression exp = parser.parseExpression("'Hello World'.concat('!')");
```

- 对于对象的操作

```
ExpressionParser parser = new SpelExpressionParser();
StandardEvaluationContext context = new StandardEvaluationContext();
Integer i = 12345;
context.setVariable("i", i);
Object value = parser.parseExpression("#i==12346").getValue(context);
return value.toString();
```

### CommandLineRunner

在使用SpringBoot构建项目时，我们通常有一些预先数据的加载。那么SpringBoot提供了一个简单的方式来实现–CommandLineRunner。

# Spring的配置

- Bean的定义信息：BeanDefinition
- 基于xml文件的方式

```
<bean id="myService" class="taiyang.balabala.MyService">
	<property name="username" value="123"/>
</bean>
```

- 基于注解的方式

```
@Component
@Controller
@Service
```



# Spring Boot

约定大于配置

maven坐标都是spring-boot-starter-

banner

spring-boot-starter-parent jar包版本，配置文件定义，资源位置定义

spring-boot-dependencies 依赖位置

spring-boot-starter-parent   springboot项目-parent 

配置文件的扫描在这里配置

sprint-boot-starter- 启动器。springboot会把这些场景都打包成这样的

@SpringBootApplication

[--@SpringBootConfiguration](http://--@SpringBootConfiguration) 配置类

[--@EnableAutoContiguration](http://--@EnableAutoContiguration)

----@AutoConfigurationPackage::AutoConfigurationPackages.Registrar.class 扫描启动器所在的包及其子包下的自定义类。由它来注册的

----@EnableAutoConfigurationImportSelector

#### springboot prifiles

Spring Profiles允许用户根据配置文件来注册bean，针对不同的环境提供不同的功能。

application-{profile}.yml

示例

application-test.yml

application-prod.yml

或

application-data.yml

application-sap.yml

默认主配置文件是application.properties

主配置文件中spring.profiles.active指定激活哪个配置文件

YMS系统中，通过maven的参数，配置将特定环境下的application-{moduleName}.yml打包

打包后通过profile的方式激活

# Spring Security

##### 认证与授权

对标shiro

spring security重量级框架

```
建一个SpringBoot项目，引入Security
进页面就要登录。默认用户名user,密码在启动控制台里
```

实现上是靠过滤器链

```
FilterSecurityInterceptor
ExceptionTransferFilter：异常转义过滤器
UsernamePasswordAuthenticationFilter：用户名密码过滤器
```

过滤器链加载方式

DelegatingFilterProxy

#### 使用

##### UserDetailService接口，自定义账户密码，封装登录对象

##### UsernamePasswordAuthenticationFilter具体进行登录的过滤器，继承它。进行自定义登录验证



#### 用户认证

登录

#### 用户授权

登录之后能够干什么

