package v2

import (
	"strings"

	parserErrors "github.com/asyncapi/parser-go/pkg/error"
	"github.com/asyncapi/parser-go/pkg/jsonpath"
	v2 "github.com/asyncapi/parser-go/pkg/schema/asyncapi/v2"

	"github.com/pkg/errors"
)

type Parser struct {
	jsonpath.RefLoader
	root               string
	documents          map[string]map[string]interface{}
	referenceTrack     map[string]bool
	blackListedPathMap map[string]bool
}

var ErrCircularDependency = errors.New("circular dependency")

func (p Parser) Parse(doc map[string]interface{}) error {
	// validate document against schema
	err := v2.Parse(doc)
	if err != nil {
		return err
	}
	ref, err := jsonpath.NewRef("#")
	if err != nil {
		return err
	}
	p.root = ref.URI()
	p.documents[p.root] = doc
	documentErrors := p.dereference(ref, p.documents[p.root])
	return parserErrors.New(documentErrors...)
}

func (p *Parser) dereferenceMap(rootRef jsonpath.Ref, v *map[string]interface{}) []error {
	var errs []error
	for key, value := range *v {
		//Here be Dragons!
		if "$ref" == key {
			// need to track visited nodes in order to avoid circular dependencies
			if reported, found := p.referenceTrack[rootRef.String()]; found {
				if !reported {
					errs = append(errs, errors.Wrap(ErrCircularDependency, rootRef.String()))
					p.referenceTrack[rootRef.String()] = true
				}
				continue
			}
			p.referenceTrack[rootRef.String()] = false

			// to allow recursive de-referencing, prepend the current root uri on local references
			if s, ok := value.(string); ok && strings.HasPrefix(s, "#") {
				value = rootRef.URI() + s
			}

			refKey, err := jsonpath.NewRef(value)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			_, ok := p.documents[refKey.URI()]
			if !ok {
				refDoc, err := p.RefLoader.Load(refKey.URI())
				if err != nil {
					errs = append(errs, err)
					continue
				}
				p.documents[refKey.URI()] = refDoc
			}
			refMap, err := jsonpath.GetRefObject(refKey.Path(), p.documents[refKey.URI()])
			if err != nil {
				errs = append(errs, errors.Wrap(err, refKey.String()))
				continue
			}
			refErrs := p.dereference(refKey, refMap)
			if len(refErrs) > 0 {
				errs = append(errs, refErrs...)
				continue
			}
			delete(*v, key)
			// inject resolved reference
			for refKey, refValue := range refMap {
				(*v)[refKey] = refValue
			}
			continue
		}
		itemRef, err := rootRef.NewChild(key)
		if err != nil {
			return append(errs, err)
		}
		if _, found := p.blackListedPathMap[itemRef.String()]; found {
			continue
		}
		errs = append(errs, p.dereference(itemRef, value)...)
	}
	return errs
}

func (p Parser) dereference(ref jsonpath.Ref, v interface{}) []error {
	switch v := v.(type) {
	case []interface{}:
		return p.dereferenceArray(ref, v)
	case map[string]interface{}:
		return p.dereferenceMap(ref, &v)
	default:
		return nil
	}
}

func (p Parser) dereferenceArray(ref jsonpath.Ref, v []interface{}) []error {
	var errs []error
	for i, v := range v {
		childRef, err := ref.NewChild(i)
		if err != nil {
			errs = append(errs, err)
		}
		errs = append(errs, p.dereference(childRef, v)...)
	}
	return errs
}

func NewParser(refLoader jsonpath.RefLoader, blackListedPaths ...string) Parser {
	blackListedPathMap := make(map[string]bool)
	for _, key := range blackListedPaths {
		blackListedPathMap[key] = true
	}
	return Parser{
		RefLoader:          refLoader,
		documents:          make(map[string]map[string]interface{}),
		referenceTrack:     make(map[string]bool),
		blackListedPathMap: blackListedPathMap,
	}
}
