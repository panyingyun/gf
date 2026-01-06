package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var filePattern string
	var filenamePattern string
	var targetDir string

	// 定义命令行参数
	flag.StringVar(&filePattern, "f", "", "文件内容搜索模式，后跟搜索字符串")
	flag.StringVar(&filenamePattern, "g", "", "文件名搜索模式，后跟搜索字符串")

	flag.Parse()

	// 检查参数
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "用法: gf -f \"pattern\" <目录> 或 gf -g \"pattern\" <目录>\n")
		os.Exit(1)
	}

	targetDir = args[0]

	// 确定搜索模式
	var searchPattern string
	var searchMode string

	if filePattern != "" {
		searchPattern = filePattern
		searchMode = "file"
	} else if filenamePattern != "" {
		searchPattern = filenamePattern
		searchMode = "filename"
	} else {
		fmt.Fprintf(os.Stderr, "错误: 必须指定 -f 或 -g 参数\n")
		os.Exit(1)
	}

	// 验证目录是否存在
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "错误: 目录不存在: %s\n", targetDir)
		os.Exit(1)
	}

	// 根据模式执行搜索
	switch searchMode {
	case "file":
		searchInFiles(searchPattern, targetDir)
	case "filename":
		searchInFilenames(searchPattern, targetDir)
	}
}

// searchInFiles 在文件内容中搜索指定字符串
func searchInFiles(pattern, rootDir string) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// 跳过无法访问的文件/目录
			return nil
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			// 跳过无法打开的文件
			return nil
		}
		defer file.Close()

		// 逐行读取文件
		scanner := bufio.NewScanner(file)
		lineNum := 0

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				fmt.Printf("%s:%d:%s\n", path, lineNum, line)
			}
		}

		if err := scanner.Err(); err != nil {
			// 读取错误，跳过该文件
			return nil
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: 遍历目录时发生错误: %v\n", err)
		os.Exit(1)
	}
}

// searchInFilenames 在文件名中搜索指定字符串
func searchInFilenames(pattern, rootDir string) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// 跳过无法访问的文件/目录
			return nil
		}

		// 检查文件名（不含路径）是否包含搜索字符串
		if strings.Contains(info.Name(), pattern) {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: 遍历目录时发生错误: %v\n", err)
		os.Exit(1)
	}
}

