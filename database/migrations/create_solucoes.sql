CREATE TABLE solucoes (
    id SERIAL PRIMARY KEY,
    demanda_id INTEGER,
    nome VARCHAR(255),
    sigla VARCHAR(50),
    versao VARCHAR(50),
    link VARCHAR(255),
    andamento VARCHAR(50),
    repositorio VARCHAR(255),
    criticidade VARCHAR(50),
    data_status VARCHAR(50)
);