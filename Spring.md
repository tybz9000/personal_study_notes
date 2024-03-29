# Spring

对spring的理解

框架，容器，基石

- #### **AOP 面向切面编程**
  
  - Spring-AOP是基于动态代理实现的
  - 创建Bean的时候，实际上是创建的Bean的代理类
  
- #### **IOC 控制反转**

- #### DI 依赖注入

  - IOC容器，放的bean，使用hashmap实现
  - 根据需要，由容器生产对象
    - 需求：XML文件
    - 需求：注解
    - 需求：java配置类


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

### Spring的流程

- 对象的定义信息
  - xml
  - 注解
  - java配置类
- 解析对象定义信息
  - BeanDefinitionReader
- Bean生命周期
  - 实例化
  - 属性设置
  - 初始化
  - 销毁
- 容器：生产对象
  - BeanDefinition：Bean定义信息
    - BeanFactoryPostProcessor：后置处理器，增强器，在工厂之前
  - BeanFactory：创建bean
    - 反射 
      - Constructor con = Class.getConstructor()；
      - Object obj = con.newInstance();
    - BeanPostProcessor：后置处理器，增强器，Bean后期增强器。工厂实例化后，初始化之前
    - 特别的，与FactoryBean的区别，
    - FactoryBean用以定制化Bean
      - 唯一的定制化的复杂Bean
  - Enviroment环境
  - 



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

##### @Primary

多个相同类型的bean，优先注册primary的，否则会报NoUniqueBeanDefinitionException

# Spring Boot

- 脚手架
- 约定大于配置
- 快速构建项目
- 内嵌web容器
- 自动配置，自动管理依赖
- 自带应用监控 spring-boot-actuator

#### 启动流程

- new SpringApplication()

```java
	//1
	public SpringApplication(ResourceLoader resourceLoader, Class<?>... primarySources) {
        //初始化配置，配置source
    	this.sources = new LinkedHashSet();
        this.bannerMode = Mode.CONSOLE;
        this.logStartupInfo = true;
        this.addCommandLineProperties = true;
        this.addConversionService = true;
        this.headless = true;
        this.registerShutdownHook = true;
        this.additionalProfiles = new HashSet();
        this.isCustomEnvironment = false;
    	//资源加载器，启动时为空
        this.resourceLoader = resourceLoader;
        Assert.notNull(primarySources, "PrimarySources must not be null");
    	//从启动类进来时primarySources就是启动类
        this.primarySources = new LinkedHashSet(Arrays.asList(primarySources));
    	//应用程序类型：REACTIVE、SERVLET(通常)、NONE
        this.webApplicationType = WebApplicationType.deduceFromClasspath(); 
		//创建初始化构造器, 得到所需工厂集合的实例 来自 META-INF/spring.factories goto1
		this.setInitializers(this.getSpringFactoriesInstances(ApplicationContextInitializer.class));
    	//创建应用监听器 来自 META-INF/spring.factories
       this.setListeners(this.getSpringFactoriesInstances(ApplicationListener.class));
    	//配置应用的主方法所在的类，找到main方法//4
    	this.mainApplicationClass = this.deduceMainApplicationClass();
    }

	//2
    private <T> Collection<T> getSpringFactoriesInstances(Class<T> type, Class<?>[] parameterTypes, Object... args) {
        //获取类加载器
        ClassLoader classLoader = this.getClassLoader();
        //从配置文件中，获取springFactory的名字 goto3
        Set<String> names = new LinkedHashSet(SpringFactoriesLoader.loadFactoryNames(type, classLoader));
        //然后实例化 反射
        List<T> instances = this.createSpringFactoriesInstances(type, parameterTypes, classLoader, args, names);
        AnnotationAwareOrderComparator.sort(instances);
        return instances;
    }

	//3
	public static List<String> loadFactoryNames(Class<?> factoryClass, @Nullable ClassLoader classLoader) {
        //这个name传了两次
        //第一次是ApplicationContextInitializer
        //第二次是ApplicationListener
        String factoryClassName = factoryClass.getName();
        //loadSpringFactories这个方法，去读了spring.factories
        //从spring.factories拿到我们这次要的factoryNames
        return (List)loadSpringFactories(classLoader).getOrDefault(factoryClassName, Collections.emptyList());
    }

	private Class<?> deduceMainApplicationClass() {
        try {
            //new了个异常，根据堆栈信息去找main方法，牛逼呀
            StackTraceElement[] stackTrace = (new RuntimeException()).getStackTrace();
            StackTraceElement[] var2 = stackTrace;
            int var3 = stackTrace.length;

            for(int var4 = 0; var4 < var3; ++var4) {
                StackTraceElement stackTraceElement = var2[var4];
                if ("main".equals(stackTraceElement.getMethodName())) {
                    return Class.forName(stackTraceElement.getClassName());
                }
            }
        } catch (ClassNotFoundException var6) {
        }

        return null;
    }

```

- spring.factories有一堆，逐个获取

```
spring.factories 一个key = 一堆value，全都拿出来
spring-boot自己的有8个
//这次要加载的就是这堆东西
org.springframework.context.ApplicationContextInitializer=\
org.springframework.boot.context.ConfigurationWarningsApplicationContextInitializer,\
org.springframework.boot.context.ContextIdApplicationContextInitializer,\
org.springframework.boot.context.config.DelegatingApplicationContextInitializer,\
org.springframework.boot.web.context.ServerPortInfoApplicationContextInitializer

```



- SpringApplication.run()
  - 启动应用

```java
public ConfigurableApplicationContext run(String... args) {
    
        //1、这个是一个计时器，没什么好说的
		StopWatch stopWatch = new StopWatch();
		stopWatch.start();
    	//应用上下文
		ConfigurableApplicationContext context = null;
    	//异常报告
		Collection<SpringBootExceptionReporter> exceptionReporters = new ArrayList<>();
		
    
        //2、这个也不是重点，就是设置了一些环境变量-->
        configureHeadlessProperty();
 
 
        //3、获取事件监听器SpringApplicationRunListener类型，并且执行starting()方法-->
		SpringApplicationRunListeners listeners = getRunListeners(args);
    	//抛出一个starting事件
		listeners.starting();
 
		try {
 
 
            //4、把参数args封装成DefaultApplicationArguments，这个了解一下就知道-->
			ApplicationArguments applicationArguments = new DefaultApplicationArguments(
					args);
 
            //5、这个很重要准备环境了，并且把环境跟spring上下文绑定好，并且执行environmentPrepared()方法-->
            //6、判断一些环境的值，并设置一些环境的值-->
			ConfigurableEnvironment environment = prepareEnvironment(listeners,
					applicationArguments);
 
            //忽略的bean
			configureIgnoreBeanInfo(environment);
 
            //7、打印banner-->
			Banner printedBanner = printBanner(environment);
 
 
            //8、创建ApplicationContext容器，即我们说的IOC容器【重要】
			context = createApplicationContext();
 
 
            //9、获取异常报告事件监听-->
			exceptionReporters = getSpringFactoriesInstances(
					SpringBootExceptionReporter.class,
					new Class[] { ConfigurableApplicationContext.class }, context);
 
 
            //10、准备上下文，执行完成后调用contextPrepared()方法,contextLoaded()方法-->
            //enviroment 一些profiles，来自application.yml
            //context 应用上下文
			prepareContext(context, environment, listeners, applicationArguments,
					printedBanner);
 
 
            //11、这个是spring启动的代码了，这里就回去里面就回去扫描并且初始化单实列bean了-->
            //这个refreshContext()加载了bean，还启动了内置web容器，需要细细的去看看
            //调用到了AbstractApplicationContext的refresh。spring的bean加载
			refreshContext(context);
 
            //12、啥事情都没有做-->
			afterRefresh(context, applicationArguments);
			stopWatch.stop();
			if (this.logStartupInfo) {
				new StartupInfoLogger(this.mainApplicationClass)
						.logStarted(getApplicationLog(), stopWatch);
			}
 
    
            //13、执行ApplicationRunListeners中的started()方法-->
			listeners.started(context);
 
            //执行Runner（ApplicationRunner和CommandLineRunner）-->
			callRunners(context, applicationArguments);
		}
		catch (Throwable ex) {
			handleRunFailure(context, listeners, exceptionReporters, ex);
			throw new IllegalStateException(ex);
		}
		listeners.running(context);
		return context;
    
    //10 准备上下文
    private void prepareContext(ConfigurableApplicationContext context, ConfigurableEnvironment environment, SpringApplicationRunListeners listeners, ApplicationArguments applicationArguments, Banner printedBanner) {
        context.setEnvironment(environment);//准备环境
        this.postProcessApplicationContext(context);//前置处理，准备一些提前需要准备的上下文内容。beanNameGenerator resourceLoader等
        this.applyInitializers(context);
        listeners.contextPrepared(context);//【事件】上下文ready了
        if (this.logStartupInfo) {
            this.logStartupInfo(context.getParent() == null);
            this.logStartupProfileInfo(context);
        }

        ConfigurableListableBeanFactory beanFactory = context.getBeanFactory();//获取关键工厂！ DefaultListableBeanFactory
        beanFactory.registerSingleton("springApplicationArguments", applicationArguments);//注册单例对象
        if (printedBanner != null) {
            beanFactory.registerSingleton("springBootBanner", printedBanner);//注册单例对象
        }

        if (beanFactory instanceof DefaultListableBeanFactory) {
            ((DefaultListableBeanFactory)beanFactory).setAllowBeanDefinitionOverriding(this.allowBeanDefinitionOverriding);
        }

        Set<Object> sources = this.getAllSources();
        Assert.notEmpty(sources, "Sources must not be empty");
        this.load(context, sources.toArray(new Object[0]));//关键方法，扫描到启动类上面的@Component注解，注册为bean
        listeners.contextLoaded(context);//【事件】主类bean加载好了
    }
```

- 怎么理解Enviroment
  - 环境，包含系统环境（jdk环境，系统环境）等
- ApplicationContext是继承BeanFactory的
  - BeanFactory定义了方法，getbean
- 怎么理解上下文
  - 容器



#### 配置类

- 广义的配置类：@Component直接或间接修饰的类，即是Spring组建
  - @Service 注解被@Component修饰
- 狭义的配置类：@Configuration类

@Configuration

一个Configuration等同于一个spring的xml配置

@Bean

一个Bean等同一个spring的xml配置里一个bean

- Auto-Configuration自动配置
- Autowire自动装配

#### maven坐标

maven坐标都是spring-boot-starter-

- spring-boot-starter-parent jar包版本，配置文件定义，资源位置定义

- spring-boot-dependencies 依赖位置

- spring-boot-starter-parent   springboot项目-parent 


配置文件的扫描在这里配置

- sprint-boot-starter- 启动器。springboot会把这些场景都打包成这样的


@SpringBootApplication

@SpringBootConfiguration配置类

@EnableAutoContiguration

@AutoConfigurationPackage::AutoConfigurationPackages.Registrar.class 扫描启动器所在的包及其子包下的自定义类。由它来注册的

@EnableAutoConfigurationImportSelector

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

