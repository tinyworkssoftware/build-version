use build_version_db;
drop table if exists tbl_active_session;
drop table if exists tbl_session_history;
drop table if exists tbl_project;
drop table if exists tbl_organisation;
drop table if exists tbl_plan_type;

create table tbl_plan_type (
   id varchar(36) primary key ,
   name varchar(100) not null unique,
   request_limit int not null default 1000,
   concurrent_session int not null default 1,
   price float not null,
   currency varchar(3) not null default 'usd',
   index (id, name)
);

create table tbl_organisation (
    id varchar(36) primary key ,
    name varchar(100) not null unique,
    created_ts timestamp not null default CURRENT_TIMESTAMP,
    updated_ts timestamp not null default CURRENT_TIMESTAMP,
    plan_type varchar(36) not null ,
    foreign key (plan_type) references tbl_plan_type(id),
    index (id, name)
);

create table tbl_project (
    id varchar(36) primary key ,
    name varchar(100) not null unique ,
    created_ts timestamp not null default CURRENT_TIMESTAMP,
    updated_ts timestamp not null default CURRENT_TIMESTAMP,
    organisation varchar(36) not null,
    exceeded_limit bool not null default false,
    access_code varchar(36) not null unique,
    foreign key (organisation) references tbl_organisation(id),
    index (id, name)
);

create table tbl_session_history (
    id varchar(36) primary key ,
    start_ts timestamp not null default CURRENT_TIMESTAMP,
    end_ts timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    associated_version varchar(36) not null ,
    associated_branch varchar(100) not null ,
    project varchar(100) not null ,
    session varchar(100) not null ,
    foreign key (project) references tbl_project(id),
    index (id)
);

create table tbl_active_session (
    id varchar(36) primary key ,
    start_ts timestamp not null default CURRENT_TIMESTAMP,
    end_ts timestamp not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    associated_version varchar(36) not null ,
    associated_branch varchar(100) not null ,
    session varchar(36) not null ,
    project varchar(100) not null ,
    foreign key (project) references tbl_project(id),
    index (id)
);

insert into tbl_plan_type (id, name, request_limit, concurrent_session, price)
values
       (uuid(), 'free', 100, 1, 0),
       (uuid(), 'basic', 1000, 5, 0.99),
       (uuid(), 'professional', 5000, 20, 3.99),
       (uuid(), 'enterprise', 9999999, 9999999, 15),
       (uuid(), 'self_hosted', 9999999, 9999999, 0);

