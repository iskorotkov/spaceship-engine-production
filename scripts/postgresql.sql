CREATE TABLE public.clients
(
    id   BIGINT NOT NULL,
    name TEXT
);

CREATE TABLE public.components
(
    id   BIGINT NOT NULL,
    name TEXT
);

CREATE TABLE public.engines
(
    id    BIGINT NOT NULL,
    name  TEXT,
    power NUMERIC
);

CREATE TABLE public.orders
(
    id           BIGINT NOT NULL,
    amount       BIGINT,
    created_at   TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    client_id    BIGINT,
    engine_id    BIGINT
);

CREATE TABLE public.providers
(
    id   BIGINT NOT NULL,
    name TEXT
);

CREATE TABLE public.requirements
(
    id           BIGINT NOT NULL,
    engine_id    BIGINT,
    component_id BIGINT
);

CREATE TABLE public.supplies
(
    id           BIGINT NOT NULL,
    component_id BIGINT,
    provider_id  BIGINT
);
