
function padIt(str,n){
  while(n>0){
    if(n%2 === 0) {
      str = str + "*";
      
    }else{
      str = "*" + str;
    }
    n --;
  }
    return str
  }

//also can be execute like that
function padIt(str,n){
  var fois = "*"
  var i = 0
do{
  str=fois+str
  i++
  if (i==n) return str
  str=str+fois
  i++
}while (n>i)
return str
}