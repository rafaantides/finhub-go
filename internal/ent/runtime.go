// Code generated by ent, DO NOT EDIT.

package ent

import (
	"finhub-go/internal/ent/category"
	"finhub-go/internal/ent/debt"
	"finhub-go/internal/ent/invoice"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/ent/schemas"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	categoryMixin := schemas.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryMixinFields1 := categoryMixin[1].Fields()
	_ = categoryMixinFields1
	categoryFields := schemas.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryMixinFields1[0].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryMixinFields1[1].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// categoryDescName is the schema descriptor for name field.
	categoryDescName := categoryFields[0].Descriptor()
	// category.NameValidator is a validator for the "name" field. It is called by the builders before save.
	category.NameValidator = categoryDescName.Validators[0].(func(string) error)
	// categoryDescColor is the schema descriptor for color field.
	categoryDescColor := categoryFields[2].Descriptor()
	// category.ColorValidator is a validator for the "color" field. It is called by the builders before save.
	category.ColorValidator = categoryDescColor.Validators[0].(func(string) error)
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryMixinFields0[0].Descriptor()
	// category.DefaultID holds the default value on creation for the id field.
	category.DefaultID = categoryDescID.Default.(func() uuid.UUID)
	debtMixin := schemas.Debt{}.Mixin()
	debtMixinFields0 := debtMixin[0].Fields()
	_ = debtMixinFields0
	debtMixinFields1 := debtMixin[1].Fields()
	_ = debtMixinFields1
	debtFields := schemas.Debt{}.Fields()
	_ = debtFields
	// debtDescCreatedAt is the schema descriptor for created_at field.
	debtDescCreatedAt := debtMixinFields1[0].Descriptor()
	// debt.DefaultCreatedAt holds the default value on creation for the created_at field.
	debt.DefaultCreatedAt = debtDescCreatedAt.Default.(func() time.Time)
	// debtDescUpdatedAt is the schema descriptor for updated_at field.
	debtDescUpdatedAt := debtMixinFields1[1].Descriptor()
	// debt.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	debt.DefaultUpdatedAt = debtDescUpdatedAt.Default.(func() time.Time)
	// debt.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	debt.UpdateDefaultUpdatedAt = debtDescUpdatedAt.UpdateDefault.(func() time.Time)
	// debtDescTitle is the schema descriptor for title field.
	debtDescTitle := debtFields[0].Descriptor()
	// debt.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	debt.TitleValidator = debtDescTitle.Validators[0].(func(string) error)
	// debtDescID is the schema descriptor for id field.
	debtDescID := debtMixinFields0[0].Descriptor()
	// debt.DefaultID holds the default value on creation for the id field.
	debt.DefaultID = debtDescID.Default.(func() uuid.UUID)
	invoiceMixin := schemas.Invoice{}.Mixin()
	invoiceMixinFields0 := invoiceMixin[0].Fields()
	_ = invoiceMixinFields0
	invoiceMixinFields1 := invoiceMixin[1].Fields()
	_ = invoiceMixinFields1
	invoiceFields := schemas.Invoice{}.Fields()
	_ = invoiceFields
	// invoiceDescCreatedAt is the schema descriptor for created_at field.
	invoiceDescCreatedAt := invoiceMixinFields1[0].Descriptor()
	// invoice.DefaultCreatedAt holds the default value on creation for the created_at field.
	invoice.DefaultCreatedAt = invoiceDescCreatedAt.Default.(func() time.Time)
	// invoiceDescUpdatedAt is the schema descriptor for updated_at field.
	invoiceDescUpdatedAt := invoiceMixinFields1[1].Descriptor()
	// invoice.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	invoice.DefaultUpdatedAt = invoiceDescUpdatedAt.Default.(func() time.Time)
	// invoice.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	invoice.UpdateDefaultUpdatedAt = invoiceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// invoiceDescTitle is the schema descriptor for title field.
	invoiceDescTitle := invoiceFields[0].Descriptor()
	// invoice.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	invoice.TitleValidator = invoiceDescTitle.Validators[0].(func(string) error)
	// invoiceDescID is the schema descriptor for id field.
	invoiceDescID := invoiceMixinFields0[0].Descriptor()
	// invoice.DefaultID holds the default value on creation for the id field.
	invoice.DefaultID = invoiceDescID.Default.(func() uuid.UUID)
	paymentstatusMixin := schemas.PaymentStatus{}.Mixin()
	paymentstatusMixinFields0 := paymentstatusMixin[0].Fields()
	_ = paymentstatusMixinFields0
	paymentstatusMixinFields1 := paymentstatusMixin[1].Fields()
	_ = paymentstatusMixinFields1
	paymentstatusFields := schemas.PaymentStatus{}.Fields()
	_ = paymentstatusFields
	// paymentstatusDescCreatedAt is the schema descriptor for created_at field.
	paymentstatusDescCreatedAt := paymentstatusMixinFields1[0].Descriptor()
	// paymentstatus.DefaultCreatedAt holds the default value on creation for the created_at field.
	paymentstatus.DefaultCreatedAt = paymentstatusDescCreatedAt.Default.(func() time.Time)
	// paymentstatusDescUpdatedAt is the schema descriptor for updated_at field.
	paymentstatusDescUpdatedAt := paymentstatusMixinFields1[1].Descriptor()
	// paymentstatus.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	paymentstatus.DefaultUpdatedAt = paymentstatusDescUpdatedAt.Default.(func() time.Time)
	// paymentstatus.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	paymentstatus.UpdateDefaultUpdatedAt = paymentstatusDescUpdatedAt.UpdateDefault.(func() time.Time)
	// paymentstatusDescName is the schema descriptor for name field.
	paymentstatusDescName := paymentstatusFields[0].Descriptor()
	// paymentstatus.NameValidator is a validator for the "name" field. It is called by the builders before save.
	paymentstatus.NameValidator = paymentstatusDescName.Validators[0].(func(string) error)
	// paymentstatusDescID is the schema descriptor for id field.
	paymentstatusDescID := paymentstatusMixinFields0[0].Descriptor()
	// paymentstatus.DefaultID holds the default value on creation for the id field.
	paymentstatus.DefaultID = paymentstatusDescID.Default.(func() uuid.UUID)
}
