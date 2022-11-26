create table user_base
(
    uid          int auto_increment
        primary key,
    email        varchar(50) not null,
    pwd          int         not null,
    xid          varchar(10) not null,
    ip           varchar(50) not null,
    client_agent varchar(50) not null,
    mobile       char(20)    not null,
    status       int         null comment '1:正常，2：冻结',
    level_bet    int         not null,
    level_pay    int         null,
    level_agent  int         not null,
    pid          int         null
);

create table user_money
(
    uid            int            not null
        primary key,
    email          varchar(50)    null,
    balance        decimal(18, 2) null,
    total_bet      decimal(18, 2) null,
    total_deposit  decimal(18, 2) not null,
    total_withdraw decimal(18, 2) not null,
    freeze         decimal(18, 2) not null
);
create table sys_conf_deposit_item
(
    id     int auto_increment
        primary key,
    title  varchar(150)  not null,
    status int           not null,
    detail varchar(1000) null,
    sid    int           not null
);

create table sys_conf_deposit
(
    sid    int auto_increment
        primary key,
    title  varchar(250) null,
    type   int          null comment '1:区块链，银行卡',
    status int          null
);

create table sys_conf
(
    code   varchar(50)  not null
        primary key,
    detail int          not null,
    status int          not null,
    type   varchar(100) null comment 'user,deposit,admin'
);
create table deposit
(
    oid          bigint         not null
        primary key,
    uid          int            null,
    pid          int            null,
    email        varchar(50)    null,
    amount       decimal(18, 2) null,
    status       int            null comment '-1：失败，0处理中 1成功',
    type         int            null comment '1,trc20,2币安，',
    time_done    int            null,
    order_detail varchar(250)   null
);



