OUT=lcs

# default option to make 
default : build 

build :
	go build -o $(OUT)

test : build
	./$(OUT) -s1=aaaabbbbcccc -s2=xxxxbbbbyyyy

clear :
	$(RM) $(OUT)

