package eventbus

var envetByName = make(map[string][]func(interface{}))

func RegisterEvent(name string, callback func(interface{})) {
	list := envetByName[name]
	list = append(list, callback)
	envetByName[name] = list
}

func CallEvent(name string, param interface{}) {
	list := envetByName[name]

	for _, callback := range list {
		callback(param)
	}
}
