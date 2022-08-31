package main

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	userpb "grpc_example/proto/gen/go"
)

func main() {
	user := userpb.User{
		Name:   "John",
		Age:    18,
		Gender: userpb.Gender_MALE,
		// 注意我们 proto 文件中定义的是下划线命名 is_verified
		// 生成出来代码自动转成了大小写命名
		IsVerified: true,
		ActiveCode: &userpb.Code{
			Name:  "coin100",
			Award: 100,
		},
		OutdatedCodes: []*userpb.Code{
			{
				Name:  "coin50",
				Award: 50,
			},
			{
				Name:  "coin30",
				Award: 30,
			},
			{
				Name:  "coin10",
				Award: 10,
			},
		},
	}

	fmt.Println(&user)
	// 输出结果如下
	// name:"John"  age:18  is_verified:true  gender:MALE  active_code:{name:"coin100"  award:100}  outdated_codes:{name:"coin50"  award:50}  outdated_codes:{name:"coin30"  award:30}  outdated_codes:{name:"coin10"  award:10}

	// 可以通过 proto.Marshal 转成 []byte
	b, err := proto.Marshal(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)
	// HEX 输出结果如下
	// 0A044A6F686E1012180120012A0B0A07636F696E3130301064320A0A06636F696E35301032320A0A06636F696E3330101E320A0A06636F696E3130100A

	// 可以通过 proto.Unmarshal 再转回来
	var user2 userpb.User
	err = proto.Unmarshal(b, &user2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&user2)
	// name:"John"  age:18  is_verified:true  gender:MALE  active_code:{name:"coin100"  award:100}  outdated_codes:{name:"coin50"  award:50}  outdated_codes:{name:"coin30"  award:30}  outdated_codes:{name:"coin10"  award:10}

	// 同样的也可以转成 JSON
	// 微服务和微服务之间可以通过二进制数据流通信, 和前端之间通过JSON通信
	b, err = json.Marshal(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	// {"name":"John","age":18,"is_verified":true,"gender":1,"active_code":{"name":"coin100","award":100},"outdated_codes":[{"name":"coin50","award":50},{"name":"coin30","award":30},{"name":"coin10","award":10}]}
}
