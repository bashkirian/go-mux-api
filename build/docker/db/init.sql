--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2022-11-12 19:36:48

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
-- TOC entry 209 (class 1259 OID 57345)
-- Name: balance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.balance (
    user_id integer NOT NULL,
    ruble_balance numeric NOT NULL,
    CONSTRAINT check_positive CHECK ((ruble_balance >= (0)::numeric))
);


ALTER TABLE public.balance OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 57364)
-- Name: reservations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reservations (
    reservation_id integer NOT NULL,
    user_id integer NOT NULL,
    service_id integer NOT NULL,
    cost numeric NOT NULL,
    reservation_time timestamp without time zone
);


ALTER TABLE public.reservations OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 57359)
-- Name: services; Type: TABLE; Schema: public; Owner: postgres
--

-- CREATE TABLE public.services (
--     service_id integer NOT NULL,
--     service_name character varying(255) NOT NULL
-- );


-- ALTER TABLE public.services OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 57352)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    transaction_id integer NOT NULL,
    date timestamp without time zone NOT NULL,
    amount numeric NOT NULL,
    user_id integer NOT NULL,
    service_id integer NOT NULL
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 3177 (class 2606 OID 57351)
-- Name: balance balance_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.balance
    ADD CONSTRAINT balance_pk PRIMARY KEY (user_id);


--
-- TOC entry 3183 (class 2606 OID 57370)
-- Name: reservations reservations_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_pk PRIMARY KEY (reservation_id);


--
-- TOC entry 3181 (class 2606 OID 57363)
-- Name: services services_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

--ALTER TABLE ONLY public.services
--    ADD CONSTRAINT services_pk PRIMARY KEY (service_id);


--
-- TOC entry 3179 (class 2606 OID 57358)
-- Name: transactions transactions_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pk PRIMARY KEY (transaction_id);


--
-- TOC entry 3186 (class 2606 OID 57381)
-- Name: reservations reservations_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reservations
    ADD CONSTRAINT reservations_fk0 FOREIGN KEY (user_id) REFERENCES public.balance(user_id);


--
-- TOC entry 3187 (class 2606 OID 57386)
-- Name: reservations reservations_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

--ALTER TABLE ONLY public.reservations
    --ADD CONSTRAINT reservations_fk1 FOREIGN KEY (service_id) REFERENCES public.services(service_id);


--
-- TOC entry 3184 (class 2606 OID 57371)
-- Name: transactions transactions_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_fk0 FOREIGN KEY (user_id) REFERENCES public.balance(user_id);


--
-- TOC entry 3185 (class 2606 OID 57376)
-- Name: transactions transactions_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

--ALTER TABLE ONLY public.transactions
   -- ADD CONSTRAINT transactions_fk1 FOREIGN KEY (service_id) REFERENCES public.services(service_id);


-- Completed on 2022-11-12 19:36:48

--
-- PostgreSQL database dump complete
--

