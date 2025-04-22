--
-- PostgreSQL database dump
--

-- Dumped from database version 17.4
-- Dumped by pg_dump version 17.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: ads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ads (
    id integer NOT NULL,
    title text NOT NULL,
    description text,
    video_url text NOT NULL,
    target_url text NOT NULL,
    thumbnail_url text,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone DEFAULT now(),
    duration integer NOT NULL
);


ALTER TABLE public.ads OWNER TO postgres;

--
-- Name: ads_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ads_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ads_id_seq OWNER TO postgres;

--
-- Name: ads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ads_id_seq OWNED BY public.ads.id;


--
-- Name: clicks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clicks (
    id integer NOT NULL,
    ad_id integer NOT NULL,
    "timestamp" timestamp with time zone NOT NULL,
    ip character varying(45) NOT NULL,
    playback_position double precision NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.clicks OWNER TO postgres;

--
-- Name: clicks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.clicks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.clicks_id_seq OWNER TO postgres;

--
-- Name: clicks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.clicks_id_seq OWNED BY public.clicks.id;


--
-- Name: impressions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.impressions (
    id integer NOT NULL,
    ad_id integer NOT NULL,
    user_id integer NOT NULL,
    "timestamp" timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.impressions OWNER TO postgres;

--
-- Name: impressions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.impressions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.impressions_id_seq OWNER TO postgres;

--
-- Name: impressions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.impressions_id_seq OWNED BY public.impressions.id;


--
-- Name: ads id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ads ALTER COLUMN id SET DEFAULT nextval('public.ads_id_seq'::regclass);


--
-- Name: clicks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clicks ALTER COLUMN id SET DEFAULT nextval('public.clicks_id_seq'::regclass);


--
-- Name: impressions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.impressions ALTER COLUMN id SET DEFAULT nextval('public.impressions_id_seq'::regclass);


--
-- Data for Name: ads; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ads (id, title, description, video_url, target_url, thumbnail_url, created_at, updated_at, duration) FROM stdin;
1	Summer Sale	Get amazing discounts on our summer collection!	https://example.com/video1.mp4	https://example.com/offer1	https://example.com/thumbnail1.jpg	2025-04-17 23:34:40.21612	2025-04-17 23:34:40.21612	30
2	New Product Launch	Introducing our latest product with cutting-edge technology.	https://example.com/video2.mp4	https://example.com/launch	https://example.com/thumbnail2.jpg	2025-04-17 23:34:40.21612	2025-04-17 23:34:40.21612	45
3	Holiday Special	Exclusive deals for the holiday season.	https://example.com/video3.mp4	https://example.com/holiday-offer	https://example.com/thumbnail3.jpg	2025-04-17 23:34:40.21612	2025-04-17 23:34:40.21612	60
4	Winter Collection	Stay warm and stylish with our new winter collection.	https://example.com/video4.mp4	https://example.com/winter-collection	https://example.com/thumbnail4.jpg	2025-04-17 23:34:40.21612	2025-04-17 23:34:40.21612	30
5	Limited Time Offer	Hurry, dont miss out on this limited time offer!	https://example.com/video5.mp4	https://example.com/limited-offer	https://example.com/thumbnail5.jpg	2025-04-17 23:34:40.21612	2025-04-17 23:34:40.21612	15
\.


--
-- Data for Name: clicks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.clicks (id, ad_id, "timestamp", ip, playback_position, user_id) FROM stdin;
20	2	2025-04-18 20:21:23+05:30	192.168.1.1	15.2	123434
21	3	2025-04-18 20:21:23+05:30	192.168.1.1	15.2	123434
22	3	2025-04-18 20:21:23+05:30	192.168.1.1	15.2	123437
23	5	2025-04-18 20:21:23+05:30	192.168.10.1	15.2	123490
24	4	2025-04-18 20:21:23+05:30	192.168.10.1	15.2	123490
25	1	2025-04-18 20:21:23+05:30	192.168.10.1	15.2	123490
26	1	2025-04-18 20:21:23+05:30	192.168.10.1	15.2	123491
28	1	2025-04-18 20:22:23+05:30	192.168.10.1	15.2	123491
30	1	2025-04-18 20:22:23+05:30	192.168.19.1	15.2	12491
\.


--
-- Data for Name: impressions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.impressions (id, ad_id, user_id, "timestamp") FROM stdin;
1	2	123434	2025-04-18 20:21:23+05:30
2	3	123434	2025-04-18 20:21:23+05:30
3	3	123437	2025-04-18 20:21:23+05:30
4	1	123456	2025-04-18 10:00:00+05:30
5	1	123457	2025-04-18 10:01:30+05:30
6	1	123458	2025-04-18 10:05:45+05:30
7	1	123459	2025-04-18 10:07:30+05:30
8	1	123460	2025-04-18 10:10:10+05:30
9	2	123461	2025-04-18 10:12:30+05:30
10	2	123462	2025-04-18 10:15:00+05:30
11	2	123463	2025-04-18 10:18:15+05:30
12	2	123464	2025-04-18 10:20:00+05:30
13	2	123465	2025-04-18 10:25:00+05:30
14	3	123466	2025-04-18 11:00:00+05:30
15	3	123467	2025-04-18 11:02:30+05:30
16	3	123468	2025-04-18 11:05:00+05:30
17	3	123469	2025-04-18 11:07:30+05:30
18	3	123470	2025-04-18 11:10:00+05:30
19	4	123471	2025-04-18 11:12:30+05:30
20	4	123472	2025-04-18 11:15:00+05:30
21	4	123473	2025-04-18 11:20:00+05:30
22	4	123474	2025-04-18 11:22:30+05:30
23	4	123475	2025-04-18 11:25:00+05:30
24	5	123476	2025-04-18 12:00:00+05:30
25	5	123477	2025-04-18 12:02:30+05:30
26	5	123478	2025-04-18 12:05:00+05:30
27	5	123479	2025-04-18 12:10:00+05:30
28	5	123480	2025-04-18 12:15:00+05:30
\.


--
-- Name: ads_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ads_id_seq', 5, true);


--
-- Name: clicks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.clicks_id_seq', 30, true);


--
-- Name: impressions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.impressions_id_seq', 28, true);


--
-- Name: ads ads_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ads
    ADD CONSTRAINT ads_pkey PRIMARY KEY (id);


--
-- Name: clicks clicks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clicks
    ADD CONSTRAINT clicks_pkey PRIMARY KEY (id);


--
-- Name: impressions impressions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.impressions
    ADD CONSTRAINT impressions_pkey PRIMARY KEY (id);


--
-- Name: idx_user_ad_clicks; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_user_ad_clicks ON public.clicks USING btree (user_id, ad_id, "timestamp");


--
-- Name: clicks clicks_ad_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clicks
    ADD CONSTRAINT clicks_ad_id_fkey FOREIGN KEY (ad_id) REFERENCES public.ads(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

