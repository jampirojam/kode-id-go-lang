--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2023-04-05 19:13:38

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
-- TOC entry 3331 (class 1262 OID 74899)
-- Name: book-store; Type: DATABASE; Schema: -; Owner: jampirojam
--

CREATE DATABASE "book-store" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Indonesia.1252';


ALTER DATABASE "book-store" OWNER TO jampirojam;

\connect -reuse-previous=on "dbname='book-store'"

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
-- TOC entry 3332 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 74901)
-- Name: book_categories; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.book_categories (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.book_categories OWNER TO jampirojam;

--
-- TOC entry 209 (class 1259 OID 74900)
-- Name: book_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.book_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_categories_id_seq OWNER TO jampirojam;

--
-- TOC entry 3333 (class 0 OID 0)
-- Dependencies: 209
-- Name: book_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.book_categories_id_seq OWNED BY public.book_categories.id;


--
-- TOC entry 212 (class 1259 OID 74912)
-- Name: books; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.books (
    id bigint NOT NULL,
    title text NOT NULL,
    author text,
    publisher text,
    year bigint,
    book_category_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.books OWNER TO jampirojam;

--
-- TOC entry 211 (class 1259 OID 74911)
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.books_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO jampirojam;

--
-- TOC entry 3334 (class 0 OID 0)
-- Dependencies: 211
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- TOC entry 3169 (class 2604 OID 74904)
-- Name: book_categories id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.book_categories ALTER COLUMN id SET DEFAULT nextval('public.book_categories_id_seq'::regclass);


--
-- TOC entry 3172 (class 2604 OID 74915)
-- Name: books id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- TOC entry 3323 (class 0 OID 74901)
-- Dependencies: 210
-- Data for Name: book_categories; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.book_categories VALUES (2, 'Komik', '2023-04-05 02:19:32.607436+07', '2023-04-05 02:19:32.607436+07', NULL, false);
INSERT INTO public.book_categories VALUES (3, 'Teknologi', '2023-04-05 02:19:42.940323+07', '2023-04-05 02:19:42.940323+07', NULL, false);
INSERT INTO public.book_categories VALUES (1, 'Sejarah', '2023-04-05 02:19:22.051648+07', '2023-04-05 02:19:22.051648+07', NULL, false);
INSERT INTO public.book_categories VALUES (4, 'Agama', '2023-04-05 02:23:33.548748+07', '2023-04-05 02:23:33.548748+07', NULL, false);
INSERT INTO public.book_categories VALUES (5, 'Fiksi', '2023-04-05 02:31:58.905569+07', '2023-04-05 02:31:58.905569+07', NULL, false);


--
-- TOC entry 3325 (class 0 OID 74912)
-- Dependencies: 212
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.books VALUES (5, 'Cinta Dalam', 'Queenqu', 'Cerita Remaja', 2007, 5, '2023-04-05 02:32:05.645835+07', '2023-04-05 02:32:05.645835+07', NULL, true);
INSERT INTO public.books VALUES (4, 'Mengenang Dunia Kapur', 'Rudi Taboti', 'Clack Press', 2019, 1, '2023-04-05 02:31:00.058927+07', '2023-04-05 02:31:00.058927+07', NULL, true);
INSERT INTO public.books VALUES (1, 'Belajar Pemrograman Java', 'Albus Griffin', 'Hacktiv8', 2015, 3, '2023-04-05 02:25:43.858911+07', '2023-04-05 17:48:07.241998+07', '2023-04-05 17:48:07.138043+07', true);
INSERT INTO public.books VALUES (7, 'Noktah Darah', 'Aneke Susatyo', 'Camar', 1997, 5, '2023-04-05 18:07:37.462835+07', '2023-04-05 18:07:37.462835+07', NULL, false);
INSERT INTO public.books VALUES (3, 'Naruto', 'Masashi Kishimoto', 'Shueisha', 1997, 2, '2023-04-05 02:29:59.865286+07', '2023-04-05 18:18:20.87153+07', '2023-04-05 18:18:20.750637+07', true);
INSERT INTO public.books VALUES (2, 'Kisah Masa Lalu', 'Rickson Wirawan', 'Kinzai Pub', 1997, 5, '2023-04-05 02:27:03.636259+07', '2023-04-05 18:21:46.686495+07', NULL, false);


--
-- TOC entry 3335 (class 0 OID 0)
-- Dependencies: 209
-- Name: book_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.book_categories_id_seq', 5, true);


--
-- TOC entry 3336 (class 0 OID 0)
-- Dependencies: 211
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.books_id_seq', 7, true);


--
-- TOC entry 3177 (class 2606 OID 74910)
-- Name: book_categories book_categories_name_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.book_categories
    ADD CONSTRAINT book_categories_name_key UNIQUE (name);


--
-- TOC entry 3179 (class 2606 OID 74908)
-- Name: book_categories book_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.book_categories
    ADD CONSTRAINT book_categories_pkey PRIMARY KEY (id);


--
-- TOC entry 3181 (class 2606 OID 74919)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- TOC entry 3182 (class 2606 OID 74920)
-- Name: books fk_book_categories_books; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT fk_book_categories_books FOREIGN KEY (book_category_id) REFERENCES public.book_categories(id);


-- Completed on 2023-04-05 19:13:39

--
-- PostgreSQL database dump complete
--

