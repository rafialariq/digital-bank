package manager

import "github.com/rafialariq/digital-bank/repository"

type RepositoryManager interface {
	RegisterRepository() repository.RegisterRepository
	LoginRepository() repository.LoginRepository
	PaymentRepository() repository.PaymentRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) RegisterRepository() repository.RegisterRepository {
	return repository.NewRegisterRepository(r.infraManager.ConnectDb())
}

func (r *repositoryManager) LoginRepository() repository.LoginRepository {
	return repository.NewLoginRepository(r.infraManager.ConnectDb())
}

func (r *repositoryManager) PaymentRepository() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.infraManager.ConnectDb())
}

func NewRepoManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
