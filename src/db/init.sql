-- ユーザーテーブル
drop table if exists users;

create table users (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null unique,
    email varchar(255) not null,
    password varchar(255) not null,
    self_introduction text check (char_length(self_introduction) <= 2000),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_users_name on users (name);

-- スキルテーブル
drop table if exists skills;

create table skills (
    id uuid primary key default gen_random_uuid(),
    tag_id uuid not null references tags(id) on delete cascade,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- ユーザースキルテーブル
drop table if exists user_skills;
create table user_skills (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references users(id),
    skill_id uuid not null references skills(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- タグテーブル
drop table if exists tags;

create table tags (
    id uuid primary key default gen_random_uuid(),
    mentor_recruitment_id uuid not null references mentor_recruitments(id),
    created_at timestamp default now(),
    updated_at timestamp default now()
)


-- 経歴テーブル
drop table if exists careers;

create table careers (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references users(id),
    detail text not null check (char_length(detail) <= 1000),
    start_year int,
    end_year int,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_careers_user_id on careers (user_id);

-- メンター募集テーブル
drop table if exists mentor_recruitments;

create table mentor_recruitments (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references users(id),
    title varchar(255) not null,
    consultation_type text not null,
    consultation_method text not null,
    description text not null check (char_length(description) <= 2000),
    budget_from int not null check (budget_from >= 1000),
    budget_to int not null check (budget_to >= 100),
    application_period not null,
    status varchar(255) not null,
    category_id uuid not null references categories(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- メンター募集提案テーブル
drop table if exists mentor_recruitment_proposals;

create table mentor_recruitment_proposals (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references users(id),
    mentor_recruitment_id uuid not null references mentor_recruitments(id),
    description text not null check (char_length(description) <= 2000),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);
-- 契約リクエストテーブル
-- 契約テーブル
-- メンタープランテーブル

-- メンター募集カテゴリテーブル
drop table if exists categories;