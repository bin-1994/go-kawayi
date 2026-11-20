package service

type WebpService struct {
}

const WEBP_PATH = "./uploads/webp"
const IMG_PATH = "./uploads"

func (s *WebpService) ToWebp(path string) error {
	//str := strings.Split(path, ".")
	//newPath := str[0] + ".webp"
	//ext := str[1]
	//
	//filename := WEBP_PATH + newPath
	//data, err := os.ReadFile(IMG_PATH + path)
	//if err != nil {
	//	return errors.New("文件不存在")
	//}
	//var buffer bytes.Buffer
	//var img image.Image
	//if ext == "jpg" {
	//	img, _ = jpeg.Decode(bytes.NewReader(data))
	//} else if ext == "png" {
	//	img, _ = png.Decode(bytes.NewReader(data))
	//}
	//
	//if err = webp.Encode(&buffer, img, &webp.Options{Quality: 80}); err != nil {
	//	return errors.New("创建webp失败")
	//}
	//
	//os.WriteFile(filename, buffer.Bytes(), 0644)

	return nil
}
