--
-- PostgreSQL database dump
--

-- Dumped from database version 12.10 (Ubuntu 12.10-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.10 (Ubuntu 12.10-0ubuntu0.20.04.1)

-- Started on 2022-05-22 13:54:20 +05

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 22074)
-- Name: connection; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.connection (
    id integer NOT NULL,
    connection_id character varying(64) NOT NULL,
    telegram character varying(250) DEFAULT 'Пока здесь ничего нет'::character varying,
    instagram character varying(250) DEFAULT 'Пока здесь ничего нет'::character varying,
    facebook character varying(250) DEFAULT 'Пока здесь ничего нет'::character varying,
    vk character varying(250) DEFAULT 'Пока здесь ничего нет'::character varying,
    tiktok character varying(250) DEFAULT 'Пока здесь ничего нет'::character varying
);


ALTER TABLE public.connection OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 22072)
-- Name: connection_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.connection_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.connection_id_seq OWNER TO postgres;

--
-- TOC entry 3059 (class 0 OID 0)
-- Dependencies: 208
-- Name: connection_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.connection_id_seq OWNED BY public.connection.id;


--
-- TOC entry 207 (class 1259 OID 22055)
-- Name: identifiers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.identifiers (
    id integer NOT NULL,
    user_id integer NOT NULL,
    connection_id character varying(64) NOT NULL,
    settings_id character varying(64) NOT NULL
);


ALTER TABLE public.identifiers OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 22053)
-- Name: identifiers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.identifiers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.identifiers_id_seq OWNER TO postgres;

--
-- TOC entry 3060 (class 0 OID 0)
-- Dependencies: 206
-- Name: identifiers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.identifiers_id_seq OWNED BY public.identifiers.id;


--
-- TOC entry 211 (class 1259 OID 22097)
-- Name: settings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.settings (
    id integer NOT NULL,
    settings_id character varying(64) NOT NULL,
    logo text DEFAULT 'not_logo.png'::text,
    banner text DEFAULT 'not_banner.png'::text,
    theme_page character varying(15) DEFAULT 'white'::character varying,
    language character varying(5) DEFAULT 'ru'::character varying,
    aboutme text[] DEFAULT ARRAY['Название неизвестно'::text, 'Пока здесь ничего нет'::text]
);


ALTER TABLE public.settings OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 22095)
-- Name: settings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.settings_id_seq OWNER TO postgres;

--
-- TOC entry 3061 (class 0 OID 0)
-- Dependencies: 210
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


--
-- TOC entry 205 (class 1259 OID 22030)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying(100) NOT NULL,
    birthday timestamp without time zone,
    "position" character varying(120)[] DEFAULT NULL::character varying[],
    main_address integer[],
    net_status character varying(8) DEFAULT 'offline'::character varying,
    subscribers integer[],
    verification boolean DEFAULT false,
    first_name character varying(50) DEFAULT 'Пока здесь ничего нет'::character varying,
    last_name character varying(50) DEFAULT 'Пока здесь ничего нет'::character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 22010)
-- Name: users_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_data (
    id integer NOT NULL,
    login character varying(100) NOT NULL,
    password text NOT NULL,
    backup_keys character varying(64)[] DEFAULT NULL::character varying[],
    token character varying(64) DEFAULT NULL::character varying,
    user_custom_id character varying(100) DEFAULT NULL::character varying,
    email character varying(70) NOT NULL
);


ALTER TABLE public.users_data OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 22008)
-- Name: users_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_data_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_data_id_seq OWNER TO postgres;

--
-- TOC entry 3062 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_data_id_seq OWNED BY public.users_data.id;


--
-- TOC entry 204 (class 1259 OID 22028)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 3063 (class 0 OID 0)
-- Dependencies: 204
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 2873 (class 2604 OID 22077)
-- Name: connection id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection ALTER COLUMN id SET DEFAULT nextval('public.connection_id_seq'::regclass);


--
-- TOC entry 2872 (class 2604 OID 22058)
-- Name: identifiers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers ALTER COLUMN id SET DEFAULT nextval('public.identifiers_id_seq'::regclass);


--
-- TOC entry 2880 (class 2604 OID 22100)
-- Name: settings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


--
-- TOC entry 2868 (class 2604 OID 22033)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2862 (class 2604 OID 22013)
-- Name: users_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data ALTER COLUMN id SET DEFAULT nextval('public.users_data_id_seq'::regclass);


--
-- TOC entry 3051 (class 0 OID 22074)
-- Dependencies: 209
-- Data for Name: connection; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (1, 'gvgbRARdwWutHywpUFtyqpuNJWYMayPRYViVnJAOfaOFHvbfxeTQXgNjMoVvUXC1', 'hhhhhhhhh', 'insta', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет');
INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (2, 'ZpVMKfQvJYbqkoIsvqSFGOxVSaQSHkaJZuYtqiklhdFNFAqMUbjNxLNqFfFpQKw2', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет');


--
-- TOC entry 3049 (class 0 OID 22055)
-- Dependencies: 207
-- Data for Name: identifiers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.identifiers (id, user_id, connection_id, settings_id) VALUES (1, 1, 'gvgbRARdwWutHywpUFtyqpuNJWYMayPRYViVnJAOfaOFHvbfxeTQXgNjMoVvUXC1', 'UQaHZBfxCtRYrzJttzYKSozbkEDncOPPPDRMAoLYNcvKmYWbFgKxoDaugRketzs1');
INSERT INTO public.identifiers (id, user_id, connection_id, settings_id) VALUES (2, 2, 'ZpVMKfQvJYbqkoIsvqSFGOxVSaQSHkaJZuYtqiklhdFNFAqMUbjNxLNqFfFpQKw2', 'ZiprnYcRmIkAjWmPjGDbpJkbYkocaPPRKwsbGlcVwqpSzItMHJMjkRLUAzhMHGs2');


--
-- TOC entry 3053 (class 0 OID 22097)
-- Dependencies: 211
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (2, 'ZiprnYcRmIkAjWmPjGDbpJkbYkocaPPRKwsbGlcVwqpSzItMHJMjkRLUAzhMHGs2', 'not_logo.png', 'id2_4k4VF9yxtuoAuLdwEMPjzPgGiEQ.png', 'white', 'en', '{"Название неизвестно","Пока здесь ничего нет"}');
INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (1, 'UQaHZBfxCtRYrzJttzYKSozbkEDncOPPPDRMAoLYNcvKmYWbFgKxoDaugRketzs1', 'id1_3SiAElxBODIuSb_zaEHkQ547DnU.jpg', 'id1_gNbA7ClTX-YhcXWEXgAqrFDxdq4.jpg', 'white', 'en', '{"Название неизвестно","Пока здесь ничего нет"}');


--
-- TOC entry 3047 (class 0 OID 22030)
-- Dependencies: 205
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name) VALUES (2, 2, '234', NULL, NULL, NULL, 'online', NULL, false, 'Пока здесь ничего нет', 'Пока здесь ничего нет');
INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name) VALUES (1, 1, 'DMITRYDVP', NULL, NULL, NULL, 'online', NULL, false, 'Пока здесь ничего нет', 'Пока здесь ничего нет');


--
-- TOC entry 3045 (class 0 OID 22010)
-- Dependencies: 203
-- Data for Name: users_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (2, '234', '289dff07669d7a23de0ef88d2f7129e7', NULL, '6fbd1c71b87527cdc2bb33d468e7d790jzV9QTAw', NULL, '234');
INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (1, '123', '202cb962ac59075b964b07152d234b70', NULL, '87336d44942c98bf00a1a260576b4727V23zUC3p', NULL, '123');


--
-- TOC entry 3064 (class 0 OID 0)
-- Dependencies: 208
-- Name: connection_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.connection_id_seq', 2, true);


--
-- TOC entry 3065 (class 0 OID 0)
-- Dependencies: 206
-- Name: identifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.identifiers_id_seq', 2, true);


--
-- TOC entry 3066 (class 0 OID 0)
-- Dependencies: 210
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.settings_id_seq', 2, true);


--
-- TOC entry 3067 (class 0 OID 0)
-- Dependencies: 202
-- Name: users_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_data_id_seq', 3, true);


--
-- TOC entry 3068 (class 0 OID 0)
-- Dependencies: 204
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- TOC entry 2908 (class 2606 OID 22089)
-- Name: connection connection_connection_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_connection_id_key UNIQUE (connection_id);


--
-- TOC entry 2910 (class 2606 OID 22087)
-- Name: connection connection_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_pkey PRIMARY KEY (id);


--
-- TOC entry 2900 (class 2606 OID 22064)
-- Name: identifiers identifiers_connection_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_connection_id_key UNIQUE (connection_id);


--
-- TOC entry 2902 (class 2606 OID 22060)
-- Name: identifiers identifiers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_pkey PRIMARY KEY (id);


--
-- TOC entry 2904 (class 2606 OID 22066)
-- Name: identifiers identifiers_settings_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_settings_id_key UNIQUE (settings_id);


--
-- TOC entry 2906 (class 2606 OID 22062)
-- Name: identifiers identifiers_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_user_id_key UNIQUE (user_id);


--
-- TOC entry 2912 (class 2606 OID 22110)
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- TOC entry 2914 (class 2606 OID 22112)
-- Name: settings settings_settings_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_settings_id_key UNIQUE (settings_id);


--
-- TOC entry 2886 (class 2606 OID 22027)
-- Name: users_data users_data_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_email_key UNIQUE (email);


--
-- TOC entry 2888 (class 2606 OID 22023)
-- Name: users_data users_data_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_login_key UNIQUE (login);


--
-- TOC entry 2890 (class 2606 OID 22021)
-- Name: users_data users_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_pkey PRIMARY KEY (id);


--
-- TOC entry 2892 (class 2606 OID 22025)
-- Name: users_data users_data_user_custom_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_user_custom_id_key UNIQUE (user_custom_id);


--
-- TOC entry 2894 (class 2606 OID 22047)
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- TOC entry 2896 (class 2606 OID 22043)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2898 (class 2606 OID 22045)
-- Name: users users_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_id_key UNIQUE (user_id);


--
-- TOC entry 2917 (class 2606 OID 22090)
-- Name: connection connection_connection_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_connection_id_fkey FOREIGN KEY (connection_id) REFERENCES public.identifiers(connection_id);


--
-- TOC entry 2916 (class 2606 OID 22067)
-- Name: identifiers identifiers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 2915 (class 2606 OID 22048)
-- Name: users users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_data(id);


-- Completed on 2022-05-22 13:54:27 +05

--
-- PostgreSQL database dump complete
--

