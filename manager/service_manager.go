package manager

import "github.com/rafialariq/digital-bank/service"

type ServiceManager interface {
	RegisterService() service.RegisterService
	LoginService() service.LoginService
}

type serviceManager struct {
	repositoryManager RepositoryManager
}

func (u *serviceManager) RegisterService() service.RegisterService {
	return service.NewRegisterService(u.repositoryManager.RegisterRepository())
}

func (u *serviceManager) LoginService() service.LoginService {
	return service.NewLoginService(u.repositoryManager.LoginRepository())
}

func NewUsecaseManager(r RepositoryManager) ServiceManager {
	return &serviceManager{
		repositoryManager: r,
	}
}
