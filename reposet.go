package idxconfig

// Reposet metadata stores index repository common metadata configuration.
type RepoSet struct{
	// read-only [system] metadata
	odmver		string		// ODM version
	schver		string		// schema version
	kind		string		// kind of content
	basetime	string	 	// window start
	maxareas	uint8	 	// maximum number of areas
	maxcats		uint8	 	// maximum number of categories
	offset		int64		// encoding relative time
	app			string		// application or service name
	// read-write [application] metadata fields in the index paramaters file
	docno		string 		// document number
	docver		string 		// document version
	// start of template specific fields - generated elsewhere
}
