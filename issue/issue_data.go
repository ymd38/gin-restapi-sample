package issue

type IssueData struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Source     string `json:"source"`
	Detail     string `json:"detail"`
	Priority   int    `json:"priority"`
	Status     int    `json:"status"`
	Limit      int64  `json:"-"`
	LimitStr   string `json:"limit" db:"-"`
	Created    int64  `json:"-"`
	CreatedStr string `json:"created,omitempty" db:"-"`
	Updated    int64  `json:"-"`
	UpdatedStr string `json:"updated,omitempty" db:"-"`
}
