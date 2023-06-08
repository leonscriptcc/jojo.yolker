package service

import (
	"fmt"
	"github.com/leonscriptcc/jojo.yolker/config"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func WriteV() (err error) {
	// 判断是否有目录文件
	var srcCells [][]string
	if config.CfgParams.WritevConfig.SrcDir.Path != "" {
		srcCells, err = getDirExcels(
			config.CfgParams.WritevConfig.SrcDir.Path,
			config.CfgParams.WritevConfig.SrcDir.Sheet,
			config.CfgParams.WritevConfig.SrcDir.Cells,
		)
		if err != nil {
			return err
		}
	}

	// 写入目标文件
	// 开启文件流
	excel, err := excelize.OpenFile(config.CfgParams.WritevConfig.DestFile.ExmPath)
	if err != nil {
		return err
	}
	defer func() {
		// release file descriptor
		if err := excel.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 插入row
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(config.CfgParams.WritevConfig.DestFile.Cell, -1)
	num, _ := strconv.Atoi(numbers[0])
	character := strings.ReplaceAll(config.CfgParams.WritevConfig.DestFile.Cell, numbers[0], "")
	//excel.SetRowHeight()
	//err = excel.InsertRows(
	//	config.CfgParams.WritevConfig.DestFile.Sheet,
	//	num,
	//	len(srcCells))
	//if err != nil {
	//	return err
	//}

	// 插入数据
	for i := range srcCells {
		excel.SetSheetRow(
			config.CfgParams.WritevConfig.DestFile.Sheet,
			fmt.Sprintf("%s%d", character, num),
			&srcCells[i],
		)
		num++
	}

	excel.SaveAs(config.CfgParams.WritevConfig.DestFile.DestPath)
	return err
}

// getDirExcels 从文件夹中遍历excel文件
func getDirExcels(path, sheet string, cells []string) (
	srcCells [][]string, err error) {
	// 读取文件见中的所有excel
	files, err := os.ReadDir(path)
	if err != nil {
		return srcCells, err
	}
	// 读取文件中信息
	for _, file := range files {
		// 开启文件流
		excel, err := excelize.OpenFile(path + file.Name())
		if err != nil {
			log.Println("open file ", file.Name(), "fail:", err)
			continue
		}

		// 按照配置循环读取cell
		var (
			v   string
			scs []string
		)
		for _, c := range cells {
			if c == "none" {
				scs = append(scs, "v")
				continue
			}
			v, err = excel.GetCellValue(sheet, c)
			if err != nil {
				continue
			}
			scs = append(scs, v)
		}

		// 加入结果集
		srcCells = append(srcCells, scs)

		// release fd
		if err := excel.Close(); err != nil {
			log.Println(err)
		}
	}

	return srcCells, err
}
