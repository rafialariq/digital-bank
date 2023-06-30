package manager

import "github.com/rafialariq/digital-bank/service"

type ServiceManager interface {
	RegisterService() service.RegisterService
	LoginService() service.LoginService
	PaymentService() service.PaymentService
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

func (u *serviceManager) PaymentService() service.PaymentService {
	return service.NewPaymentService(u.repositoryManager.PaymentRepository())
}

func NewUsecaseManager(r RepositoryManager) ServiceManager {
	return &serviceManager{
		repositoryManager: r,
	}
}
