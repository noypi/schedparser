package gen

import "fmt"

func (this *SchedLex) Lex(lval *SchedSymType) (ret int) { // begin main()

	var err error
	var c byte = ' '

	defer func() {
		this.back()
	}()

yystate0:

	if nil != this.buf {
		this.buf = this.buf[len(this.buf)-1:]
	}

	goto yystart1

	goto yystate1 // silence unused label error
yystate1:
	c = this.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == ',':
		goto yystate5
	case c == '0':
		goto yystate6
	case c == '1':
		goto yystate23
	case c == '2':
		goto yystate25
	case c == '3':
		goto yystate26
	case c == '\n':
		goto yystate4
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c == 'a':
		goto yystate27
	case c == 'd':
		goto yystate36
	case c == 'e':
		goto yystate43
	case c == 'f':
		goto yystate48
	case c == 'h':
		goto yystate62
	case c == 'j':
		goto yystate67
	case c == 'm':
		goto yystate73
	case c == 'n':
		goto yystate84
	case c == 'o':
		goto yystate86
	case c == 's':
		goto yystate90
	case c == 't':
		goto yystate102
	case c == 'w':
		goto yystate109
	case c >= '4' && c <= '9':
		goto yystate8
	}

yystate2:
	c = this.next()
	goto yyrule16

yystate3:
	c = this.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	}

yystate4:
	c = this.next()
	switch {
	default:
		goto yyrule1
	case c == '\n':
		goto yystate4
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate3
	}

yystate5:
	c = this.next()
	goto yyrule15

yystate6:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate7
	case c == '1':
		goto yystate20
	case c == '2':
		goto yystate21
	case c == '3':
		goto yystate22
	}

yystate7:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == ':':
		goto yystate17
	case c == 't':
		goto yystate11
	}

yystate8:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == 't':
		goto yystate11
	}

yystate9:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '2' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate10
	case c == 's':
		goto yystate13
	case c == 't':
		goto yystate11
	}

yystate10:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '2' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate10
	case c == 't':
		goto yystate11
	}

yystate11:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'h':
		goto yystate12
	}

yystate12:
	c = this.next()
	goto yyrule4

yystate13:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 't':
		goto yystate12
	}

yystate14:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == 'n':
		goto yystate16
	case c == 't':
		goto yystate11
	}

yystate15:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == 'r':
		goto yystate16
	case c == 't':
		goto yystate11
	}

yystate16:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate12
	}

yystate17:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '6':
		goto yystate18
	}

yystate18:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate19
	}

yystate19:
	c = this.next()
	goto yyrule7

yystate20:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '2' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate10
	case c == ':':
		goto yystate17
	case c == 's':
		goto yystate13
	case c == 't':
		goto yystate11
	}

yystate21:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == ':':
		goto yystate17
	case c == 'n':
		goto yystate16
	case c == 't':
		goto yystate11
	}

yystate22:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == ':':
		goto yystate17
	case c == 'r':
		goto yystate16
	case c == 't':
		goto yystate11
	}

yystate23:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '2' && c <= '9':
		goto yystate7
	case c == '1':
		goto yystate24
	case c == 's':
		goto yystate13
	}

yystate24:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '2' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate10
	case c == ':':
		goto yystate17
	case c == 't':
		goto yystate11
	}

yystate25:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c == '4':
		goto yystate7
	case c == '1':
		goto yystate20
	case c == '2':
		goto yystate21
	case c == '3':
		goto yystate22
	case c == 'n':
		goto yystate16
	case c >= '5' && c <= '9':
		goto yystate8
	}

yystate26:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == '0' || c >= '4' && c <= '9':
		goto yystate8
	case c == '1':
		goto yystate9
	case c == '2':
		goto yystate14
	case c == '3':
		goto yystate15
	case c == 'r':
		goto yystate16
	}

yystate27:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'p':
		goto yystate28
	case c == 'u':
		goto yystate32
	}

yystate28:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate29
	}

yystate29:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'i':
		goto yystate30
	}

yystate30:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'l':
		goto yystate31
	}

yystate31:
	c = this.next()
	goto yyrule6

yystate32:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'g':
		goto yystate33
	}

yystate33:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'u':
		goto yystate34
	}

yystate34:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate35
	}

yystate35:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 't':
		goto yystate31
	}

yystate36:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate37
	}

yystate37:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'c':
		goto yystate38
	}

yystate38:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'e':
		goto yystate39
	}

yystate39:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'm':
		goto yystate40
	}

yystate40:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'b':
		goto yystate41
	}

yystate41:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate42
	}

yystate42:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate31
	}

yystate43:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'v':
		goto yystate44
	}

yystate44:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate45
	}

yystate45:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate46
	}

yystate46:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'y':
		goto yystate47
	}

yystate47:
	c = this.next()
	goto yyrule8

yystate48:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate49
	case c == 'r':
		goto yystate55
	}

yystate49:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'b':
		goto yystate50
	}

yystate50:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'r':
		goto yystate51
	}

yystate51:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate52
	}

yystate52:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate53
	}

yystate53:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate54
	}

yystate54:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'y':
		goto yystate31
	}

yystate55:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'i':
		goto yystate56
	case c == 'o':
		goto yystate60
	}

yystate56:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 'd':
		goto yystate57
	}

yystate57:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate58
	}

yystate58:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'y':
		goto yystate59
	}

yystate59:
	c = this.next()
	goto yyrule5

yystate60:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'm':
		goto yystate61
	}

yystate61:
	c = this.next()
	goto yyrule9

yystate62:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate63
	case c == 'r':
		goto yystate65
	}

yystate63:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate64
	}

yystate64:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate65
	}

yystate65:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate66
	}

yystate66:
	c = this.next()
	goto yyrule12

yystate67:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate68
	case c == 'u':
		goto yystate70
	}

yystate68:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'n':
		goto yystate69
	}

yystate69:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'u':
		goto yystate52
	}

yystate70:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'l':
		goto yystate71
	case c == 'n':
		goto yystate72
	}

yystate71:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'y':
		goto yystate31
	}

yystate72:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'e':
		goto yystate31
	}

yystate73:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate74
	case c == 'i':
		goto yystate77
	case c == 'o':
		goto yystate83
	}

yystate74:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate75
	case c == 'y':
		goto yystate31
	}

yystate75:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'c':
		goto yystate76
	}

yystate76:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'h':
		goto yystate31
	}

yystate77:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'n':
		goto yystate78
	}

yystate78:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate79
	case c == 'u':
		goto yystate80
	}

yystate79:
	c = this.next()
	goto yyrule13

yystate80:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 't':
		goto yystate81
	}

yystate81:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate82
	}

yystate82:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate79
	}

yystate83:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'n':
		goto yystate56
	}

yystate84:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate85
	}

yystate85:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'v':
		goto yystate38
	}

yystate86:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'c':
		goto yystate87
	case c == 'f':
		goto yystate89
	}

yystate87:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 't':
		goto yystate88
	}

yystate88:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 'o':
		goto yystate40
	}

yystate89:
	c = this.next()
	goto yyrule11

yystate90:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate91
	case c == 'e':
		goto yystate95
	case c == 'u':
		goto yystate83
	}

yystate91:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 't':
		goto yystate92
	}

yystate92:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 'u':
		goto yystate93
	}

yystate93:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'r':
		goto yystate94
	}

yystate94:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate57
	}

yystate95:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'c':
		goto yystate96
	case c == 'p':
		goto yystate101
	}

yystate96:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate97
	case c == 's':
		goto yystate100
	}

yystate97:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'n':
		goto yystate98
	}

yystate98:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate99
	}

yystate99:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate100
	}

yystate100:
	c = this.next()
	goto yyrule14

yystate101:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c == 't':
		goto yystate38
	}

yystate102:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'h':
		goto yystate103
	case c == 'o':
		goto yystate106
	case c == 'u':
		goto yystate107
	}

yystate103:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate104
	}

yystate104:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 'r':
		goto yystate105
	}

yystate105:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 's':
		goto yystate94
	}

yystate106:
	c = this.next()
	goto yyrule10

yystate107:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate108
	}

yystate108:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 's':
		goto yystate94
	}

yystate109:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate110
	}

yystate110:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'd':
		goto yystate111
	}

yystate111:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 'n':
		goto yystate112
	}

yystate112:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate105
	}

yyrule1: // [\n]+
	{
		fmt.Println("returning newline")
		return NEWLINE
		goto yystate0
	}
yyrule2: // [ \t\n\r]+

	goto yystate0
yyrule3: // {D}+
	{
		if lval.item, err = this.getInt(); nil != err {
			return -1
		} else {
			return DECIMAL_DIGIT
		}

		goto yystate0
	}
yyrule4: // {ORD}
	{
		if lval.item, err = this.getOrdinal(); nil != err {
			return -1
		} else {
			return ORDINAL
		}
		goto yystate0
	}
yyrule5: // {DY}
	{
		lval.item = this.getDay()
		return DAY
		goto yystate0
	}
yyrule6: // {MN}
	{
		lval.item = this.getMonth()
		return MONTH

		goto yystate0
	}
yyrule7: // {H}":"{M}
	{
		lval.item = this.getStr()
		return TIME

		goto yystate0
	}
yyrule8: // "every"
	{
		lval.item = this.getStr()
		return EVERY
		goto yystate0
	}
yyrule9: // "from"
	{
		lval.item = this.getStr()
		return FROM
		goto yystate0
	}
yyrule10: // "to"
	{
		lval.item = this.getStr()
		return TO
		goto yystate0
	}
yyrule11: // "of"
	{
		lval.item = this.getStr()
		return OF
		goto yystate0
	}
yyrule12: // "hours"|"hrs"
	{
		lval.item = this.getClockLbl()
		return HRS
		goto yystate0
	}
yyrule13: // "minutes"|"mins"
	{
		lval.item = this.getClockLbl()
		return MINS
		goto yystate0
	}
yyrule14: // "seconds"|"secs"
	{
		lval.item = this.getClockLbl()
		return SECS
		goto yystate0
	}
yyrule15: // ","
	{
		lval.item = ","
		return COMMA
		goto yystate0
	}
yyrule16: // \0
	{
		return -1
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized

	return -1

} // ends main()
