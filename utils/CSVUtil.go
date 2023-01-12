package util

import (
	"encoding/csv"
	"fmt"
	errors "github.com/pkg/errors"
	"os"
)

func SimpleWriteCSV(filepath string, data [][]string) error {

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if os.IsPermission(err) {
		return errors.Errorf("权限不足, 打开CSV文件[%s]异常", filepath)
	}
	if err != nil {
		fmt.Printf("打开CSV文件[%s]异常, [%s}\n", filepath, err)
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, dataItem := range data {
		if err = csvWriter.Write(dataItem); err != nil {
			fmt.Printf("数据写入文件[%s]异常, [%s]\n", filepath, err)
			return errors.Errorf("数据写入文件[%s]异常", filepath)
		}
	}

	csvWriter.Flush()

	return nil
}

// WriteCSVMap 将指定数据写入文件中，可自定义结构体的转换
func WriteCSVMap[T any](filepath string, data []T, mapFn func(t T, i int) []string) error {

	s := make([][]string, len(data))
	for index, datum := range data {
		s[index] = mapFn(datum, index)
	}

	return SimpleWriteCSV(filepath, s)
}

// SimpleReadCSV 基础读取CSV文件
func SimpleReadCSV(filepath string) ([][]string, error) {

	// 判断是否存在
	if FileIsNotExist(filepath) {
		return nil, errors.Errorf("文件[%s]是不存在", filepath)
	}

	// 打开文件
	file, err := os.Open(filepath)

	// 无法访问文件
	if os.IsPermission(err) {
		return nil, errors.Errorf("无法访问文件[%s]", filepath)
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件
	csvReader := csv.NewReader(file)
	csvCont, err := csvReader.ReadAll()
	return csvCont, err
}

// ReadCSVMap 读取CSV文件，并调用处理函数，将结果置换成结构体
func ReadCSVMap[R any](filepath string, mapFn func([]string, int) R) ([]R, error) {

	csvCont, err := SimpleReadCSV(filepath)

	if err != nil {
		return nil, err
	}

	result := make([]R, len(csvCont))
	for i, row := range csvCont {
		result[i] = mapFn(row, i)
	}

	return result, nil
}

// ReadCSVFlatMap 读取CSV文件，将处理函数中返回的数组扁平化处理
func ReadCSVFlatMap[R any](filepath string, mapFn func([]string, int) []R) ([]R, error) {

	csvCont, err := SimpleReadCSV(filepath)

	if err != nil {
		return nil, err
	}

	var result []R
	for i, row := range csvCont {
		result = append(result, mapFn(row, i)...)
	}

	return result, nil
}
