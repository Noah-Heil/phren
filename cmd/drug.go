package cmd

import (
	"bytes"
	"encoding/xml"
	"time"
)

// Actionlisttype list
type Actionlisttype struct {
	Action []string `xml:"action,omitempty"`
}

// Affectedorganismlisttype list
type Affectedorganismlisttype struct {
	Affectedorganism []string `xml:"affected-organism,omitempty"`
}

// Ahfscodelisttype list
type Ahfscodelisttype struct {
	Ahfscode []string `xml:"ahfs-code,omitempty"`
}

// Articletype container
type Articletype struct {
	Pubmedid string `xml:"pubmed-id"`
	Citation string `xml:"citation"`
}

// Articlelisttype list
type Articlelisttype struct {
	Article []Articletype `xml:"article,omitempty"`
}

// Atccodetype container
type Atccodetype struct {
	Level []Atccodeleveltype `xml:"level"`
	Code  string             `xml:"code,attr,omitempty"`
}

// Atccodeleveltype container
type Atccodeleveltype struct {
	Value string `xml:",chardata"`
	Code  string `xml:"code,attr"`
}

// Atccodelisttype list
type Atccodelisttype struct {
	Atccode []Atccodetype `xml:"atc-code,omitempty"`
}

// Calculatedpropertykindtype May be one of logP, logS, Water Solubility, IUPAC Name, Traditional IUPAC Name, Molecular Weight, Monoisotopic Weight, SMILES, Molecular Formula, InChI, InChIKey, Polar Surface Area (PSA), Refractivity, Polarizability, Rotatable Bond Count, H Bond Acceptor Count, H Bond Donor Count, pKa (strongest acidic), pKa (strongest basic), Physiological Charge, Number of Rings, Bioavailability, Rule of Five, Ghose Filter, MDDR-Like Rule
type Calculatedpropertykindtype string

// Calculatedpropertylisttype list
type Calculatedpropertylisttype struct {
	Property []Calculatedpropertytype `xml:"property,omitempty"`
}

// Calculatedpropertysourcetype May be one of ChemAxon, ALOGPS
type Calculatedpropertysourcetype string

// Calculatedpropertytype container
type Calculatedpropertytype struct {
	Kind   Calculatedpropertykindtype   `xml:"kind"`
	Value  string                       `xml:"value"`
	Source Calculatedpropertysourcetype `xml:"source"`
}

// Carriertype container
type Carriertype struct {
	ID          string            `xml:"id"`
	Name        string            `xml:"name"`
	Organism    string            `xml:"organism"`
	Actions     Actionlisttype    `xml:"actions"`
	References  Referencelisttype `xml:"references"`
	Knownaction Knownactiontype   `xml:"known-action"`
	Polypeptide []Polypeptidetype `xml:"polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

// Carrierlisttype list
type Carrierlisttype struct {
	Carrier []Carriertype `xml:"carrier,omitempty"`
}

// Categorytype container
type Categorytype struct {
	Category string `xml:"category"`
	Meshid   string `xml:"mesh-id"`
}

// Categorylisttype list
type Categorylisttype struct {
	Category []Categorytype `xml:"category,omitempty"`
}

// Classificationtype Drug classification is obtained from ClassyFire (http://classyfire.wishartlab.com).
type Classificationtype struct {
	Description       string   `xml:"description"`
	Directparent      string   `xml:"direct-parent"`
	Kingdom           string   `xml:"kingdom"`
	Superclass        string   `xml:"superclass"`
	Class             string   `xml:"class"`
	Subclass          string   `xml:"subclass"`
	Alternativeparent []string `xml:"alternative-parent,omitempty"`
	Substituent       []string `xml:"substituent,omitempty"`
}

// Cost container
type Cost struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

// Dosagetype container
type Dosagetype struct {
	Form     string `xml:"form"`
	Route    string `xml:"route"`
	Strength string `xml:"strength"`
}

// Dosagelisttype list
type Dosagelisttype struct {
	Dosage []Dosagetype `xml:"dosage,omitempty"`
}

// Drugbankdrugsaltidtype The DrugBank ID is used to uniquely identify a drug or salt entry. There is a primary ID and several secondary IDs that come from older ID formats or merged entries.
type Drugbankdrugsaltidtype struct {
	Drugbankdrugsaltidvalue Drugbankdrugsaltidvalue `xml:",chardata"`
	Primary                 bool                    `xml:"primary,attr,omitempty"`
}

// UnmarshalXML unloads XML from Drugbankdrugsaltidtype structure
func (t *Drugbankdrugsaltidtype) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Drugbankdrugsaltidtype
	var overlay struct {
		*T
		Primary *bool `xml:"primary,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Primary = (*bool)(&overlay.T.Primary)
	return d.DecodeElement(&overlay, &start)
}

// Drugbankdrugsaltidvalue Must match the pattern DB[0-9]{5}|DBSALT[0-9]{6}|APRD[0-9]{5}|BIOD[0-9]{5}|BTD[0-9]{5}|EXPT[0-9]{5}|NUTR[0-9]{5}
type Drugbankdrugsaltidvalue string

// Drugbankmetaboliteidtype The metabolite DrugBank ID uniquely identifies a metabolite entry. Multiple IDs indicate a merged entry.
type Drugbankmetaboliteidtype struct {
	Drugbankmetaboliteidvalue Drugbankmetaboliteidvalue `xml:",chardata"`
	Primary                   bool                      `xml:"primary,attr,omitempty"`
}

// UnmarshalXML unloads XML from Drugbankmetaboliteidtype structure
func (t *Drugbankmetaboliteidtype) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Drugbankmetaboliteidtype
	var overlay struct {
		*T
		Primary *bool `xml:"primary,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Primary = (*bool)(&overlay.T.Primary)
	return d.DecodeElement(&overlay, &start)
}

// Drugbankmetaboliteidvalue Must match the pattern DBMET[0-9]{5}
type Drugbankmetaboliteidvalue string

// Drugbanktype This is the root element type for the DrugBank database schema.
type Drugbanktype struct {
	Drug       []Drugtype `xml:"drug"`
	Version    string     `xml:"version,attr"`
	Exportedon time.Time  `xml:"exported-on,attr"`
}

// MarshalXML loads xml into drugbanktype structure
func (t *Drugbanktype) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Drugbanktype
	var layout struct {
		*T
		Exportedon *xsdDate `xml:"exported-on,attr"`
	}
	layout.T = (*T)(t)
	layout.Exportedon = (*xsdDate)(&layout.T.Exportedon)
	return e.EncodeElement(layout, start)
}

// UnmarshalXML unloads XML into drugbanktype structure
func (t *Drugbanktype) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Drugbanktype
	var overlay struct {
		*T
		Exportedon *xsdDate `xml:"exported-on,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Exportedon = (*xsdDate)(&overlay.T.Exportedon)
	return d.DecodeElement(&overlay, &start)
}

// Druginteractionlisttype struct for a list of drug interaction types
type Druginteractionlisttype struct {
	Druginteraction []Druginteractiontype `xml:"drug-interaction,omitempty"`
}

// Druginteractiontype struct for drug interaction types
type Druginteractiontype struct {
	Drugbankid  Drugbankdrugsaltidtype `xml:"drugbank-id"`
	Name        string                 `xml:"name"`
	Description string                 `xml:"description"`
}

// Drugtype Main drug datastructure
type Drugtype struct {
	Drugbankid              []Drugbankdrugsaltidtype       `xml:"drugbank-id"`
	Name                    string                         `xml:"name"`
	Description             string                         `xml:"description"`
	Casnumber               string                         `xml:"cas-number"`
	Unii                    string                         `xml:"unii"`
	Averagemass             float32                        `xml:"average-mass,omitempty"`
	Monoisotopicmass        float32                        `xml:"monoisotopic-mass,omitempty"`
	State                   Statetype                      `xml:"state,omitempty"`
	Groups                  Grouplisttype                  `xml:"groups"`
	Generalreferences       Referencelisttype              `xml:"general-references"`
	Synthesisreference      string                         `xml:"synthesis-reference"`
	Indication              string                         `xml:"indication"`
	Pharmacodynamics        string                         `xml:"pharmacodynamics"`
	Mechanismofaction       string                         `xml:"mechanism-of-action"`
	Toxicity                string                         `xml:"toxicity"`
	Metabolism              string                         `xml:"metabolism"`
	Absorption              string                         `xml:"absorption"`
	Halflife                string                         `xml:"half-life"`
	Proteinbinding          string                         `xml:"protein-binding"`
	Routeofelimination      string                         `xml:"route-of-elimination"`
	Volumeofdistribution    string                         `xml:"volume-of-distribution"`
	Clearance               string                         `xml:"clearance"`
	Classification          Classificationtype             `xml:"classification,omitempty"`
	Salts                   Saltlisttype                   `xml:"salts"`
	Synonyms                Synonymlisttype                `xml:"synonyms"`
	Products                Productlisttype                `xml:"products"`
	Internationalbrands     Internationalbrandlisttype     `xml:"international-brands"`
	Mixtures                Mixturelisttype                `xml:"mixtures"`
	Packagers               Packagerlisttype               `xml:"packagers"`
	Manufacturers           Manufacturerlisttype           `xml:"manufacturers"`
	Prices                  Pricelisttype                  `xml:"prices"`
	Categories              Categorylisttype               `xml:"categories"`
	Affectedorganisms       Affectedorganismlisttype       `xml:"affected-organisms"`
	Dosages                 Dosagelisttype                 `xml:"dosages"`
	Atccodes                Atccodelisttype                `xml:"atc-codes"`
	Ahfscodes               Ahfscodelisttype               `xml:"ahfs-codes"`
	Pdbentries              Pdbentrylisttype               `xml:"pdb-entries"`
	Fdalabel                string                         `xml:"fda-label,omitempty"`
	Msds                    string                         `xml:"msds,omitempty"`
	Patents                 Patentlisttype                 `xml:"patents"`
	Foodinteractions        Foodinteractionlisttype        `xml:"food-interactions"`
	Druginteractions        Druginteractionlisttype        `xml:"drug-interactions"`
	Sequences               Sequencelisttype               `xml:"sequences,omitempty"`
	Calculatedproperties    Calculatedpropertylisttype     `xml:"calculated-properties,omitempty"`
	Experimentalproperties  Experimentalpropertylisttype   `xml:"experimental-properties"`
	Externalidentifiers     Externalidentifierlisttype     `xml:"external-identifiers"`
	Externallinks           Externallinklisttype           `xml:"external-links"`
	Pathways                Pathwaylisttype                `xml:"pathways"`
	Reactions               Reactionlisttype               `xml:"reactions"`
	Snpeffects              Snpeffectlisttype              `xml:"snp-effects"`
	Snpadversedrugreactions Snpadversedrugreactionlisttype `xml:"snp-adverse-drug-reactions"`
	Targets                 Targetlisttype                 `xml:"targets"`
	Enzymes                 Enzymelisttype                 `xml:"enzymes"`
	Carriers                Carrierlisttype                `xml:"carriers"`
	Transporters            Transporterlisttype            `xml:"transporters"`
	Type                    Type                           `xml:"type,attr"`
	Created                 time.Time                      `xml:"created,attr"`
	Updated                 time.Time                      `xml:"updated,attr"`
}

// MarshalXML Used to marshal drugtype into XML
func (t *Drugtype) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Drugtype
	var layout struct {
		*T
		Created *xsdDate `xml:"created,attr"`
		Updated *xsdDate `xml:"updated,attr"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDate)(&layout.T.Created)
	layout.Updated = (*xsdDate)(&layout.T.Updated)
	return e.EncodeElement(layout, start)
}

// UnmarshalXML Used to UnmarshalXML into go structs
func (t *Drugtype) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Drugtype
	var overlay struct {
		*T
		Created *xsdDate `xml:"created,attr"`
		Updated *xsdDate `xml:"updated,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDate)(&overlay.T.Created)
	overlay.Updated = (*xsdDate)(&overlay.T.Updated)
	return d.DecodeElement(&overlay, &start)
}

// Enzymelisttype list of enzyme types
type Enzymelisttype struct {
	Enzyme []Enzymetype `xml:"enzyme,omitempty"`
}

// Enzymetype a type for enzymes
type Enzymetype struct {
	ID                 string            `xml:"id"`
	Name               string            `xml:"name"`
	Organism           string            `xml:"organism"`
	Actions            Actionlisttype    `xml:"actions"`
	References         Referencelisttype `xml:"references"`
	Knownaction        Knownactiontype   `xml:"known-action"`
	Polypeptide        []Polypeptidetype `xml:"polypeptide,omitempty"`
	Inhibitionstrength string            `xml:"inhibition-strength"`
	Inductionstrength  string            `xml:"induction-strength"`
	Position           int               `xml:"position,attr,omitempty"`
}

// Experimentalpropertytype container
type Experimentalpropertytype struct {
	Kind   Experimentalpropertykindtype `xml:"kind"`
	Value  string                       `xml:"value"`
	Source string                       `xml:"source"`
}

//Experimentalpropertykindtype May be one of Water Solubility, Melting Point, Boiling Point, logP, logS, Hydrophobicity, Isoelectric Point, caco2 Permeability, pKa, Molecular Weight, Molecular Formula
type Experimentalpropertykindtype string

// Experimentalpropertylisttype list of experimental property
type Experimentalpropertylisttype struct {
	Property []Experimentalpropertytype `xml:"property,omitempty"`
}

// Externalidentifierresourcetype May be one of UniProtKB, Wikipedia, ChEBI, ChEMBL, PubChem Compound, PubChem Substance, Drugs Product Database (DPD), KEGG Compound, KEGG Drug, ChemSpider, BindingDB, National Drug Code Directory, GenBank, Therapeutic Targets Database, PharmGKB, PDB, IUPHAR, Guide to Pharmacology
type Externalidentifierresourcetype string

// Externalidentifierlisttype list
type Externalidentifierlisttype struct {
	Externalidentifier []Externalidentifiertype `xml:"external-identifier,omitempty"`
}

// Externalidentifiertype type for external identifiers
type Externalidentifiertype struct {
	Resource   Externalidentifierresourcetype `xml:"resource"`
	Identifier string                         `xml:"identifier"`
}

// Externallinkresourcetype May be one of RxList, PDRhealth, Drugs.com
type Externallinkresourcetype string

// Externallinktype container
type Externallinktype struct {
	Resource Externallinkresourcetype `xml:"resource"`
	URL      string                   `xml:"url"`
}

// Externallinklisttype list type for external lists
type Externallinklisttype struct {
	Externallink []Externallinktype `xml:"external-link,omitempty"`
}

// Foodinteractionlisttype list
type Foodinteractionlisttype struct {
	Foodinteraction []string `xml:"food-interaction,omitempty"`
}

// Goclassifiertype container
type Goclassifiertype struct {
	Category    string `xml:"category"`
	Description string `xml:"description"`
}

// Goclassifierlisttype list
type Goclassifierlisttype struct {
	Goclassifier []Goclassifiertype `xml:"go-classifier,omitempty"`
}

// Grouptype May be one of approved, illicit, experimental, withdrawn, nutraceutical, investigational, vet_approved
type Grouptype string

// Grouplisttype list
type Grouplisttype struct {
	Group Grouptype `xml:"group"`
}

// Internationalbrandtype container
type Internationalbrandtype struct {
	Name    string `xml:"name"`
	Company string `xml:"company"`
}

// Internationalbrandlisttype list
type Internationalbrandlisttype struct {
	Internationalbrand []Internationalbrandtype `xml:"international-brand,omitempty"`
}

// Knownactiontype May be one of yes, no, unknown
type Knownactiontype string

// Linktype container
type Linktype struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
}

// Linklisttype list
type Linklisttype struct {
	Link []Linktype `xml:"link,omitempty"`
}

// Manufacturertype container
type Manufacturertype struct {
	Value   string `xml:",chardata"`
	Generic bool   `xml:"generic,attr,omitempty"`
	URL     string `xml:"url,attr,omitempty"`
}

// Manufacturerlisttype list
type Manufacturerlisttype struct {
	Manufacturer []Manufacturertype `xml:"manufacturer,omitempty"`
}

// Mixturetype container
type Mixturetype struct {
	Name        string `xml:"name"`
	Ingredients string `xml:"ingredients"`
}

// Mixturelisttype list
type Mixturelisttype struct {
	Mixture []Mixturetype `xml:"mixture,omitempty"`
}

// Organism container
type Organism struct {
	Value          string `xml:",chardata"`
	Ncbitaxonomyid string `xml:"ncbi-taxonomy-id,attr,omitempty"`
}

// Packagertype container
type Packagertype struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

// Packagerlisttype list
type Packagerlisttype struct {
	Packager []Packagertype `xml:"packager,omitempty"`
}

// Patenttype container
type Patenttype struct {
	Number             string `xml:"number"`
	Country            string `xml:"country"`
	Approved           string `xml:"approved"`
	Expires            string `xml:"expires"`
	Pediatricextension bool   `xml:"pediatric-extension"`
}

// Patentlisttype list
type Patentlisttype struct {
	Patent []Patenttype `xml:"patent,omitempty"`
}

// Pathwaytype container
type Pathwaytype struct {
	Smpdbid  string                `xml:"smpdb-id"`
	Name     string                `xml:"name"`
	Category string                `xml:"category"`
	Drugs    Pathwaydruglisttype   `xml:"drugs"`
	Enzymes  Pathwayenzymelisttype `xml:"enzymes"`
}

// Pathwaydrugtype container
type Pathwaydrugtype struct {
	Drugbankid Drugbankdrugsaltidtype `xml:"drugbank-id"`
	Name       string                 `xml:"name"`
}

// Pathwaydruglisttype list
type Pathwaydruglisttype struct {
	Drug []Pathwaydrugtype `xml:"drug"`
}

// Pathwayenzymelisttype list
type Pathwayenzymelisttype struct {
	Uniprotid []string `xml:"uniprot-id,omitempty"`
}

// Pathwaylisttype list
type Pathwaylisttype struct {
	Pathway []Pathwaytype `xml:"pathway,omitempty"`
}

// Pdbentrylisttype list
type Pdbentrylisttype struct {
	Pdbentry []string `xml:"pdb-entry,omitempty"`
}

// Polypeptidetype container
type Polypeptidetype struct {
	Name                 string                                `xml:"name"`
	Generalfunction      string                                `xml:"general-function"`
	Specificfunction     string                                `xml:"specific-function"`
	Genename             string                                `xml:"gene-name"`
	Locus                string                                `xml:"locus"`
	Cellularlocation     string                                `xml:"cellular-location"`
	Transmembraneregions string                                `xml:"transmembrane-regions"`
	Signalregions        string                                `xml:"signal-regions"`
	Theoreticalpi        string                                `xml:"theoretical-pi"`
	Molecularweight      string                                `xml:"molecular-weight"`
	Chromosomelocation   string                                `xml:"chromosome-location"`
	Organism             Organism                              `xml:"organism"`
	Externalidentifiers  Polypeptideexternalidentifierlisttype `xml:"external-identifiers"`
	Synonyms             Polypeptidesynonymlisttype            `xml:"synonyms"`
	Aminoacidsequence    Sequencetype                          `xml:"amino-acid-sequence"`
	Genesequence         Sequencetype                          `xml:"gene-sequence"`
	Pfams                Pfamlisttype                          `xml:"pfams"`
	Goclassifiers        Goclassifierlisttype                  `xml:"go-classifiers"`
	Source               string                                `xml:"source,attr"`
}

// Pfamtype container
type Pfamtype struct {
	Identifier string `xml:"identifier"`
	Name       string `xml:"name"`
}

// Pfamlisttype list
type Pfamlisttype struct {
	Pfam []Pfamtype `xml:"pfam,omitempty"`
}

// Polypeptideexternalidentifierresourcetype May be one of UniProtKB, UniProt Accession, HUGO Gene Nomenclature Committee (HGNC), Human Protein Reference Database (HPRD), GenAtlas, GeneCards, GenBank Gene Database, GenBank Protein Database, ChEMBL, IUPHAR, Guide to Pharmacology
type Polypeptideexternalidentifierresourcetype string

// Polypeptideexternalidentifierlisttype list
type Polypeptideexternalidentifierlisttype struct {
	Externalidentifier []Polypeptideexternalidentifiertype `xml:"external-identifier,omitempty"`
}

// Polypeptideexternalidentifiertype container
type Polypeptideexternalidentifiertype struct {
	Resource   Polypeptideexternalidentifierresourcetype `xml:"resource"`
	Identifier string                                    `xml:"identifier"`
}

// Polypeptidesynonymlisttype list
type Polypeptidesynonymlisttype struct {
	Synonym []string `xml:"synonym,omitempty"`
}

// Polypeptidelisttype list
type Polypeptidelisttype struct {
	Polypeptide []Polypeptidetype `xml:"polypeptide,omitempty"`
}

// Pricetype The price for the given drug in US or Canadian currency.
type Pricetype struct {
	Description string `xml:"description"`
	Cost        Cost   `xml:"cost"`
	Unit        string `xml:"unit"`
}

// Pricelisttype list
type Pricelisttype struct {
	Price []Pricetype `xml:"price,omitempty"`
}

// Producttype container for products
type Producttype struct {
	Name                 string             `xml:"name"`
	Labeller             string             `xml:"labeller"`
	Ndcid                string             `xml:"ndc-id"`
	Ndcproductcode       string             `xml:"ndc-product-code"`
	Dpdid                string             `xml:"dpd-id,omitempty"`
	Emaproductcode       string             `xml:"ema-product-code,omitempty"`
	Emamanumber          string             `xml:"ema-ma-number,omitempty"`
	Startedmarketingon   string             `xml:"started-marketing-on"`
	Endedmarketingon     string             `xml:"ended-marketing-on"`
	Dosageform           string             `xml:"dosage-form"`
	Strength             string             `xml:"strength"`
	Route                string             `xml:"route"`
	Fdaapplicationnumber string             `xml:"fda-application-number"`
	Generic              bool               `xml:"generic"`
	Overthecounter       bool               `xml:"over-the-counter"`
	Approved             bool               `xml:"approved"`
	Country              Productcountrytype `xml:"country"`
	Source               Productsourcetype  `xml:"source"`
}

// Productsourcetype May be one of FDA NDC, DPD, EMA
type Productsourcetype string

// Productcountrytype May be one of US, Canada, EU
type Productcountrytype string

// Productlisttype list
type Productlisttype struct {
	Product []Producttype `xml:"product,omitempty"`
}

// Reactiontype container
type Reactiontype struct {
	Sequence     string                 `xml:"sequence"`
	Leftelement  Reactionelementtype    `xml:"left-element"`
	Rightelement Reactionelementtype    `xml:"right-element"`
	Enzymes      Reactionenzymelisttype `xml:"enzymes"`
}

// Reactionlisttype list
type Reactionlisttype struct {
	Reaction []Reactiontype `xml:"reaction,omitempty"`
}

// Reactionelementtype container
type Reactionelementtype struct {
	Drugbankid string `xml:"drugbank-id"`
	Name       string `xml:"name"`
}

// Reactionenzymetype container
type Reactionenzymetype struct {
	Drugbankid string `xml:"drugbank-id"`
	Name       string `xml:"name"`
	Uniprotid  string `xml:"uniprot-id"`
}

// Reactionenzymelisttype list
type Reactionenzymelisttype struct {
	Enzyme []Reactionenzymetype `xml:"enzyme,omitempty"`
}

// Referencelisttype list
type Referencelisttype struct {
	Articles  Articlelisttype  `xml:"articles"`
	Textbooks Textbooklisttype `xml:"textbooks"`
	Links     Linklisttype     `xml:"links"`
}

// Salttype container
type Salttype struct {
	Drugbankid       []Drugbankdrugsaltidtype `xml:"drugbank-id,omitempty"`
	Name             string                   `xml:"name"`
	Unii             string                   `xml:"unii"`
	Casnumber        string                   `xml:"cas-number"`
	Inchikey         string                   `xml:"inchikey"`
	Averagemass      float32                  `xml:"average-mass,omitempty"`
	Monoisotopicmass float32                  `xml:"monoisotopic-mass,omitempty"`
}

// Saltlisttype list
type Saltlisttype struct {
	Salt []Salttype `xml:"salt,omitempty"`
}

// Sequence container
type Sequence struct {
	Value  string `xml:",chardata"`
	Format string `xml:"format,attr,omitempty"`
}

// Sequencetype container
type Sequencetype struct {
	Value  string `xml:",chardata"`
	Format string `xml:"format,attr,omitempty"`
}

// Sequencelisttype list
type Sequencelisttype struct {
	Sequence []Sequence `xml:"sequence,omitempty"`
}

// Snpadversedrugreactiontype container
type Snpadversedrugreactiontype struct {
	Proteinname     string `xml:"protein-name"`
	Genesymbol      string `xml:"gene-symbol"`
	Uniprotid       string `xml:"uniprot-id"`
	Rsid            string `xml:"rs-id"`
	Allele          string `xml:"allele"`
	Adversereaction string `xml:"adverse-reaction"`
	Description     string `xml:"description"`
	Pubmedid        string `xml:"pubmed-id"`
}

// Snpadversedrugreactionlisttype list
type Snpadversedrugreactionlisttype struct {
	Reaction []Snpadversedrugreactiontype `xml:"reaction,omitempty"`
}

// Snpeffecttype container
type Snpeffecttype struct {
	Proteinname    string `xml:"protein-name"`
	Genesymbol     string `xml:"gene-symbol"`
	Uniprotid      string `xml:"uniprot-id"`
	Rsid           string `xml:"rs-id"`
	Allele         string `xml:"allele"`
	Definingchange string `xml:"defining-change"`
	Description    string `xml:"description"`
	Pubmedid       string `xml:"pubmed-id"`
}

// Snpeffectlisttype list
type Snpeffectlisttype struct {
	Effect []Snpeffecttype `xml:"effect,omitempty"`
}

// Statetype May be one of solid, liquid, gas
type Statetype string

// Synonymtype type to contain synonymtype
type Synonymtype struct {
	Value    string `xml:",chardata"`
	Language string `xml:"language,attr,omitempty"`
	Coder    string `xml:"coder,attr,omitempty"`
}

// Synonymlisttype list
type Synonymlisttype struct {
	Synonym []Synonymtype `xml:"synonym,omitempty"`
}

// Targettype target type container
type Targettype struct {
	ID          string            `xml:"id"`
	Name        string            `xml:"name"`
	Organism    string            `xml:"organism"`
	Actions     Actionlisttype    `xml:"actions"`
	References  Referencelisttype `xml:"references"`
	Knownaction Knownactiontype   `xml:"known-action"`
	Polypeptide []Polypeptidetype `xml:"polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

// Targetlisttype list of targetlisttype
type Targetlisttype struct {
	Target []Targettype `xml:"target,omitempty"`
}

// Textbooktype type to contain textbook entries
type Textbooktype struct {
	Isbn     string `xml:"isbn"`
	Citation string `xml:"citation"`
}

// Textbooklisttype list type for textbooktype
type Textbooklisttype struct {
	Textbook []Textbooktype `xml:"textbook,omitempty"`
}

// Transportertype type to contain transporter construct
type Transportertype struct {
	ID          string            `xml:"id"`
	Name        string            `xml:"name"`
	Organism    string            `xml:"organism"`
	Actions     Actionlisttype    `xml:"actions"`
	References  Referencelisttype `xml:"references"`
	Knownaction Knownactiontype   `xml:"known-action"`
	Polypeptide []Polypeptidetype `xml:"polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

// Transporterlisttype list type for transporter
type Transporterlisttype struct {
	Transporter []Transportertype `xml:"transporter,omitempty"`
}

// Type May be one of small molecule, biotech
type Type string

type xsdDate time.Time

// UnmarshalText unmarshals time
func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}

// MarshalText marhsals time
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}

// MarshalXML Marshals XML for time
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}

// MarshalXMLAttr unmarshals time
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
