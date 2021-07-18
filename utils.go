package main

func isACK(buf []byte) bool {
	return buf[0] == ACK_FORMAT[0] &&
		buf[1] == ACK_FORMAT[1] &&
		buf[2] == ACK_FORMAT[2]
}

func isREQ(buf []byte) bool {
	return buf[0] == REQ_FORMAT[0] &&
		buf[1] == REQ_FORMAT[1] &&
		buf[2] == REQ_FORMAT[2]
}
