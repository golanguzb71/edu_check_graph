-- Enable foreign key constraints
PRAGMA foreign_keys = ON;

-- Create students table
CREATE TABLE IF NOT EXISTS students (
                                        id           INTEGER PRIMARY KEY AUTOINCREMENT,
                                        phone_number VARCHAR NOT NULL,
                                        full_name    VARCHAR NOT NULL,
                                        created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS collections (
                                           id         INTEGER PRIMARY KEY AUTOINCREMENT,
                                           name       VARCHAR NOT NULL,
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS questions (
                                         id             INTEGER PRIMARY KEY AUTOINCREMENT,
                                         question_field TEXT    NOT NULL,
                                         collection_id  INTEGER NOT NULL,
                                         created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         FOREIGN KEY (collection_id) REFERENCES collections (id)
                                             ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS answers (
                                       id           INTEGER PRIMARY KEY AUTOINCREMENT,
                                       is_true      BOOLEAN NOT NULL,
                                       question_id  INTEGER NOT NULL,
                                       answer_field TEXT    NOT NULL,
                                       created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       FOREIGN KEY (question_id) REFERENCES questions (id)
                                           ON DELETE CASCADE
);


-- Create groups table
CREATE TABLE IF NOT EXISTS groups (
                                      id           INTEGER PRIMARY KEY AUTOINCREMENT,
                                      name         VARCHAR NOT NULL,
                                      teacher_name VARCHAR NOT NULL,
                                      level        VARCHAR NOT NULL CHECK (level IN ('BEGINNER', 'ELEMENTARY', 'PRE_INTERMEDIATE', 'INTERMEDIATE',
                                                                                     'UPPER_INTERMEDIATE', 'ADVANCED', 'PROFICIENT')),
                                      created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
