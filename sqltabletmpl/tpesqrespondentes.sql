CREATE TABLE tpesqrespondentes (
  
  prp_id serial NOT NULL,

  prp_uid VARCHAR(256) NOT NULL,

  prp_cod integer NOT NULL,

  prp_cli_cod integer NOT NULL,

  prp_pes_cod integer NOT NULL,
  
  prp_email VARCHAR(500) NOT NULL,

  prp_telefone  VARCHAR(60) NOT NULL,
  
  prp_MessengerID  VARCHAR(1000) NOT NULL,
  
  ppr_dtcadastro  date NOT NULL,
  
  ppr_dtaltera  date NOT NULL,
  
  prp_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (prp_id)

  );

ALTER TABLE ONLY tpesqrespondentes ADD CONSTRAINT tpesqrespondentes_prp_id UNIQUE (prp_uid);