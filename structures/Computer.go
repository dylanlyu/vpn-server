package structures

import (
	"time"
)

/**
CREATE TABLE `web_computer`
(
  `id` bigint NOT NULL AUTO_INCREMENT,
  `computer` varchar(50) NOT NULL DEFAULT '',
  `city` varchar(50) NOT NULL DEFAULT '',
  `code` varchar(50) NOT NULL DEFAULT '',
  `ip` varchar(50) NOT NULL DEFAULT '',
  `time` datetime DEFAULT NULL,
  `status` integer NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE (computer),
  FULLTEXT  INDEX (city)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1

CREATE TABLE `web_usercomputer`
(
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '',
  `code` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  UNIQUE (name)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1
*/

type WebComputer struct {
	Id       int64        `json:"id"`
	Computer string       `json:"computer"`
	City     string       `json:"city"`
	Code     string       `json:"code"`
	Ip       string       `json:"ip"`
	Time     time.Time    `json:"time"`
	Status   int          `json:"status"`
}

type UserComputer struct {
	Code 	string        `json:"code"`
	Name	string        `json:"name"`
}

type Webusercomputer struct {
	Id 	int64         `json:"id"`
	Name	string        `json:"name"`
	Code 	string        `json:"code"`

}