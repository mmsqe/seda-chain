package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// sedad config set client chain-id test
	// sedad config set client keyring-backend test
	// sedad config set client broadcast-mode sync
	// sedad init test --chain-id test --overwrite
	// sedad keys add validator --keyring-backend test
	// sedad add-genesis-account validator 1500000000000000000000stake,1500000000000000000000aseda --keyring-backend test
	// sedad gentx validator 1000000000000000000000stake --chain-id test --moniker="validator" --min-self-delegation="1000000" --details dh4S5hSmReBoE1VHYrOVlDT3Ek1ZIIg8kZynpmM3Teg= --ip="127.0.0.1" --keyring-backend test
	// sedad collect-gentxs
	// sedad start

	// genesis.json: max_gas "81500000"
	// app.toml: pruning = "everything"
	// config.toml timeout_commit = "1.5s"

	errMsg := "Value missing for key"
	home := "/Users/mavis/.sedad"
	for i := 0; i < 1; i++ {
		err := os.RemoveAll(home)
		if err != nil {
			fmt.Println("err when remove sedad:", err)
			continue
		}

		err = os.Mkdir(home, 0777)
		if err != nil {
			fmt.Println("err when make sedad:", err)
			continue
		}

		cmd := exec.Command("cp", "-R", "/Users/mavis/Desktop/sedad/", home)
		fmt.Println("mm-cp")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("err when cp:", err)
			continue
		}

		l := "cmd.log"
		if _, err := os.Stat(l); err == nil {
			err = os.Remove(l)
			if err != nil {
				panic(err)
			}
		}

		file, err := os.OpenFile(l, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error when open file:", err)
			continue
		}

		fmt.Println("mm-start1")
		sedadCmd := exec.Command("sedad", "start")
		go sedadCmd.Run()
		time.Sleep(time.Second * 3)
		fmt.Println("mm-kill1")
		sedadCmd.Process.Kill()
		sedadCmd.Process.Wait()

		for i := 0; i < 15; i++ {
			fmt.Println("mm-start2")
			sedadCmd = exec.Command("sedad", "start")
			sedadCmd.Stdout = file
			sedadCmd.Stderr = os.Stderr
			go sedadCmd.Run()
			time.Sleep(time.Second * 3)

			cmd := exec.Command("sedad", "tx", "data-proxy", "register", "seda14xdjsea6w9wwfm7sgatp56pkspklgfcjwk7eus", "200000000000000aseda", "03f8624d228ee68a7cd1fc116a85cc43ff45a0ad07738320b32aeb3452687b2e1e", "86f9528f3107972c8cd85dc03e2b1343be50ac473b37cb15e71be1fc6796fdd308ff8735ee3f96144aa927e1fa869abf59132431d85c0068cb467ba0bdd89b74", "--from", "validator", "--keyring-backend", "test", "--gas", "auto", "--gas-adjustment", "1.5", "--gas-prices", "10000000000aseda", "--chain-id", "test", "-y")
			fmt.Println("mm-tx", i)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("mm-tx-err", i, err)

				fmt.Println("killing proc")
				sedadCmd.Process.Kill()
				fmt.Println("killed proc")
				sedadCmd.Process.Wait()
				fmt.Println("proc end")
				continue
			}

			errCount := 0
			maxErrCount := 2
			for {
				time.Sleep(time.Second * 1)
				cmd = exec.Command("sedad", "query", "data-proxy", "data-proxy-config", "03f8624d228ee68a7cd1fc116a85cc43ff45a0ad07738320b32aeb3452687b2e1e")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("mm-query-err", errCount)
					if errCount += 1; errCount < maxErrCount {
						continue
					}
				} else {
					fmt.Println("mm-query-success")
					time.Sleep(time.Second * 2)

					fmt.Println("killing proc")
					sedadCmd.Process.Kill()
					fmt.Println("killed proc")
					sedadCmd.Process.Wait()
					fmt.Println("proc end")
				}
				break
			}

			if errCount >= maxErrCount {
				fmt.Println("Error found stop")
				break
			}
		}
		fmt.Println("killing proc")
		sedadCmd.Process.Kill()
		fmt.Println("killed proc")
		sedadCmd.Process.Wait()
		fmt.Println("proc end")
		file.Close()
		fmt.Println("log closed")

		log, err := os.ReadFile(l)
		if err != nil {
			fmt.Println("read log err:", err)
		} else {
			fmt.Println("checking log")
			if strings.Contains(string(log), errMsg) {
				fmt.Println("mm-found")
				break
			}
		}
	}
}
