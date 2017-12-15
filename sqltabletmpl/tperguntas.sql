CREATE TABLE tperguntas (
  
  per_id serial NOT NULL,

  per_uid VARCHAR(256) NOT NULL,

  per_cod integer NOT NULL,

  per_titulo VARCHAR(256) NOT NULL,

  per_desc text NOT NULL,
  
  per_cli_cod integer NOT NULL,
  
  per_spe_cod date NOT NULL,

  per_keywords VARCHAR(256) NOT NULL,

  per_tipo VARCHAR(100) NOT NULL,

  per_dtcadastro date NOT NULL,

  per_dtaltera date NOT NULL,

  per_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (per_id)

  );

ALTER TABLE ONLY tperguntas ADD CONSTRAINT tpergunta_sper_id UNIQUE (per_uid);