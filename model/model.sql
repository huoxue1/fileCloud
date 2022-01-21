create table file
(
    id        TEXT not null
        constraint file_pk
            primary key,
    file_name TEXT not null,
    is_dir    integer default 0,
    content   blob,
    upload_time INTEGER,
    parent_id TEXT default ''
);

create unique index file_id_uindex
    on file (id);

