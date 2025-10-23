package pdf

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func ConfigBuilder(author, CreatedBy, title string, pageNumber props.PageNumber) *entity.Config {
	return config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithAuthor(author, true).
		WithCreator(CreatedBy, true).
		WithTitle(title, true).
		WithPageNumber(pageNumber).
		WithLeftMargin(20).
		WithRightMargin(20).
		WithTopMargin(10).
		WithBottomMargin(10).
		Build()
}

func ConfigPageNumber(r, g, b int) props.PageNumber {
	return props.PageNumber{
		Pattern: "{current}/{total}",
		Place:   props.Bottom,
		Family:  fontfamily.Courier,
		Style:   fontstyle.Normal,
		Size:    0,
		Color: &props.Color{
			Red:   r,
			Green: g,
			Blue:  b,
		},
	}
}
