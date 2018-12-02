package idxconfig

// Metadata stores index repository metadata configuration.
type Metadata struct {
	Offsetannotationhint 	string	// = {unordered, ordered}
    Kind 		[]Kind
}

type Kind struct {
	Name		string
	Field		[]string
}

/*
Metadata: Metadata{
	Field: [
		// read-only [system] metadata
		"odmver",		// ODM version
		"schver",		// schema version
		"base-time", 	// window start
		"max-areas", 	// maximum number of areas
		"max-cats", 	// maximum number of categories
		"app", 			// application or service name
		// read-write [application] metadata fields in the index paramaters file
		"docno", 		// document number
		"docver", 		// document version
		// start of template specific fields - generated elsewhere
	],
	Forward: [
		// read-only [system] metadata
		"odmver",		// ODM version
		"schver",		// schema version
		"base-time", 	// window start
		"max-areas", 	// maximum number of areas
		"max-cats", 	// maximum number of categories
		"app", 			// application or service name
		// read-write [application] metadata fields in the index paramaters file
		"docno", 		// document number
		"docver", 		// document version
	],
	Backword: [
		// read-only [system] metadata
		"odmver",		// ODM version
		"schver",		// schema version
		"base-time", 	// window start
		"max-areas", 	// maximum number of areas
		"max-cats", 	// maximum number of categories
		"app", 			// application or service name
		// read-write [application] metadata fields in the index paramaters file
		"docno", 		// document number
		"docver", 		// document version
	],
	Offsetannotationhint "ordered",	// = {unordered, ordered}
},
*/
