package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type InviteCardsRp struct {
	PgPool *pgxpool.Pool
}

type InviteCard struct {
	ID         uuid.UUID  `json:"id,omitempty"`
	BusinessId *uuid.UUID `json:"business_id,omitempty"`

	SlugCode, PassCode string `json:"-"`

	Business BusinessJSON `json:"business,omitempty"`
	Social   SocialJSON   `json:"social,omitempty"`
	Content  ContentJSON  `json:"content,omitempty"`
}

type BusinessJSON struct {
	Name        string `json:"name,omitempty"`
	CorpPhrase  string `json:"corp_phrase,omitempty"`
	Description string `json:"description,omitempty"`
	LogoImage   string `json:"logo_image,omitempty"`
}

type ContentJSON struct {
	BackgroundImage string `json:"background_image,omitempty"`
	BackgroundColor string `json:"background_color,omitempty"`
}

type SocialJSON struct {
	Links []SocialItem `json:"links,omitempty"`
}

type SocialItem struct {
	Title      string `json:"title,omitempty"`
	SubTitle   string `json:"sub_title,omitempty"`
	BackLink   string `json:"back_link,omitempty"`
	CoverImage string `json:"cover_image,omitempty"`
}

const sqlSelectInviteCardBySlug = `
	SELECT Id, BusinessId, SlugCode, PassCode, Business, Social, Content FROM InviteCard WHERE SlugCode = $1
	`

func (repo *InviteCardsRp) BySlugCode(ctx context.Context, slugCode string) (*InviteCard, error) {
	row := repo.PgPool.QueryRow(ctx, sqlSelectInviteCardBySlug, slugCode)
	return repo.ScanInviteCard(ctx, row)
}

func (repo *InviteCardsRp) ScanInviteCards(ctx context.Context, row pgx.Rows) ([]*InviteCard, error) {
	inviteCards := make([]*InviteCard, 0)

	for row.Next() {
		inviteCard, err := repo.ScanInviteCard(ctx, row)
		if err != nil {
			return nil, err
		}

		inviteCards = append(inviteCards, inviteCard)
	}

	return inviteCards, nil
}

func (repo *InviteCardsRp) ScanInviteCard(ctx context.Context, row pgx.Row) (*InviteCard, error) {
	inviteCard := &InviteCard{}
	err := row.Scan(&inviteCard.ID, &inviteCard.BusinessId, &inviteCard.SlugCode, &inviteCard.PassCode, &inviteCard.Business, &inviteCard.Social, &inviteCard.Content)
	return inviteCard, err
}
