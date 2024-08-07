CREATE TABLE IF NOT EXISTS students
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    phone_number VARCHAR NOT NULL,
    full_name    VARCHAR NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS collections
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS questions
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    question_field TEXT                                NOT NULL,
    collection_id  INTEGER REFERENCES collections (id) NOT NULL,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS answers
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    is_true      BOOLEAN                           NOT NULL,
    question_id  INTEGER REFERENCES questions (id) NOT NULL,
    answer_field TEXT                              NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    name         VARCHAR NOT NULL,
    teacher_name VARCHAR NOT NULL,
    level        VARCHAR NOT NULL CHECK (level IN ('BEGINNER', 'ELEMENTARY', 'PRE INTERMEDIATE', 'INTERMEDIATE',
                                                   'UPPER INTERMEDIATE', 'ADVANCED', 'PROFICIENT')),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
