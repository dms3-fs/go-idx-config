package idxconfig

// Indexer stores index repository server configuration.
type Indexer struct {
	Path		string 	// path to index, ex: repo/base
	Memory		string
	Stemmer		string
	Normalize	bool
	Stopper		[]string
	Corpus 		Corpus
}

type Corpus struct {
    Path		string	// path to repository, ex: repo/base/corpus
    Class		string
    Annotations	string
    Metadata	string	// path to metadata, ex: repo/base/meta
}
