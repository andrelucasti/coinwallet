create table wallet (
	uuid uuid primary key default uuid_generate_v4(),
	name varchar(200) not null,
	user_id uuid not null,
	value BIGINT not null default 0,
	CREATED_DATE TIMESTAMP WITH TIME ZONE NOT null,
    LAST_MODIFIED_DATE TIMESTAMP WITH TIME ZONE NOT null default now()
);
