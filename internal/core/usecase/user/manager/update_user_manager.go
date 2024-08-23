package manager

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateUserManagerUseCase struct {
	userRepository port.IUserRepository
}

func NewUpdateUserManagerUseCase(userRepository port.IUserRepository) *UpdateUserManagerUseCase {
	return &UpdateUserManagerUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserManagerUseCase) Execute(user *domain.User) error {
	return uc.userRepository.Update(user)
}
