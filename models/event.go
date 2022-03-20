package models

type Event struct {
	Event string `json:"event" db:"event"`
	Data Answer `json:"data" db:"data"`
}

type EventOutput struct {
	Event string `json:"event" db:"event"`
	Data JSONB `json:"data" db:"data"`
}
