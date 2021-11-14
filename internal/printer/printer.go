package printer

import (
	"context"
	"fmt"
	"log"

	"github.com/iskorotkov/spaceship-engine-production/api/report-printer"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const floatPrecision = 2

var _ = (report_printer.ReportPrinterServer)(Printer{}) //nolint:exhaustivestruct

type Printer struct {
	report_printer.UnimplementedReportPrinterServer
}

func New() *Printer {
	log.Printf("printer created")

	return &Printer{} //nolint:exhaustivestruct
}

func (s Printer) Print(_ context.Context, req *report_printer.PrintRequest) (*report_printer.PrintResponse, error) {
	sheetName := "Report"

	f := excelize.NewFile()
	sheet := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	const nameWidth = 40
	_ = f.SetColWidth(sheetName, "A", "A", nameWidth)

	const valueWidth = 15
	_ = f.SetColWidth(sheetName, "B", "C", valueWidth)

	_ = f.SetCellStr(sheetName, "A1", "Total orders:")
	_ = f.SetCellInt(sheetName, "B1", int(req.TotalOrders))

	_ = f.MergeCell(sheetName, "A3", "B3")
	_ = f.SetCellStr(sheetName, "A3", "Top clients")
	_ = f.SetCellStr(sheetName, "A4", "Name")
	_ = f.SetCellStr(sheetName, "B4", "Total orders")

	row := 5

	for _, client := range req.TopClients {
		_ = f.SetCellStr(sheetName, fmt.Sprintf("A%d", row), client.Name)
		_ = f.SetCellInt(sheetName, fmt.Sprintf("B%d", row), int(client.Orders))
		row++
	}

	row++

	_ = f.MergeCell(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("C%d", row))
	_ = f.SetCellStr(sheetName, fmt.Sprintf("A%d", row), "Top engine models")

	row++

	_ = f.SetCellStr(sheetName, fmt.Sprintf("A%d", row), "Name")
	_ = f.SetCellStr(sheetName, fmt.Sprintf("B%d", row), "Power")
	_ = f.SetCellStr(sheetName, fmt.Sprintf("C%d", row), "Total orders")

	row++

	for _, engine := range req.TopEngines {
		_ = f.SetCellStr(sheetName, fmt.Sprintf("A%d", row), engine.Name)
		_ = f.SetCellFloat(sheetName, fmt.Sprintf("B%d", row), engine.Power, floatPrecision, 64) //nolint:gomnd
		_ = f.SetCellInt(sheetName, fmt.Sprintf("C%d", row), int(engine.Orders))
		row++
	}

	if err := f.SaveAs(req.Filepath); err != nil {
		return nil, status.Errorf(codes.Internal, "error saving file: %v", err)
	}

	log.Printf("printer wrote file %q", req.Filepath)

	return &report_printer.PrintResponse{}, nil
}
