CREATE TABLE tpesqrespostas (
  
  prs_id serial NOT NULL,

  prs_uid VARCHAR(256) NOT NULL,

  prs_cod integer NOT NULL,

  prs_prp_cod integer NOT NULL,

  prs_res_cod integer NOT NULL,
  
  prs_texto text NOT NULL,

  prs_dtcadastro  date NOT NULL,
  
  prs_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (prs_id)

  );

ALTER TABLE ONLY tpesqrespostas ADD CONSTRAINT tpesqrespostas_prs_id UNIQUE (prs_uid);