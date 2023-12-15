package predicate

import (
	"encoding/json"
	"fmt"
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/shared"
	"golang.org/x/exp/slices"
	"strings"
)

type WherePredicates struct {
	Predicate Predicate
	Params    ParamBinds
}

func (w *WherePredicates) Evaluate(data schema.Schema) bool {
	return Evaluate(w.Predicate, data, w.Params)
}

func Where(query string, params ParamBinds) (*WherePredicates, error) {
	if query == "" {
		return nil, nil
	}

	predicates, err := Parse(query)
	if err != nil {
		return nil, err
	}

	var missingParams, extraParams []string

	paramsInPredicate := bindValuesFromPredicate(predicates, nil)
	for _, param := range paramsInPredicate {
		if _, ok := params[param]; !ok {
			missingParams = append(missingParams, param)
		}
	}

	// find params that are in params but not in predicate
	for param := range params {
		if !slices.Contains(paramsInPredicate, param) {
			extraParams = append(extraParams, param)
		}
	}

	if len(extraParams) > 0 || len(missingParams) > 0 {
		message := strings.Builder{}
		if len(missingParams) > 0 {
			message.WriteString(fmt.Sprintf(`missing params: "%s"`, strings.Join(missingParams, `", "`)))
		}
		if len(extraParams) > 0 {
			if message.Len() > 0 {
				message.WriteString(", ")
			}
			message.WriteString(fmt.Sprintf(`extra params: "%s"`, strings.Join(extraParams, `", "`)))
		}

		return nil, fmt.Errorf(message.String())
	}

	return &WherePredicates{
		Predicate: predicates,
		Params:    params,
	}, nil
}

func MustWhere(query string, params ParamBinds) *WherePredicates {
	where, err := Where(query, params)
	if err != nil {
		panic(err)
	}
	return where
}

func bindValuesFromPredicate(predicate Predicate, params []string) []string {
	return MustMatchPredicate(
		predicate,
		func(x *And) []string {
			for _, p := range x.L {
				params = bindValuesFromPredicate(p, params)
			}
			return params
		},
		func(x *Or) []string {
			for _, p := range x.L {
				params = bindValuesFromPredicate(p, params)
			}
			return params
		},
		func(x *Not) []string {
			return bindValuesFromPredicate(x.P, params)
		},
		func(x *Compare) []string {
			if bind, ok := x.BindValue.(*BindValue); ok {
				return append(params, bind.BindName)
			}

			return params
		},
	)
}

var (
	_ json.Unmarshaler = (*WherePredicates)(nil)
	_ json.Marshaler   = (*WherePredicates)(nil)
)

func (w *WherePredicates) UnmarshalJSON(bytes []byte) error {
	return shared.JSONParseObject(bytes, func(key string, value []byte) error {
		switch key {
		case "Predicate":
			var err error
			w.Predicate, err = PredicateFromJSON(value)
			if err != nil {
				return err
			}
		case "Params":
			return json.Unmarshal(value, &w.Params)
		default:
			return fmt.Errorf("predicate.WherePredicates: unknown key %s", key)
		}
		return nil
	})
}

func (w *WherePredicates) MarshalJSON() ([]byte, error) {
	result := map[string]json.RawMessage{}
	field_Predicate, err := PredicateToJSON(w.Predicate)
	if err != nil {
		return nil, err
	}
	if field_Predicate != nil {
		result["Predicate"] = field_Predicate
	}

	field_Params, err := json.Marshal(w.Params)
	if err != nil {
		return nil, err
	}
	if field_Params != nil {
		result["Params"] = field_Params
	}

	return json.Marshal(result)
}
