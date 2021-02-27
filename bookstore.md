> 参考的是官网文档 https://www.yuque.com/tal-tech/go-zero/rm435c



## 目的

1. 构建一个对外的api 提供书的新增和查询
2. 构建rpc服务提供给api调用
3. 写入数据库
4. 服务用etcd发现



## 开始

### api服务

项目目录bookstore

新建api文件夹 存放api服务文件

```
cd bookstore
gomod init bookstore
cd api
goctl api -o bookstore.api
```

bookstore.api

```
syntax = "v1"

info(
	title: "title"
	desc: "desc"
	author: "brady"
	email: "brady.wang@qq.com"
)

type (
	AddBookReq {
		Name  string `form:"name"`
		Price int64  `form:"price"`
	}

	AddBookResp {
		Ok bool `json:"ok"`
	}
)

type (
	QueryByNameReq {
		Name string `form:"name"`
	}

	QueryByNameResp {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Price int64  `json:"price"`
	}
)

type (
	QueryAllReq {
	}

	QueryAllResp {
		BookList []QueryByNameResp `json:"bookList"`
	}
)

service bookstore-api {

	@handler AddBookHandler
	get /addBook (AddBookReq) returns (AddBookResp)
	
	@handler QueryByNameHandler
	get /queryByName (QueryByNameReq) returns (QueryByNameResp)
	
	@handler QueryAllHandler
	get /queryAll (QueryAllReq) returns (QueryAllResp)
}
```

根据api生成服务

goctl api go -api bookstore.api -dir .

go mod tidy

运行api服务 

go run bookstore.go -f etc/bookstore-api.yaml

### rpc服务

进入上级目录 创建rpc服务

```
cd ..
mkdir rpc
cd rpc
mkdir book
cd book
goctl rpc template -o book.proto

```

book.proto

```
syntax = "proto3";

package book;

message AddRequest {
  string name = 1;
  int64 price = 2;
}

message AddResponse {
  bool Ok = 1;
}


message QueryByNameRequest{
  string name = 1;
}

message QueryByNameResponse{
  int64 Id = 1;
  string Name = 2;
  int64 Price = 3;
}

message QueryAllRequest{

}

message QueryAllResponse{
  repeated QueryByNameResponse List = 1 ;
}


service Book {
  rpc AddBook(AddRequest) returns(AddResponse);
  rpc QueryByName(QueryByNameRequest) returns (QueryByNameResponse);
  rpc QueryAll(QueryAllRequest) returns (QueryAllResponse);
}

```

生成rpc服务

goctl rpc proto -src book.proto -dir .

### model

   

1. 生成model

mkdir model

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/go_zero" -table="book"  -dir="./model”

### 代码修改

修改代码

bookstore-api.yaml里面添加etcd发现服务

```
Book:
  Etcd:
    Hosts:
      - localhost:2379
    Key: book.rpc
```

internal/config添加代码

```
type Config struct {
	rest.RestConf
	Book zrpc.RpcClientConf // 手动添加
}
```

svc目录下文件添加

```
type ServiceContext struct {
	Config config.Config
	Book    bookclient.Book
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Book: bookclient.NewBook(zrpc.MustNewClient(c.Book)),
	}
}

```

addbooklogic.go添加逻辑代码 调用rpc服务

```
func (l *AddBookLogic) AddBook(req types.AddBookReq) (*types.AddBookResp, error) {
	// 手动代码开始
	resp, err :=l.svcCtx.Book.AddBook(l.ctx,&book.AddRequest{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddBookResp{
		Ok: resp.Ok,
	}, nil
	// 手动代码结束
}

```

queryByNamelogic文件添加逻辑代码 调用rpc

```
// 手动代码开始
	resp, err :=l.svcCtx.Book.QueryByName(l.ctx,&book.QueryByNameRequest{Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &types.QueryByNameResp{
		Id: resp.GetId(),
		Name: resp.GetName(),
		Price:resp.GetPrice(),
	}, nil
	// 手动代码结束
```



至此api目录改完 下面改rpc的 

etc/book.yaml添加mysql redis 端口等

```
Name: book.rpc
ListenOn: 127.0.0.1:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: book.rpc

DataSource: root:123456@tcp(localhost:3306)/go_zero
Table: book
Cache:
  - Host: localhost:6379
```

config.go

```
type Config struct {
	zrpc.RpcServerConf
	DataSource string             // 手动代码
	Cache      cache.CacheConf    // 手动代码
}

```

svc下文件

```
type ServiceContext struct {
	Config config.Config
	Model model.BookModel   // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource)), // 手动代码
	}
}

```

addbooklogic.go

```
func (l *AddBookLogic) AddBook(in *book.AddRequest) (*book.AddResponse, error) {
	// 手动代码开始
	_, err := l.svcCtx.Model.Insert(model.Book{
		Name:  in.Name,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &book.AddResponse{
		Ok: true,
	}, nil
	// 手动代码结束
}

```

querybynamelogic.go

```
// 手动代码开始
	bookInfo, err := l.svcCtx.Model.FindByName(in.Name)
	if err != nil {
		return nil, err
	}

	return &book.QueryByNameResponse{
		Id:    bookInfo.Id,
		Name:  bookInfo.Name,
		Price: bookInfo.Price,
	}, nil
	// 手动代码结束
```

### 启动服务

首先启动rpc 再启动api

进入rpc/book目录

go run book.go -f etc/book.yaml

```
╭─mac@macdeMacBook-Pro /www/go/bookstore/rpc/book ‹master*› 
╰─$ go run book.go -f etc/book.yaml
Starting rpc server at 127.0.0.1:8080...
```



进入api目录 

```
╭─mac@macdeMacBook-Pro /www/go/bookstore/api ‹master*› 
╰─$ go run bookstore.go -f etc/bookstore-api.yaml                           1 ↵
Starting server at 0.0.0.0:8888...
```



## 结果

       访问浏览器
    
       http://localhost:8888/addBook?name=wang&price=333
    
       返回
    
       ```
       {
       ok: true
       }
       ```
    
       访问
    
       http://localhost:8888/queryByName?name=wang
    
       返回
    
       ```
       {
       id: 6,
       name: "wang",
       price: 333,
       }
       ```

查看数据库已经存储了数据