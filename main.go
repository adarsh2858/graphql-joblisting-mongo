package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/graphql-go/graphql"
// )

// func main() {
// 	fields := graphql.Fields{
// 		"hello": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				return "world", nil
// 			},
// 		},
// 	}

// 	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
// 	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

// 	schema, err := graphql.NewSchema(schemaConfig)
// 	if err != nil {
// 		log.Fatalf("error which instantiating schema %+v", err)
// 	}

// 	query := `
// 		{
// 			hello
// 		}
// 	`
// 	params := graphql.Params{Schema: schema, RequestString: query}
// 	r := graphql.Do(params)
// 	if len(r.Errors) > 0 {
// 		log.Fatalf("something went wrong while querying %+v", r.Errors)
// 	}
// 	rJSON, _ := json.Marshal(r)
// 	fmt.Printf("Output:  %s\n", rJSON)
// }
