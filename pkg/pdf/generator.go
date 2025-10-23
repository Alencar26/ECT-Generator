package pdf

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GetMaroto() core.Maroto {

	pageNumber := ConfigPageNumber(0, 0, 0)
	configBuilder := ConfigBuilder(
		"E2E Test Evidence",
		"E2E Test Evidence Generator",
		"E2E Test Evidence",
		pageNumber,
	)

	mrt := maroto.New(configBuilder)
	m := maroto.NewMetricsDecorator(mrt)
	return configFooter(m)

}

func configFooter(m core.Maroto) core.Maroto {
	m.RegisterFooter(
		m.AddRow(5,
			text.NewCol(12, "E2E Test Evidence | Software XPTO", props.Text{Align: align.Center}),
		),
		m.AddRow(5,
			text.NewCol(12, "All rights reserved.", props.Text{Align: align.Center}),
		),
	)
	return m
}
