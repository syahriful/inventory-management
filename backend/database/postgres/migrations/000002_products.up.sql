CREATE TABLE IF NOT EXISTS products
(
   id                       SERIAL,
   code                     VARCHAR(100) NOT NULL UNIQUE,
   name                     VARCHAR(100) NOT NULL,
   unit_mass_acronym        VARCHAR(20)  NOT NULL,
   unit_mass_description    VARCHAR(50)  NOT NULL,
   created_at               TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at               TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
)