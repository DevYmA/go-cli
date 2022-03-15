package regional

func GetRegionalGiftMessage(region string) string {
	message := ""

	switch region {
	case "Asia":
		message = "We added asian culture gift to your pack"
	case "Europe":
		message = "We added Europe culture gift to your pack"
	case "EME":
		message = "We added EME culture gift to your pack"
	default:

		message = "Enterd region not eligible for culural gift"
	}
	return message
}
