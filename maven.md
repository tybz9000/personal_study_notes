# MAVEN

- 自动化构建
- 抽象构建过程
- 依赖管理工具
- 项目管理工具

#### 常用命令

- mvn clean
  - 清理，将target文件夹中的数据清理
- mvn compile
  - 编译，编译代码到target文件夹，生成class文件
- mvn test
  - 运行测试
- mvn package
  - 打包，生成jar文件
- mvn install
  - 清理和安装，将打好的包放入本地仓库，方便其他项目调用
- mvn deploy
  - 发布，发布到私服上去

#### setting.xml

配置仓库、路径、用户名密码等

#### pom文件

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <groupId>分组</groupId>
    <artifactId>坐标量</artifactId>
    <version>版本号</version>
    <name>springBootDemo1</name>
</project>
```

- 根元素

  - project
- 依赖

```xml
<dependencies>
    <dependency>
        <groupId>实际项目</groupId>
　　　　 <artifactId>模块</artifactId>
　　　　 <version>版本</version>
　　　　 <type>依赖类型，默认为jar</type>
　　　　 <scope>依赖范围
         compile，默认，编译、测试、运行三种都有效
         test，测试依赖范围，只对测试classpath有效，编译主代码或者运行时，无法使用此依赖，例如Junit
         provided，编译测试classpath有效，运行时无效，例如servlet-api
         runtime，运行时依赖，例如JDBC驱动实现
        
        </scope>
　　　　 <optional>依赖是否可选</optional>
　　　　 <!—主要用于排除传递性依赖-->
　　　　 <exclusions>
　　　　     <exclusion>
　　　　　　　    <groupId>…</groupId>
　　　　　　　　　 <artifactId>…</artifactId>
　　　　　　　</exclusion>
　　　　 </exclusions>
　　</dependency>
<dependencies>
```

#### 依赖传递

- test依赖不传递

- compile正常传递

- provided，只传递同样为provided的依赖

#### 依赖调解

- 最短路径优先

- 最先声明优先

#### 属性管理

```
<properties>
            <tomcat.version>8.0.42</tomcat.version>
            <env>prod</env>
            </properties>
</properties>
```

  使用${}引用

#### 仓库

- 中央仓库

- 镜像：一般情况下拥有远程仓库的copy

- 本地仓库

#### Modules

- 一个父级项目parent聚合很多子项目
- 除了要打包的项目，其他module（包括parent）应该设置<packaging>pom</packaging>
- 继承
  - 通过<parent></parent>的方式来**继承**
  - 子模块从父模块继承一切东西，包括依赖，插件配置等
  - 主项目的配置作用就是用以给子项目继承
- DepencyManagement & dependencies的区别
  - DepencyManagement只声明了依赖，并不实现引入，因此子项目需要声明要用的依赖
  - dependencies会被继承

  

  
