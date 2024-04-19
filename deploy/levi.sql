create schema account;
create schema product;
create schema trade;

create table account.account
(
    account_id bigint generated always as identity primary key,
    email      varchar(100) unique,
    name       varchar(100),
    avatar_url varchar(255),
    is_violate bool default false,
    create_at  date default current_date
);

create type account.gender as enum ('male', 'female', 'other');
create table account.detail
(
    account_id bigint primary key,
    phone      varchar(50)  default '',
    address    varchar(100) default '',
    gender     account.gender
);

create type account.bank as enum ('viettin', 'techcom', 'mb');
create table account.payment
(
    account_id bigint primary key,
    bank       account.bank,
    number     varchar(50) default ''
);

create table account.social
(
    account_id         bigint primary key,
    booth_following_id bigint[] default array []::integer[],
    booth_like_id      bigint[] default array []::integer[],
    friend_request_id  bigint[] default array []::integer[],
    friend_id          bigint[] default array []::integer[]
);

create table account.cart
(
    account_id       bigint primary key,
    order_delivering int default 0,
    order_paid       int default 0
);

create table account.booth
(
    account_id            bigint primary key,
    order_id_wait_confirm bigint[] default array []::integer[],
    picture_avatar_url    varchar(100),
    picture_cover_url     varchar(100),
    description           text,
    is_authentic          bool     default false
);

create table product.product
(
    product_id           bigint generated always as identity primary key,
    account_id           bigint,
    name                 text,
    picture_avatar_url   text,
    pictures_description text[],
    description          text,
    price                real,
    total_score          real,
    create_at            timestamp default current_timestamp
);

create table product.category
(
    product_id bigint,
    l1         varchar(50),
    l2         varchar(50),
    l3         varchar(50)
);

create table product.sale
(
    product_id bigint,
    create_at  timestamp default current_timestamp,
    expire     timestamp check (expire > current_timestamp),
    discount   real check (discount > 0)
);

create table product.sku
(
    product_id bigint,
    size       int,
    name       varchar(50),
    price      money
);

create table product.rating
(
    product_id  bigint,
    account_id  bigint,
    comment     text,
    picture_url varchar(100),
    score       real check (score > 0),
    create_at   timestamp default current_timestamp
);

create table product.inventory
(
    product_id bigint,
    total      bigint,
    sold       bigint
);

create type trade.order_status as enum ('pending', 'delivering', 'paid');
create table trade.order
(
    order_id     bigint generated always as identity primary key,
    account_id   bigint,
    product_id   bigint,
    total_amount money,
    status       trade.order_status default 'pending',
    last_update  timestamp          default current_timestamp,
    create_at    timestamp          default current_timestamp
);

create type trade.shipping_unit as enum ('express', 'standard');
create type trade.shipping_state as enum ('taking', 'comming_to_warehouse', 'shipping');
create table trade.ship
(
    order_id       bigint primary key,
    price          money,
    shipping_unit  trade.shipping_unit  default 'standard',
    shipping_state trade.shipping_state default 'taking',
    description    text,
    last_update    timestamp            default current_timestamp
);

