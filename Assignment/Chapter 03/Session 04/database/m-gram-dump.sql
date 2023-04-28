--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2023-04-20 05:07:36

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3363 (class 1262 OID 99475)
-- Name: m-gram; Type: DATABASE; Schema: -; Owner: jampirojam
--

CREATE DATABASE "m-gram" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Indonesia.1252';


ALTER DATABASE "m-gram" OWNER TO jampirojam;

\connect -reuse-previous=on "dbname='m-gram'"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3364 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 99527)
-- Name: comments; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.comments (
    id bigint NOT NULL,
    message text NOT NULL,
    url text NOT NULL,
    user_id bigint NOT NULL,
    photo_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.comments OWNER TO jampirojam;

--
-- TOC entry 215 (class 1259 OID 99526)
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.comments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comments_id_seq OWNER TO jampirojam;

--
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 215
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- TOC entry 214 (class 1259 OID 99510)
-- Name: photos; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.photos (
    id bigint NOT NULL,
    title text NOT NULL,
    caption text NOT NULL,
    url text NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.photos OWNER TO jampirojam;

--
-- TOC entry 213 (class 1259 OID 99509)
-- Name: photos_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.photos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.photos_id_seq OWNER TO jampirojam;

--
-- TOC entry 3366 (class 0 OID 0)
-- Dependencies: 213
-- Name: photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.photos_id_seq OWNED BY public.photos.id;


--
-- TOC entry 212 (class 1259 OID 99493)
-- Name: social_media; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.social_media (
    id bigint NOT NULL,
    name text NOT NULL,
    url text NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.social_media OWNER TO jampirojam;

--
-- TOC entry 211 (class 1259 OID 99492)
-- Name: social_media_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.social_media_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.social_media_id_seq OWNER TO jampirojam;

--
-- TOC entry 3367 (class 0 OID 0)
-- Dependencies: 211
-- Name: social_media_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.social_media_id_seq OWNED BY public.social_media.id;


--
-- TOC entry 210 (class 1259 OID 99477)
-- Name: users; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text NOT NULL,
    user_name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    age bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO jampirojam;

--
-- TOC entry 209 (class 1259 OID 99476)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO jampirojam;

--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 209
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3191 (class 2604 OID 99530)
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- TOC entry 3187 (class 2604 OID 99513)
-- Name: photos id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.photos ALTER COLUMN id SET DEFAULT nextval('public.photos_id_seq'::regclass);


--
-- TOC entry 3183 (class 2604 OID 99496)
-- Name: social_media id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.social_media ALTER COLUMN id SET DEFAULT nextval('public.social_media_id_seq'::regclass);


--
-- TOC entry 3179 (class 2604 OID 99480)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3357 (class 0 OID 99527)
-- Dependencies: 216
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.comments VALUES (2, 'biasa aja kok', 'http://m-gram/com/photo/7/5d02b2ba-75ab-4a54-96f4-e52b2c814a46/a7762e20-ff81-41ac-8324-4edee2ebe3ae', 7, 1, '2023-04-19 17:32:11.998605+07', '2023-04-19 17:32:11.998605+07', NULL, false);
INSERT INTO public.comments VALUES (1, 'cakep banget', 'http://m-gram/com/photo/7/5d02b2ba-75ab-4a54-96f4-e52b2c814a46/d8d62e20-098a-41ac-8324-4edee2ebe3ae', 3, 1, '2023-04-19 17:16:08.266662+07', '2023-04-19 17:16:08.266662+07', NULL, false);
INSERT INTO public.comments VALUES (4, 'ngikut aje lo', 'http://m-gram.com/photo/10/98792416-7f4d-4722-89c1-f9dd6734300b/d8d62e20-ff81-41ac-8324-4edee2ebe3ae', 7, 3, '2023-04-20 04:55:53.553984+07', '2023-04-20 04:55:53.553984+07', NULL, false);
INSERT INTO public.comments VALUES (3, 'kapan healing bareng?', 'http://m-gram/com/photo/7/5d02b2ba-75ab-4a54-96f4-e52b2c814a46/d8d95e20-098a-41ac-8324-4edee2ebe3ae', 10, 1, '2023-04-20 04:53:45.151716+07', '2023-04-20 04:53:45.151716+07', NULL, false);


--
-- TOC entry 3355 (class 0 OID 99510)
-- Dependencies: 214
-- Data for Name: photos; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.photos VALUES (3, 'Senang', 'Menarilah', 'http://m-gram.com/photo/10/98792416-7f4d-4722-89c1-f9dd6734300b/', 10, '2023-04-20 04:55:14.505688+07', '2023-04-20 04:55:14.505688+07', NULL, false);
INSERT INTO public.photos VALUES (1, 'Semangat', 'Menyemangati diri dengan bernyanyi', 'http://m-gram/com/photo/7/5d02b2ba-75ab-4a54-96f4-e52b2c814a46/', 7, '2023-04-19 17:12:31.73284+07', '2023-04-19 17:12:31.73284+07', NULL, false);
INSERT INTO public.photos VALUES (2, 'Senang', 'Menarilah', 'http://m-gram.com/photo/10/e34948ce-95f2-4f81-8810-3a2100001089/', 10, '2023-04-19 17:47:55.025086+07', '2023-04-19 17:47:55.025086+07', NULL, false);


--
-- TOC entry 3353 (class 0 OID 99493)
-- Dependencies: 212
-- Data for Name: social_media; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.social_media VALUES (1, 'Gram Account', 'http://m-gram.com/social-media/7/113f8ffd-f07a-4024-b7a1-30181fbee80f/', 7, '2023-04-19 17:34:14.566722+07', '2023-04-19 17:34:14.566722+07', NULL, false);
INSERT INTO public.social_media VALUES (2, 'Gram Alter', 'http://m-gram.com/social-media/7/0b6351e4-e874-4d9b-abaf-0efb6d56981f/', 7, '2023-04-19 17:34:28.578939+07', '2023-04-19 17:34:28.578939+07', NULL, false);
INSERT INTO public.social_media VALUES (3, 'G Account', 'http://m-gram.com/social-media/10/b48d7420-bb88-4c6b-a9c9-70c4f7cb002c/', 10, '2023-04-19 17:47:02.860307+07', '2023-04-19 17:47:02.860307+07', NULL, false);


--
-- TOC entry 3351 (class 0 OID 99477)
-- Dependencies: 210
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.users VALUES (7, 'Ojam', 'ojam', 'ojam@jd.com', '$2a$08$hrUy8PF6uS1kajECim5NAuRSeYgPdWNbF6.kubjYY56VzN9sHqamy', 26, '2023-04-19 16:05:57.324834+07', '2023-04-19 16:05:57.324834+07', NULL, false);
INSERT INTO public.users VALUES (1, 'Jandara Khab', 'jandara', 'jandara@jd.com', '$2a$08$rkg7iHKlfEamLevHQKgGxukF6qP0D1BBlFROZtoetWbdiSM/MWNEe', 24, '2023-04-19 15:06:34.754578+07', '2023-04-19 15:06:34.754578+07', NULL, false);
INSERT INTO public.users VALUES (3, 'Jason', 'jason', 'jason@jd.com', '$2a$08$39PX.VYrU2GXSfurkrPipOteTp2eFWzW56i6tzquEfQuz1xnqyVUa', 19, '2023-04-19 15:48:08.228571+07', '2023-04-19 15:48:08.228571+07', NULL, false);
INSERT INTO public.users VALUES (5, 'Jason Wang', 'jwang', 'jasson@jd.com', '$2a$08$2aM7R8gIlUWTScTyA5KjdeK0g8c2wruG7Klq5plOt2WF9y/o4AVbm', 35, '2023-04-19 15:53:59.255857+07', '2023-04-19 15:53:59.255857+07', NULL, false);
INSERT INTO public.users VALUES (8, 'Jessica Syah', 'jessy', 'jasonz@jd.com', '$2a$08$WR/qPEPD/uYAwCaAmF4ADeXLZlm3X5t1JgiHXnZ7EFzGrqBBr.jRq', 33, '2023-04-19 17:37:18.31235+07', '2023-04-19 17:37:18.31235+07', NULL, false);
INSERT INTO public.users VALUES (10, 'Wira Negara', 'wira', 'gilang@jd.com', '$2a$08$1KzqXxanGdT0nORJ6LqKUe27o7Vx1zWBBwtE3VOcGS/o.HOSIxnX2', 28, '2023-04-19 17:46:29.998725+07', '2023-04-19 17:46:29.998725+07', NULL, false);


--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 215
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.comments_id_seq', 4, true);


--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 213
-- Name: photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.photos_id_seq', 3, true);


--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 211
-- Name: social_media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.social_media_id_seq', 3, true);


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 209
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.users_id_seq', 10, true);


--
-- TOC entry 3206 (class 2606 OID 99537)
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- TOC entry 3204 (class 2606 OID 99520)
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 99503)
-- Name: social_media social_media_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.social_media
    ADD CONSTRAINT social_media_pkey PRIMARY KEY (id);


--
-- TOC entry 3196 (class 2606 OID 99491)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3198 (class 2606 OID 99487)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3200 (class 2606 OID 99489)
-- Name: users users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);


--
-- TOC entry 3209 (class 2606 OID 99538)
-- Name: comments fk_photos_comment; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_photos_comment FOREIGN KEY (photo_id) REFERENCES public.photos(id);


--
-- TOC entry 3210 (class 2606 OID 99543)
-- Name: comments fk_users_comments; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_users_comments FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3208 (class 2606 OID 99521)
-- Name: photos fk_users_photos; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT fk_users_photos FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- TOC entry 3207 (class 2606 OID 99504)
-- Name: social_media fk_users_social_media; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.social_media
    ADD CONSTRAINT fk_users_social_media FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2023-04-20 05:07:36

--
-- PostgreSQL database dump complete
--

