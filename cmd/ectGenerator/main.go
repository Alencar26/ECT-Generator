package main

import (
	"github.com/Alencar26/ECT-Generator/internal/service"
	"github.com/Alencar26/ECT-Generator/pkg/files"
)

func main() {
	path := "assets/evidences"
	evidences := files.GetEvidences(path)

	ect := service.NewECTService("ECT_TS0000")
	ect.SetEvidences(evidences)
	ect.GetCover("Tester Name")
	ect.GenerateBody()
	ect.SavePDF()
}

// func main() {
// 	ectFile, err := files.NewECTFile()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		fmt.Println("Renamed the file. Press ENTER to exit.")
// 		fmt.Scanln()
// 	}

// 	fmt.Println("Generating PDF...")

// 	ect := service.NewECTService(fmt.Sprintf("%s_%s", ectFile.ECTType, ectFile.TestSuite))
// 	ect.SetEvidences(ectFile.Evicences)
// 	ect.GetCover(ectFile.Author)
// 	ect.GenerateBody()
// 	ect.SavePDF()

// 	fmt.Println("PDF generated successfully! Press ENTER to exit.")
// 	fmt.Scanln()
// }
