%{
	
package gen

import "fmt"

func (this *SchedLex) Lex(lval *SchedSymType) (ret int) { // begin main()

	var err error
	var c byte = ' '
	
	defer func(){
		this.back()
	}()


%}

D			[0-9]
H			[01]{D}|[2][0-4]
M			[0-6]{D}

DY			"monday"|"mon"|"tuesday"|"tue"|"wednesday"|"wed"|"thursday"|"thu"|"friday"|"fri"|"saturday"|"sat"|"sunday"|"sun"
MN 			"january"|"jan"|"february"|"feb"|"march"|"mar"|"april"|"apr"|"may"|"may"|"june"|"jun"|"july"|"jul"|"august"|"aug"|"september"|"sep"|"sept"|"october"|"oct"|"november"|"nov"|"december"|"dec"

ORD1		{D}*[02-9]+
ORD2		[4-9]+
ORD			{ORD2}"th"|{ORD1}*1"st"|{ORD1}*2"nd"|{ORD1}*3"rd"|{D}+{D}"th"

%yyc c
%yyn c = this.next()

%%
		if nil!=this.buf {
			this.buf = this.buf[len(this.buf)-1:]
		}

[\n]+				fmt.Println("returning newline"); return NEWLINE

[ \t\n\r]+	

{D}+				if lval.item, err = this.getInt(); nil!=err {
						return -1
					} else {
						return DECIMAL_DIGIT
					}
					
{ORD}				if lval.item, err = this.getOrdinal(); nil!=err {
						return -1
					} else {
						return ORDINAL		
					}

{DY}		lval.item = this.getDay(); return DAY
{MN}		lval.item = this.getMonth(); return MONTH

					
{H}":"{M}		lval.item = this.getStr(); return TIME
				
"every"				lval.item = this.getStr(); return EVERY
"from"				lval.item = this.getStr(); return FROM
"to"				lval.item = this.getStr(); return TO
"of"				lval.item = this.getStr(); return OF
"hours"|"hrs"			lval.item = this.getClockLbl(); return HRS
"minutes"|"mins"		lval.item = this.getClockLbl(); return MINS
"seconds"|"secs"		lval.item = this.getClockLbl(); return SECS
","						lval.item = ","; return COMMA

\0					return -1

%%

return -1

} // ends main()
