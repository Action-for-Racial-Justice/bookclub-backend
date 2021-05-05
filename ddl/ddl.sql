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

CREATE TABLE IF NOT EXISTS club (
  entryID VARCHAR(255) NOT NULL,
  leaderID VARCHAR(255) NOT NULL,
  clubName VARCHAR(255) NOT NULL,
  bookID VARCHAR(255) NULL,
  description VARCHAR(1000) NOT NULL,
  CONSTRAINT c_fk_1 FOREIGN KEY (leaderID) REFERENCES user (id),
  PRIMARY KEY (entryID),
  UNIQUE INDEX clubName_UNIQUE (clubName ASC)
);

INSERT INTO club VALUES
("1","1","Keaton Club","1", "test description 1"),
("2","2","Ethen Club","2", "test description 2"),
("3","3","Meghan Club","3", "Here is a long description of a club Here");

CREATE TABLE IF NOT EXISTS club_member (
  entryID VARCHAR(255) NOT NULL,
  userID VARCHAR(255) NOT NULL,
  clubID VARCHAR(255) NOT NULL,
  CONSTRAINT cm_fk_1 FOREIGN KEY (userID) REFERENCES user (id),
  CONSTRAINT cm_fk_2 FOREIGN KEY (clubID) REFERENCES club (entryID),
  PRIMARY KEY (entryID)
);
