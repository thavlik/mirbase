DROP TABLE IF EXISTS confidence;
CREATE TABLE confidence (
  mirna_id varchar(20) NOT NULL DEFAULT '',
  auto_mirna int NOT NULL DEFAULT '0',
  exp_count int NOT NULL DEFAULT '0',
  _5p_count double NOT NULL DEFAULT '0',
  _5p_raw_count float NOT NULL DEFAULT '0',
  _3p_count float NOT NULL DEFAULT '0',
  _3p_raw_count float NOT NULL DEFAULT '0',
  _5p_consistent float NOT NULL DEFAULT '0',
  _5p_mature_consistent decimal(4,0) NOT NULL DEFAULT '0',
  _3p_consistent float NOT NULL DEFAULT '0',
  _3p_mature_consistent decimal(4,0) NOT NULL DEFAULT '0',
  _5p_overhang int DEFAULT NULL,
  _3p_overhang int DEFAULT NULL,
  energy_precursor float DEFAULT '0',
  energy_by_length float NOT NULL,
  paired_hairpin float NOT NULL DEFAULT '0',
  mirdeep_score double NOT NULL DEFAULT '0'
);

DROP TABLE IF EXISTS confidence_score;
CREATE TABLE confidence_score (
  auto_mirna int NOT NULL,
  confidence int NOT NULL DEFAULT '0'
);

DROP TABLE IF EXISTS dead_mirna;
CREATE TABLE dead_mirna (
  mirna_acc varchar(9) NOT NULL DEFAULT '',
  mirna_id varchar(40) NOT NULL DEFAULT '',
  previous_id varchar(100) DEFAULT NULL,
  forward_to varchar(20) DEFAULT NULL,
  comment mediumtext
);

DROP TABLE IF EXISTS literature_references;
CREATE TABLE literature_references (
  auto_lit INTEGER PRIMARY KEY,
  medline int DEFAULT NULL,
  title tinytext,
  author tinytext,
  journal tinytext
);

DROP TABLE IF EXISTS mature_database_links;
CREATE TABLE mature_database_links (
  auto_mature int NOT NULL DEFAULT '0',
  auto_db int NOT NULL DEFAULT '0',
  link tinytext NOT NULL,
  display_name tinytext NOT NULL
);

DROP TABLE IF EXISTS mature_database_url;
CREATE TABLE mature_database_url (
  auto_db INTEGER PRIMARY KEY,
  display_name tinytext NOT NULL,
  url tinytext NOT NULL,
  type smallint DEFAULT NULL
);

DROP TABLE IF EXISTS mirna;
CREATE TABLE mirna (
  auto_mirna int PRIMARY KEY,
  mirna_acc varchar(9) NOT NULL DEFAULT '',
  mirna_id varchar(40) NOT NULL DEFAULT '',
  previous_mirna_id text NOT NULL,
  description varchar(100) DEFAULT NULL,
  sequence blob,
  comment longtext,
  auto_species int NOT NULL DEFAULT '0',
  dead_flag tinyint NOT NULL
);

CREATE VIRTUAL TABLE mirna_search USING fts5(mirna_acc, mirna_id, description, sequence, comment);

DROP TABLE IF EXISTS mirna_2_prefam;
CREATE TABLE mirna_2_prefam (
  auto_mirna int NOT NULL DEFAULT '0',
  auto_prefam int NOT NULL DEFAULT '0',
  PRIMARY KEY (auto_mirna,auto_prefam)
);

DROP TABLE IF EXISTS mirna_chromosome_build;
CREATE TABLE mirna_chromosome_build (
  auto_mirna int NOT NULL DEFAULT '0',
  xsome varchar(20) DEFAULT NULL,
  contig_start bigint DEFAULT NULL,
  contig_end bigint DEFAULT NULL,
  strand char DEFAULT NULL
);

DROP TABLE IF EXISTS mirna_context;
CREATE TABLE mirna_context (
  auto_mirna int NOT NULL DEFAULT '0',
  transcript_id varchar(50) DEFAULT NULL,
  overlap_sense char(2) DEFAULT NULL,
  overlap_type varchar(20) DEFAULT NULL,
  number int DEFAULT NULL,
  transcript_source varchar(50) DEFAULT NULL,
  transcript_name varchar(50) DEFAULT NULL
);

DROP TABLE IF EXISTS mirna_database_links;
CREATE TABLE mirna_database_links (
  auto_mirna int NOT NULL DEFAULT '0',
  auto_db int DEFAULT NULL,
  link tinytext NOT NULL,
  display_name tinytext NOT NULL
);

DROP TABLE IF EXISTS mirna_database_url;
CREATE TABLE mirna_database_url (
  auto_db int PRIMARY KEY,
  display_name tinytext NOT NULL,
  url tinytext NOT NULL
);

DROP TABLE IF EXISTS mirna_literature_references;
CREATE TABLE mirna_literature_references (
  auto_mirna int NOT NULL DEFAULT '0',
  auto_lit int NOT NULL DEFAULT '0',
  comment mediumtext,
  order_added tinyint DEFAULT NULL
);

DROP TABLE IF EXISTS mirna_mature;
CREATE TABLE mirna_mature (
  auto_mature int PRIMARY KEY,
  mature_name varchar(40) NOT NULL DEFAULT '',
  previous_mature_id text NOT NULL,
  mature_acc varchar(20) NOT NULL DEFAULT '',
  evidence mediumtext,
  experiment mediumtext,
  similarity mediumtext,
  dead_flag int NOT NULL
);

DROP TABLE IF EXISTS mirna_pre_mature;
CREATE TABLE mirna_pre_mature (
  auto_mirna int NOT NULL DEFAULT '0',
  auto_mature int NOT NULL DEFAULT '0',
  mature_from varchar(4) DEFAULT NULL,
  mature_to varchar(4) DEFAULT NULL
);

DROP TABLE IF EXISTS mirna_prefam;
CREATE TABLE mirna_prefam (
  auto_prefam int PRIMARY KEY,
  prefam_acc varchar(15) NOT NULL DEFAULT '',
  prefam_id varchar(40) NOT NULL DEFAULT '',
  description text
);

DROP TABLE IF EXISTS mirna_species;
CREATE TABLE mirna_species (
  auto_id bigint PRIMARY KEY,
  organism varchar(10) DEFAULT NULL,
  division varchar(10) DEFAULT NULL,
  name varchar(100) DEFAULT NULL,
  taxon_id bigint DEFAULT NULL,
  taxonomy varchar(200) DEFAULT NULL,
  genome_assembly varchar(50) DEFAULT '',
  genome_accession varchar(50) DEFAULT '',
  ensembl_db varchar(50) DEFAULT NULL
);