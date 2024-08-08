CREATE TABLE IF NOT EXISTS students (
                                        id SERIAL PRIMARY KEY,
                                        phone_number VARCHAR NOT NULL,
                                        full_name VARCHAR NOT NULL,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS collections (
                                           id SERIAL PRIMARY KEY,
                                           name VARCHAR NOT NULL,
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS questions (
                                         id SERIAL PRIMARY KEY,
                                         question_field TEXT NOT NULL,
                                         collection_id INTEGER NOT NULL,
                                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                         FOREIGN KEY (collection_id) REFERENCES collections (id)
                                             ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS answers (
                                       id SERIAL PRIMARY KEY,
                                       is_true BOOLEAN NOT NULL,
                                       question_id INTEGER NOT NULL,
                                       answer_field TEXT NOT NULL,
                                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       FOREIGN KEY (question_id) REFERENCES questions (id)
                                           ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS groups (
                                      id SERIAL PRIMARY KEY,
                                      name VARCHAR NOT NULL,
                                      teacher_name VARCHAR NOT NULL,
                                      level VARCHAR NOT NULL CHECK (level IN ('BEGINNER', 'ELEMENTARY', 'PRE_INTERMEDIATE', 'INTERMEDIATE',
                                                                              'UPPER_INTERMEDIATE', 'ADVANCED', 'PROFICIENT')),
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
