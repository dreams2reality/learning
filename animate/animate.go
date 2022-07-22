package animate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	gunASCII = `
 |=#=#=#=#=#=#=#=#\
|=================() %s %s %s %s %s %s %s %s %s %s
|#=#=#=#=#=#=#=#=#/
  |  |)  |
  |  |___/
  |  |
 #=#=#=#=#
 |######|
 | ammo |
 |      |
 |######|
`
	war = `






 V___V_________|          ▄▄▄▄▄▄[███████   
  O-() ULTRA ==O %s %s %s %s %s %s %s %s %s %s ▂▄▅█████████▅▄▃▂      
 |_|===========|/         ███████████████████].
 \| /|-(                   ◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤..
| /|
|/O\\
`
	warDone = `
                                            ***#*##*#
                                           **#SOS*#*##
                                            #**#*##*
                                              #***##
                                              *#**
                                             ***
 V___V_________|                    ▄▄▄▄▄▄[███████
  O-() ULTRA ==O                       ▂▄▅█████████▅▄▃▂
 |_|===========|/                    ███████████████████].
 \| /|-(                     ◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤..
| /|
|/O\\
`
	warPlaneASCII = `%s|\
%s|-\
%s#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#\
%s|         |O===           |__\__\
%s#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=#=/
%s  [|]     ||    \[ | ]/     [|]
%s          ||      (|)===
%s       |===+===)
`
)

func Clear() {
	cmd := exec.Command("Clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to Clear screen: %v", err)
	}
}

func Gun() {
	Clear()
	for i := 0; i < 11; i++ {
		time.Sleep(time.Second)
		Clear()
		switch i {
		case 0:
			fmt.Printf(gunASCII, "", "", "", "", "", "", "", "", "", "")
		case 1:
			fmt.Printf(gunASCII, "*", "", "", "", "", "", "", "", "", "")
		case 2:
			fmt.Printf(gunASCII, "", "*", "", "", "", "", "", "", "", "")
		case 3:
			fmt.Printf(gunASCII, "", "", "*", "", "", "", "", "", "", "")
		case 4:
			fmt.Printf(gunASCII, "", "", "", "*", "", "", "", "", "", "")
		case 5:
			fmt.Printf(gunASCII, "", "", "", "", "*", "", "", "", "", "")
		case 6:
			fmt.Printf(gunASCII, "", "", "", "", "", "*", "", "", "", "")
		case 7:
			fmt.Printf(gunASCII, "", "", "", "", "", "", "*", "", "", "")
		case 8:
			fmt.Printf(gunASCII, "", "", "", "", "", "", "", "*", "", "")
		case 9:
			fmt.Printf(gunASCII, "", "", "", "", "", "", "", "", "*", "")
		case 10:
			fmt.Printf(gunASCII, "", "", "", "", "", "", "", "", "", "*")
		}
	}
}

func WarPlane() {
	Clear()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		Clear()
		s := ""
		for j := 0; j < i; j++ {
			s = s + "  "
		}
		fmt.Printf(warPlaneASCII, s, s, s, s, s, s, s, s)
	}
}

func War() {
	Clear()
	for i := 0; i < 11; i++ {
		time.Sleep(time.Millisecond * 250)
		Clear()
		switch i {
		case 0:
			fmt.Printf(war, " ", "", "", "", "", "", "", "", "", "")
		case 1:
			fmt.Printf(war, "*", "", "", "", "", "", "", "", "", "")
		case 2:
			fmt.Printf(war, "", "*", "", "", "", "", "", "", "", "")
		case 3:
			fmt.Printf(war, "", "", "*", "", "", "", "", "", "", "")
		case 4:
			fmt.Printf(war, "", "", "", "*", "", "", "", "", "", "")
		case 5:
			fmt.Printf(war, "", "", "", "", "*", "", "", "", "", "")
		case 6:
			fmt.Printf(war, "", "", "", "", "", "*", "", "", "", "")
		case 7:
			fmt.Printf(war, "", "", "", "", "", "", "*", "", "", "")
		case 8:
			fmt.Printf(war, "", "", "", "", "", "", "", "*", "", "")
		case 9:
			fmt.Printf(war, "", "", "", "", "", "", "", "", "*", "")
		case 10:
			fmt.Printf(war, "", "", "", "", "", "", "", "", "", "*")
		}
	}
	Clear()
	fmt.Print(warDone)
}
