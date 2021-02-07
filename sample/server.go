package main

import (
	"fmt"
	gosoap "github.com/jackyshaw2017/golang-soap"
)

// 简单的
type User struct {
	In struct {
		Id int64
	}
	Out struct {
		ReturnInfo string
	}
}

func (u *User) Action() *gosoap.SoapFault {

	if u.In.Id != 100 {
		return gosoap.NewSoapFault("输入参数错误", "id必须是100", "")
	}

	u.Out.ReturnInfo = "test"
	return nil
}

// 返回带列表的
type DataList struct {
	In struct {
		Page    int
		PerPage int
		Search  string
	}
	Out struct {
		ReturnInfo string
	}
}

func (d *DataList) Action() *gosoap.SoapFault {
	fmt.Printf("%+v\n", d.In)
	d.Out.ReturnInfo = "test2"
	return nil
}

func main() {
	s := gosoap.NewServer("192.168.2.50", "people")
	s.Register(new(User), new(DataList))

	panic(s.Service("8080"))
}
