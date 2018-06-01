package cmd

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Actionlisttype struct {
	Action []string `xml:"http://www.drugbank.ca action,omitempty"`
}

type Affectedorganismlisttype struct {
	Affectedorganism []string `xml:"http://www.drugbank.ca affected-organism,omitempty"`
}

type Ahfscodelisttype struct {
	Ahfscode []string `xml:"http://www.drugbank.ca ahfs-code,omitempty"`
}

type Articlelisttype struct {
	Article []Articletype `xml:"http://www.drugbank.ca article,omitempty"`
}

type Articletype struct {
	Pubmedid string `xml:"http://www.drugbank.ca pubmed-id"`
	Citation string `xml:"http://www.drugbank.ca citation"`
}

type Atccodeleveltype struct {
	Value string `xml:",chardata"`
	Code  string `xml:"code,attr"`
}

type Atccodelisttype struct {
	Atccode []Atccodetype `xml:"http://www.drugbank.ca atc-code,omitempty"`
}

type Atccodetype struct {
	Level []Atccodeleveltype `xml:"http://www.drugbank.ca level"`
	Code  string             `xml:"code,attr,omitempty"`
}

// May be one of logP, logS, Water Solubility, IUPAC Name, Traditional IUPAC Name, Molecular Weight, Monoisotopic Weight, SMILES, Molecular Formula, InChI, InChIKey, Polar Surface Area (PSA), Refractivity, Polarizability, Rotatable Bond Count, H Bond Acceptor Count, H Bond Donor Count, pKa (strongest acidic), pKa (strongest basic), Physiological Charge, Number of Rings, Bioavailability, Rule of Five, Ghose Filter, MDDR-Like Rule
type Calculatedpropertykindtype string

type Calculatedpropertylisttype struct {
	Property []Calculatedpropertytype `xml:"http://www.drugbank.ca property,omitempty"`
}

// May be one of ChemAxon, ALOGPS
type Calculatedpropertysourcetype string

type Calculatedpropertytype struct {
	Kind   Calculatedpropertykindtype   `xml:"http://www.drugbank.ca kind"`
	Value  string                       `xml:"http://www.drugbank.ca value"`
	Source Calculatedpropertysourcetype `xml:"http://www.drugbank.ca source"`
}

type Carrierlisttype struct {
	Carrier []Carriertype `xml:"http://www.drugbank.ca carrier,omitempty"`
}

type Carriertype struct {
	Id          string            `xml:"http://www.drugbank.ca id"`
	Name        string            `xml:"http://www.drugbank.ca name"`
	Organism    string            `xml:"http://www.drugbank.ca organism"`
	Actions     Actionlisttype    `xml:"http://www.drugbank.ca actions"`
	References  Referencelisttype `xml:"http://www.drugbank.ca references"`
	Knownaction Knownactiontype   `xml:"http://www.drugbank.ca known-action"`
	Polypeptide []Polypeptidetype `xml:"http://www.drugbank.ca polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

type Categorylisttype struct {
	Category []Categorytype `xml:"http://www.drugbank.ca category,omitempty"`
}

type Categorytype struct {
	Category string `xml:"http://www.drugbank.ca category"`
	Meshid   string `xml:"http://www.drugbank.ca mesh-id"`
}

// Drug classification is obtained from ClassyFire (http://classyfire.wishartlab.com).
type Classificationtype struct {
	Description       string   `xml:"http://www.drugbank.ca description"`
	Directparent      string   `xml:"http://www.drugbank.ca direct-parent"`
	Kingdom           string   `xml:"http://www.drugbank.ca kingdom"`
	Superclass        string   `xml:"http://www.drugbank.ca superclass"`
	Class             string   `xml:"http://www.drugbank.ca class"`
	Subclass          string   `xml:"http://www.drugbank.ca subclass"`
	Alternativeparent []string `xml:"http://www.drugbank.ca alternative-parent,omitempty"`
	Substituent       []string `xml:"http://www.drugbank.ca substituent,omitempty"`
}

type Cost struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Dosagelisttype struct {
	Dosage []Dosagetype `xml:"http://www.drugbank.ca dosage,omitempty"`
}

type Dosagetype struct {
	Form     string `xml:"http://www.drugbank.ca form"`
	Route    string `xml:"http://www.drugbank.ca route"`
	Strength string `xml:"http://www.drugbank.ca strength"`
}

// The DrugBank ID is used to uniquely identify a drug or salt entry. There is a primary ID and several secondary IDs that come from older ID formats or merged entries.
type Drugbankdrugsaltidtype struct {
	Drugbankdrugsaltidvalue Drugbankdrugsaltidvalue `xml:",chardata"`
	Primary                 bool                    `xml:"primary,attr,omitempty"`
}

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
	Drug       []Drugtype `xml:"http://www.drugbank.ca drug"`
	Version    string     `xml:"version,attr"`
	Exportedon time.Time  `xml:"exported-on,attr"`
}

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
	Druginteraction []Druginteractiontype `xml:"http://www.drugbank.ca drug-interaction,omitempty"`
}

// Druginteractiontype struct for drug interaction types
type Druginteractiontype struct {
	Drugbankid  Drugbankdrugsaltidtype `xml:"http://www.drugbank.ca drugbank-id"`
	Name        string                 `xml:"http://www.drugbank.ca name"`
	Description string                 `xml:"http://www.drugbank.ca description"`
}

// Drugtype Main drug datastructure
type Drugtype struct {
	Drugbankid              []Drugbankdrugsaltidtype       `xml:"http://www.drugbank.ca drugbank-id"`
	Name                    string                         `xml:"http://www.drugbank.ca name"`
	Description             string                         `xml:"http://www.drugbank.ca description"`
	Casnumber               string                         `xml:"http://www.drugbank.ca cas-number"`
	Unii                    string                         `xml:"http://www.drugbank.ca unii"`
	Averagemass             float32                        `xml:"http://www.drugbank.ca average-mass,omitempty"`
	Monoisotopicmass        float32                        `xml:"http://www.drugbank.ca monoisotopic-mass,omitempty"`
	State                   Statetype                      `xml:"http://www.drugbank.ca state,omitempty"`
	Groups                  Grouplisttype                  `xml:"http://www.drugbank.ca groups"`
	Generalreferences       Referencelisttype              `xml:"http://www.drugbank.ca general-references"`
	Synthesisreference      string                         `xml:"http://www.drugbank.ca synthesis-reference"`
	Indication              string                         `xml:"http://www.drugbank.ca indication"`
	Pharmacodynamics        string                         `xml:"http://www.drugbank.ca pharmacodynamics"`
	Mechanismofaction       string                         `xml:"http://www.drugbank.ca mechanism-of-action"`
	Toxicity                string                         `xml:"http://www.drugbank.ca toxicity"`
	Metabolism              string                         `xml:"http://www.drugbank.ca metabolism"`
	Absorption              string                         `xml:"http://www.drugbank.ca absorption"`
	Halflife                string                         `xml:"http://www.drugbank.ca half-life"`
	Proteinbinding          string                         `xml:"http://www.drugbank.ca protein-binding"`
	Routeofelimination      string                         `xml:"http://www.drugbank.ca route-of-elimination"`
	Volumeofdistribution    string                         `xml:"http://www.drugbank.ca volume-of-distribution"`
	Clearance               string                         `xml:"http://www.drugbank.ca clearance"`
	Classification          Classificationtype             `xml:"http://www.drugbank.ca classification,omitempty"`
	Salts                   Saltlisttype                   `xml:"http://www.drugbank.ca salts"`
	Synonyms                Synonymlisttype                `xml:"http://www.drugbank.ca synonyms"`
	Products                Productlisttype                `xml:"http://www.drugbank.ca products"`
	Internationalbrands     Internationalbrandlisttype     `xml:"http://www.drugbank.ca international-brands"`
	Mixtures                Mixturelisttype                `xml:"http://www.drugbank.ca mixtures"`
	Packagers               Packagerlisttype               `xml:"http://www.drugbank.ca packagers"`
	Manufacturers           Manufacturerlisttype           `xml:"http://www.drugbank.ca manufacturers"`
	Prices                  Pricelisttype                  `xml:"http://www.drugbank.ca prices"`
	Categories              Categorylisttype               `xml:"http://www.drugbank.ca categories"`
	Affectedorganisms       Affectedorganismlisttype       `xml:"http://www.drugbank.ca affected-organisms"`
	Dosages                 Dosagelisttype                 `xml:"http://www.drugbank.ca dosages"`
	Atccodes                Atccodelisttype                `xml:"http://www.drugbank.ca atc-codes"`
	Ahfscodes               Ahfscodelisttype               `xml:"http://www.drugbank.ca ahfs-codes"`
	Pdbentries              Pdbentrylisttype               `xml:"http://www.drugbank.ca pdb-entries"`
	Fdalabel                string                         `xml:"http://www.drugbank.ca fda-label,omitempty"`
	Msds                    string                         `xml:"http://www.drugbank.ca msds,omitempty"`
	Patents                 Patentlisttype                 `xml:"http://www.drugbank.ca patents"`
	Foodinteractions        Foodinteractionlisttype        `xml:"http://www.drugbank.ca food-interactions"`
	Druginteractions        Druginteractionlisttype        `xml:"http://www.drugbank.ca drug-interactions"`
	Sequences               Sequencelisttype               `xml:"http://www.drugbank.ca sequences,omitempty"`
	Calculatedproperties    Calculatedpropertylisttype     `xml:"http://www.drugbank.ca calculated-properties,omitempty"`
	Experimentalproperties  Experimentalpropertylisttype   `xml:"http://www.drugbank.ca experimental-properties"`
	Externalidentifiers     Externalidentifierlisttype     `xml:"http://www.drugbank.ca external-identifiers"`
	Externallinks           Externallinklisttype           `xml:"http://www.drugbank.ca external-links"`
	Pathways                Pathwaylisttype                `xml:"http://www.drugbank.ca pathways"`
	Reactions               Reactionlisttype               `xml:"http://www.drugbank.ca reactions"`
	Snpeffects              Snpeffectlisttype              `xml:"http://www.drugbank.ca snp-effects"`
	Snpadversedrugreactions Snpadversedrugreactionlisttype `xml:"http://www.drugbank.ca snp-adverse-drug-reactions"`
	Targets                 Targetlisttype                 `xml:"http://www.drugbank.ca targets"`
	Enzymes                 Enzymelisttype                 `xml:"http://www.drugbank.ca enzymes"`
	Carriers                Carrierlisttype                `xml:"http://www.drugbank.ca carriers"`
	Transporters            Transporterlisttype            `xml:"http://www.drugbank.ca transporters"`
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
	Enzyme []Enzymetype `xml:"http://www.drugbank.ca enzyme,omitempty"`
}

// Enzymetype a type for enzymes
type Enzymetype struct {
	Id                 string            `xml:"http://www.drugbank.ca id"`
	Name               string            `xml:"http://www.drugbank.ca name"`
	Organism           string            `xml:"http://www.drugbank.ca organism"`
	Actions            Actionlisttype    `xml:"http://www.drugbank.ca actions"`
	References         Referencelisttype `xml:"http://www.drugbank.ca references"`
	Knownaction        Knownactiontype   `xml:"http://www.drugbank.ca known-action"`
	Polypeptide        []Polypeptidetype `xml:"http://www.drugbank.ca polypeptide,omitempty"`
	Inhibitionstrength string            `xml:"http://www.drugbank.ca inhibition-strength"`
	Inductionstrength  string            `xml:"http://www.drugbank.ca induction-strength"`
	Position           int               `xml:"position,attr,omitempty"`
}

//Experimentalpropertykindtype May be one of Water Solubility, Melting Point, Boiling Point, logP, logS, Hydrophobicity, Isoelectric Point, caco2 Permeability, pKa, Molecular Weight, Molecular Formula
type Experimentalpropertykindtype string

// Experimentalpropertylisttype list of experimental property
type Experimentalpropertylisttype struct {
	Property []Experimentalpropertytype `xml:"http://www.drugbank.ca property,omitempty"`
}

type Experimentalpropertytype struct {
	Kind   Experimentalpropertykindtype `xml:"http://www.drugbank.ca kind"`
	Value  string                       `xml:"http://www.drugbank.ca value"`
	Source string                       `xml:"http://www.drugbank.ca source"`
}

type Externalidentifierlisttype struct {
	Externalidentifier []Externalidentifiertype `xml:"http://www.drugbank.ca external-identifier,omitempty"`
}

// Externalidentifierresourcetype May be one of UniProtKB, Wikipedia, ChEBI, ChEMBL, PubChem Compound, PubChem Substance, Drugs Product Database (DPD), KEGG Compound, KEGG Drug, ChemSpider, BindingDB, National Drug Code Directory, GenBank, Therapeutic Targets Database, PharmGKB, PDB, IUPHAR, Guide to Pharmacology
type Externalidentifierresourcetype string

// Externalidentifiertype type for external identifiers
type Externalidentifiertype struct {
	Resource   Externalidentifierresourcetype `xml:"http://www.drugbank.ca resource"`
	Identifier string                         `xml:"http://www.drugbank.ca identifier"`
}

// Externallinklisttype list type for external lists
type Externallinklisttype struct {
	Externallink []Externallinktype `xml:"http://www.drugbank.ca external-link,omitempty"`
}

// Externallinkresourcetype May be one of RxList, PDRhealth, Drugs.com
type Externallinkresourcetype string

type Externallinktype struct {
	Resource Externallinkresourcetype `xml:"http://www.drugbank.ca resource"`
	Url      string                   `xml:"http://www.drugbank.ca url"`
}

type Foodinteractionlisttype struct {
	Foodinteraction []string `xml:"http://www.drugbank.ca food-interaction,omitempty"`
}

type Goclassifierlisttype struct {
	Goclassifier []Goclassifiertype `xml:"http://www.drugbank.ca go-classifier,omitempty"`
}

type Goclassifiertype struct {
	Category    string `xml:"http://www.drugbank.ca category"`
	Description string `xml:"http://www.drugbank.ca description"`
}

type Grouplisttype struct {
	Group Grouptype `xml:"http://www.drugbank.ca group"`
}

// May be one of approved, illicit, experimental, withdrawn, nutraceutical, investigational, vet_approved
type Grouptype string

type Internationalbrandlisttype struct {
	Internationalbrand []Internationalbrandtype `xml:"http://www.drugbank.ca international-brand,omitempty"`
}

type Internationalbrandtype struct {
	Name    string `xml:"http://www.drugbank.ca name"`
	Company string `xml:"http://www.drugbank.ca company"`
}

// May be one of yes, no, unknown
type Knownactiontype string

type Linklisttype struct {
	Link []Linktype `xml:"http://www.drugbank.ca link,omitempty"`
}

type Linktype struct {
	Title string `xml:"http://www.drugbank.ca title"`
	Url   string `xml:"http://www.drugbank.ca url"`
}

type Manufacturerlisttype struct {
	Manufacturer []Manufacturertype `xml:"http://www.drugbank.ca manufacturer,omitempty"`
}

type Manufacturertype struct {
	Value   string `xml:",chardata"`
	Generic bool   `xml:"generic,attr,omitempty"`
	Url     string `xml:"url,attr,omitempty"`
}

type Mixturelisttype struct {
	Mixture []Mixturetype `xml:"http://www.drugbank.ca mixture,omitempty"`
}

type Mixturetype struct {
	Name        string `xml:"http://www.drugbank.ca name"`
	Ingredients string `xml:"http://www.drugbank.ca ingredients"`
}

type Organism struct {
	Value          string `xml:",chardata"`
	Ncbitaxonomyid string `xml:"ncbi-taxonomy-id,attr,omitempty"`
}

type Packagerlisttype struct {
	Packager []Packagertype `xml:"http://www.drugbank.ca packager,omitempty"`
}

type Packagertype struct {
	Name string `xml:"http://www.drugbank.ca name"`
	Url  string `xml:"http://www.drugbank.ca url"`
}

type Patentlisttype struct {
	Patent []Patenttype `xml:"http://www.drugbank.ca patent,omitempty"`
}

type Patenttype struct {
	Number             string `xml:"http://www.drugbank.ca number"`
	Country            string `xml:"http://www.drugbank.ca country"`
	Approved           string `xml:"http://www.drugbank.ca approved"`
	Expires            string `xml:"http://www.drugbank.ca expires"`
	Pediatricextension bool   `xml:"http://www.drugbank.ca pediatric-extension"`
}

type Pathwaydruglisttype struct {
	Drug []Pathwaydrugtype `xml:"http://www.drugbank.ca drug"`
}

type Pathwaydrugtype struct {
	Drugbankid Drugbankdrugsaltidtype `xml:"http://www.drugbank.ca drugbank-id"`
	Name       string                 `xml:"http://www.drugbank.ca name"`
}

type Pathwayenzymelisttype struct {
	Uniprotid []string `xml:"http://www.drugbank.ca uniprot-id,omitempty"`
}

type Pathwaylisttype struct {
	Pathway []Pathwaytype `xml:"http://www.drugbank.ca pathway,omitempty"`
}

type Pathwaytype struct {
	Smpdbid  string                `xml:"http://www.drugbank.ca smpdb-id"`
	Name     string                `xml:"http://www.drugbank.ca name"`
	Category string                `xml:"http://www.drugbank.ca category"`
	Drugs    Pathwaydruglisttype   `xml:"http://www.drugbank.ca drugs"`
	Enzymes  Pathwayenzymelisttype `xml:"http://www.drugbank.ca enzymes"`
}

type Pdbentrylisttype struct {
	Pdbentry []string `xml:"http://www.drugbank.ca pdb-entry,omitempty"`
}

type Pfamlisttype struct {
	Pfam []Pfamtype `xml:"http://www.drugbank.ca pfam,omitempty"`
}

type Pfamtype struct {
	Identifier string `xml:"http://www.drugbank.ca identifier"`
	Name       string `xml:"http://www.drugbank.ca name"`
}

type Polypeptideexternalidentifierlisttype struct {
	Externalidentifier []Polypeptideexternalidentifiertype `xml:"http://www.drugbank.ca external-identifier,omitempty"`
}

// May be one of UniProtKB, UniProt Accession, HUGO Gene Nomenclature Committee (HGNC), Human Protein Reference Database (HPRD), GenAtlas, GeneCards, GenBank Gene Database, GenBank Protein Database, ChEMBL, IUPHAR, Guide to Pharmacology
type Polypeptideexternalidentifierresourcetype string

type Polypeptideexternalidentifiertype struct {
	Resource   Polypeptideexternalidentifierresourcetype `xml:"http://www.drugbank.ca resource"`
	Identifier string                                    `xml:"http://www.drugbank.ca identifier"`
}

type Polypeptidelisttype struct {
	Polypeptide []Polypeptidetype `xml:"http://www.drugbank.ca polypeptide,omitempty"`
}

type Polypeptidesynonymlisttype struct {
	Synonym []string `xml:"http://www.drugbank.ca synonym,omitempty"`
}

type Polypeptidetype struct {
	Name                 string                                `xml:"http://www.drugbank.ca name"`
	Generalfunction      string                                `xml:"http://www.drugbank.ca general-function"`
	Specificfunction     string                                `xml:"http://www.drugbank.ca specific-function"`
	Genename             string                                `xml:"http://www.drugbank.ca gene-name"`
	Locus                string                                `xml:"http://www.drugbank.ca locus"`
	Cellularlocation     string                                `xml:"http://www.drugbank.ca cellular-location"`
	Transmembraneregions string                                `xml:"http://www.drugbank.ca transmembrane-regions"`
	Signalregions        string                                `xml:"http://www.drugbank.ca signal-regions"`
	Theoreticalpi        string                                `xml:"http://www.drugbank.ca theoretical-pi"`
	Molecularweight      string                                `xml:"http://www.drugbank.ca molecular-weight"`
	Chromosomelocation   string                                `xml:"http://www.drugbank.ca chromosome-location"`
	Organism             Organism                              `xml:"http://www.drugbank.ca organism"`
	Externalidentifiers  Polypeptideexternalidentifierlisttype `xml:"http://www.drugbank.ca external-identifiers"`
	Synonyms             Polypeptidesynonymlisttype            `xml:"http://www.drugbank.ca synonyms"`
	Aminoacidsequence    Sequencetype                          `xml:"http://www.drugbank.ca amino-acid-sequence"`
	Genesequence         Sequencetype                          `xml:"http://www.drugbank.ca gene-sequence"`
	Pfams                Pfamlisttype                          `xml:"http://www.drugbank.ca pfams"`
	Goclassifiers        Goclassifierlisttype                  `xml:"http://www.drugbank.ca go-classifiers"`
	Source               string                                `xml:"source,attr"`
}

type Pricelisttype struct {
	Price []Pricetype `xml:"http://www.drugbank.ca price,omitempty"`
}

// The price for the given drug in US or Canadian currency.
type Pricetype struct {
	Description string `xml:"http://www.drugbank.ca description"`
	Cost        Cost   `xml:"http://www.drugbank.ca cost"`
	Unit        string `xml:"http://www.drugbank.ca unit"`
}

// May be one of US, Canada, EU
type Productcountrytype string

type Productlisttype struct {
	Product []Producttype `xml:"http://www.drugbank.ca product,omitempty"`
}

// May be one of FDA NDC, DPD, EMA
type Productsourcetype string

type Producttype struct {
	Name                 string             `xml:"http://www.drugbank.ca name"`
	Labeller             string             `xml:"http://www.drugbank.ca labeller"`
	Ndcid                string             `xml:"http://www.drugbank.ca ndc-id"`
	Ndcproductcode       string             `xml:"http://www.drugbank.ca ndc-product-code"`
	Dpdid                string             `xml:"http://www.drugbank.ca dpd-id,omitempty"`
	Emaproductcode       string             `xml:"http://www.drugbank.ca ema-product-code,omitempty"`
	Emamanumber          string             `xml:"http://www.drugbank.ca ema-ma-number,omitempty"`
	Startedmarketingon   string             `xml:"http://www.drugbank.ca started-marketing-on"`
	Endedmarketingon     string             `xml:"http://www.drugbank.ca ended-marketing-on"`
	Dosageform           string             `xml:"http://www.drugbank.ca dosage-form"`
	Strength             string             `xml:"http://www.drugbank.ca strength"`
	Route                string             `xml:"http://www.drugbank.ca route"`
	Fdaapplicationnumber string             `xml:"http://www.drugbank.ca fda-application-number"`
	Generic              bool               `xml:"http://www.drugbank.ca generic"`
	Overthecounter       bool               `xml:"http://www.drugbank.ca over-the-counter"`
	Approved             bool               `xml:"http://www.drugbank.ca approved"`
	Country              Productcountrytype `xml:"http://www.drugbank.ca country"`
	Source               Productsourcetype  `xml:"http://www.drugbank.ca source"`
}

type Reactionelementtype struct {
	Drugbankid string `xml:"http://www.drugbank.ca drugbank-id"`
	Name       string `xml:"http://www.drugbank.ca name"`
}

type Reactionenzymelisttype struct {
	Enzyme []Reactionenzymetype `xml:"http://www.drugbank.ca enzyme,omitempty"`
}

type Reactionenzymetype struct {
	Drugbankid string `xml:"http://www.drugbank.ca drugbank-id"`
	Name       string `xml:"http://www.drugbank.ca name"`
	Uniprotid  string `xml:"http://www.drugbank.ca uniprot-id"`
}

type Reactionlisttype struct {
	Reaction []Reactiontype `xml:"http://www.drugbank.ca reaction,omitempty"`
}

type Reactiontype struct {
	Sequence     string                 `xml:"http://www.drugbank.ca sequence"`
	Leftelement  Reactionelementtype    `xml:"http://www.drugbank.ca left-element"`
	Rightelement Reactionelementtype    `xml:"http://www.drugbank.ca right-element"`
	Enzymes      Reactionenzymelisttype `xml:"http://www.drugbank.ca enzymes"`
}

type Referencelisttype struct {
	Articles  Articlelisttype  `xml:"http://www.drugbank.ca articles"`
	Textbooks Textbooklisttype `xml:"http://www.drugbank.ca textbooks"`
	Links     Linklisttype     `xml:"http://www.drugbank.ca links"`
}

type Saltlisttype struct {
	Salt []Salttype `xml:"http://www.drugbank.ca salt,omitempty"`
}

type Salttype struct {
	Drugbankid       []Drugbankdrugsaltidtype `xml:"http://www.drugbank.ca drugbank-id,omitempty"`
	Name             string                   `xml:"http://www.drugbank.ca name"`
	Unii             string                   `xml:"http://www.drugbank.ca unii"`
	Casnumber        string                   `xml:"http://www.drugbank.ca cas-number"`
	Inchikey         string                   `xml:"http://www.drugbank.ca inchikey"`
	Averagemass      float32                  `xml:"http://www.drugbank.ca average-mass,omitempty"`
	Monoisotopicmass float32                  `xml:"http://www.drugbank.ca monoisotopic-mass,omitempty"`
}

type Sequence struct {
	Value  string `xml:",chardata"`
	Format string `xml:"format,attr,omitempty"`
}

type Sequencelisttype struct {
	Sequence []Sequence `xml:"http://www.drugbank.ca sequence,omitempty"`
}

type Sequencetype struct {
	Value  string `xml:",chardata"`
	Format string `xml:"format,attr,omitempty"`
}

type Snpadversedrugreactionlisttype struct {
	Reaction []Snpadversedrugreactiontype `xml:"http://www.drugbank.ca reaction,omitempty"`
}

type Snpadversedrugreactiontype struct {
	Proteinname     string `xml:"http://www.drugbank.ca protein-name"`
	Genesymbol      string `xml:"http://www.drugbank.ca gene-symbol"`
	Uniprotid       string `xml:"http://www.drugbank.ca uniprot-id"`
	Rsid            string `xml:"http://www.drugbank.ca rs-id"`
	Allele          string `xml:"http://www.drugbank.ca allele"`
	Adversereaction string `xml:"http://www.drugbank.ca adverse-reaction"`
	Description     string `xml:"http://www.drugbank.ca description"`
	Pubmedid        string `xml:"http://www.drugbank.ca pubmed-id"`
}

type Snpeffectlisttype struct {
	Effect []Snpeffecttype `xml:"http://www.drugbank.ca effect,omitempty"`
}

type Snpeffecttype struct {
	Proteinname    string `xml:"http://www.drugbank.ca protein-name"`
	Genesymbol     string `xml:"http://www.drugbank.ca gene-symbol"`
	Uniprotid      string `xml:"http://www.drugbank.ca uniprot-id"`
	Rsid           string `xml:"http://www.drugbank.ca rs-id"`
	Allele         string `xml:"http://www.drugbank.ca allele"`
	Definingchange string `xml:"http://www.drugbank.ca defining-change"`
	Description    string `xml:"http://www.drugbank.ca description"`
	Pubmedid       string `xml:"http://www.drugbank.ca pubmed-id"`
}

// Statetype May be one of solid, liquid, gas
type Statetype string

// Synonymtype type to contain synonymtype
type Synonymtype struct {
	Value    string `xml:",chardata"`
	Language string `xml:"language,attr,omitempty"`
	Coder    string `xml:"coder,attr,omitempty"`
}

// Targettype target type container
type Targettype struct {
	Id          string            `xml:"http://www.drugbank.ca id"`
	Name        string            `xml:"http://www.drugbank.ca name"`
	Organism    string            `xml:"http://www.drugbank.ca organism"`
	Actions     Actionlisttype    `xml:"http://www.drugbank.ca actions"`
	References  Referencelisttype `xml:"http://www.drugbank.ca references"`
	Knownaction Knownactiontype   `xml:"http://www.drugbank.ca known-action"`
	Polypeptide []Polypeptidetype `xml:"http://www.drugbank.ca polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

// Targetlisttype list of targetlisttype
type Targetlisttype struct {
	Target []Targettype `xml:"http://www.drugbank.ca target,omitempty"`
}

// Textbooktype type to contain textbook entries
type Textbooktype struct {
	Isbn     string `xml:"http://www.drugbank.ca isbn"`
	Citation string `xml:"http://www.drugbank.ca citation"`
}

// Textbooklisttype list type for textbooktype
type Textbooklisttype struct {
	Textbook []Textbooktype `xml:"http://www.drugbank.ca textbook,omitempty"`
}

// Transportertype type to contain transporter construct
type Transportertype struct {
	Id          string            `xml:"http://www.drugbank.ca id"`
	Name        string            `xml:"http://www.drugbank.ca name"`
	Organism    string            `xml:"http://www.drugbank.ca organism"`
	Actions     Actionlisttype    `xml:"http://www.drugbank.ca actions"`
	References  Referencelisttype `xml:"http://www.drugbank.ca references"`
	Knownaction Knownactiontype   `xml:"http://www.drugbank.ca known-action"`
	Polypeptide []Polypeptidetype `xml:"http://www.drugbank.ca polypeptide,omitempty"`
	Position    int               `xml:"position,attr,omitempty"`
}

// Transporterlisttype list type for transporter
type Transporterlisttype struct {
	Transporter []Transportertype `xml:"http://www.drugbank.ca transporter,omitempty"`
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
