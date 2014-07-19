all: sched.yy.go sched.y sched.y.go

sched.yy.go: sched.l
	golex -t -o sched.yy.go sched.l | gofmt > sched.yy.go 
	
sched.y.go: sched.y
	if [ -e sched.y.go ];then cp -vf sched.y.go sched.y.go.bak; fi
	go tool yacc -p Sched -o sched.y.go sched.y
	
#sched.y: sched.ebnf
#	if [ -e sched.y ];then cp -vf sched.y sched.y.bak; fi
#	ebnf2y -pkg gen -start Sched sched.ebnf > sched.y
	
clean:
	#rm -f sched.yy.go sched.y.go sched.y y.output
	rm -f sched.yy.go



