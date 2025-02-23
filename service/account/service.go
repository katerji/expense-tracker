package account

type Service struct {
	repo repo
}

var instance *Service

func GetServiceInstance() *Service {
	if instance == nil {
		instance = &Service{
			repo: repo{},
		}
	}

	return instance
}
