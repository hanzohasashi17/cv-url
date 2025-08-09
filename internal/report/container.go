package report

// DI контейнер отдельно
// var ReportRepo = NewRepo()
var ReportService = NewService()
var ReportHandler = NewHandler(ReportService)