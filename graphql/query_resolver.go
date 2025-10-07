package main

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts() AccountResolver {
	return &queryResolver{server: r.server}
}
