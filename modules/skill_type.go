package modules

import (
	"github.com/graphql-go/graphql"
)

var SkillType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Skill",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if skill, ok := p.Source.(*Skill); ok == true {
					return skill.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if skill, ok := p.Source.(*Skill); ok == true {
					return skill.Name, nil
				}
				return nil, nil
			},
		},
		"level": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if skill, ok := p.Source.(*Skill); ok == true {
					return skill.Level, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	SkillType.AddFieldConfig("skill", &graphql.Field{
		Type: SkillType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Skill ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if skill, ok := p.Source.(*Skill); ok == true {
				return GetSkillByID(skill.ID)
			}
			return nil, nil
		},
	})
}
