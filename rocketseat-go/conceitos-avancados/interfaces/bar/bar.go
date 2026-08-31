package bar

import "github.com/luizandrends/interfaces/foo"

func TakeFoo(i foo.Interface) {
	i.Interface()
}