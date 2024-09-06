package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

type Postgres struct {
	*sql.DB
}

func New(connectString string) (*Postgres, error){
	db , err :=sql.Open("postgres", connectString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}
	return &Postgres{db}, nil
}

func (p *Postgres) Ping (ctx context.Context)error{
	return p.DB.PingContext(ctx)
}

func (p *Postgres) Close(){
	p.DB.Close()
}