package xxx

var name string

func init() {
	name = "init"
}

//允许一个文件多个init
func init() {
	name = "init2"
}

func Fun() {
	name = "init"
}
