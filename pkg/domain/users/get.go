package users

import "context"

func (s *service) GetAllUser(ctx context.Context) error {
	err := s.userRepo.GetAllUser(ctx) // Suponiendo que `GetAll` esté en tu repositorio
	if err != nil {
		return err
	}
	return nil
}
