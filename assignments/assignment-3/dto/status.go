package dto

import "assignment_3/entity"

type StatusDetail struct {
	Value       int    `json:"value"`
	Description string `json:"description"`
}

type StatusResponse struct {
	Water StatusDetail `json:"water"`
	Wind  StatusDetail `json:"wind"`
}

func (s *StatusResponse) ParseStatusDescription(status entity.Status) {
	s.Water.Value = status.Water
	s.Wind.Value = status.Wind

	if status.Water <= 5 {
		s.Water.Description = "Aman"
	} else if status.Water >= 6 && status.Water <= 8 {
		s.Water.Description = "Siaga"
	} else if status.Water > 8 {
		s.Water.Description = "Bahaya"
	}

	if status.Wind <= 6 {
		s.Wind.Description = "Aman"
	} else if status.Wind >= 7 && status.Wind <= 15 {
		s.Wind.Description = "Siaga"
	} else if status.Wind > 15 {
		s.Wind.Description = "Bahaya"
	}
}
