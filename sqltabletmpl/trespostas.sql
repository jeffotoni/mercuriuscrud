CREATE TABLE trespostas (
  
  rsp_id serial NOT NULL,

  rsp_uid VARCHAR(256) NOT NULL,

  rsp_cod integer NOT NULL,

  rsp_per_cod integer NOT NULL,

  rsp_titulo VARCHAR(256) NOT NULL,
  
  rsp_correta integer NOT NULL,
  
  rsp_dtcadastro date NOT NULL,

  rsp_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (rsp_id)

  );

ALTER TABLE ONLY trespostas ADD CONSTRAINT trespostas_rsp_uid UNIQUE (rsp_uid);