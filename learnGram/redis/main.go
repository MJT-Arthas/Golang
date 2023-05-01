package main

type Dog struct {
	name string
}

type A interface {
	GetName()
}

func main() {

	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // 没有密码，默认值
	//	DB:       0,  // 默认DB 0
	//})
	//ctx := context.Background()
	//val, _ := rdb.Get(ctx, "hello").Result()
	//fmt.Println(val)
}

func (d *Dog) GetName() string {
	return d.name
}
