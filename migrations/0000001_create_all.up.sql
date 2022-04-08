begin;

drop table t_user;
drop table t_file_object;
drop table t_projects;
drop table t_news;
drop table t_competitions;
drop table t_faculties;
drop table t_study_place;

create table if not exists t_user
(
    id              serial  not null primary key,
    sname           character varying(30),
    name            character varying(30),
    pname           character varying(30),
    email           text,
    password        text,
    birthday        date,
    is_active       boolean not null default true,
    is_admin        boolean not null default false,
    is_member       boolean not null default false,
    create_date     date    not null default now(),
    file_object_id  integer,
    competitions_id integer[],
    faculty_id    integer,
    study_place_id    integer
)
    WITH (OIDS = FALSE);

-- alter table t_user owner to postgres;
-- grant all on table t_user to postgres;
-- grant select on table t_user to public;
-- comment on table t_user is 'Таблица пользователей';

comment on column t_user.id is 'Код';
comment on column t_user.sname is 'Фамилия';
comment on column t_user.name is 'Имя';
comment on column t_user.pname is 'Отчество';
comment on column t_user.is_active is 'Пользователь активен';
comment on column t_user.is_admin is 'Пользователь админ';
comment on column t_user.is_member is 'Пользователь участник лаборатории';
COMMENT ON COLUMN t_user.file_object_id IS 'Коды файлов';

-- grant all on sequence t_user_id_seq to postgres;
-- grant select on sequence t_user_id_seq to public;
-- comment on sequence t_user_id_seq
--     is 'Таблица пользователей';

create table if not exists t_faculties
(
    id   serial not null primary key,
    name character varying
)
    WITH (OIDS = FALSE);
create table if not exists t_study_place
(
    id   serial not null primary key,
    name character varying
)
    WITH (OIDS = FALSE);

create table if not exists t_news
(
    id             serial  not null primary key,
    name           character varying,
    body           character varying,
    author_id      integer,
    is_active      boolean not null default true,
    is_main        boolean not null default false,
    create_date    date    not null,
    file_object_id integer[]
)
    WITH (OIDS = FALSE);

-- alter table t_news owner to postgres;
-- grant all on table t_news TO postgres;
-- GRANT select on table t_news to public;
-- COMMENT ON table t_news IS 'Таблица пользователей';

COMMENT ON COLUMN t_news.id IS 'Код';
COMMENT ON COLUMN t_news.name IS 'Заголовок новости';
COMMENT ON COLUMN t_news.body IS 'Текст новости';
COMMENT ON COLUMN t_news.author_id IS 'Код автора новости';
COMMENT ON COLUMN t_news.is_active IS 'Новость активна';
COMMENT ON COLUMN t_news.is_main IS 'Новость главная';
COMMENT ON COLUMN t_news.create_date IS 'Дата создания';
COMMENT ON COLUMN t_news.file_object_id IS 'Коды файлов';
--
--
--
--
create table if not exists t_projects
(
    id             serial  not null primary key,
    name           character varying,
    body           character varying,
    is_active      boolean not null default true,
    create_date    date    not null,
    author_id      integer,
    file_object_id integer[]
)
    WITH (OIDS = FALSE);

-- alter table t_projects owner to postgres;
-- grant all on table t_projects TO postgres;
-- GRANT select on table t_projects to public;
-- COMMENT ON table t_projects IS 'Таблица пользователей';

COMMENT ON COLUMN t_projects.id IS 'Код';
COMMENT ON COLUMN t_projects.name IS 'Название проекта';
COMMENT ON COLUMN t_projects.body IS 'Текст проекта';
COMMENT ON COLUMN t_projects.is_active IS 'Проект активен';
COMMENT ON COLUMN t_projects.create_date IS 'Дата создания';
COMMENT ON COLUMN t_projects.file_object_id IS 'Коды файлов';
--
--
--
--
--
create table if not exists t_file_object
(
    id          serial  not null primary key,
    file_path   character varying,
    comment     character varying,
    is_active   boolean not null default true,
    create_date date    not null
)
    WITH (OIDS = FALSE);

-- alter table t_file_object
--     owner to postgres;
-- grant all on table t_file_object TO postgres;
-- GRANT select on table t_file_object to public;
-- COMMENT ON table t_file_object IS 'Таблица пользователей';

COMMENT ON COLUMN t_file_object.id IS 'Код';
COMMENT ON COLUMN t_file_object.file_path IS 'Название проекта';
COMMENT ON COLUMN t_file_object.comment IS 'Текст проекта';
COMMENT ON COLUMN t_file_object.is_active IS 'Проект активен';
COMMENT ON COLUMN t_file_object.create_date IS 'Дата создания';

create table if not exists t_competitions
(
    id             serial  not null primary key,
    name           character varying,
    body           character varying,
    color          text,
    is_active      boolean not null default true,
    create_date    date    not null default now(),
    file_object_id integer[]
)
    WITH (OIDS = FALSE);

alter table t_competitions
    owner to postgres;
grant all on table t_competitions TO postgres;
GRANT select on table t_competitions to public;
COMMENT ON table t_competitions IS 'Таблица компетенций';

COMMENT ON COLUMN t_competitions.id IS 'Код';
COMMENT ON COLUMN t_competitions.name IS 'Название компетенции';
COMMENT ON COLUMN t_competitions.body IS 'Описание компетенции';
COMMENT ON COLUMN t_competitions.is_active IS 'Компетенция активна';
COMMENT ON COLUMN t_competitions.create_date IS 'Дата создания';
COMMENT ON COLUMN t_competitions.file_object_id IS 'Коды файлов';

DELETE
FROM t_news;
DELETE
FROM t_projects;
DELETE
FROM t_file_object;
DELETE
FROM t_user;
DELETE
FROM t_competitions;
DELETE
FROM t_faculties;
DELETE
FROM t_study_place;

INSERT INTO t_user(sname, name, pname, is_active, is_admin, is_member, create_date, file_object_id, email, password,
                   birthday, competitions_id,faculty_id,study_place_id)
VALUES ('Тестирующий', 'Тест', 'Тестов', true, true, true, '2022-04-01', 3, 'test1@yandex.ru',
        '$2y$10$uhceKdu4UobO1pwShyZTUOqEjBHkuwDPy7T9ZE/I39jz8IR6WW1Ku', '2001-01-18', '{1,2,3}',1,1),
       ('Ломающий', 'Лом', 'Ломов', true, false, true, '2022-03-01', 4, 'test2@yandex.ru',
        '$2y$10$uhceKdu4UobO1pwShyZTUOqEjBHkuwDPy7T9ZE/I39jz8IR6WW1Ku', '2001-01-20', '{4,5}',2,2);
INSERT INTO t_news(name, body, author_id, is_active, is_main, create_date, file_object_id)
VALUES ('News №1', 'Some cool news 1', 1, false, false, '2022-04-01', '{1}'),
       ('News №2', 'Some cool news 2', 1, false, false, '2022-03-30', '{1}'),
       ('News №3', 'Some bad news 1', 1, false, false, '2022-03-29', '{1,2}');
INSERT INTO t_projects(name, body, is_active, create_date, file_object_id, author_id)
VALUES ('test1', 'testbody', true, '2022-04-01', '{1,2}', 1),
       ('project1', 'body1', false, '2022-04-01', '{1}', 1),
       ('project2', 'body2', false, '2022-03-30', '{1,2}', 1),
       ('project3', 'body3', false, '2022-03-29', '{1}', 1),
       ('project4', 'body4', false, '2022-03-28', '{1,2}', 1),
       ('project5', 'body5', false, '2022-03-27', '{1}', 1),
       ('project6', 'body6', false, '2022-03-26', '{1,2}', 1),
       ('project7', 'body7', false, '2022-03-25', '{1}', 1);
INSERT INTO t_file_object(file_path, comment, is_active, create_date)
VALUES ('/photo/boston1.jpg', 'test', true, '2022-04-01'),
       ('/photo/boston2.jpg', 'test', true, '2022-04-01'),
       ('/photo/luntik.jpeg', 'test', true, '2022-04-01'),
       ('/photo/man1.png', 'test', true, '2022-04-01'),
       ('/photo/comp1.jpg', 'test', true, '2022-04-08'),
       ('/photo/comp2.jpg', 'test', true, '2022-04-08'),
       ('/photo/comp3.jpg', 'test', true, '2022-04-08'),
       ('/photo/comp4.jpg', 'test', true, '2022-04-08'),
       ('/photo/comp5.jpg', 'test', true, '2022-04-08');

INSERT INTO t_competitions(name, body, color, is_active, create_date, file_object_id)
VALUES ('Электроника', 'test1', 'yellow', true, '2022-04-08', '{5}'),
       ('Программирование МК', 'test2', 'indigo', true, '2022-04-08', '{9}'),
       ('Программирование ВУ', 'test3', 'orange', true, '2022-04-08', '{6}'),
       ('Нейросети', 'test4', 'lime', true, '2022-04-08', '{7}'),
       ('Конструирование', 'test5', 'red', true, '2022-04-08', '{8}');
INSERT INTO t_faculties(name)
VALUES ('ФМКН'),
       ('ФТФ'),
       ('ФПМ');

INSERT INTO t_study_place(name)
VALUES ('КубГУ'),('КубГАУ');

commit;