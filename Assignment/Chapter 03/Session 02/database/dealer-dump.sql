--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2023-04-16 03:30:34

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
-- TOC entry 3364 (class 1262 OID 83091)
-- Name: dealer; Type: DATABASE; Schema: -; Owner: jampirojam
--

CREATE DATABASE dealer WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_Indonesia.1252';


ALTER DATABASE dealer OWNER TO jampirojam;

\connect dealer

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
-- TOC entry 3365 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 83165)
-- Name: product_types; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.product_types (
    id bigint NOT NULL,
    type text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.product_types OWNER TO jampirojam;

--
-- TOC entry 209 (class 1259 OID 83164)
-- Name: product_types_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.product_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_types_id_seq OWNER TO jampirojam;

--
-- TOC entry 3366 (class 0 OID 0)
-- Dependencies: 209
-- Name: product_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.product_types_id_seq OWNED BY public.product_types.id;


--
-- TOC entry 212 (class 1259 OID 83179)
-- Name: products; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    name text NOT NULL,
    product_type_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false,
    price bigint NOT NULL
);


ALTER TABLE public.products OWNER TO jampirojam;

--
-- TOC entry 211 (class 1259 OID 83178)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_id_seq OWNER TO jampirojam;

--
-- TOC entry 3367 (class 0 OID 0)
-- Dependencies: 211
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- TOC entry 214 (class 1259 OID 83196)
-- Name: user_roles; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.user_roles (
    id bigint NOT NULL,
    role_type text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.user_roles OWNER TO jampirojam;

--
-- TOC entry 213 (class 1259 OID 83195)
-- Name: user_roles_id_seq; Type: SEQUENCE; Schema: public; Owner: jampirojam
--

CREATE SEQUENCE public.user_roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_roles_id_seq OWNER TO jampirojam;

--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 213
-- Name: user_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.user_roles_id_seq OWNED BY public.user_roles.id;


--
-- TOC entry 216 (class 1259 OID 91286)
-- Name: users; Type: TABLE; Schema: public; Owner: jampirojam
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text NOT NULL,
    user_name text NOT NULL,
    password text NOT NULL,
    product_type_id bigint,
    user_role_id bigint,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    deleted boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO jampirojam;

--
-- TOC entry 215 (class 1259 OID 91285)
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
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jampirojam
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3179 (class 2604 OID 83168)
-- Name: product_types id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.product_types ALTER COLUMN id SET DEFAULT nextval('public.product_types_id_seq'::regclass);


--
-- TOC entry 3183 (class 2604 OID 83182)
-- Name: products id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- TOC entry 3187 (class 2604 OID 83199)
-- Name: user_roles id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.user_roles ALTER COLUMN id SET DEFAULT nextval('public.user_roles_id_seq'::regclass);


--
-- TOC entry 3191 (class 2604 OID 91289)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3352 (class 0 OID 83165)
-- Dependencies: 210
-- Data for Name: product_types; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.product_types VALUES (1, 'CAR', '2023-04-11 23:00:34.862037+07', '2023-04-11 23:00:34.862037+07', NULL, false);
INSERT INTO public.product_types VALUES (2, 'MOTOR', '2023-04-11 23:00:44.978942+07', '2023-04-11 23:00:44.978942+07', NULL, false);
INSERT INTO public.product_types VALUES (3, 'CYCLE', '2023-04-11 23:00:53.132842+07', '2023-04-11 23:00:53.132842+07', NULL, false);


--
-- TOC entry 3354 (class 0 OID 83179)
-- Dependencies: 212
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.products VALUES (7, 'Sepeda BMX New Era', 3, '2023-04-14 01:11:26.915544+07', '2023-04-14 01:11:26.915544+07', NULL, false, 8500000);
INSERT INTO public.products VALUES (8, 'Sepeda Maximo New Era', 3, '2023-04-14 01:23:18.400107+07', '2023-04-14 01:23:18.400107+07', NULL, false, 2500000);
INSERT INTO public.products VALUES (9, 'Sepeda Ardiless', 3, '2023-04-14 01:27:10.67403+07', '2023-04-14 01:27:10.67403+07', NULL, false, 4500000);
INSERT INTO public.products VALUES (10, 'Sepeda Phoenix', 3, '2023-04-14 02:58:07.863101+07', '2023-04-14 02:58:07.863101+07', NULL, false, 1259900);
INSERT INTO public.products VALUES (11, 'Sepeda Poligon', 3, '2023-04-14 02:58:23.756726+07', '2023-04-14 02:58:23.756726+07', NULL, false, 965000);
INSERT INTO public.products VALUES (13, 'Sepeda Bromthink', 3, '2023-04-14 02:59:55.257011+07', '2023-04-14 02:59:55.257011+07', NULL, false, 42500000);
INSERT INTO public.products VALUES (3, 'Mobil Phanter', 1, '2023-04-14 03:29:54.19926+07', '2023-04-14 03:29:54.19926+07', NULL, false, 0);
INSERT INTO public.products VALUES (14, 'Mobil Albatros', 1, '2023-04-14 03:50:57.599704+07', '2023-04-14 03:50:57.599704+07', NULL, false, 134500000);
INSERT INTO public.products VALUES (15, 'Mobil Albatros', 1, '2023-04-14 09:01:34.63753+07', '2023-04-14 09:01:34.63753+07', NULL, false, 134500000);
INSERT INTO public.products VALUES (12, 'Sepeda Bromton', 3, '2023-04-14 02:58:32.601569+07', '2023-04-15 23:27:52.156201+07', '2023-04-15 23:27:52.052346+07', true, 49500000);
INSERT INTO public.products VALUES (1, 'Motor Honda Beat', 2, '2023-04-14 00:12:42.254729+07', '2023-04-14 00:12:42.254729+07', NULL, false, 17560000);
INSERT INTO public.products VALUES (2, 'Motor Yamaha Mio', 2, '2023-04-14 00:29:22.110774+07', '2023-04-14 00:29:22.110774+07', NULL, false, 15500000);
INSERT INTO public.products VALUES (4, 'Motor Honda Scoopy', 2, '2023-04-14 00:42:08.547306+07', '2023-04-14 00:42:08.547306+07', NULL, false, 15500000);
INSERT INTO public.products VALUES (5, 'Mobil Honda Brio', 1, '2023-04-14 00:42:39.398476+07', '2023-04-14 00:42:39.398476+07', NULL, false, 215500000);
INSERT INTO public.products VALUES (6, 'Motor Suzuki Satria FU', 2, '2023-04-14 00:44:40.172586+07', '2023-04-14 00:44:40.172586+07', NULL, false, 19500000);


--
-- TOC entry 3356 (class 0 OID 83196)
-- Dependencies: 214
-- Data for Name: user_roles; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.user_roles VALUES (1, 'HEAD', '2023-04-11 23:01:38.262895+07', '2023-04-11 23:01:38.262895+07', NULL, false);
INSERT INTO public.user_roles VALUES (2, 'ADMIN_MOBIL', '2023-04-11 23:01:57.416406+07', '2023-04-11 23:01:57.416406+07', NULL, false);
INSERT INTO public.user_roles VALUES (3, 'ADMIN_MOTOR', '2023-04-11 23:02:06.003888+07', '2023-04-11 23:02:06.003888+07', NULL, false);
INSERT INTO public.user_roles VALUES (4, 'ADMIN_SEPEDA', '2023-04-11 23:02:15.894553+07', '2023-04-11 23:02:15.894553+07', NULL, false);


--
-- TOC entry 3358 (class 0 OID 91286)
-- Dependencies: 216
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: jampirojam
--

INSERT INTO public.users VALUES (1, 'Aliy Shiro', 'aliys', '$2a$08$c2jDk1bdoVQeCPrf/ISBmeGmFetn3NBs2KqQcLezHIVlaoNbE9tFG', 2, 1, '2023-04-11 23:51:00.196891+07', '2023-04-11 23:51:00.196891+07', NULL, false);
INSERT INTO public.users VALUES (2, 'Bellatrix Potter', 'bellatrix', '$2a$08$a8b0/LtDo57oaQoYFZPNk.Xq./Bx.24EmJKuAYQrXtCKubxkeVc3q', 2, 2, '2023-04-14 00:43:52.761475+07', '2023-04-14 00:43:52.761475+07', NULL, false);
INSERT INTO public.users VALUES (3, 'Ojam', 'ojam', '$2a$08$cIOzbjuhO.IdorFfhxUe0O.od7p1yxMHEqYl9Sac8wQ5Db5bI7Wcm', 3, 4, '2023-04-14 01:09:56.673161+07', '2023-04-14 01:09:56.673161+07', NULL, false);
INSERT INTO public.users VALUES (4, 'Machrush', 'mach', '$2a$08$WMxrOh587wZe4yMExYEu6.w9syeKhgs5tVh3vOaIVq7YGP9ejs.qq', 1, 2, '2023-04-14 02:57:45.145455+07', '2023-04-14 02:57:45.145455+07', NULL, false);
INSERT INTO public.users VALUES (5, 'Jandara Khab', 'jandara', '$2a$08$3SHqYvIdPuPIjuGeXOhnlOgzYUYWdW6qTqY02jfimqeqpNLF688qm', 3, 1, '2023-04-14 03:49:56.859343+07', '2023-04-14 03:49:56.859343+07', NULL, false);


--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 209
-- Name: product_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.product_types_id_seq', 3, true);


--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 211
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.products_id_seq', 15, true);


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 213
-- Name: user_roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.user_roles_id_seq', 4, true);


--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: jampirojam
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- TOC entry 3196 (class 2606 OID 83175)
-- Name: product_types product_types_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.product_types
    ADD CONSTRAINT product_types_pkey PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 83177)
-- Name: product_types product_types_type_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.product_types
    ADD CONSTRAINT product_types_type_key UNIQUE (type);


--
-- TOC entry 3200 (class 2606 OID 83189)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3202 (class 2606 OID 83206)
-- Name: user_roles user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (id);


--
-- TOC entry 3204 (class 2606 OID 83208)
-- Name: user_roles user_roles_role_type_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_role_type_key UNIQUE (role_type);


--
-- TOC entry 3206 (class 2606 OID 91296)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3208 (class 2606 OID 91298)
-- Name: users users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);


--
-- TOC entry 3209 (class 2606 OID 83190)
-- Name: products fk_product_types_products; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_product_types_products FOREIGN KEY (product_type_id) REFERENCES public.product_types(id);


--
-- TOC entry 3210 (class 2606 OID 91299)
-- Name: users fk_product_types_users; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_product_types_users FOREIGN KEY (product_type_id) REFERENCES public.product_types(id);


--
-- TOC entry 3211 (class 2606 OID 91304)
-- Name: users fk_user_roles_users; Type: FK CONSTRAINT; Schema: public; Owner: jampirojam
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_user_roles_users FOREIGN KEY (user_role_id) REFERENCES public.user_roles(id);


-- Completed on 2023-04-16 03:30:35

--
-- PostgreSQL database dump complete
--

