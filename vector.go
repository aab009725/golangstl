package mystl
type Vector struct {
	grow int
	size int
	data[] interface{}
}
func (self *Vector)del(num int){
	self.setvalue(num,nil)
	self.size--
}
func (self *Vector)append(value interface{}){
	if self.size+1==len(self.data){
		self.begrow()
	}
	self.setvalue(self.size+1,value)
	self.size++
}
func (self *Vector)insert(num int,value interface{}){
	if self.size+1==len(self.data) {
		self.begrow()
	}
	if num<self.size {
		a := self.data[num : len(self.data)-1]
		b := len(self.data)
		self.data[num] = value
		//self.data[num+1:len(self.data)-1]=a
		for pass := num + 1; pass < b; pass++ {
			self.data[pass+num+1] = a[pass]
		}
	}else{
		self.setvalue(num,value)
	}
	self.size++
}
func(self *Vector)setvalue(num int,value interface{}){
	self.data[num]=value
}
func (self *Vector)setgrow(num int){
	self.grow=num
}
func (self *Vector)getvalue(num int)interface{}{
	return self.data[num]
}
func(self *Vector)begrow(){
	mylen:=len(self.data)
	a:=make([]interface{},mylen+self.grow)
	for i:=0;i<mylen;i++{
		a[i]=self.data[i]
	}
	self.data=a
}
func makevector(size int)Vector{
	var re Vector
	re.grow=10
	re.data=make([]interface{},size)
	re.size=0
	return re
}
func copyvectorwitharray(indata[] interface{})Vector{
	var re Vector
	num:=0
	re.data=make([]interface{},len(indata)+re.grow)
	for _,value:=range indata {
		re.data[num] = value
		num++
	}
	re.size=num+1
	return re
}