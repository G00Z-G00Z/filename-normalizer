package caseformaters

var caseFormattersMetaData caseFormatterMetaDataMap = caseFormatterMetaDataMap{
	SnakeCase: {
		Name:           "snake-case",
		Case2snakeCase: ToSnakeCase,
		IsCase:         isSnakeCase,
	},
	CammelCase: {
		Name:           "cammel-case",
		Case2snakeCase: ToCammelCase,
		IsCase:         isCammelCase,
	},
	SpacesCase: {
		Name:           "spaces-case",
		Case2snakeCase: ToSpaces,
		IsCase:         isSpaces,
	},
}
