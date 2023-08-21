create user robert_admin
    superuser
    createdb
    createrole
    replication
    bypassrls;


create table orders
(
    orderuid          varchar not null
        unique,
    tracknumber       varchar not null
        primary key,
    entry             varchar,
    name              varchar,
    phone             varchar,
    zip               varchar,
    city              varchar,
    address           varchar,
    region            varchar,
    email             varchar,
    transaction       varchar,
    requestid         varchar,
    currency          varchar,
    provider          varchar,
    amount            integer,
    paymentdt         integer,
    bank              varchar,
    deliverycost      integer,
    goodstotal        integer,
    customfee         integer,
    locale            varchar,
    internalsignature varchar,
    customerid        varchar,
    deliveryservice   varchar,
    shardkey          varchar,
    smid              integer,
    datecreated       timestamp,
    oofshard          varchar
);

alter table orders
    owner to robert_admin;

create table order_item
(
    chrtid      integer,
    tracknumber varchar
        constraint tracknumberfk
            references orders,
    price       integer,
    rid         varchar,
    name        varchar,
    sale        integer,
    size        varchar,
    totalprice  integer,
    nmid        integer,
    brand       varchar,
    status      integer
);

alter table order_item
    owner to robert_admin;

