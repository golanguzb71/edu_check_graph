CREATE TABLE IF NOT EXISTS students
(
    id
        SERIAL
        PRIMARY
            KEY,
    phone_number
        VARCHAR
        NOT
            NULL
        UNIQUE,
    full_name
        VARCHAR
        NOT
            NULL,
    created_at
        TIMESTAMP
        DEFAULT
            NOW(),
    updated_at
        TIMESTAMP
        DEFAULT
            NOW()
);


CREATE TABLE IF NOT EXISTS collections
(
    id
        SERIAL
        PRIMARY
            KEY,
    name
        VARCHAR
        NOT
            NULL,
    created_at
        TIMESTAMP
        DEFAULT
            NOW(),
    updated_at
        TIMESTAMP
        DEFAULT
            NOW()
);


CREATE TABLE IF NOT EXISTS questions
(
    id
        SERIAL
        PRIMARY
            KEY,
    question_field
        TEXT
        NOT
            NULL,
    collection_id
        INTEGER
        NOT
            NULL,
    created_at
        TIMESTAMP
        DEFAULT
            NOW(),
    updated_at
        TIMESTAMP
        DEFAULT
            NOW(),
    FOREIGN
        KEY
        (
         collection_id
            ) REFERENCES collections
        (
         id
            )
        ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS answers
(
    id
        SERIAL
        PRIMARY
            KEY,
    is_true
        BOOLEAN
        NOT
            NULL,
    question_id
        INTEGER
        NOT
            NULL,
    answer_field
        TEXT
        NOT
            NULL,
    created_at
        TIMESTAMP
        DEFAULT
            NOW(),
    updated_at
        TIMESTAMP
        DEFAULT
            NOW(),
    FOREIGN
        KEY
        (
         question_id
            ) REFERENCES questions
        (
         id
            )
        ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS groups
(
    id
               SERIAL
        PRIMARY
            KEY,
    name
               VARCHAR
        NOT
            NULL,
    teacher_name
               VARCHAR
        NOT
            NULL,
    level
               VARCHAR
        NOT
            NULL
        CHECK (
            level
                IN
            (
             'BEGINNER',
             'ELEMENTARY',
             'PRE_INTERMEDIATE',
             'INTERMEDIATE',
             'UPPER_INTERMEDIATE',
             'ADVANCED',
             'PROFICIENT'
                )),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE if not exists student_collection
(
    id            SERIAL PRIMARY KEY,
    student_id    INT references students (id)    NOT NULL,
    collection_id INT references collections (id) NOT NULL,
    created_at    timestamp DEFAULT NOW(),
    update_at     timestamp DEFAULT NOW()
);


CREATE TABLE student_answer
(
    id                    SERIAL PRIMARY KEY,
    student_collection_id INT references student_collection (id) NOT NULL,
    question_id           INT references questions (id)          NOT NULL,
    answer_id             INT references answers (id)            NOT NULL,
    is_true               BOOLEAN                                NOT NULL,
    created_at            timestamp DEFAULT NOW(),
    update_at             timestamp DEFAULT NOW()
);