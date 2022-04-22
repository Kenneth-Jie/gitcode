package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init()  {
	pool=&redis.Pool{
		MaxActive:0,
		MaxIdle:8,
		IdleTimeout:100,
		Dial:func ()(redis.Conn,error)  {
			return redis.Dial("tcp","localhost:6379")
		},
	}
}


func main()  {
	conn:=pool.Get()
	defer conn.Close()

	// c,err:=redis.Dial("tcp","127.0.0.1:6379")
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer c.Close()

	_,err:=conn.Do("Set","name","john")
	if err!=nil{
		fmt.Println(err)
		return
	}

	r,err:=redis.String(conn.Do("Get","name"))
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(r)

	conn2:=pool.Get()
	
	_,err=conn2.Do("Set","name2","john2")
	if err!=nil{
		fmt.Println(err)
		return
	}

	r2,err:=redis.String(conn2.Do("Get","name2"))
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(r2)

}
#1111111111
