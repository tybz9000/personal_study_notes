````
### JSP

- 老玩意，被各种模板引擎替代了
- Java Server Pages
- ${} EL表达式 **可以方便地读取对象中的属性、提交的参数、JavaBean、甚至集合**
- jstl：**JSP Standard Tag Library 即JSP标准标签库**

### Mapping

可以精确比配、路径匹配、扩展名匹配、缺省匹配。优先级递减

路径匹配匹配长的

### ServletContext

Servletshang上下文

### Filter

servlet定义个一个接口

```
public interface Filter {
    void init(FilterConfig var1) throws ServletException;

    void doFilter(ServletRequest var1, ServletResponse var2, FilterChain var3) throws IOException, ServletException;

    void destroy();
}
```

### TomCat

Apache提供的轻量级javaWeb服务器

支持jsp、servlet

**目录**

bin：可执行程序

conf：配置文件

lib：依赖，jar包

log：日志

temp：临时文件

webapps：存放tomcat工程

work：tomcat运行时的数据，session钝化信息，jsp源码等

## SpringMVC

#### MVC

- model：数据、模型
  - dao：数据访问层
  - service：服务层
- view：视图、页面
- controller：控制器
````