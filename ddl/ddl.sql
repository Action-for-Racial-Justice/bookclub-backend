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
  entryID VARCHAR(255) NOT NULL,
  userID VARCHAR(255) NOT NULL,
  token VARCHAR(255) NOT NULL,
  CONSTRAINT s_fk_1 FOREIGN KEY (userID) REFERENCES user (id),
  PRIMARY KEY (entryID)
);

CREATE TABLE IF NOT EXISTS book (
  entryID VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  author VARCHAR(255) NOT NULL,
  isActive boolean,
  PRIMARY KEY (entryID)
);

INSERT INTO book VALUES
("1","The Divine Comedy","Dante Alighieri",1),
("2","SQL For Dummies","Allen G. Taylor",1),
("3","Inactive Book","Not A. Author",1);

CREATE TABLE IF NOT EXISTS club (
  entryID VARCHAR(255) NOT NULL,
  leaderID VARCHAR(255) NOT NULL,
  clubName VARCHAR(255) NOT NULL,
  bookID VARCHAR(255) NOT NULL,
  CONSTRAINT c_fk_1 FOREIGN KEY (leaderID) REFERENCES user (id),
  CONSTRAINT c_fk_2 FOREIGN KEY (bookID) REFERENCES book (entryID),
  PRIMARY KEY (entryID, clubName)
);

INSERT INTO club VALUES
("1","1","Keaton Club","1"),
("2","2","Ethen Club","2"),
("3","3","Meghan Club","3");

CREATE TABLE IF NOT EXISTS club_member (
  entryID VARCHAR(255) NOT NULL,
  userID VARCHAR(255) NOT NULL,
  clubID VARCHAR(255) NOT NULL,
  CONSTRAINT cm_fk_1 FOREIGN KEY (userID) REFERENCES user (id),
  CONSTRAINT cm_fk_2 FOREIGN KEY (clubID) REFERENCES club (entryID),
  PRIMARY KEY (entryID)
);

CREATE TABLE IF NOT EXISTS user_interest (
    id int(5) NOT NULL AUTO_INCREMENT,
    uid VARCHAR(255) NOT NULL,
    interestId int(5) NOT NULL,
    ranking int(5) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    lastModified TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE uid_ik (uid,interestId),

    CONSTRAINT ui_fk_1 FOREIGN KEY (uid) REFERENCES user (id),
    CONSTRAINT ui_fk_2 FOREIGN KEY (interestId) REFERENCES interest (id)
) ENGINE=INNODB;