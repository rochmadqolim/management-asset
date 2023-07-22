create type act_st as enum(
	'input',
	'sold',
	'retur'
);

create type act_wh as enum(
	'input',
	'output'
);


--================================================================Warehouse Part=============================================================================
create table product_wh(
id varchar(50) primary key,
product_name varchar(50),
price int not null,
product_category varchar(50) not null,
stock int
)
create table report_trx_wh(
	id serial ,
	product_wh_id varchar(50),
	stock integer,
	product_name varchar(50),
	act act_wh,
	last_stock int,
	created_at date
)

create table trx_wh(
id serial primary key,
	product_wh_id varchar(50),
	stock integer,
	act act_wh,
	created_at date
)

select * from  product_wh
select * from trx_wh
select * from report_trx_wh
ALTER table trx_wh ADD CONSTRAINT fk_trx_in_wh_product_wh foreign key (product_wh_id) references product_wh(id)
alter table product_wh alter column stock set default 0;
--====================================================Store Part =================================================================
create table product_st(
id varchar(50) primary key,
product_name varchar(50),
price int not null,
product_category varchar(50) not null,
stock int
)

create table report_trx_st(
	id serial ,
	product_st_id varchar(50),
	stock_in integer,
	product_name varchar(50),
	act act_st,
	last_stock int,
	created_at date
)

create table trx_st(
id varchar(50) primary key,
	product_st_id varchar(50),
	stock_in integer,
	act act_st,
	created_at date
)
select * from  product_st
select * from trx_st
select * from report_trx_st

alter table product_st alter column stock set default 0;

ALTER table trx_st ADD CONSTRAINT fk_trx_in_st_product_st foreign key (product_st_id) references product_st(id)

-- ============================================ Inventory Control Part or Product So =============================================================================================

create table product_so(
id serial primary key,
product_st_id varchar(50) unique,
stock int,
diff_stock int,
diff_price int
)
ALTER table product_so ADD CONSTRAINT fk_product_so_product_st foreign key (product_st_id) references product_st(id)

alter table product_so alter column stock set default 0;
alter table product_so alter column diff_stock set default 0;
alter table product_so alter column diff_price set default 0;

create table trx_so(
id serial primary key,
product_so_st_id varchar,
stock int
)
ALTER table trx_so ADD CONSTRAINT fk_trx_so_product_st_so_id foreign key (product_so_st_id) references product_so(product_st_id)

create table report_so_detail(
id serial ,
total_loss int not null,
product_min varchar not null,
total_min int not null,
product_max varchar not null,
total_max int not null,
created_at date
)


create table interim_so_report(
id serial ,
total_loss int not null,
product_min varchar not null,
total_min int not null,
product_max varchar not null,
total_max int not null,
created_at date
)
select * from trx_so
select * from interim_so_report
select * from report_so_detail
select * from product_so
select * from product_st
select * from st_team

--=============================================================================================================================================
--role

create table st_team(
id varchar(200) primary key,
	name varchar(200) unique,
	email varchar(200),
	password varchar(200),
	phone varchar(200),
	photo varchar(200)
)


create table ic_team(
id varchar(200) primary key,
	name varchar(200),
	email varchar(200) unique,
	password varchar(200),
	phone varchar(200),
	photo varchar(200)
)
select * from st_team;

create table admin_wh(
id varchar(200) primary key,
	name varchar(200),
	email varchar(200) unique,
	password varchar(200),
	phone varchar(200),
	photo varchar(200)
)