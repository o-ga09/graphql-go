CREATE TABLE users (
  id   BIGINT  NOT NULL AUTO_INCREMENT,
  user_id varchar(255) UNIQUE NOT NULL,
  username varchar(255) UNIQUE NOT NULL,
  displayname varchar(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE notes (
  id   BIGINT  NOT NULL AUTO_INCREMENT,
  note_id varchar(255) UNIQUE NOT NULL,
  title varchar(255) NOT NULL,
  tags varchar(255) NOT NULL,
  content text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user_notes (
  user_id varchar(255) NOT NULL,
  note_id varchar(255) NOT NULL,
  PRIMARY KEY (user_id, note_id),
  FOREIGN KEY (user_id) REFERENCES users (user_id),
  FOREIGN KEY (note_id) REFERENCES notes (note_id)
);

INSERT INTO users (user_id, username, displayname) VALUES ('1', 'admin', 'admin');
INSERT INTO users (user_id, username, displayname) VALUES ('2', 'user', 'user');
INSERT INTO notes (note_id, title, tags, content) VALUES ('1', 'First Note', 'first,note', 'This is the first note');
INSERT INTO notes (note_id, title, tags, content) VALUES ('2', 'Second Note', 'second,note', 'This is the second note');
INSERT INTO notes (note_id, title, tags, content) VALUES ('3', 'Third Note', 'third,note', 'This is the third note');
INSERT INTO notes (note_id, title, tags, content) VALUES ('4', 'Fourth Note', 'fourth,note', 'This is the fourth note');
INSERT INTO user_notes (user_id, note_id) VALUES ('1', '1');  
INSERT INTO user_notes (user_id, note_id) VALUES ('1', '2');
INSERT INTO user_notes (user_id, note_id) VALUES ('2', '3');
INSERT INTO user_notes (user_id, note_id) VALUES ('2', '4');
