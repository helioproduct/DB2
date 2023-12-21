package model

type Booking struct {
	BookingID       int
	ClientLastName  string
	ClientFirstName string
	RoomNumber      string
	CheckInDate     string
	CheckOutDate    string
	TotalPrice      float64
}
