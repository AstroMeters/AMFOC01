package sdcard

const (
	CMD0_GO_IDLE_STATE              = 0
	CMD1_SEND_OP_CND                = 1
	CMD2_ALL_SEND_CID               = 2
	CMD3_SEND_RELATIVE_ADDR         = 3
	CMD4_SET_DSR                    = 4
	CMD6_SWITCH_FUNC                = 6
	CMD7_SELECT_DESELECT_CARD       = 7
	CMD8_SEND_IF_COND               = 8
	CMD9_SEND_CSD                   = 9
	CMD10_SEND_CID                  = 10
	CMD12_STOP_TRANSMISSION         = 12
	CMD13_SEND_STATUS               = 13
	CMD15_GO_INACTIVE_STATE         = 15
	CMD16_SET_BLOCKLEN              = 16
	CMD17_READ_SINGLE_BLOCK         = 17
	CMD18_READ_MULTIPLE_BLOCK       = 18
	CMD24_WRITE_BLOCK               = 24
	CMD25_WRITE_MULTIPLE_BLOCK      = 25
	CMD27_PROGRAM_CSD               = 27
	CMD28_SET_WRITE_PROT            = 28
	CMD29_CLR_WRITE_PROT            = 29
	CMD30_SEND_WRITE_PROT           = 30
	CMD32_ERASE_WR_BLK_START_ADDR   = 32
	CMD33_ERASE_WR_BLK_END_ADDR     = 33
	CMD38_ERASE                     = 38
	CMD42_LOCK_UNLOCK               = 42
	CMD55_APP_CMD                   = 55
	CMD56_GEN_CMD                   = 56
	CMD58_READ_OCR                  = 58
	CMD59_CRC_ON_OFF                = 59
	ACMD6_SET_BUS_WIDTH             = 6
	ACMD13_SD_STATUS                = 13
	ACMD22_SEND_NUM_WR_BLOCKS       = 22
	ACMD23_SET_WR_BLK_ERASE_COUNT   = 23
	ACMD41_SD_APP_OP_COND           = 41
	ACMD42_SET_CLR_CARD_DETECT      = 42
	ACMD51_SEND_SCR                 = 51
	ACMD18_SECURE_READ_MULTI_BLOCK  = 18
	ACMD25_SECURE_WRITE_MULTI_BLOCK = 25
	ACMD26_SECURE_WRITE_MKB         = 26
	ACMD38_SECURE_ERASE             = 38
	ACMD43_GET_MKB                  = 43
	ACMD44_GET_MID                  = 44
	ACMD45_SET_CER_RN1              = 45
	ACMD46_SET_CER_RN2              = 46
	ACMD47_SET_CER_RES2             = 47
	ACMD48_SET_CER_RES1             = 48
	ACMD49_CHANGE_SECURE_AREA       = 49
)
