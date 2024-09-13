package main

import "encoding/xml"

type Response struct {
	XMLName      xml.Name `xml:"Response"`
	Text         string   `xml:",chardata"`
	ResponseDate string   `xml:"responseDate"`
	Request      struct {
		Text           string `xml:",chardata"`
		Verb           string `xml:"verb,attr"`
		Set            string `xml:"set,attr"`
		MetadataPrefix string `xml:"metadataPrefix,attr"`
	} `xml:"request"`
	Error struct {
		Text string `xml:",chardata"`
		Code string `xml:"code,attr"`
	} `xml:"error"`
	GetRecord struct {
		Text   string `xml:",chardata"`
		Record struct {
			Text   string `xml:",chardata"`
			Header struct {
				Text   string `xml:",chardata"`
				Status string `xml:"status,attr"`
			} `xml:"header"`
			Metadata string `xml:"metadata"`
			About    string `xml:"about"`
		} `xml:"record"`
	} `xml:"GetRecord"`
	Identify        string `xml:"Identify"`
	ListIdentifiers struct {
		Text            string `xml:",chardata"`
		ResumptionToken struct {
			Text             string `xml:",chardata"`
			CompleteListSize string `xml:"completeListSize,attr"`
			Cursor           string `xml:"cursor,attr"`
			ExpirationDate   string `xml:"expirationDate,attr"`
		} `xml:"resumptionToken"`
	} `xml:"ListIdentifiers"`
	ListMetadataFormats string `xml:"ListMetadataFormats"`
	ListRecords         struct {
		Text   string `xml:",chardata"`
		Record []struct {
			Text   string `xml:",chardata"`
			Xmlns  string `xml:"xmlns,attr"`
			Header struct {
				Text       string   `xml:",chardata"`
				Status     string   `xml:"status,attr"`
				Identifier string   `xml:"identifier"`
				Datestamp  string   `xml:"datestamp"`
				SetSpec    []string `xml:"setSpec"`
			} `xml:"header"`
			Metadata struct {
				Text string `xml:",chardata"`
				DIDL struct {
					Text           string `xml:",chardata"`
					Didl           string `xml:"didl,attr"`
					Dii            string `xml:"dii,attr"`
					Dip            string `xml:"dip,attr"`
					Dcterms        string `xml:"dcterms,attr"`
					DIDLDocumentId string `xml:"DIDLDocumentId,attr"`
					SchemaLocation string `xml:"schemaLocation,attr"`
					Xsi            string `xml:"xsi,attr"`
					Item           struct {
						Text       string `xml:",chardata"`
						Descriptor []struct {
							Text      string `xml:",chardata"`
							Statement struct {
								Text       string `xml:",chardata"`
								MimeType   string `xml:"mimeType,attr"`
								Identifier string `xml:"Identifier"`
								Modified   string `xml:"modified"`
							} `xml:"Statement"`
						} `xml:"Descriptor"`
						Component struct {
							Text     string `xml:",chardata"`
							Resource struct {
								Text     string `xml:",chardata"`
								MimeType string `xml:"mimeType,attr"`
								Ref      string `xml:"ref,attr"`
							} `xml:"Resource"`
						} `xml:"Component"`
						Item []struct {
							Text       string `xml:",chardata"`
							Descriptor struct {
								Text      string `xml:",chardata"`
								Statement struct {
									Text       string `xml:",chardata"`
									MimeType   string `xml:"mimeType,attr"`
									ObjectType string `xml:"ObjectType"`
								} `xml:"Statement"`
							} `xml:"Descriptor"`
							Component struct {
								Text     string `xml:",chardata"`
								Resource struct {
									Text     string `xml:",chardata"`
									MimeType string `xml:"mimeType,attr"`
									Ref      string `xml:"ref,attr"`
									Dc       struct {
										Text           string   `xml:",chardata"`
										OaiDc          string   `xml:"oai_dc,attr"`
										Dc             string   `xml:"dc,attr"`
										SchemaLocation string   `xml:"schemaLocation,attr"`
										Relation       string   `xml:"relation"`
										Title          string   `xml:"title"`
										Creator        []string `xml:"creator"`
										Subject        []string `xml:"subject"`
										Description    string   `xml:"description"`
										Publisher      string   `xml:"publisher"`
										Date           string   `xml:"date"`
										Type           []string `xml:"type"`
										Format         []string `xml:"format"`
										Language       []string `xml:"language"`
										Identifier     []string `xml:"identifier"`
										Contributor    []string `xml:"contributor"`
									} `xml:"dc"`
								} `xml:"Resource"`
							} `xml:"Component"`
						} `xml:"Item"`
					} `xml:"Item"`
				} `xml:"DIDL"`
			} `xml:"metadata"`
			About string `xml:"about"`
		} `xml:"record"`
		ResumptionToken struct {
			Text             string `xml:",chardata"`
			CompleteListSize string `xml:"completeListSize,attr"`
			Cursor           string `xml:"cursor,attr"`
			ExpirationDate   string `xml:"expirationDate,attr"`
		} `xml:"resumptionToken"`
	} `xml:"ListRecords"`
	ListSets struct {
		Text            string `xml:",chardata"`
		ResumptionToken struct {
			Text             string `xml:",chardata"`
			CompleteListSize string `xml:"completeListSize,attr"`
			Cursor           string `xml:"cursor,attr"`
			ExpirationDate   string `xml:"expirationDate,attr"`
		} `xml:"resumptionToken"`
	} `xml:"ListSets"`
}
