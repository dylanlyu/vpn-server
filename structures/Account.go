package structures

/**
ALTER TABLE `web_account` ADD `isauto` INT NOT NULL DEFAULT '1' AFTER `modified_by`, ADD `isuserset` INT NOT NULL DEFAULT '0' AFTER `isauto`, ADD `gas` INT NOT NULL DEFAULT '50' AFTER `isuserset`, ADD `action` INT NOT NULL DEFAULT '1' AFTER `gas`, ADD `iserror` INT NOT NULL DEFAULT '0' AFTER `action`, ADD `updatetime` TIMESTAMP NOT NULL AFTER `iserror`;
ALTER TABLE `web_account`
  DROP `updatetime`;
  ALTER TABLE `web_account` ADD `update_time` INT NOT NULL DEFAULT '1457318200' AFTER `iserror`;
ALTER TABLE `web_account` ADD `istask` INT NOT NULL DEFAULT '1' AFTER `update_time`;
ALTER TABLE `web_account` ADD `israise` INT NOT NULL DEFAULT '0' AFTER `istask`;
ALTER TABLE `web_account` ADD `isaction` INT NOT NULL DEFAULT '0' AFTER `israise`;
 */

import (
	"time"
)

type WebAccount struct {
	Id              int       `json:"id"`
	ForumId         uint      `json:"forumid"`
	IdentityId      uint      `json:"identityid"`
	Account         string    `json:"acciunt"`
	Passwd          string    `json:"passwd"`
	IsAlive         uint8     `json:"isalive"`
	RegisterDate    time.Time `json:"registerdate"`
	Email           string    `json:"email"`
	IsAuthEmail     uint8     `json:"isauthemail"`
	IsAuthPhone     uint8     `json:"isauthphone"`
	AuthPhoneDate   time.Time `json:"authphonedate"`
	AuthPhoneYearId uint      `json:"auth_phone_year_id"`
	AuthPhoneEdate  time.Time `json:"authphoneedate"`
	Background      string    `json:"background"`
	LoginNum        uint32    `json:"loginnum"`
	Level           string    `json:"level"`
	SecureQuestion1 string    `json:"securequestion1"`
	SecureQuestion2 string    `json:"securequestion2"`
	SecureQuestion3 string    `json:"securequestion3"`
	SecureQuestion4 string    `json:"securequestion4"`
	SecureQuestion5 string    `json:"securequestion5"`
	Source          string    `json:"source"`
	Memo            string    `json:"memo"`
	State           int8      `json:"state"`
	Ordering        uint      `json:"ordering"`
	PublishUp       time.Time `json:"publishup"`
	PublishDown     time.Time `json:"publishdown"`
	Created         time.Time `json:"created"`
	CreatedBy       uint      `json:"createdby"`
	Modified        time.Time `json:"modified"`
	ModifiedBy      uint      `json:"modifiedby"`
	IsAuto          int        `json:"isauto"`
	IsUserSet       int        `json:"isuserset"`
	Gas             int        `json:"gas"`
	Action          int        `json:"action"`
	IsError         int        `json:"iserror"`
	UpdateTime      int        `json:"updatetime"`
	Istask          int        `json:"istask"`
	Israise         int        `json:"israise"`
	Isaction	int        `json:"isaction"`
}

type ClientAccount struct {
	Id         int        `json:"id"`
	Account    string     `json:"account"`
	Pass       string     `json:"pass"`
	IsError    int        `json:"iserror"`
	IsAuto     int        `json:"isauto"`
	IsUserSet  int        `json:"isuserset"`
	Gas        int        `json:"gas"`
	Action     int        `json:"action"`
	UpdateTime int        `json:"updatetime"`
	Primary    string     `json:"primary"`
	Secondary  string     `json:"secondary"`
	Tertiary   string     `json:"tertiary"`
}
