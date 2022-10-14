# 中国餐饮数据大屏显示

---

![p1](./picture/p1.png)

---


## 1. 功能

### 1. 登录页

![p2](./picture/p2.png)

*   (免登录直接跳转数据页)
*   提示预先写入的用户数据(独立模块可拆卸)
*   收集登录信息，提供 3 天没登录选项，提供登录按键
*   提供注册按键

### 2. 注册页

*   登录页的另一种形式

![p3](./picture/p3.png)

*   收集用户信息，提供注册按键
*   提供返回登录按键

### 3. 数据页

![p4](./picture/p4.png)

*   数据显示
*   用户信息显示
*   提供注销按键



## 2. 存在的问题和优化方向

### 1. 功能上

#### 1. 用户类型

*   预设了用户类型管理员，但是没有设计对应的用户管理等功能
    *   `TODO` 单独设置管理员的用户管理功能

#### 2. 数据显示

*   `TODO`  火锅梯队分布的数据可以按照分段具体显示

*   其余数据可以按照具体的时间访问进行访问，但是没有设置

    *   `TODO` 额外添加 input 收集开始和结束的具体时间，进行动态显示

    

### 2. 前端上

*   几乎所有的布局都被设置在一个固定的位置，融资数据的显示采用的 table 没有定制化的
    *   `TODO`优化当前的布局和显示

### 3. 后端上

*   所有数据交互都是线性的
    *   `TODO` 并发执行多个 manager 层函数，配合 Redis 完成对数据库的异步操作
*   `TODO` 建立 Redis 缓存，减少访问mysql的频率