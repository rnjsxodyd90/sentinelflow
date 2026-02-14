package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// 1. 감시할 파일 이름을 정합니다
	targetFile := "data.csv"
	fmt.Println("--- SentinelFlow 감시 시작 (Go) ---")

	for {
		// 2.	 파일이 존재하는지 확인합니다
		_, err := os.Stat(targetFile)

		if err == nil {
			// 파일이 발견되었다면!
			fmt.Printf("\n[이벤트 발생] '%s' 파일이 감지되었습니다!\n", targetFile)

			// 3. 파이썬 워커를 실행합니다
			fmt.Println("파이썬 데이터 프로세서를 호출합니다...")
			cmd := exec.Command("python", "worker.py")
			
			// 파이썬이 출력하는 내용을 화면에 보여줍니다
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			
			err := cmd.Run()
			if err != nil {
				fmt.Println("에러 발생:", err)
			}

			// 4. 처리가 끝났으니 파일을 삭제하거나 이름을 바꿉니다 (반복 방지)
			os.Rename(targetFile, "processed_data.csv")
			fmt.Println("--- 처리 완료 및 대기 중 ---")
		}

		// 5. 2초마다 한 번씩 확인합니다 (시스템 부하 방지)
		time.Sleep(2 * time.Second)
	}
}