package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"
	"errors"
	"example/graphql/graph/model"
	"example/graphql/utils"
	"fmt"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, link model.NewLink) (*model.Link, error) {
	user, ok := ctx.Value("username").(string)
	if !ok || user == "" {
		return nil, errors.New("unauthorized access")
	}
	id, err := r.Usr.UserByName(user)
	if err != nil || id == 0 {
		return nil, errors.New("unauthorized access - invalid token ")
	}

	l, err := r.Link.Create(&model.Link{
		Title:   link.Title,
		Address: link.Address,
		User: &model.User{
			ID:   fmt.Sprintf("%d", id),
			Name: user,
		},
	})

	if err != nil {
		return nil, err
	}
	publishLink(l)
	return l, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (string, error) {
	if _, err := r.Usr.Create(user); err != nil {
		return "", err
	}
	if token := utils.GenerateJWT(user.Username); token != "" {
		return token, nil
	}
	return "", errors.New("failed to generate token")
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	if !r.Usr.AuthenticateUser(input.Username, input.Password) {
		return "", errors.New("unauthorized access ")
	}

	token := utils.GenerateJWT(input.Username)
	return token, nil
}

// Refreshtoken is the resolver for the refreshtoken field.
func (r *mutationResolver) Refreshtoken(ctx context.Context, token model.RefreshTokenInput) (string, error) {
	username, err := utils.ParseToken(token.Token)
	if err != nil {
		return "", err
	}
	return utils.GenerateJWT(username), nil
}

// VoteLink is the resolver for the voteLink field.
func (r *mutationResolver) VoteLink(ctx context.Context, input model.VoteInput) (int32, error) {
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return 0, errors.New("unauthorized access")
	}
	userid, err := r.Usr.UserByName(username)
	if err != nil || userid == 0 {
		return 0, errors.New("unauthorized access - invalid token ")
	}
	return r.Link.VoteLink(input, userid)
}

// AllLinks is the resolver for the allLinks field.
func (r *queryResolver) AllLinks(ctx context.Context) ([]*model.Link, error) {
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return nil, errors.New("unauthorized access")
	}
	userid, err := r.Usr.UserByName(username)
	if err != nil || userid == 0 {
		return nil, errors.New("unauthorized access - invalid token ")
	}
	return r.Link.All(userid)
}

// LinkAdded is the resolver for the linkAdded field.
func (r *subscriptionResolver) LinkAdded(ctx context.Context) (<-chan *model.Link, error) {
	linkSubs.mu.Lock()
	defer linkSubs.mu.Unlock()
	linkChan := make(chan *model.Link, 1)
	linkSubs.clients = append(linkSubs.clients, linkChan)

	go func() {
		<-ctx.Done() //client disconnected
		linkSubs.mu.Lock()
		defer linkSubs.mu.Unlock()
		for i, v := range linkSubs.clients {
			if v == linkChan {
				linkSubs.clients = append(linkSubs.clients[:i], linkSubs.clients[i+1:]...)
				close(v)
				break
			}
		}
	}()
	return linkChan, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
