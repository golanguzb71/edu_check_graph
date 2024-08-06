CREATE TABLE students
(
    id           bigserial primary key,
    phone_number varchar NOT NULL,
    full_name    varchar NOT NULL,
    created_at   timestamp DEFAULT NOW(),
    updated_at   timestamp DEFAULT NOW()
);

CREATE TABLE collections
(
    id         bigserial primary key,
    name       varchar NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW()
);


CREATE TABLE questions
(
    id             bigserial primary key,
    question_field text                               NOT NULL,
    collection_id  bigint references collections (id) NOT NULL,
    created_at     timestamp DEFAULT NOW(),
    updated_at     timestamp DEFAULT NOW()
);


CREATE TABLE answers
(
    id           bigserial primary key,
    is_true      boolean                          NOT NULL,
    question_id  bigint references questions (id) NOT NULL,
    answer_field text                             NOT NULL,
    created_at   timestamp DEFAULT NOW(),
    updated_at   timestamp DEFAULT NOW()
);

CREATE TABLE groups
(
    id           bigserial primary key,
    name         varchar NOT NULL,
    teacher_name varchar Not Null,
    level        varchar NOT NULL check ( level in ('BEGINNER', 'ELEMENTARY', 'PRE INTERMEDIATE', 'INTERMEDIATE',
                                                    'UPPER INTERMEDIATE', 'ADVANCED', 'PROFICIENT')),
    created_at   timestamp DEFAULT NOW(),
    updated_at   timestamp DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON students
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON collections
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON questions
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON answers
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON groups
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
