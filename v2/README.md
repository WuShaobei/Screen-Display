<!--
 * @Author: WuShaobei
 * @Date: 2022-10-10 10:14:23
 * @LastEditTime: 2022-10-10 10:17:33
 * @FilePath: /Screen-Display/v2/README.md
 * @Description: 
-->
# 中国餐饮数据大屏显示

## 后端

### 用户管理功能

#### GORM 模型
```go
type ChineseCateringUser struct {
	Id       uint   `gorm:"primaryKey;type:bigint UNSIGNED not null AUTO_INCREMENT"`
	Username string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Password string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	RealName string `gorm:"VARCHAR(255) NULL DEFAULT NULL"`
	Identity string `gorm:"enum(0,1,2,3) NULL DEFAULT NULL"`
}
```


