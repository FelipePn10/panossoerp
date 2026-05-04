package request

type CreateItemCalendarDayDTO struct {
	ItemCode    int64   `json:"item_code"`
	Mask        string  `json:"mask"`
	Year        int     `json:"year"`
	Month       int     `json:"month"`
	Day         int     `json:"day"`
	IsWorkday   bool    `json:"is_workday"`
	Description *string `json:"description,omitempty"`
}
