CREATE TABLE users
(
    id integer NOT NULL,
    login character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    pass character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    PRIMARY KEY (id)
);