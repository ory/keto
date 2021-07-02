-- we can not fizz this migration because there is no check constraint support in fizz
-- and it gets cumbersome with sqlite; having a working SQL version is actually way easier
CREATE TABLE keto_networks
(
    limiter    INTEGER   NOT NULL DEFAULT 0 UNIQUE,
    network_id UUID      NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,

    PRIMARY KEY (network_id),

    CONSTRAINT chk_keto_networks_limit CHECK (limiter = 0)
);
