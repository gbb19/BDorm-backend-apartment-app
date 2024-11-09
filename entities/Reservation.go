package entities

type Reservation struct {
	ReservationID         int     `json:"reservation_id"`
	MoveInDateTime        string  `json:"move_in_date_time"` // ใช้ string สำหรับ DateTime เพื่อความสะดวกในการจัดการ JSON
	ReservationRoomNumber int     `json:"reservation_room_number"`
	ReservationStatus     int     `json:"reservation_status"`
	TenantUsername        string  `json:"tenant_username"`
	ManagerUsername       *string `json:"manager_username,omitempty"` // ใช้ *string เพื่อให้เป็น null ได้ หากยังไม่มีผู้จัดการกำกับ
	BillID                *int    `json:"bill_id,omitempty"`          // ใช้ *int เพื่อรองรับกรณีไม่มี bill
}

type ReservationCreate struct {
	MoveInDateTime        string  `json:"move_in_date_time"`
	ReservationRoomNumber int     `json:"reservation_room_number"`
	TenantUsername        string  `json:"tenant_username"`
	ManagerUsername       *string `json:"manager_username,omitempty"` // ใช้ *string เพื่อให้เป็น null ได้
	BillID                *int    `json:"bill_id,omitempty"`          // ใช้ *int เพื่อรองรับกรณีไม่มี bill
}

type UpdateReservationStatus struct {
	ReservationID     int `json:"reservation_id"`
	ReservationStatus int `json:"reservation_status"`
}

type UpdateReservationDetails struct {
	ReservationID   int    `json:"reservation_id"`
	BillID          int    `json:"bill_id"`
	ManagerUsername string `json:"manager_username"`
}
