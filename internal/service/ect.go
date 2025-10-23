package service

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/Alencar26/ECT-Generator/pkg/pdf"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

//go:embed assets/logo.png
var logo []byte

type ECTService struct {
	pdf           core.Maroto
	name          string
	pathEvidences []string
	logo          []byte
}

func NewECTService(name string) *ECTService {
	return &ECTService{
		pdf:  pdf.GetMaroto(),
		name: name,
		logo: logo,
	}
}

func (e *ECTService) SetEvidences(pathEvidences []string) {
	e.pathEvidences = pathEvidences
}

func (e *ECTService) GetCover(author string) {
	e.pdf.AddRow(80,
		text.NewCol(12, "", props.Text{Top: 100}),
	)
	e.pdf.AddRow(50,
		image.NewFromBytesCol(12, e.logo, extension.Png, props.Rect{Center: true}),
	)
	e.pdf.AddRow(15,
		text.NewCol(12, "ECT - Evidência de Cenário de Teste", props.Text{Align: align.Center, Size: 20}),
	)
	e.pdf.AddRow(10,
		text.NewCol(12, fmt.Sprintf("Autor: %s", author), props.Text{Align: align.Center, Size: 12}),
	)
	e.pdf.AddRow(10,
		text.NewCol(12, fmt.Sprintf("Data: %s", time.Now().Format("02/01/2006")), props.Text{Align: align.Center, Size: 12}),
	)
	e.pdf.AddRow(85) //blank space
}

func (e *ECTService) GenerateBody() {
	for _, img := range e.pathEvidences {
		e.pdf.AddRow(120,
			image.NewFromFileCol(12, img, props.Rect{Center: true, JustReferenceWidth: true}),
		)
	}
}

func (e *ECTService) SavePDF() {
	document, err := e.pdf.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = document.Save(fmt.Sprintf("%s.pdf", e.name))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("PDF '%s.pdf' gerado com sucesso!", e.name)
}
