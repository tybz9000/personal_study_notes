### Http

- 超本文传输协议

- 无状态：http协议对于事务处理没有记忆能力，服务器不知道客户端是什么状态。打开网页和上一次打开没有任何联系。每次请求都是独立的。

- ##### http与tcp

  - 无状态与有链接
  - Http一直是无状态的
  - TCP一直是有链接的
  - HTTP/1.0 首部字段Connection默认为close,每次HTTP请求都要建立TCP链接
  - HTTP/1.1 首部字段Connection默认为keep-alive，链接可以复用，只要发送端，接口端都没有提出断开链接，则保持tcp链接状态
  - 并行？一个tcp连接同一时刻只能处理一个请求，请求生命周期不可重叠
    - HTTP/1.1添加了管线化
    - HTTP2.0 添加了多路复用

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

### HTTP2

- 同一个域名，浏览器最多同时创建6-8个TCP链接
- 每个TCP连接同时只能处理一个请求-响应。FIFO处理一个再处理一个
  - 管线化，pipelining，批量提交，不等响应就继续提交。但还是一个个提价的
- Http2的优势：
  - 二进制分帧层：以帧为数据传输最小单位，以二进制传输代替明文传输。
  - 在一个TCP连接上，向对方不断发送帧，每个帧标识自己的流，接收后再拼接。
  - 把HTTP/1.1每个请求当做一个流，多个请求变成多个流，交错发给对方。就是HTTP/2的多路复用
  - HTTP/2只需要一个TCP连接
  - 服务端推送，浏览器发送请求，服务器主动向浏览器推送与这个请求相关的资源。
  - Header压缩
  - 应用层重置连接
  - 请求优先级控制
  - 流量控制

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
