CREATE TABLE public.user
(
    id serial,
    username text,
    password text,
    email text,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.user
    OWNER to postgres;



CREATE TABLE public.review
(
    id integer NOT NULL DEFAULT nextval('"User_id_seq"'::regclass),
    "user" integer,
    title text,
    content text,
    PRIMARY KEY (id),
    CONSTRAINT fk_review_user FOREIGN KEY ("user")
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

ALTER TABLE IF EXISTS public.review
    OWNER to postgres;