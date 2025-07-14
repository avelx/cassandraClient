CREATE TABLE Users (
                       id          int,
                       firstName   varchar(50),
                       lastName    varchar(60),
                       CONSTRAINT production UNIQUE(id)
);