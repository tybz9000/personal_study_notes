### Http

- 超本文传输协议

- 无状态：http协议对于事务处理没有记忆能力，服务器不知道客户端是什么状态。打开网页和上一次打开没有任何联系。每次请求都是独立的。

- ```
  无状态与有链接
  Http一直是无状态的
  TCP一直是有链接的
  HTTP/1.0 首部字段Connection默认为close,每次HTTP请求都要建立TCP链接
  HTTP/1.1 首部字段Connection默认为keep-alive，链接可以复用，只要发送端，接口端都没有提出断开链接，则保持tcp链接状态
  ```

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
- 同一个域名下cookie数量也有限制
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
