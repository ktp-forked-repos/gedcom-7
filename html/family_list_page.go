package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type FamilyListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewFamilyListPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *FamilyListPage {
	return &FamilyListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *FamilyListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{
		core.NewTableHead("Husband", "Date", "Wife"),
	}

	for _, family := range c.document.Families() {
		familyInList := NewFamilyInList(c.document, family, c.options.LivingVisibility)
		table = append(table, familyInList)
	}

	column := core.NewColumn(core.EntireRow, core.NewTable("", table...))
	header := NewPublishHeader(c.document, "", selectedFamiliesTab, c.options, c.indexLetters)
	components := core.NewComponents(header, core.NewRow(column))

	return core.NewPage("Families", components, c.googleAnalyticsID).
		WriteHTMLTo(w)
}
