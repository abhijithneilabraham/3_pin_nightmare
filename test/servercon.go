func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}

			temp := strings.TrimSpace(string(netData))
			if temp == "STOP" {
					break
			}

			result := strconv.Itoa(random()) + "\n"
			c.Write([]byte(string(result)))
	}
	c.Close()
}