create table User(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `username` VARCHAR(100) not NULL,
    `password` VARCHAR(100) not null,
    PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;