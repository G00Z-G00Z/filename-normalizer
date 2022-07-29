package caseformaters

var caseFormattersMetaData caseFormatterMetaDataMap = caseFormatterMetaDataMap{
	SnakeCase: {
		Name:           "snake-case",
		Case2snakeCase: ToSnakeCase,
		SnakeCase2Case: func(s string) string { return s },
		IsCase:         isSnakeCase,
	},
	CammelCase: {
		Name:           "cammel-case",
		Case2snakeCase: ToCammelCase,
		SnakeCase2Case: snakeCase2CammelCase,
		IsCase:         isCammelCase,
	},
	SpacesCase: {
		Name:           "spaces-case",
		Case2snakeCase: ToSpaces,
		SnakeCase2Case: spaces2SnakeCase,
		IsCase:         isSpaces,
	},
}
