package request

type Terminal struct {
	Name string `json:"name" validate:"required"`
}
