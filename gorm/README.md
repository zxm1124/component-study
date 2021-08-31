# gorm数据库连接配置
## 实现思路
1. 安装gorm命令
2. 封装实体以及创建引擎方法,方法内设置响应参数
3. init初始化引擎

### 1.安装gorm
`go get -u github.com/jinzhu/gorm@v1.9.12`
### 2.封装实体以及创建引擎方法
```go
// 可将DB配置信息抽取出来，封装到全局变量实体中
// 作为参数传入函数，读取配置信息
func NewDBEngine() (*gorm.DB, error){
	url := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root", "root",3306, "dorm", "utf8", "True")
	db, err := gorm.Open("mysql", url)
	if err != nil{
		return nil, err
	}
	// 判断是否未debug模式，如果是则启用
	db.LogMode(true)
	return db, err
}
```

3. init初始化
```go
func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine()
	if err != nil{
		return err
	}
	return nil
}
// 初始化
func init(){
    err := setupDBEngine()
    if err != nil{
    	log.Fatalf("init.setupDBEngine err: %v", err)
    }
}
```