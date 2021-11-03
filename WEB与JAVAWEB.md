# WEB与JAVAWEB

# WEB

### 请求

Get、Post、等

Get：参数数量有限，uri明文显示，不安全

Post：参数数量没有限制，安全

### OSI七层

应用层：Telnet、FTP、HTTP等提供访问网络服务的接口

表示层：提供数据格式的转化服务

会话层：提供端连接冰提供访问验证和会话管理

传输层：提供应用进程之间的逻辑通信TCP、UDP、Socket

网络层：路由器、交换机、Ip

数据链路层：通信实体间建立数据链路连接

物理层：物理通路

#### TCP

- 有连接
- Transmission Control Protocol
- 三次握手

```
SYN：客户端向服务端要求建立连接
SYN&ACK: 服务端要向客户端建立连接，同意客户端的连接。此时客户端明白他的消息可以到服务端，且服务端可以应答
ACK 客户端告知服务端，我能收到你的消息
```

- 全双工通信
- 四次挥手

```
类似三次握手，但是中间两个不同合在一起
客户端：我要停
服务端：我知道你要停了，我还有点事，你等等
服务端：干活ing
服务端：我干完了，我也要停
客户端：我知道你要停了，停吧
```



#### UDP

- UDP User Datagram Protocol
- 无连接

### 会话

一次**会话**：浏览器第一次给服务器发送请求，会话建立。直到有一方断开为止

**会话**技术：解决HTTP无状态的问题

cookie是客户端会话技术

##### forward与redirect

都是请求的转发

forward是直接转发, 转发后再返回

redirect是间接转发, 通知浏览器二次请求

### cookie

```
new Cookie(String name, String value)
httpServletResponse.addCookie();
httpServletRequert.getCookie();
cookie.getName();
cookie.getValue();
cookie.setMaxAge();//正数 持久化 seconds  负数，关浏览器就关。  0，删除cookie
cookie.setPath();//设置作用范围
// /youpin/ticket
cookie.setDomain();//设置域名
// .be.xiaomi.com
```

##### 原理：

- 响应头里写set-cookie头

- 请求返回体里会带cookie

- cookie默认浏览器关闭销毁
- cookie大小在浏览器有限制：一般4kb
- ​	同一个域名下cookie数量也有限制
- cookie一般存储少量不太敏感的数据
- 一般做不登录环境下的身份识别

### session

```
//获取session
request.getSession();
//使用session
Object getAttribute(String name);
void setAttribute(String name, Object value);
void removeAttribute(String name);
```

- 服务器端会话技术
- 一次会话内多次请求是同一个session
- session依赖coookie
  - 第一次的时候，没cookie，创建session对象。将session-id写入cookie
  - 之后由cookie携带的session-id获取session对象
- 客户端关闭，服务端不关闭。session还是同一个么？ 看cookie
- 服务端关闭，session怎么处理，tomcat会处理
  - 钝化：存储到硬盘里
  - 活化：存储在内存中
- session销毁
  - 服务器关闭
  - session调用自杀方法
  - 时间到期

### Http

- 超本文传输协议

- 请求响应协议，应用层

- 默认端口80，https 443

- http 1.1 引入长连接

  **General**

  ```
  Request Url:
  Request Method:'POST'
  Status Code:200 OK
  Remote Address:
  ```

  **Response Header**

  ```
  Cache-Control
  Connection:Keep-Alive 1.1长连接
  Content-Encoding
  Content-Type
  
  ```

  **Request Header**

  ```
  Accept:接受的格式
  Accept-Encoding
  Accept-Language
  Cache-Control
  Connection
  Cookie
  COOKIE_SSION
  HOST
  Refresh:刷新（Js就是作用在这些上了）
  Location：重新定位
  ```

- 响应码

  ```
  200 成功
  3** 重定向
  404 找不到
  5xx 服务器代码错误
  	502 网关错误
  ```

##### content-type

Content-Type（内容类型），一般是指网页中存在的 Content-Type，用于定义网络文件的类型和网页的编码，决定浏览器将以什么形式、什么编码读取这个文件

**application/x-www-form-urlencoded**: form中的内容，以key-value的形式，被吐到服务器

**multipart/form-data**: 上传文件的时候，一般是这个格式

**application/json**: json数据格式

### Servlet

- sun公司用于开发动态web的技术

- jsp编译之后就是servlet

- sun公司提供的一个接口Servlet

- ```
  public interface Servlet {
      void init(ServletConfig var1) throws ServletException;
  
      ServletConfig getServletConfig();
  
      void service(ServletRequest var1, ServletResponse var2) throws ServletException, IOException;
  
      String getServletInfo();
  
      void destroy();
      }
  ```

- HttpServlet继承Servlet  里面有doGet doPost

  ```
  class HttpServlet extends GenericServlet
  
  class MyServlet extends HttpServlet {
  	public void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
         
      }
      
      public void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
         
      }
  }
  ```

- web.xml进行配置

  ```
  <servlet>
  	<servlet-name>MyServlet</servlet-name>
  </servlet>
  <servlet-mapping>
  	<servlet-name>MyServlet</servlet-name>
  	<url-pattern>/mypage/*</url-pattern>
  </servlet-mapping>
  ```

- 运行原理：

- 浏览器发送Http请求，首次访问产生Servlet

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

![](D:\学习文件\图片\v2-debc2a494efdf6242720d2fbb1087286_b.jpg)