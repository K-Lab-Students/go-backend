begin;

create table if not exists t_user(
                             id serial not null primary key,
                             sname character varying(30),
                             name character varying(30),
                             pname character varying(30),
                             is_active boolean not null default true,
                             is_admin boolean not null default false,
                             is_member boolean not null default false,
                             create_date date not null,
                             file_object_id integer
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

create table if not exists t_news(
                       id serial not null primary key,
                       name character varying,
                       body character varying,
                       author_id integer,
                       is_active boolean not null default true,
                       is_main boolean not null default false,
                       create_date date not null,
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
create table if not exists t_projects(
                           id serial not null primary key,
                           name character varying,
                           body character varying,
                           is_active boolean not null default true,
                           create_date date not null,
                           author_id integer,
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

DELETE FROM t_news;
DELETE FROM t_projects;
DELETE FROM t_file_object;
DELETE FROM t_user;

INSERT INTO t_user(sname, name, pname,is_active,is_admin,is_member, create_date, file_object_id) VALUES('Тестирующий','Тест','Тестов',true,true,true,'2022-04-01',3),('Ломающий','Лом','Ломов',true,false,false,'2022-03-01',4);
INSERT INTO t_news(name, body, author_id, is_active,is_main,create_date, file_object_id)  VALUES('News №1','Some cool news 1',1,false,false,'2022-04-01','{1}'),('News №2','Some cool news 2',1,false,false,'2022-03-30','{1}'),('News №3','Some bad news 1',1,false,false,'2022-03-29','{1,2}');
INSERT INTO t_projects(name,body,is_active,create_date,file_object_id,author_id) VALUES ('test1','testbody',true,'2022-04-01','{1,2}',1 ),('project1','body1',false,'2022-04-01','{1}',1),('project2','body2',false,'2022-03-30','{1,2}',1),('project3','body3',false,'2022-03-29','{1}',1),('project4','body4',false,'2022-03-28','{1,2}',1),('project5','body5',false,'2022-03-27','{1}',1),('project6','body6',false,'2022-03-26','{1,2}',1),('project7','body7',false,'2022-03-25','{1}',1);
INSERT INTO t_file_object(file_path, comment, is_active,create_date) VALUES ('/photo/boston1.jpg','test',true,'2022-04-01'), ('/photo/boston2.jpg','test',true,'2022-04-01'),('/photo/luntik.jpeg','test',true,'2022-04-01'),('/photo/man1.png','test',true,'2022-04-01');

commit ;