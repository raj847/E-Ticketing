package response

import (
	"eticketing/entity"
)

type Terminal struct {
	TerminalID uint   `json:"terminal_id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

func BuildTerminal(t entity.Terminal) Terminal {
	return Terminal{
		TerminalID: t.ID,
		Name:       t.Name,
		CreatedAt:  t.CreatedAt.String(),
		UpdatedAt:  t.UpdatedAt.String(),
		DeletedAt:  t.DeletedAt.Time.String(),
	}
}

func BuildTerminals(terminals []entity.Terminal) (res []Terminal) {
	for _, v := range terminals {
		terminal := BuildTerminal(v)
		res = append(res, terminal)
	}
	return res
}
