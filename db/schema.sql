CREATE TABLE users (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id varchar(255) UNIQUE NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE notes (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  note_id varchar(255) UNIQUE NOT NULL,
  title varchar(255) NOT NULL,
  content text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE user_notes (
  user_id varchar(255) NOT NULL,
  note_id varchar(255) NOT NULL,
  PRIMARY KEY (user_id, note_id),
  FOREIGN KEY (user_id) REFERENCES users (user_id),
  FOREIGN KEY (note_id) REFERENCES notes (note_id)
);
