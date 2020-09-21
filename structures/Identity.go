package structures

import (
	"time"
)

type WebIdentity struct {
	Id          int       `json:"id"`
	Code        uint      `json:"code"`
	Title       string    `json:"title"`
	Nickname    string    `json:"nickname"`
	Sex         string    `json:"sex"`
	Birth       time.Time `json:"birth"`
	Idnum       string    `json:"idnum"`
	Address     string    `json:"address"`
	Tel         string    `json:"tel"`
	Org         string    `json:"org"`
	CityId      uint      `json:"cityid"`
	SkillArea   string    `json:"skillarea"`
	AuthPhone   string    `json:"authphone"`
	AuthIdnum   string    `json:"authidnum"`
	Email       string    `json:"email"`
	SubEmail_1   string    `json:"subemail1"`
	SubEmail_2   string    `json:"subemail2"`
	SubEmail_3   string    `json:"subemail3"`
	CaseId      uint      `json:"caseid"`
	IsCase      uint8     `json:"iscase"`
	FreezeDate  time.Time `json:"freezedate"`
	IsFreeze    uint8     `json:"isfreeze"`
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
