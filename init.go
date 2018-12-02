package idxconfig

import (
	//"encoding/base64"
	//"errors"
	//"fmt"
	"io"
	//"time"

)

func Init(out io.Writer) (*IdxConfig, error) {

	//datastore := DefaultDatastoreConfig()

	iconf := &IdxConfig{
		Indexer: Indexer{
			Path: 		"repo/base", // path to index, ex: repo/base
			Memory: 	"100M",
			Stemmer: 	"krovetz",
			Normalize: 	true,
			Stopper: 	[]string{
				"a",
				"an",
				"the",
				"as",
			},
			Corpus: 	Corpus{
				Path: 			"repo/base/corpus", // path to repository, ex: repo/base/corpus
				Class:			"html",
				Annotations: 	"",
				Metadata: 		"repo/base/meta", 	// path to metadata, ex: repo/base/meta
			},
		},
		Metadata: Metadata{
    		Kind: []Kind{{
				Name: "blog",
				Field: []string{
					"About",
	            	"Address",
	            	"Affiliation",
	            	"Author",
	            	"Brand",
	            	"Citation",
	            	"Description",
	            	"Email",
	            	"Headline",
	            	"Keywords",
	            	"Language",
	            	"Name",
	            	"Telephone",
	            	"Version",
					},
				},
			},
			Offsetannotationhint: "",	// = {unordered, ordered}
  		},

		Retriever: Retriever{
			MaxResultCount: 100,
		},

		// Note: do not configure reposet proprties, the system
		// will internally set these when creating a reposet.
		// these are common index repository metadata
		RepoSet: RepoSet{
			odmver:		"",		// ODM version
			schver:		"",		// schema version
			kind:		"",		// kind of content
			basetime:	"",	 	// window start
			maxareas:	0,	 	// maximum number of areas
			maxcats:	0,	 	// maximum number of categories
			offset:		0,		// encoding relative time
			app:		"",		// application or service name
			docno:		"", 	// document number
			docver:		"", 	// document version
		},
	}

	return iconf, nil
}
