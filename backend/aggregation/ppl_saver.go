package aggreation

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lagrange92/Haechi/model"
)

// savePpl : save ppl data to file
func savePpl(ppls []model.PpltnData) {
	// 파일 생성 및 열기
	filename := "./current_people_" + time.Now().Format("060102_1504") + ".txt" // ex) current_people_YYMMDD_HHmm.txt
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}

	defer file.Close()

	for idx, ppl := range ppls {
		_, err := fmt.Fprintf(file, "[%d] %s %s %s %s %s\n", idx+1, ppl.AreaName, ppl.AreaCode, ppl.AreaLatitude, ppl.AreaLongitude, ppl.AreaAvgPpltn)
		if err != nil {
			log.Fatal("Error writing to file: ", err)
		}
	}

	fmt.Println("Data saved to ", filename)
}
