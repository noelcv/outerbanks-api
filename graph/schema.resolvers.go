package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cat-turner/outerbanks-api/graph/generated"
	"cat-turner/outerbanks-api/graph/model"
	"context"
	"fmt"
	"strconv"
)

// UpsertCharacter is the resolver for the upsertCharacter field.
func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	id := input.ID
	var character model.Character
	character.Name = input.Name
	character.CliqueType = input.CliqueType

	n := len(r.Resolver.CharacterStore)
	if n == 0 {
		r.Resolver.CharacterStore = make(map[string]model.Character)
	}

	if id != nil {
		cs, ok := r.Resolver.CharacterStore[*id]
		if !ok {
			return nil, fmt.Errorf("not found")
		}
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		} else {
			character.IsHero = cs.IsHero
		}
		r.Resolver.CharacterStore[*id] = character
	} else {
		//generate unique id
		nid := strconv.Itoa(n + 1)
		character.ID = nid
		if input.IsHero != nil {
			character.IsHero = *input.IsHero
		}
		r.Resolver.CharacterStore[nid] = character
	}
	return &character, nil
}

// Character is the resolver for the character field.
func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	character, ok := r.Resolver.CharacterStore[id]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &character, nil
}

// Characters is the resolver for the characters field.
func (r *queryResolver) Characters(ctx context.Context, cliqueType model.CliqueType) ([]*model.Character, error) {
  characters := make([]*model.Character, 0)
	for idx := range r.Resolver.CharacterStore {
		character := r.Resolver.CharacterStore[idx]
		//segregate by cliqueType --> pogues && kooks
		if character.CliqueType == cliqueType {
			characters = append(characters, &character)
		}
	}
	return characters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
