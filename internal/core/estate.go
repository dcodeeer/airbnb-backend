package core

type Estate struct {
	Id         int    `json:"id"`
	OwnerId    int    `json:"owner_id"`
	PriceNight int    `json:"price_night"`
	PriceWeek  int    `json:"price_week"`
	Area       int    `json:"area"`
	Rooms      int    `json:"rooms"`
	Showers    int    `json:"showers"`
	BabyRooms  int    `json:"baby_rooms"`
	CreatedAt  string `json:"created_at"`
}
