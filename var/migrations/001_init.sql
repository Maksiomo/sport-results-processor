-- +goose Up
-- +goose StatementBegin

-- Основные таблицы справочников

CREATE TABLE sport (
    id            BIGSERIAL PRIMARY KEY,
    name          TEXT       NOT NULL UNIQUE,
    min_team_size INT        NOT NULL,
    max_team_size INT        NOT NULL,
    description   TEXT,
    created_at    TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash   BYTEA,
    tx_hash       TEXT
);
CREATE INDEX idx_sport_name ON sport(name);


CREATE TABLE country (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT       NOT NULL UNIQUE,
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW()
);


CREATE TABLE location (
    id            BIGSERIAL PRIMARY KEY,
    country_id    BIGINT     NOT NULL REFERENCES country(id),
    state         TEXT,
    city          TEXT,
    address       TEXT,
    full_address  TEXT       NOT NULL UNIQUE,
    created_at    TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP  NOT NULL DEFAULT NOW()
);


CREATE TABLE role (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT       NOT NULL UNIQUE,
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW()
);


CREATE TABLE competition_level (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT       NOT NULL UNIQUE,
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW()
);


-- Таблица валют для призов
CREATE TABLE currency (
    code   TEXT PRIMARY KEY,
    name   TEXT NOT NULL,
    symbol TEXT UNIQUE
);


-- Основные бизнес-сущности

CREATE TABLE person (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT       NOT NULL UNIQUE,
    birth_date  TIMESTAMP  NOT NULL,
    country_id  BIGINT     NOT NULL REFERENCES country(id),
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash BYTEA,
    tx_hash     TEXT
);

CREATE TABLE person_sport (
    id          BIGSERIAL PRIMARY KEY,
    person_id   BIGINT     NOT NULL REFERENCES person(id),
    sport_id    BIGINT     NOT NULL REFERENCES sport(id),
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    UNIQUE(person_id, sport_id)
);


CREATE TABLE team (
    id              BIGSERIAL PRIMARY KEY,
    name            TEXT       NOT NULL UNIQUE,
    country_id      BIGINT     NOT NULL REFERENCES country(id),
    foundation_date TIMESTAMP  NOT NULL,
    created_at      TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash     BYTEA,
    tx_hash         TEXT
);


CREATE TABLE team_person (
    id          BIGSERIAL PRIMARY KEY,
    team_id     BIGINT     NOT NULL REFERENCES team(id),
    person_id   BIGINT     NOT NULL REFERENCES person(id),
    role_id     BIGINT     NOT NULL REFERENCES role(id),
    joined_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    left_at     TIMESTAMP,
    created_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash BYTEA,
    tx_hash     TEXT,
    UNIQUE(team_id, person_id, role_id, joined_at)
);


CREATE TABLE competition (
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT       NOT NULL UNIQUE,
    sport_id     BIGINT     NOT NULL REFERENCES sport(id),
    location_id  BIGINT     NOT NULL REFERENCES location(id),
    level_id     BIGINT     NOT NULL REFERENCES competition_level(id),
    created_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash  BYTEA,
    tx_hash      TEXT
);

CREATE TABLE stage (
    id             BIGSERIAL PRIMARY KEY,
    competition_id BIGINT     NOT NULL REFERENCES competition(id),
    name           TEXT       NOT NULL,
    start_time     TIMESTAMP  NOT NULL,
    end_time       TIMESTAMP,
    created_at     TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP  NOT NULL DEFAULT NOW(),
    UNIQUE(competition_id, name)
);

CREATE TABLE match (
    id           BIGSERIAL PRIMARY KEY,
    stage_id     BIGINT     NOT NULL REFERENCES stage(id),
    match_time   TIMESTAMP  NOT NULL,
    location_id  BIGINT REFERENCES location(id),
    metadata     JSONB,
    created_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash  BYTEA,
    tx_hash      TEXT,
    UNIQUE(stage_id, match_time, location_id)
);

CREATE TABLE match_participant (
    id           BIGSERIAL PRIMARY KEY,
    match_id     BIGINT     NOT NULL REFERENCES match(id),
    team_id      BIGINT     NOT NULL REFERENCES team(id),
    score        INT        NOT NULL,
    is_winner    BOOLEAN    NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash  BYTEA,
    tx_hash      TEXT,
    UNIQUE(match_id, team_id)
);

CREATE TABLE prize (
    id              BIGSERIAL PRIMARY KEY,
    competition_id  BIGINT     NOT NULL REFERENCES competition(id),
    place_bracket   TEXT       NOT NULL,
    currency_code   TEXT       NOT NULL REFERENCES currency(code),
    prize_amount    BIGINT     NOT NULL,
    created_at      TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash     BYTEA,
    tx_hash         TEXT,
    UNIQUE(competition_id, place_bracket)
);

CREATE TABLE competition_teams (
    id             BIGSERIAL PRIMARY KEY,
    team_id        BIGINT     NOT NULL REFERENCES team(id),
    competition_id BIGINT     NOT NULL REFERENCES competition(id),
    created_at     TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash    BYTEA,
    tx_hash        TEXT,
    UNIQUE(team_id, competition_id)
);

CREATE TABLE team_achievements (
    id           BIGSERIAL PRIMARY KEY,
    team_id      BIGINT     NOT NULL REFERENCES team(id),
    prize_id     BIGINT     NOT NULL REFERENCES prize(id),
    created_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT NOW(),
    record_hash  BYTEA,
    tx_hash      TEXT,
    UNIQUE(team_id, prize_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists team_achievements;
drop table if exists competition_teams;
drop table if exists prize;
drop table if exists match_participant;
drop table if exists match;
drop table if exists stage;
drop table if exists competition;
drop table if exists competition_level;
drop table if exists team_person;
drop table if exists role;
drop table if exists team;
drop table if exists person_sport;
drop table if exists person;
drop table if exists currency;
drop table if exists location;
drop table if exists country;
drop table if exists sport;

-- +goose StatementEnd