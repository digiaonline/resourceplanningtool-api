package modules

import (
	"github.com/graphql-go/graphql"
)

var NewsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "News",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(*News); ok == true {
					return news.ID, nil
				}
				return nil, nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(*News); ok == true {
					return news.URL, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if news, ok := p.Source.(*News); ok == true {
					return news.Description, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	NewsType.AddFieldConfig("news", &graphql.Field{
		Type: NewsType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "News ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
	})
}
