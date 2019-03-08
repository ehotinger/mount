package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// tail -f /var/log/kern.log
	// http://man7.org/linux/man-pages/man2/mount.2.html

	mountLocation := ""

	exists, err := mountExists(mountLocation)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("mount exists:", exists)

}

func mountExists(mount string) (bool, error) {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return false, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		if err := s.Err(); err != nil {
			return false, err
		}
		text := s.Text()
		if strings.Contains(text, mount) {
			return true, nil
		}
	}

	return false, nil
}
