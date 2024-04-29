create table game_open
(
    id                 bigint       not null
        primary key,
    status             int          not null,
    open_result        varchar(50)  not null,
    open_source        varchar(100) not null,
    open_at            datetime     not null,
    issue              bigint       not null,
    game_code          int          not null,
    game_name          varchar(50)  not null,
    close_at           datetime     not null,
    open_result_detail varchar(500) null,
    won_result         varchar(250) not null,
    issue_detail       bigint       not null
);

create table game_open_item
(
    id           bigint      not null
        primary key,
    game_open_id bigint      not null,
    play_code    int         not null,
    play_id      int         not null,
    play_name    varchar(50) not null,
    rate         bigint      not null,
    sort         int         not null,
    status       int         not null,
    game_name    varchar(50) not null,
    game_code    int         not null,
    play_type    varchar(50) not null
);

create table order_balance
(
    order_no          bigint       not null
        primary key,
    uid               bigint       not null,
    account           varchar(50)  not null,
    pid               bigint       not null,
    balance_code      int          not null,
    title             varchar(50)  not null,
    balance_before    bigint       not null,
    balance_after     bigint       not null,
    created_at        datetime     not null,
    note              varchar(500) not null,
    order_no_relation bigint       not null,
    tramper_no        varchar(50)  not null,
    parent_path       varchar(500) not null,
    balance           bigint       not null
);

create table order_bet
(
    order_no    bigint        not null
        primary key,
    uid         bigint        not null,
    pid         bigint        not null,
    account     varchar(50)   not null,
    game_code   int           not null,
    game_type   varchar(20)   not null,
    amount      bigint        not null,
    status      int           not null,
    game_name   varchar(50)   not null,
    won         bigint        not null,
    play_code   int           not null,
    play_name   varchar(50)   not null,
    title       varchar(50)   not null,
    parent_path varchar(1000) not null,
    open_result varchar(50)   not null,
    create_at   datetime      not null,
    settle_at   datetime      not null,
    rate        bigint        not null,
    bet_content varchar(50)   not null
);

create table order_bet_settle
(
    order_no               bigint        not null
        primary key,
    uid                    bigint        not null,
    pid                    bigint        not null,
    account                varchar(50)   not null,
    game_code              int           not null,
    game_type              varchar(20)   not null,
    amount                 bigint        not null,
    status                 int           not null,
    game_name              varchar(50)   not null,
    won                    bigint        not null,
    play_code              int           not null,
    play_name              varchar(50)   not null,
    title                  varchar(50)   not null,
    parent_path            varchar(1000) not null,
    open_result            varchar(50)   not null,
    create_at              datetime      not null,
    settle_at              datetime      not null,
    rate                   bigint        not null,
    bet_content            varchar(50)   not null,
    game_open_id           bigint        not null,
    game_open_issue        bigint        not null,
    game_open_issue_detail bigint        not null
);

create table order_deposit
(
    order_no         bigint        not null,
    account          varchar(50)   not null,
    uid              bigint        not null,
    pid              bigint        not null,
    status           int           not null,
    finish_at        datetime      not null,
    detail           varchar(500)  not null,
    status_remark    varchar(250)  not null,
    amount           bigint        not null,
    sys_remark       varchar(250)  not null,
    address          varchar(100)  not null,
    net              varchar(100)  null,
    amount_item_code int           not null,
    currency         varchar(20)   not null,
    protocol         varchar(20)   not null,
    parnet_path      varchar(1000) not null
);

create table order_withdraw
(
    order_no         bigint       not null
        primary key,
    account          varchar(50)  not null,
    uid              bigint       not null,
    pid              bigint       not null,
    status           int          not null,
    finish_at        datetime     not null,
    detail           varchar(500) not null,
    status_remark    varchar(250) not null,
    amount           bigint       not null,
    sys_remark       varchar(250) not null,
    address          varchar(100) not null,
    amount_finally   bigint       not null,
    fee              bigint       not null,
    created_at       datetime     not null,
    net              varchar(100) not null,
    amount_item_code int          not null,
    currency         varchar(20)  not null,
    protocol         varchar(20)  not null
);

create table sys_amount
(
    id       int auto_increment
        primary key,
    title    varchar(250) null,
    category varchar(50)  not null comment '1:区块链，银行卡',
    status   int          null,
    type     varchar(50)  not null comment 'deposit,withdraw'
);

create table sys_amount_item
(
    code      int auto_increment
        primary key,
    title     varchar(150)  not null,
    status    int           not null,
    detail    varchar(1000) null,
    amount_id int           not null,
    net       varchar(50)   not null comment 'erc20 trc20',
    min       bigint        not null,
    max       bigint        not null,
    fee       bigint        not null,
    type      varchar(20)   not null,
    logo      varchar(200)  not null,
    sort      int           not null,
    category  varchar(20)   not null,
    country   varchar(50)   not null,
    currency  varchar(20)   not null,
    protocol  varchar(20)   not null
);

create table sys_balance_code
(
    code   int          not null
        primary key,
    remark varchar(500) not null,
    status int          not null,
    type   varchar(50)  not null,
    title  varchar(50)  null
);

create table sys_conf
(
    code   varchar(50)  not null
        primary key,
    detail varchar(500) not null,
    status int          not null,
    type   varchar(100) not null comment 'user,deposit,admin',
    remark varchar(250) not null
);

create table sys_game
(
    code             int         not null
        primary key,
    name             varchar(50) not null,
    status           int         not null,
    created_at       datetime    not null,
    start_time       varchar(20) not null,
    end_time         varchar(20) not null,
    total_issue      int         not null,
    interval_seconds int         not null comment '间隔秒数',
    type             varchar(20) not null,
    sort             int         not null
);

create table sys_game_play
(
    id        int auto_increment
        primary key,
    game_code int         not null,
    game_name varchar(50) null,
    play_name varchar(50) not null,
    status    int         not null,
    play_code int         not null,
    play_type varchar(50) not null,
    sort      int         not null,
    bet_min   bigint      not null,
    bet_max   bigint      not null,
    `group`   varchar(20) not null
);

create table sys_game_type
(
    code   varchar(20)  not null
        primary key,
    status int          not null,
    logo   varchar(200) not null,
    name   varchar(20)  not null,
    remark varchar(250) not null
);

create table sys_play
(
    code   int         not null
        primary key,
    name   varchar(20) not null,
    status int         not null,
    type   varchar(50) not null
);

create table sys_play_type
(
    code   varchar(20)  not null
        primary key,
    status int          not null,
    logo   varchar(200) not null,
    name   varchar(50)  not null
);

create table user_address
(
    id      int auto_increment
        primary key,
    uid     bigint       not null,
    account varchar(50)  not null,
    address varchar(128) not null,
    net     varchar(50)  not null,
    status  int          not null
);

create table user_amount
(
    uid            bigint auto_increment
        primary key,
    email          varchar(50) not null,
    balance        bigint      not null,
    total_bet      bigint      not null,
    total_deposit  bigint      not null,
    total_withdraw bigint      not null,
    freeze         bigint      not null,
    account        varchar(50) not null,
    parent_path    int         null,
    total_profit   bigint      not null,
    total_gift     bigint      not null
);

create table user_base
(
    uid          bigint auto_increment
        primary key,
    account      varchar(100) default '' not null,
    email        varchar(50)             not null,
    password     varchar(128)            not null,
    status       int                     null comment '1:正常，2：冻结',
    xid          varchar(10)             not null comment 'short code',
    ip           varchar(50)             not null,
    client_agent varchar(50)             not null,
    mobile       char(20)                not null,
    level_bet    int                     not null,
    level_pay    int                     null,
    level_agent  int                     not null,
    pid          bigint                  null,
    created_at   datetime                not null,
    updated_at   datetime                not null,
    parent_path  varchar(1000)           not null,
    country      varchar(50)             not null
);

create table user_digital
(
    id            int auto_increment
        primary key,
    address       varchar(100) not null,
    net           varchar(20)  not null,
    status        int          not null,
    count         int          not null,
    private_key   varchar(250) not null,
    total_deposit bigint       not null,
    uid           bigint       not null,
    account       varchar(120) not null
);

