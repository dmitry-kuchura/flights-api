package models

type Flight struct {
	Number              string     `json:"Number"`
	Info                Plane      `json:"Info"`
	DepartureTrafficHub TrafficHub `json:"DepartureTrafficHub"`
	ArrivalTrafficHub   TrafficHub `json:"ArrivalTrafficHub"`
	DepartureTime       string     `json:"DepartureTime"`
	ArrivalTime         string     `json:"ArrivalTime"`
	BoardStatus         string     `json:"BoardStatus"`
	IsCharter           bool       `json:"IsCharter"`
}
