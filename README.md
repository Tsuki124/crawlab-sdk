# crawlab-db

适配 crawlab-worker 封装的SDK

基本使用

#### 一、Mongo
##### 1. 增
``` Go
//InsertOne
func TestMongoInsert(t *testing.T) {
	vuln := &Vuln{
		Id: 1,
		Title: "反序列化",
		Cve: "CVE-2022-0001",
		Cnvd: "CNVD-2022-0001",
		Cnnvd: "CNNVD-2022-0001",

	}

	table := sdk.Mongo.Db("test").TB("test")
	if err:=table.Insert(vuln);err!=nil {
		trace.PrintError(err)
	}

}
```
##### 2. 删
``` Go
//DeleteALL
func TestMongoDelete(t *testing.T) {
	vuln := map[string]string{
		"cve":"CVE-2022-0001",
	}

	table := sdk.Mongo.Db("test").TB("test")
	if err:=table.Delete(vuln);err!=nil {
		trace.PrintError(err)
	}
}
```
##### 3. 改
``` Go
// UpdateALL
func TestMongoUpdate(t *testing.T) {
	condi := map[string]string{
		"cve":"CVE-2022-0001",
	}
	table := sdk.Mongo.Db("test").TB("test")

	//修改结构体全部变量，未赋初值则置空
	vuln1 := &Vuln{
		Id: 1111,
	}
	if err:=table.Update(vuln1,condi);err!=nil {
		trace.PrintError(err)
	}

	//只修改map中设置的变量
	vuln2 := map[string]interface{}{
		"id":9999,
	}
	if err:=table.Update(vuln2,condi);err!=nil {
		trace.PrintError(err)
	}
}

//UpdateOne or InsertOne
func TestMongoUpsert(t *testing.T) {
	newVuln := &Vuln{
		Id: 1,
		Title: "反序列化",
		Cve: "CVE-2022-0001",
		Cnvd: "CNVD-2022-0001",
		Cnnvd: "CNNVD-2022-0001",

	}
	condi := map[string]string{
		"cve":"CVE-2022-0001",
	}

	table := sdk.Mongo.Db("test").TB("test")
	if err:=table.Upsert(newVuln,condi);err!=nil {
		trace.PrintError(err)
	}
}
```
##### 4. 查
``` Go
//FindOne
func TestMongoFindOne(t *testing.T) {
	condi := map[string]string{
		"cve":"CVE-2022-0001",
	}

	result := &Vuln{}
	table := sdk.Mongo.Db("test").TB("test")
	if err:=table.FindOne(result,condi);err!=nil {
		trace.PrintError(err)
	}else {
		fmt.Println(result)
	}
}

//FindALL
func TestMongoFindALL(t *testing.T) {
	condi := map[string]string{
		"cve":"CVE-2022-0001",
	}

	results := make([]*Vuln,0)
	table := sdk.Mongo.Db("test").TB("test")
	if err:=table.FindALL(&results,condi);err!=nil {
		trace.PrintError(err)
	}else {
		for _,result := range results {
			fmt.Println(result)
		}
	}
}
```

#### 二、SQL
##### 1. 增
``` Go
//Insert
func TestSQLInsert(t *testing.T) {
	vuln := &Vuln{
		Title: "Test SQL Inject",
		Cve: "CVE-2022-0001",
		Cnvd: "CNVD-2022-0001",
		Cnnvd: "CNNVD-2022-0001",
	}

	table := sdk.SQL.Db("test").TB("vuln")
	if err:=table.Insert(vuln);err!=nil {
		trace.PrintError(err)
	}
}
```
##### 2. 删
``` Go
//DeleteALL
func TestSQLDelete(t *testing.T) {
	table := sdk.SQL.Db("test").TB("vuln")
	if err:=table.Delete("id=?",999);err!=nil {
		trace.PrintError(err)
	}
}
```
##### 3. 改
``` Go
//UpdateALL
func TestSQLUpdate(t *testing.T) {
	vuln := &Vuln{
		Cve: "CVE-2022-0001",
	}

	table := sdk.SQL.Db("test").TB("vuln")
	if err:=table.Update(vuln,"id=?",1);err!=nil {
		trace.PrintError(err)
	}
}

```
##### 4. 查
``` Go
//FindOne
func TestSQLFindOne(t *testing.T) {
	vuln := &Vuln{}

	table := sdk.SQL.Db("test").TB("vuln")
	if err:=table.FindOne(vuln,"id=?",1);err!=nil {
		trace.PrintError(err)
	}else {
		fmt.Println(vuln)
	}
}

//FindALL
func TestSQLFind(t *testing.T) {
	vulns := make([]*Vuln,0)
	table := sdk.SQL.Db("test").TB("vuln")
	if err:=table.FindALL(&vulns,"cve=?","CVE-2022-0001");err!=nil {
		trace.PrintError(err)
	}else {
		for _,vuln := range vulns {
			fmt.Println(vuln)
		}
	}
}
```
##### 5. 其他
``` Go
//create the table
func TestSQLCreateTB(t *testing.T) {
	if err:=sdk.SQL.Db("test").CreateTB(&Vuln{});err!=nil {
		trace.PrintError(err)
	}
}

//use the gorm (Not recommended)
func TestUseGorm(t *testing.T) {
	table := sdk.SQL.Db("test").TB("vuln")

	var count int64
	err := table.UseGorm(func(tx *gorm.DB) error {
		return tx.Where("cve=?","CVE-2022-0001").Count(&count).Error
	})
	if err!=nil {
		panic(err)
	}else {
		fmt.Println(count)
	}
}
```


#### 三、Redis
##### 消息订阅与推送
``` Go
func TestRedisPublish(t *testing.T) {
	go func() {
		sdk.Redis.Db().Subscribe("topic", func(msg interfaces.RedisMsg) {
			fmt.Println(msg.GetChannel())
			fmt.Println(msg.GetPattern())
			fmt.Println(msg.GetPayload())
		})
	}()

	for id:=1;id<=10000;id++ {
		_,err := sdk.Redis.Db().Publish("topic",fmt.Sprint(id))
		if err!=nil {
			logrus.Error(id,err)
		}
	}
	time.Sleep(time.Second*3)
}
```

#### SeaweedFS
##### 文件枚举
``` Go
func TestFsLIST(t *testing.T) {
	paths,err := sdk.FS.Path("/vuln").List("cve")
	if err!=nil {
		panic(err)
	}

	for _,path := range paths {
		fmt.Println(path)
	}
}
```

##### 文件下载
``` Go
func TestFsDownload(t *testing.T) {
	content,err := sdk.FS.Path("/vuln/cve").Download("cve/nvdcve-1.1-2022.json")
	if err!=nil {
		panic(err)
	}

	fmt.Println(string(content))
}
```

##### 文件上传
``` Go
func TestFsUpload(t *testing.T) {
	content,_ := os.ReadFile("nvdcve-1.1-2022.json")
	err := sdk.FS.Path("/vuln/cve").Upload("nvdcve-1.1-2022.json",content)
	if err!=nil {
		panic(err)
	}
}
```

##### 文件删除
``` Go
func TestFsDelete(t *testing.T) {
	err := sdk.FS.Path("/vuln/cve").Delete("nvdcve-1.1-2022.json")
	if err!=nil {
		panic(err)
	}
}
```

##### 文件信息
``` Go
func TestFsInfo(t *testing.T) {
	file,err := sdk.FS.Path("/vuln/cve").Info("nvdcve-1.1-2022.json")
	if err!=nil {
		panic(err)
	}

	fmt.Println(file)
}
```