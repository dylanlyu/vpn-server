package structures

import "time"

type WebCity struct {
	Id          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Memo        string    `json:"memo"`
	State       int8      `json:"state"`
	Ordering    uint      `json:"ordering"`
	PublishUp   time.Time `json:"publishup"`
	PublishDown time.Time `json:"publishdown"`
	Created     time.Time `json:"created"`
	CreatedBy   uint      `json:"createdby"`
	Modified    time.Time `json:"modified"`
	ModifiedBy  uint      `json:"modifiedby"`
}
