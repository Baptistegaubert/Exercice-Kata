function findNb(m) {
  let n=0; 
  let t;
  for(t=0;t<m;){
      n++
      t+=n**3
    }
  if (t==m){
    return n 
  }return -1
  }

