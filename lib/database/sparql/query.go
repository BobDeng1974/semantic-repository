package sparql

import "net/url"

func (*Database) getConstructQueryWithoutProperties(s string, p string, o string) (string) {
	if s == "" {
		s = " ?s "
	} else {
		s = " <" + s + "> "
	}

	if p == "" {
		p = " ?p "
	} else {
		p = " <" + p + "> "
	}

	if o == "" {
		o = " ?o "
	} else {
		o = " <" + o + "> "
	}
	return url.QueryEscape("CONSTRUCT {?s ?p ?o} " +
		"WHERE {" + "?s ?p ?o ." +
		s + p + o + "." +
		"?s <http://www.w3.org/2000/01/rdf-schema#label> ?label" +
		"}" +
		"ORDER BY ASC(?label)")
}

func (*Database) getSubjectWithAllPropertiesQuery(subject string) (string) {

	return url.QueryEscape("prefix x: <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> " +
		"construct { ?s ?p ?o } " +
		"where {<" + subject + "> (x:|!x:)* ?s . ?s ?p ?o . }")
}

func (*Database) getDeleteDeviceTypeQuery(subject string) (string) {

	return url.QueryEscape(
		"PREFIX ses: <https://senergy.infai.org/ontology/> " +
			"PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>" +
			"delete {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service ?p ?o." +
			"} where {" +
			"<" + subject + "> rdf:type ?type;" +
			"rdfs:label ?label;" +
			"ses:hasDeviceClass ?deviceclass;" +
			"ses:hasService ?service ." +
			"?service ?p ?o." +
			"}")
}
