# Viper 读取配置文件

## 实现思路

1. 导入viper
   ` go get -u github.com/spf13/viper@v1.4.0`
2. 创建Setting结构体，把 *viper.Viper 作为成员
3. 创建New方法，配置viper
4. 创建一个成员函数，作为读取yaml各扇区中的配置信息
### 1.创建viper结构体以及相关方法
```go
type Setting struct{
	viper *viper.Viper
}
func NewSetting() (*Setting, error){
    vp := viper.New()
    // 配置文件名
    vp.SetConfigName("config")
    // 配置路径
    vp.AddConfigPath("configs")
    // 配置后缀
    vp.SetConfigType("yaml")
    // 读取配置文件
    err := vp.ReadInConfig()
    if err != nil{
    return nil, err
    }
    return &Setting{
    viper: vp,
    }, nil
}
func (s Setting) ReadSection(configName string, v interface{}) error {
    err := s.viper.UnmarshalKey(configName, v)
    if err != nil{
    return err
    }
    return nil
}
```

### 2.创建全局变量
1. 在全局变量包中的section定义要写入信息的结构体
2. 在global中声明全局变量
```go
// global/section.go中定义结构体
type APP struct {
    Name string
    Version string
    Port int
    Suffix string
}
// global/global.go中声明全局变量
var(
    APPSetting *APP
)
```

### 3.在main入口定义相关函数
1. 定义setupSetting主入口函数，用于读取各扇区的配置信息
2. 在init中启用setupSetting函数
```go
// 读取配置文件信息
func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil{
		return err
	}
	// 读取APP配置文件信息
	err = setting.ReadSection("APP", &global.APPSetting)
	if err != nil{
		log.Fatalf("setupSetting err: %v",err)
		return err
	}
	return nil
}

// init函数
func init(){
    err := setupSetting()
    if err != nil{
    	log.Fatalf("setupSetting err: %v",err)
    }
}

// main函数打印是否注入
func main() {
    fmt.Println(global.APPSetting)
}
```