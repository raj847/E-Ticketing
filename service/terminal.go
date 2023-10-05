package service

import (
	"context"
	"eticketing/entity"
	"eticketing/repository"
)

type TerminalService struct {
	terminalRepo *repository.TerminalRepository
}

func NewTerminalService(terminalRepo *repository.TerminalRepository) *TerminalService {
	return &TerminalService{
		terminalRepo: terminalRepo,
	}
}

func (m *TerminalService) GetAllTerminal(ctx context.Context) ([]entity.Terminal, error) {
	res, err := m.terminalRepo.GetTerminal(ctx)
	if err != nil {
		return []entity.Terminal{}, err
	}
	return res, nil
}

func (m *TerminalService) AddTerminal(ctx context.Context, terminal *entity.Terminal) error {
	err := m.terminalRepo.CreateTerminal(ctx, terminal)
	if err != nil {
		return err
	}
	return nil
}

