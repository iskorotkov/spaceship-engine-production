CREATE TABLE clients
(
    id   BIGSERIAL
        CONSTRAINT clients_pkey
            PRIMARY KEY,
    name TEXT
);

CREATE TABLE components
(
    id   BIGSERIAL
        CONSTRAINT components_pkey
            PRIMARY KEY,
    name TEXT
);

CREATE TABLE engines
(
    id    BIGSERIAL
        CONSTRAINT engines_pkey
            PRIMARY KEY,
    name  TEXT,
    power NUMERIC
);

CREATE TABLE orders
(
    id           BIGSERIAL
        CONSTRAINT orders_pkey
            PRIMARY KEY,
    amount       BIGINT,
    created_at   TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    client_id    BIGINT
        CONSTRAINT fk_orders_client
            REFERENCES clients,
    engine_id    BIGINT
        CONSTRAINT fk_orders_engine
            REFERENCES engines
);

CREATE TABLE providers
(
    id   BIGSERIAL
        CONSTRAINT providers_pkey
            PRIMARY KEY,
    name TEXT
);

CREATE TABLE requirements
(
    id           BIGSERIAL
        CONSTRAINT requirements_pkey
            PRIMARY KEY,
    engine_id    BIGINT
        CONSTRAINT fk_requirements_engine
            REFERENCES engines,
    component_id BIGINT
        CONSTRAINT fk_requirements_component
            REFERENCES components
);

CREATE TABLE supplies
(
    id           BIGSERIAL
        CONSTRAINT supplies_pkey
            PRIMARY KEY,
    component_id BIGINT
        CONSTRAINT fk_supplies_component
            REFERENCES components,
    provider_id  BIGINT
        CONSTRAINT fk_supplies_provider
            REFERENCES providers
);
