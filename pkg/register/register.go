package register

var services []Service

func Server(name string, run RunFn) {
	services = append(services, Service{
		name: name,
		run:  run,
	})
}

func GetServer() []Service {
	return services
}

var initers []Service

func Initer(name string, run RunFn) {
	initers = append(initers, Service{
		name: name,
		run:  run,
	})
}

func GetIniter() []Service {
	return services
}
