CREATE TABLE tstatusperguntas (
  
  spe_id serial NOT NULL,

  spe_uid VARCHAR(256) NOT NULL,

  spe_cod  integer NOT NULL,

  spe_nome VARCHAR(256) NOT NULL,

  spe_dtcadastro date NOT NULL,
  
  spe_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (spe_id)

  );

ALTER TABLE ONLY tstatusperguntas ADD CONSTRAINT tstatusperguntas_spe_uid UNIQUE (spe_uid);