CREATE TABLE tpesqpaginas (
  
  ppq_id serial NOT NULL,

  ppq_uid VARCHAR(256) NOT NULL,

  ppq_cod integer NOT NULL,

  ppq_pes_cod integer NOT NULL,

  ppq_nropasso integer NOT NULL,
  
  ppq_descricao TEXT NOT NULL,
  
  ppq_datetime VARCHAR(256) NOT NULL,

  PRIMARY KEY (ppq_id)

  );

ALTER TABLE ONLY tpesqpaginas ADD CONSTRAINT tpesqpaginas_ppq_id UNIQUE (ppq_id);