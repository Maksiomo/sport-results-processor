-- +goose Up
-- +goose StatementBegin

create table sport(
    id                  bigserial primary key,
    name                text not null,
    min_team_size       int not null,
    max_team_size       text not null,
    description         text,
    unique(name)
);
create index idx_sport_name on sport(name);

create table country(
    id                  bigserial primary key,
    name                text not null,
    unique(name)
);

create table location(
    id                  bigserial primary key,
    country_id          bigint references country(id) not null,
    state               text,
    city                text,
    address             text,
    full_address        text not null,
    unique(full_address)
);

create table language(
    id                  bigserial primary key,
    name                text not null,
    unique(name)
);

create table person(
    id                  bigserial primary key,
    name                text not null,
    birth_date          timestamp without time zone not null,
    country_id          bigint references country(id) not null,
    unique(name)
);

create table person_sport(
    id                  bigserial primary key,
    person_id           bigint references person(id) not null,
    sport_id            bigint references sport(id) not null,
    unique (person_id, sport_id)
);

create table team(
    id                  bigserial primary key,
    name                text not null,
    country_id          bigint references country(id) not null,
    foundation_date     timestamp without time zone not null,
    unique(name)
);

create table role(
    id                  bigserial primary key,
    name                text not null,
    unique(name)
);

create table team_person(
    id                  bigserial primary key,
    person_id           bigint references person(id) not null,
    role_id             bigint references role(id) not null
);

create table competition_level(
    id                  bigserial primary key,
    name                text not null,
    unique(name)
);

create table competition(
    id                  bigserial primary key,
    name                text not null,
    sport_id            bigint references sport(id) not null,
    location_id         bigint references location(id) not null,
    level_id            bigint references competition_level(id) not null,
    unique(name)
);

create table prize(
    id                  bigserial primary key,
    competition_id      bigint references competition(id) not null,
    place_bracket       text not null,
    currency_code       text not null,
    prize_amount        bigint not null
);

create table competition_teams(
    id                  bigserial primary key,
    team_id             bigint references team(id) not null,
    competition_id      bigint references competition(id) not null
);

create table team_achievements(
    id                  bigserial primary key,
    team_id             bigint references team(id) not null,
    prize_id            bigint references prize(id) not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists team_achievements;
drop table if exists competition_teams;
drop table if exists prize;
drop table if exists competition;
drop table if exists competition_level;
drop table if exists team_person;
drop table if exists role;
drop table if exists team;
drop table if exists person_sport;
drop table if exists person;
drop table if exists language;
drop table if exists location;
drop table if exists country;
drop table if exists sport;

-- +goose StatementEnd