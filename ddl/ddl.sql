CREATE TABLE IF NOT EXISTS user (
  id VARCHAR(255) NOT NULL,
  fullName VARCHAR(255) NOT NULL,
  clubAssigned tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);

INSERT INTO user VALUES
("1","Keaton Currie",0),
("2","Ethen Pociask",0),
("3","Meghan Johnson",0),
("4","Stephanie Grasso",0);

CREATE TABLE IF NOT EXISTS session (
  id VARCHAR(255) NOT NULL,
  uid VARCHAR(255) NOT NULL,
  token VARCHAR(255) NOT NULL,
  CONSTRAINT s_fk_1 FOREIGN KEY (uid) REFERENCES user (id),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS book (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  isActive boolean,
  PRIMARY KEY (id)
);

INSERT INTO book VALUES
("1","The Divine Comedy","Dante Alighieri",1),
("2","SQL For Dummies","Allen G. Taylor",1),
("3","Inactive Book","Not A. Author",1);

CREATE TABLE IF NOT EXISTS club (
  id VARCHAR(255) NOT NULL,
  leaderId VARCHAR(255) NOT NULL,
  clubName VARCHAR(255) NOT NULL,
  bookId VARCHAR(255) NOT NULL,
  CONSTRAINT c_fk_1 FOREIGN KEY (leaderId) REFERENCES user (id),
  CONSTRAINT c_fk_2 FOREIGN KEY (bookId) REFERENCES book (id),
  PRIMARY KEY (id)
);

INSERT INTO club VALUES
("1","1","Keaton Club","1"),
("2","2","Ethen Club","2"),
("3","3","Meghan Club","3");

CREATE TABLE IF NOT EXISTS club_member (
  id VARCHAR(255) NOT NULL,
  uid VARCHAR(255) NOT NULL,
  clubId VARCHAR(255) NOT NULL,
  CONSTRAINT cm_fk_1 FOREIGN KEY (uid) REFERENCES user (id),
  CONSTRAINT cm_fk_2 FOREIGN KEY (clubId) REFERENCES club (id),
  PRIMARY KEY (id)
);
