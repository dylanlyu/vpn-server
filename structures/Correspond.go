package structures

/**
CREATE TABLE `web_correspond`(
	`id` bigint NOT NULL AUTO_INCREMENT,
	`first` varchar(50) NOT NULL DEFAULT '',
	`second` varchar(50) NOT NULL DEFAULT '',
	`third` varchar(50) NOT NULL DEFAULT '',
	PRIMARY KEY (id),
    	UNIQUE (first),
	FULLTEXT  INDEX (second)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 AUTO_INCREMENT=1


INSERT INTO `web_correspond`(`id`, `first`, `second`, `third`) VALUES
(1,'TPE','TXG','MLI'),
(2,'NTP','TXG','MLI'),
(3,'TNN','KHH','TXG'),
(4,'KHH','TNN','TXG'),
(5,'HCU','MLI','TYN'),
(6,'MLI','TYN','HCU'),
(7,'TYN','MLI','HCU'),
(8,'CHA','TXG','TPE'),
(9,'TXG','CHA','TPE')

*/


type WebCorrespond struct {
	Id     int64        `json:"id"`
	First  string       `json:"first"`
	Second string       `json:"second"`
	Third  string       `json:"third"`
}
