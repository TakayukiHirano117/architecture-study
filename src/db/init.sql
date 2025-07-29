-- ユーザーテーブル
drop table if exists users;

create table users (
    id uuid primary key not null,
    name varchar(255) not null unique,
    email varchar(255) not null,
    password varchar(255) not null check (
        char_length(password) >= 12 AND
        password ~ '[a-zA-Z]' AND
        password ~ '[0-9]'
    ),
    self_introduction text check (char_length(self_introduction) <= 2000),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_users_name on users (name);

-- スキルテーブル
drop table if exists skills;

create table skills (
    id uuid primary key not null,
    tag_id uuid not null references tags(id) on delete cascade,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- ユーザーとスキルの中間テーブル
drop table if exists user_skills;
create table user_skills (
    id uuid primary key not null,
    user_id uuid not null references users(id),
    skill_id uuid not null references skills(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- タグテーブル
drop table if exists tags;

create table tags (
    id uuid primary key not null,
    name varchar(255) not null unique,
    created_at timestamp default now(),
    updated_at timestamp default now()
)

-- 経歴テーブル
drop table if exists careers;

create table careers (
    id uuid primary key not null,
    user_id uuid not null references users(id),
    detail text not null,
    start_year int,
    end_year int,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_careers_user_id on careers (user_id);

-- メンター募集テーブル
drop table if exists mentor_recruitments;

create table mentor_recruitments (
    id uuid primary key not null,
    user_id uuid not null references users(id),
    title varchar(255) not null,
    consultation_type varchar(255) not null,
    consultation_method varchar(255) not null,
    description text not null,
    budget_from int not null,
    budget_to int not null,
    application_period date not null,
    status varchar(255) not null,
    category_id uuid not null references categories(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- メンター募集とタグの中間テーブル
drop table if exists mentor_recruitment_tags;

create table mentor_recruitment_tags (
    id uuid primary key not null,
    mentor_recruitment_id uuid not null references mentor_recruitments(id),
    tag_id uuid references tags(id),
    tag_text varchar(255),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- メンター募集提案テーブル
drop table if exists mentor_recruitment_proposals;

create table mentor_recruitment_proposals (
    id uuid primary key not null,
    user_id uuid not null references users(id),
    mentor_recruitment_id uuid not null references mentor_recruitments(id),
    description text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- 契約リクエストテーブル
drop table if exists contract_requests;

create table contract_requests (
    id uuid primary key not null,
    message text not null,
    user_id uuid not null references users(id),
    price_at_request int not null,
    plan_id uuid not null references plans(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_contract_requests_user_id on contract_requests (user_id);

-- 契約テーブル
drop table if exists contracts;

create table contracts (
    id uuid primary key not null,
    user_id uuid not null references users(id),
    plan_id uuid not null references plans(id),
    message text not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
)

-- メンタープランテーブル
drop table if exists plans;

create table plans (
    id uuid primary key not null,
    title varchar(255) not null,
    category_id uuid not null references categories(id),
    description text not null,
    status varchar(255) not null,
    consultation_type varchar(255) not null,
    price int not null,
    consultation_method varchar(255),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- メンタープランとタグの中間テーブル
drop table if exists plan_tags;

create table plan_tags (
    id uuid primary key not null,
    plan_id uuid not null references plans(id),
    tag_id uuid not null references tags(id),
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- メンター募集カテゴリテーブル
drop table if exists categories;

create table categories (
    id uuid primary key not null,
    name varchar(255) not null unique,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

insert into categories (id, name) values (uuid_generate_v4(), 'プログラミング');
insert into categories (id, name) values (uuid_generate_v4(), 'マーケティング');
insert into categories (id, name) values (uuid_generate_v4(), 'デザイン');
insert into categories (id, name) values (uuid_generate_v4(), 'ライティング');
insert into categories (id, name) values (uuid_generate_v4(), '動画・映像');
insert into categories (id, name) values (uuid_generate_v4(), 'ビジネス');
insert into categories (id, name) values (uuid_generate_v4(), '語学');
insert into categories (id, name) values (uuid_generate_v4(), 'ライフスタイル');

-- 相談形式・相談方式のマスターを作ってもいいかも
-- タグは選択と記入両方いけるので、マスターのidを参照させるのではなくてテキストを直で中間テーブルに入れる。
-- 公開・中止の型どうするか