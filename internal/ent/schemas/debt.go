package schemas

import (
	"finhub-go/internal/utils/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Debt struct {
	ent.Schema
}

func (Debt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.TimestampsMixin{},
		mixins.MoneyMixin{Name: "amount"},
	}
}

func (Debt) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255),
		field.Time("purchase_date"),
		field.Time("due_date").Nillable().Optional(),
	}
}

func (Debt) Edges() []ent.Edge {
	return []ent.Edge{
		// TODO: deixar obrigatorio
		edge.To("invoice", Invoice.Type).Unique().StorageKey(edge.Column("invoice_id")).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("category", Category.Type).Unique().StorageKey(edge.Column("category_id")),
		edge.To("status", PaymentStatus.Type).Unique().StorageKey(edge.Column("status_id")),
	}
}
