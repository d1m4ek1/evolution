--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

-- Started on 2022-10-08 23:42:22 +05

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
-- TOC entry 209 (class 1259 OID 16499)
-- Name: albums; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.albums (
    id integer NOT NULL,
    album_id character varying NOT NULL,
    creator text NOT NULL,
    cover_link text DEFAULT '/beats/cover_album/no_cover.png'::text,
    folowers_count integer DEFAULT 0,
    private boolean DEFAULT false
);


ALTER TABLE public.albums OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 16507)
-- Name: albums_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.albums_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.albums_id_seq OWNER TO postgres;

--
-- TOC entry 3479 (class 0 OID 0)
-- Dependencies: 210
-- Name: albums_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.albums_id_seq OWNED BY public.albums.id;


--
-- TOC entry 211 (class 1259 OID 16508)
-- Name: tracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tracks (
    id integer NOT NULL,
    album_id character varying(128) NOT NULL,
    track_id character varying(64) NOT NULL,
    name text NOT NULL,
    sound_link text NOT NULL,
    price double precision NOT NULL,
    folowers_count integer DEFAULT 0,
    release_date date NOT NULL,
    track_type character varying(20) NOT NULL,
    "position" integer
);


ALTER TABLE public.tracks OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 16514)
-- Name: beats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.beats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.beats_id_seq OWNER TO postgres;

--
-- TOC entry 3480 (class 0 OID 0)
-- Dependencies: 212
-- Name: beats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.beats_id_seq OWNED BY public.tracks.id;


--
-- TOC entry 213 (class 1259 OID 16515)
-- Name: chats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.chats (
    id integer NOT NULL,
    user_id_one integer NOT NULL,
    user_id_two integer NOT NULL,
    chat_id integer NOT NULL,
    messages json[],
    new_messages json[],
    user_data_one json NOT NULL,
    user_data_two json NOT NULL
);


ALTER TABLE public.chats OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16520)
-- Name: chats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.chats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.chats_id_seq OWNER TO postgres;

--
-- TOC entry 3481 (class 0 OID 0)
-- Dependencies: 214
-- Name: chats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.chats_id_seq OWNED BY public.chats.id;


--
-- TOC entry 215 (class 1259 OID 16521)
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
-- TOC entry 216 (class 1259 OID 16531)
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
-- TOC entry 3482 (class 0 OID 0)
-- Dependencies: 216
-- Name: connection_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.connection_id_seq OWNED BY public.connection.id;


--
-- TOC entry 217 (class 1259 OID 16532)
-- Name: identifiers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.identifiers (
    id integer NOT NULL,
    user_id integer NOT NULL,
    connection_id character varying(64) NOT NULL,
    settings_id character varying(64) NOT NULL,
    album_id character varying(128)
);


ALTER TABLE public.identifiers OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16535)
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
-- TOC entry 3483 (class 0 OID 0)
-- Dependencies: 218
-- Name: identifiers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.identifiers_id_seq OWNED BY public.identifiers.id;


--
-- TOC entry 219 (class 1259 OID 16536)
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
-- TOC entry 220 (class 1259 OID 16546)
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
-- TOC entry 3484 (class 0 OID 0)
-- Dependencies: 220
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.settings_id_seq OWNED BY public.settings.id;


--
-- TOC entry 221 (class 1259 OID 16547)
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
    last_name character varying(50) DEFAULT 'Пока здесь ничего нет'::character varying,
    subscriptions integer[]
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 16557)
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
-- TOC entry 223 (class 1259 OID 16565)
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
-- TOC entry 3485 (class 0 OID 0)
-- Dependencies: 223
-- Name: users_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_data_id_seq OWNED BY public.users_data.id;


--
-- TOC entry 224 (class 1259 OID 16566)
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
-- TOC entry 3486 (class 0 OID 0)
-- Dependencies: 224
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3247 (class 2604 OID 16567)
-- Name: albums id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums ALTER COLUMN id SET DEFAULT nextval('public.albums_id_seq'::regclass);


--
-- TOC entry 3250 (class 2604 OID 16568)
-- Name: chats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chats ALTER COLUMN id SET DEFAULT nextval('public.chats_id_seq'::regclass);


--
-- TOC entry 3256 (class 2604 OID 16569)
-- Name: connection id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection ALTER COLUMN id SET DEFAULT nextval('public.connection_id_seq'::regclass);


--
-- TOC entry 3257 (class 2604 OID 16570)
-- Name: identifiers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers ALTER COLUMN id SET DEFAULT nextval('public.identifiers_id_seq'::regclass);


--
-- TOC entry 3263 (class 2604 OID 16571)
-- Name: settings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings ALTER COLUMN id SET DEFAULT nextval('public.settings_id_seq'::regclass);


--
-- TOC entry 3249 (class 2604 OID 16572)
-- Name: tracks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks ALTER COLUMN id SET DEFAULT nextval('public.beats_id_seq'::regclass);


--
-- TOC entry 3269 (class 2604 OID 16573)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3273 (class 2604 OID 16574)
-- Name: users_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data ALTER COLUMN id SET DEFAULT nextval('public.users_data_id_seq'::regclass);


--
-- TOC entry 3458 (class 0 OID 16499)
-- Dependencies: 209
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3462 (class 0 OID 16515)
-- Dependencies: 213
-- Data for Name: chats; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.chats (id, user_id_one, user_id_two, chat_id, messages, new_messages, user_data_one, user_data_two) VALUES (1, 1, 5, 15, '{"{\"date\":\"01.08.2022\",\"message\":\"Лт\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"Лл\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"Ло\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"Ир\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"Лл\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"Ио\",\"sender_id\":5}","{\"date\":\"01.08.2022\",\"message\":\"6867\",\"sender_id\":1}"}', NULL, '{"userId":1,"logo":"id1_3SiAElxBODIuSb_zaEHkQ547DnU.jpg","banner":"id1_4NhCGhH-iLNQ9esceKXA4ilwASY.jpg","name":"dmitrydvp"}', '{"userId":5,"logo":"not_logo.png","banner":"not_banner.png","name":"456"}');


--
-- TOC entry 3464 (class 0 OID 16521)
-- Dependencies: 215
-- Data for Name: connection; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (2, 'ZpVMKfQvJYbqkoIsvqSFGOxVSaQSHkaJZuYtqiklhdFNFAqMUbjNxLNqFfFpQKw2', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет');
INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (1, 'gvgbRARdwWutHywpUFtyqpuNJWYMayPRYViVnJAOfaOFHvbfxeTQXgNjMoVvUXC1', '567567', 'insta', 'есть', 'Пока здесь ничего нет', 'Пока здесь ничего нет');
INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (3, 'rpeoMpUbNUkApVbdkRZKFdUNjfLAyiyOuZmzZnIirLictZyyiLFOFKtmkecsJXa4', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет');
INSERT INTO public.connection (id, connection_id, telegram, instagram, facebook, vk, tiktok) VALUES (4, 'stPaGmKkvGyrBeYQuzgknRuJxWmaTAauZMHphYtZKcdyqwmgRQcOnBnRDXpCnyL5', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет', 'Пока здесь ничего нет');


--
-- TOC entry 3466 (class 0 OID 16532)
-- Dependencies: 217
-- Data for Name: identifiers; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.identifiers (id, user_id, connection_id, settings_id, album_id) VALUES (1, 1, 'gvgbRARdwWutHywpUFtyqpuNJWYMayPRYViVnJAOfaOFHvbfxeTQXgNjMoVvUXC1', 'UQaHZBfxCtRYrzJttzYKSozbkEDncOPPPDRMAoLYNcvKmYWbFgKxoDaugRketzs1', NULL);
INSERT INTO public.identifiers (id, user_id, connection_id, settings_id, album_id) VALUES (2, 2, 'ZpVMKfQvJYbqkoIsvqSFGOxVSaQSHkaJZuYtqiklhdFNFAqMUbjNxLNqFfFpQKw2', 'ZiprnYcRmIkAjWmPjGDbpJkbYkocaPPRKwsbGlcVwqpSzItMHJMjkRLUAzhMHGs2', NULL);
INSERT INTO public.identifiers (id, user_id, connection_id, settings_id, album_id) VALUES (3, 4, 'rpeoMpUbNUkApVbdkRZKFdUNjfLAyiyOuZmzZnIirLictZyyiLFOFKtmkecsJXa4', 'WGccjntKSsLQjuaACyOpEOiCUYKHxagiYszHnkUeFHNVqLQxUDzSXZOiavTetAd4', NULL);
INSERT INTO public.identifiers (id, user_id, connection_id, settings_id, album_id) VALUES (4, 5, 'stPaGmKkvGyrBeYQuzgknRuJxWmaTAauZMHphYtZKcdyqwmgRQcOnBnRDXpCnyL5', 'wyjjkZAuhMYskqXPwIAKlySarIEwqIHXycaBGJxfRqhdLrUNHpKsJJuKLPITdGs5', NULL);


--
-- TOC entry 3468 (class 0 OID 16536)
-- Dependencies: 219
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (2, 'ZiprnYcRmIkAjWmPjGDbpJkbYkocaPPRKwsbGlcVwqpSzItMHJMjkRLUAzhMHGs2', 'not_logo.png', 'id2_4k4VF9yxtuoAuLdwEMPjzPgGiEQ.png', 'white', 'en', '{t,3214234}');
INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (3, 'WGccjntKSsLQjuaACyOpEOiCUYKHxagiYszHnkUeFHNVqLQxUDzSXZOiavTetAd4', 'not_logo.png', 'not_banner.png', 'white', 'ru', '{"Название неизвестно","Пока здесь ничего нет"}');
INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (4, 'wyjjkZAuhMYskqXPwIAKlySarIEwqIHXycaBGJxfRqhdLrUNHpKsJJuKLPITdGs5', 'not_logo.png', 'not_banner.png', 'white', 'ru', '{"Название неизвестно","Пока здесь ничего нет"}');
INSERT INTO public.settings (id, settings_id, logo, banner, theme_page, language, aboutme) VALUES (1, 'UQaHZBfxCtRYrzJttzYKSozbkEDncOPPPDRMAoLYNcvKmYWbFgKxoDaugRketzs1', 'id1_3SiAElxBODIuSb_zaEHkQ547DnU.jpg', 'id1_4NhCGhH-iLNQ9esceKXA4ilwASY.jpg', 'dark', 'ru', '{123,567567567567567567567567567567}');


--
-- TOC entry 3460 (class 0 OID 16508)
-- Dependencies: 211
-- Data for Name: tracks; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3470 (class 0 OID 16547)
-- Dependencies: 221
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name, subscriptions) VALUES (2, 2, '234', NULL, NULL, NULL, 'offline', NULL, false, 'Пока здесь ничего нет', 'Пока здесь ничего нет', NULL);
INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name, subscriptions) VALUES (4, 4, '1234', NULL, NULL, NULL, 'offline', NULL, false, 'Пока здесь ничего нет', 'Пока здесь ничего нет', NULL);
INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name, subscriptions) VALUES (1, 1, 'dmitrydvp', NULL, NULL, NULL, 'online', '{5}', false, 'Пока здесь ничего нет', 'Пока здесь ничего нет', '{}');
INSERT INTO public.users (id, user_id, name, birthday, "position", main_address, net_status, subscribers, verification, first_name, last_name, subscriptions) VALUES (5, 5, '456', NULL, NULL, NULL, 'offline', '{}', false, 'Пока здесь ничего нет', 'Пока здесь ничего нет', '{1}');


--
-- TOC entry 3471 (class 0 OID 16557)
-- Dependencies: 222
-- Data for Name: users_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (5, '456', '250cf8b51c773f3f8dc8b4be867a9a02', NULL, 'bd6652e42777b08dc8d969f0990648f70bLr4x8V', NULL, '456');
INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (4, '1234', '81dc9bdb52d04dc20036dbd8313ed055', NULL, 'b50ccb88c4be7cc6bdbd8d3a046cd372JazTveIN', NULL, '1234');
INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (2, '234', '289dff07669d7a23de0ef88d2f7129e7', NULL, '60d6c9e170030d82fbdf5c8b13cee387zepmkVEz', NULL, '234');
INSERT INTO public.users_data (id, login, password, backup_keys, token, user_custom_id, email) VALUES (1, '123', '827ccb0eea8a706c4c34a16891f84e7b', '{202cb962ac59075b964b07152d234b70,289dff07669d7a23de0ef88d2f7129e7,d81f9c1be2e08964bf9f24b15f0e4900,250cf8b51c773f3f8dc8b4be867a9a02}', 'qlAstYpg106UzXbE1pqxfhxyghcTHUicAJFpFKsZVFh5REpB8WrhFQlC32kghA3Z', NULL, 'dmitrydvp@outlook.com');


--
-- TOC entry 3487 (class 0 OID 0)
-- Dependencies: 210
-- Name: albums_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.albums_id_seq', 1, false);


--
-- TOC entry 3488 (class 0 OID 0)
-- Dependencies: 212
-- Name: beats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.beats_id_seq', 1, false);


--
-- TOC entry 3489 (class 0 OID 0)
-- Dependencies: 214
-- Name: chats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.chats_id_seq', 4, true);


--
-- TOC entry 3490 (class 0 OID 0)
-- Dependencies: 216
-- Name: connection_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.connection_id_seq', 4, true);


--
-- TOC entry 3491 (class 0 OID 0)
-- Dependencies: 218
-- Name: identifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.identifiers_id_seq', 4, true);


--
-- TOC entry 3492 (class 0 OID 0)
-- Dependencies: 220
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.settings_id_seq', 4, true);


--
-- TOC entry 3493 (class 0 OID 0)
-- Dependencies: 223
-- Name: users_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_data_id_seq', 5, true);


--
-- TOC entry 3494 (class 0 OID 0)
-- Dependencies: 224
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- TOC entry 3275 (class 2606 OID 16576)
-- Name: albums albums_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_pkey PRIMARY KEY (id);


--
-- TOC entry 3277 (class 2606 OID 16578)
-- Name: tracks beats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT beats_pkey PRIMARY KEY (id);


--
-- TOC entry 3279 (class 2606 OID 16580)
-- Name: chats chats_chat_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chats
    ADD CONSTRAINT chats_chat_id_key UNIQUE (chat_id);


--
-- TOC entry 3281 (class 2606 OID 16582)
-- Name: chats chats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chats
    ADD CONSTRAINT chats_pkey PRIMARY KEY (id);


--
-- TOC entry 3283 (class 2606 OID 16584)
-- Name: chats chats_user_id_one_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chats
    ADD CONSTRAINT chats_user_id_one_key UNIQUE (user_id_one);


--
-- TOC entry 3285 (class 2606 OID 16586)
-- Name: chats chats_user_id_two_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.chats
    ADD CONSTRAINT chats_user_id_two_key UNIQUE (user_id_two);


--
-- TOC entry 3287 (class 2606 OID 16588)
-- Name: connection connection_connection_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_connection_id_key UNIQUE (connection_id);


--
-- TOC entry 3289 (class 2606 OID 16590)
-- Name: connection connection_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_pkey PRIMARY KEY (id);


--
-- TOC entry 3291 (class 2606 OID 16592)
-- Name: identifiers identifiers_connection_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_connection_id_key UNIQUE (connection_id);


--
-- TOC entry 3293 (class 2606 OID 16594)
-- Name: identifiers identifiers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_pkey PRIMARY KEY (id);


--
-- TOC entry 3295 (class 2606 OID 16596)
-- Name: identifiers identifiers_settings_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_settings_id_key UNIQUE (settings_id);


--
-- TOC entry 3297 (class 2606 OID 16598)
-- Name: identifiers identifiers_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_user_id_key UNIQUE (user_id);


--
-- TOC entry 3299 (class 2606 OID 16600)
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- TOC entry 3301 (class 2606 OID 16602)
-- Name: settings settings_settings_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.settings
    ADD CONSTRAINT settings_settings_id_key UNIQUE (settings_id);


--
-- TOC entry 3309 (class 2606 OID 16604)
-- Name: users_data users_data_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_email_key UNIQUE (email);


--
-- TOC entry 3311 (class 2606 OID 16606)
-- Name: users_data users_data_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_login_key UNIQUE (login);


--
-- TOC entry 3313 (class 2606 OID 16608)
-- Name: users_data users_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_pkey PRIMARY KEY (id);


--
-- TOC entry 3315 (class 2606 OID 16610)
-- Name: users_data users_data_user_custom_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_data
    ADD CONSTRAINT users_data_user_custom_id_key UNIQUE (user_custom_id);


--
-- TOC entry 3303 (class 2606 OID 16612)
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- TOC entry 3305 (class 2606 OID 16614)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3307 (class 2606 OID 16616)
-- Name: users users_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_id_key UNIQUE (user_id);


--
-- TOC entry 3316 (class 2606 OID 16617)
-- Name: connection connection_connection_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.connection
    ADD CONSTRAINT connection_connection_id_fkey FOREIGN KEY (connection_id) REFERENCES public.identifiers(connection_id);


--
-- TOC entry 3317 (class 2606 OID 16622)
-- Name: identifiers identifiers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.identifiers
    ADD CONSTRAINT identifiers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 3318 (class 2606 OID 16627)
-- Name: users users_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users_data(id);


-- Completed on 2022-10-08 23:42:22 +05

--
-- PostgreSQL database dump complete
--

