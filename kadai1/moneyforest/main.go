package main

import(
    "bufio"
    "fmt"
    "path/filepath"
    "os"
    "image"
    _ "image/jpeg"
    "image/jpeg"
    "image/png"
) 

func main() {
    fmt.Println("指定したディレクトリ以下の画像ファイルの拡張子を変換します")

    fmt.Println("ディレクトリを指定して下さい")
    dir := scanDirStr()

    fmt.Println("変換元となる拡張子を入力して下さい")
    fromExt := scanImgExt()

    fmt.Println("変換先となる拡張子を入力して下さい")
    toExt := scanImgExt()

    imgList := imgList(dir, fromExt)
    convertImg(imgList, toExt)
}

// ディレクトリ
func scanDirStr() (dir string) {
    sc := bufio.NewScanner(os.Stdin)

    if sc.Scan() {
        dir = sc.Text()
    }

    if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
        panic("ディレクトリが存在しません")
    }

    return
}

// 画像形式のファイルのリスト
func imgList(dir, pickUpFormat string) (imgList []string) {
    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		
        if err != nil {
			return err
		}

        f, err := os.Open(path)
        defer f.Close()

        _, format, err := image.DecodeConfig(f)

        if err == nil && format == pickUpFormat {
            imgList = append(imgList, f.Name())
        }

		return nil
	})

	if err != nil {
		panic(err)
	}

    return
}

// 拡張子
func scanImgExt() (ext string) {
    sc := bufio.NewScanner(os.Stdin)

    if sc.Scan() {
        ext = sc.Text()
    }

    if ext != "png" && ext != "jpeg" {
        panic("拡張子が不正です")
    }

    return
}

// 画像ファイルの変換
func convertImg(imgList []string, ext string) {
    for _, img := range imgList {
        f, err := os.Open(img)
        defer f.Close()

        if err != nil {
            panic(err)
        }

        fName, _ := os.Create(convertExt(f.Name(), "." + ext))
        defer fName.Close()

        img, _, _ :=  image.Decode(f)

        switch ext {
            case "png" :
                png.Encode(fName, img)
            case "jpeg" :
                jpeg.Encode(fName, img, nil)
        }
    }
}

func convertExt(fileName, ext string) string {
    fBaseName := fileName[:len(fileName) - len(filepath.Ext(fileName))]
    return fBaseName + ext
}
