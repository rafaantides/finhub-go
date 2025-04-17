// Code generated by ent, DO NOT EDIT.

package debt

import (
	"finhub-go/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldUpdatedAt, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldAmount, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldTitle, v))
}

// PurchaseDate applies equality check predicate on the "purchase_date" field. It's identical to PurchaseDateEQ.
func PurchaseDate(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldPurchaseDate, v))
}

// DueDate applies equality check predicate on the "due_date" field. It's identical to DueDateEQ.
func DueDate(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldDueDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldUpdatedAt, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldAmount, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Debt {
	return predicate.Debt(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Debt {
	return predicate.Debt(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Debt {
	return predicate.Debt(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Debt {
	return predicate.Debt(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Debt {
	return predicate.Debt(sql.FieldContainsFold(FieldTitle, v))
}

// PurchaseDateEQ applies the EQ predicate on the "purchase_date" field.
func PurchaseDateEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldPurchaseDate, v))
}

// PurchaseDateNEQ applies the NEQ predicate on the "purchase_date" field.
func PurchaseDateNEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldPurchaseDate, v))
}

// PurchaseDateIn applies the In predicate on the "purchase_date" field.
func PurchaseDateIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldPurchaseDate, vs...))
}

// PurchaseDateNotIn applies the NotIn predicate on the "purchase_date" field.
func PurchaseDateNotIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldPurchaseDate, vs...))
}

// PurchaseDateGT applies the GT predicate on the "purchase_date" field.
func PurchaseDateGT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldPurchaseDate, v))
}

// PurchaseDateGTE applies the GTE predicate on the "purchase_date" field.
func PurchaseDateGTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldPurchaseDate, v))
}

// PurchaseDateLT applies the LT predicate on the "purchase_date" field.
func PurchaseDateLT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldPurchaseDate, v))
}

// PurchaseDateLTE applies the LTE predicate on the "purchase_date" field.
func PurchaseDateLTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldPurchaseDate, v))
}

// DueDateEQ applies the EQ predicate on the "due_date" field.
func DueDateEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldEQ(FieldDueDate, v))
}

// DueDateNEQ applies the NEQ predicate on the "due_date" field.
func DueDateNEQ(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNEQ(FieldDueDate, v))
}

// DueDateIn applies the In predicate on the "due_date" field.
func DueDateIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldIn(FieldDueDate, vs...))
}

// DueDateNotIn applies the NotIn predicate on the "due_date" field.
func DueDateNotIn(vs ...time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldNotIn(FieldDueDate, vs...))
}

// DueDateGT applies the GT predicate on the "due_date" field.
func DueDateGT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGT(FieldDueDate, v))
}

// DueDateGTE applies the GTE predicate on the "due_date" field.
func DueDateGTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldGTE(FieldDueDate, v))
}

// DueDateLT applies the LT predicate on the "due_date" field.
func DueDateLT(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLT(FieldDueDate, v))
}

// DueDateLTE applies the LTE predicate on the "due_date" field.
func DueDateLTE(v time.Time) predicate.Debt {
	return predicate.Debt(sql.FieldLTE(FieldDueDate, v))
}

// DueDateIsNil applies the IsNil predicate on the "due_date" field.
func DueDateIsNil() predicate.Debt {
	return predicate.Debt(sql.FieldIsNull(FieldDueDate))
}

// DueDateNotNil applies the NotNil predicate on the "due_date" field.
func DueDateNotNil() predicate.Debt {
	return predicate.Debt(sql.FieldNotNull(FieldDueDate))
}

// HasInvoice applies the HasEdge predicate on the "invoice" edge.
func HasInvoice() predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InvoiceTable, InvoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInvoiceWith applies the HasEdge predicate on the "invoice" edge with a given conditions (other predicates).
func HasInvoiceWith(preds ...predicate.Invoice) predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := newInvoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatus applies the HasEdge predicate on the "status" edge.
func HasStatus() predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusWith applies the HasEdge predicate on the "status" edge with a given conditions (other predicates).
func HasStatusWith(preds ...predicate.PaymentStatus) predicate.Debt {
	return predicate.Debt(func(s *sql.Selector) {
		step := newStatusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Debt) predicate.Debt {
	return predicate.Debt(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Debt) predicate.Debt {
	return predicate.Debt(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Debt) predicate.Debt {
	return predicate.Debt(sql.NotPredicates(p))
}
