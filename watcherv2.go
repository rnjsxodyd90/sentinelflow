package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
	"os"
)

// function that runs python
func runPythonWorker(filePath string) {
	fmt.Printf("[시작]%s 처리 중.\n", filePath)
	cmd := exec.Command("python", "worker.py", filePath)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Printf("[에러] %s 처리 실패: %v\n", filePath, err)
		return
	}

	// 처리가 끝나면 processed 폴더로 이동 (없으면 생성)
	os.Mkdir("processed", 0755)
	newName := filepath.Join("processed", filepath.Base(filePath))
	os.Rename(filePath, newName)
	
	fmt.Printf("[완료] %s 처리가 끝났습니다.\n", filePath)
}

func main() {
	inbox := "./inbox"
	fmt.Println("--- SentinelFlow V2: 병렬 감시 시작 ---")

	for {
		// 1. inbox 폴더 안의 모든 파일을 읽습니다
		files, _ := filepath.Glob(filepath.Join(inbox, "*.csv"))

		for _, file := range files {
			// 2.  'go' 키워드를 붙여서 별도의 스레드(고루틴)로 실행합니다.
			// 이제 이 파일이 처리되는 동안 메인 루프는 기다리지 않고 다음 파일로 넘어갑니다.
			go runPythonWorker(file)
		}

		time.Sleep(2 * time.Second)
	}
}
