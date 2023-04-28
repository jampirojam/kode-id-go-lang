package constant

import "time"

const (
	PORT_SERVER = "localhost:8188"
)

// Security Constant
const (
	API_SECRET = "6CEA9568-42DD-4E50-880E-C1FF4566F8EE"
	KEY_SECRET = "k4t/\nyA_!nd0nes1@"
)

var TOKEN_LIFESPAN = time.Now().Local().UTC().Add(time.Hour * 8).Unix()
var REFRESH_TOKEN_LIFESPAN = time.Now().Local().UTC().Add(time.Hour * 720).Unix()
