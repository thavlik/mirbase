package mirbase

type Confidence struct {
	///mirna_id varchar
	MiRNAID string `json:"mirna_id"`

	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///exp_count int
	ExpCount int64 `json:"exp_count"`

	///_5p_count double
	Z_5pCount float64 `json:"5p_count"`

	///_5p_raw_count float
	Z_5pRawCount float64 `json:"5p_raw_count"`

	///_3p_count float
	Z_3pCount float64 `json:"3p_count"`

	///_3p_raw_count float
	Z_3pRawCount float64 `json:"3p_raw_count"`

	///_5p_consistent float
	Z_5pConsistent float64 `json:"5p_consistent"`

	///_5p_mature_consistent decimal
	Z_5pMatureConsistent float64 `json:"5p_mature_consistent"`

	///_3p_consistent float
	Z_3pConsistent float64 `json:"3p_consistent"`

	///_3p_mature_consistent decimal
	Z_3pMatureConsistent float64 `json:"3p_mature_consistent"`

	///_5p_overhang int
	Z_5pOverhang int64 `json:"5p_overhang"`

	///_3p_overhang int
	Z_3pOverhang int64 `json:"3p_overhang"`

	///energy_precursor float
	EnergyPrecursor *float64 `json:"energy_precursor,omitempty"`

	///energy_by_length float
	EnergyByLength float64 `json:"energy_by_length"`

	///paired_hairpin float
	PairedHairpin float64 `json:"paired_hairpin"`

	///mirdeep_score double
	MirdeepScore float64 `json:"mirdeep_score"`
}

type ConfidenceScore struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///confidence int
	Confidence int64 `json:"confidence"`
}

type DeadMiRNA struct {
	///mirna_acc varchar
	MiRNAAcc string `json:"mirna_acc"`

	///mirna_id varchar
	MiRNAID string `json:"mirna_id"`

	///previous_id varchar
	PreviousID string `json:"previous_id"`

	///forward_to varchar
	ForwardTo string `json:"forward_to"`

	///comment mediumtext
	Comment string `json:"comment,omitempty"`
}

type LiteratureReferences struct {
	///auto_lit INTEGER
	AutoLit int64 `json:"auto_lit"`

	///medline int
	Medline *int64 `json:"medline,omitempty"`

	///title tinytext
	Title string `json:"title"`

	///author tinytext
	Author string `json:"author"`

	///journal tinytext
	Journal string `json:"journal"`
}

type MatureDatabaseLinks struct {
	///auto_mature int
	AutoMature int64 `json:"auto_mature"`

	///auto_db int
	AutoDB int64 `json:"auto_db"`

	///link tinytext
	Link string `json:"link"`

	///display_name tinytext
	DisplayName string `json:"display_name"`
}

type MatureDatabaseUrl struct {
	///auto_db INTEGER
	AutoDB int64 `json:"auto_db"`

	///display_name tinytext
	DisplayName string `json:"display_name"`

	///url tinytext
	URL string `json:"url"`

	///type smallint
	Type int64 `json:"type,omitempty"`
}

type MiRNA struct {
	///auto_mirna INTEGER
	AutoMiRNA int64 `json:"auto_mirna"`

	///mirna_acc varchar
	MiRNAAcc string `json:"mirna_acc"`

	///mirna_id varchar
	MiRNAID string `json:"mirna_id"`

	///previous_mirna_id text
	PreviousMiRNAID string `json:"previous_mirna_id"`

	///description varchar
	Description string `json:"description,omitempty"`

	///sequence blob
	Sequence []byte `json:"sequence,omitempty"`

	///comment longtext
	Comment string `json:"comment,omitempty"`

	///auto_species int
	AutoSpecies int64 `json:"auto_species"`

	///dead_flag tinyint
	DeadFlag bool `json:"dead_flag"`
}

type MiRNA2Prefam struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///auto_prefam int
	AutoPrefam int64 `json:"auto_prefam"`
}

type MiRNAChromosomeBuild struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///xsome varchar
	Xsome string `json:"xsome"`

	///contig_start bigint
	ContigStart int64 `json:"contig_start"`

	///contig_end bigint
	ContigEnd int64 `json:"contig_end"`

	///strand char
	Strand string `json:"strand"`
}

type MiRNAContext struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///transcript_id varchar
	TranscriptID string `json:"transcript_id"`

	///overlap_sense char
	OverlapSense string `json:"overlap_sense"`

	///overlap_type varchar
	OverlapType string `json:"overlap_type"`

	///number int
	Number int64 `json:"number"`

	///transcript_source varchar
	TranscriptSource *string `json:"transcript_source,omitempty"`

	///transcript_name varchar
	TranscriptName *string `json:"transcript_name,omitempty"`
}

type MiRNADatabaseLinks struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///auto_db int
	AutoDB int64 `json:"auto_db"`

	///link tinytext
	Link string `json:"link"`

	///display_name tinytext
	DisplayName string `json:"display_name"`
}

type MiRNADatabaseUrl struct {
	///auto_db int
	AutoDB int64 `json:"auto_db"`

	///display_name tinytext
	DisplayName string `json:"display_name"`

	///url tinytext
	URL string `json:"url"`
}

type MiRNALiteratureReferences struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///auto_lit int
	AutoLit int64 `json:"auto_lit"`

	///comment mediumtext
	Comment string `json:"comment"`

	///order_added tinyint
	OrderAdded int64 `json:"order_added"`
}

type MiRNAMature struct {
	///auto_mature int
	AutoMature int64 `json:"auto_mature"`

	///mature_name varchar
	MatureName string `json:"mature_name"`

	///previous_mature_id text
	PreviousMatureID string `json:"previous_mature_id"`

	///mature_acc varchar
	MatureAcc string `json:"mature_acc"`

	///evidence mediumtext
	Evidence string `json:"evidence"`

	///experiment mediumtext
	Experiment string `json:"experiment"`

	///similarity mediumtext
	Similarity string `json:"similarity"`

	///dead_flag int
	DeadFlag bool `json:"dead_flag"`
}

type MiRNAPreMature struct {
	///auto_mirna int
	AutoMiRNA int64 `json:"auto_mirna"`

	///auto_mature int
	AutoMature int64 `json:"auto_mature"`

	///mature_from varchar
	MatureFrom string `json:"mature_from"`

	///mature_to varchar
	MatureTo string `json:"mature_to"`
}

type MiRNAPrefam struct {
	///auto_prefam int
	AutoPrefam int64 `json:"auto_prefam"`

	///prefam_acc varchar
	PrefamAcc string `json:"prefam_acc"`

	///prefam_id varchar
	PrefamID string `json:"prefam_id"`

	Description *string `json:"description,omitempty"`
}

type MiRNASpecies struct {
	///auto_id bigint
	AutoID int64 `json:"auto_id"`

	///organism varchar
	Organism *string `json:"organism"`

	///division varchar
	Division *string `json:"division"`

	///name varchar
	Name *string `json:"name"`

	///taxon_id bigint
	TaxonID *int64 `json:"taxon_id,omitempty"`

	///taxonomy varchar
	Taxonomy *string `json:"taxonomy,omitempty"`

	///genome_assembly varchar
	GenomeAssembly string `json:"genome_assembly"`

	///genome_accession varchar
	GenomeAccession string `json:"genome_accession"`

	///ensembl_db varchar
	EnsemblDB *string `json:"ensembl_db"`
}
