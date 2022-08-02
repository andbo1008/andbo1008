CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    lastname varchar(255) not null,
    email varchar(255) not null unique,
    password varchar(255) not null,
    account boolean not null default false
);

CREATE TABLE accounts(
    id serial not null unique,
    currency varchar(255) not null,
    total float not null default 0.00
);
CREATE TABLE accounts_list(
    id serial not null unique,
    users_id int references users(id) on delete cascade not null,
    account_id int references users(id) on delete cascade not null
);
CREATE TABLE transaction_list(
    id serial not null unique,
    users_id_sender int not null,
    account_id int not null,
    user_id_geter int not null,
    currency varchar(255) not null,
    total float not null,
    sendsum float not null,
    date timestamp not null
);