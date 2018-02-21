package modules

import (
	"github.com/graphql-go/graphql"
)

var TechnologyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Technology",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if technology, ok := p.Source.(*Technology); ok == true {
					return technology.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if technology, ok := p.Source.(*Technology); ok == true {
					return technology.Name, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if technology, ok := p.Source.(*Technology); ok == true {
					return technology.Description, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	TechnologyType.AddFieldConfig("technology", &graphql.Field{
		Type: TechnologyType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Technology ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
	})
}
