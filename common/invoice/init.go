package invoice

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	"github.com/fogleman/gg"
	"github.com/spf13/viper"
)

type InvoiceManager struct {
	Config Config
}

func NewInvoiceManager(invoicePath string) *InvoiceManager {
	var config Config
	viper.SetConfigFile("json")
	viper.AddConfigPath(invoicePath)
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	textName := []string{
		"date", "team_name", "leader_name",
		"member1_name", "member2_name", "member3_name", "member4_name", "member5_name",
		"competition_name",
		"competition_price",
		"competition_total_price",
		"total_payment",
	}
	textProperties := make([]Properties, 12)
	for i := 0; i < len(textProperties); i++ {
		textProperties[i] = Properties{
			FontPath:  viper.GetString("config." + textName[i] + ".font_path"),
			FontSize:  viper.GetFloat64("config." + textName[i] + ".font_size"),
			PositionX: viper.GetFloat64("config." + textName[i] + ".position_x"),
			PositionY: viper.GetFloat64("config." + textName[i] + ".position_y"),
		}
		rgbColors := strings.Split(viper.GetString("config."+textName[i]+".color_rgb"), ",")
		textProperties[i].Color = color.RGBA{
			R: StringToUint8(rgbColors[0]),
			G: StringToUint8(rgbColors[1]),
			B: StringToUint8(rgbColors[2]),
			A: 255,
		}
	}
	config = Config{
		EventName:      viper.GetString("config.event_name"),
		BackgroundPath: viper.GetString("config.background_img"),
		TextProperties: textProperties,
	}

	return &InvoiceManager{Config: config}
}

func (i InvoiceManager) CreateInvoice(team teamDomain.Team) (err error, filePath string, fileName string) {
	invoiceDetails := CreateInvoiceDetails(team)
	img, err := i.writeTextOnImage(invoiceDetails, i.Config)
	if err != nil {
		return err, filePath, fileName
	}
	fileName = fmt.Sprintf("%s_%s.jpg", team.Name, "Invoice")
	filePath = fmt.Sprintf("./out/%s", fileName)
	err = gg.SaveJPG(filePath, img, 100)
	if err != nil {
		return err, filePath, fileName
	}
	return err, filePath, fileName
}

func (i InvoiceManager) writeTextOnImage(details InvoiceDetails, config Config) (image.Image, error) {
	loadImage, err := gg.LoadImage(config.BackgroundPath)
	if err != nil {
		return nil, err
	}

	// Get width and height from loaded image
	imgWidth, imgHeight := loadImage.Bounds().Dx(), loadImage.Bounds().Dy()
	// Create new canvas and put loaded image on canvas
	var canvas = gg.NewContext(imgWidth, imgHeight)
	canvas.DrawImage(loadImage, 0, 0)
	latestImage := canvas.Image()
	text := InvoiceDetails.MapToArray(details)
	for i := 0; i < len(config.TextProperties); i++ {
		canvas = gg.NewContext(imgWidth, imgHeight)
		canvas.DrawImage(latestImage, 0, 0)

		// Decrease font size 15% if name character greater than 25 characters
		fontSize := config.TextProperties[i].FontSize
		// if len(person.Name) > 25 {
		// 	fontSize -= math.Round(fontSize * 0.15)
		// }

		err = canvas.LoadFontFace(config.TextProperties[i].FontPath, fontSize)
		if err != nil {
			return nil, err
		}

		canvas.SetColor(config.TextProperties[i].Color)
		canvas.DrawString(text[i], config.TextProperties[i].PositionX, config.TextProperties[i].PositionY)
		latestImage = canvas.Image()
	}

	// Load font from path and set font size

	// Set maximal width of text and color

	// Write text on image
	// canvas.DrawStringWrapped(person.Name, config.Name.PositionX, config.Name.PositionY, 0.5, 0.5,
	// 	maxWidth, 1.3, gg.AlignLeft)

	// if config for code is set, then write text again
	// if config.Code.FontPath != "" {
	// 	canvas2 := gg.NewContext(imgWidth, imgHeight)
	// 	canvas2.DrawImage(canvas.Image(), 0, 0)

	// 	err = canvas2.LoadFontFace(config.Code.FontPath, config.Code.FontSize)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	canvas2.SetColor(config.Code.Color)
	// 	canvas2.DrawStringWrapped(person.Code, config.Code.PositionX, config.Code.PositionY, 0.5, 0.5,
	// 		maxWidth, 1.3, gg.AlignCenter)

	// 	return canvas2.Image(), nil
	// } else {
	// 	return canvas.Image(), nil
	// }
	return latestImage, nil
}
