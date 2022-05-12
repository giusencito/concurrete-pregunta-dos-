package main



func valid(board [] int,row int,col int ) bool{

	for i := 0;i< row;i++{
       var col_i=board[i]
	   var delta=row-i
       var check [3] int
	   check[0]=col_i
	   check[1]=col_i+delta
	   check[2]=col_i-delta
       for i := 0;i<len(check);i++{
           if(check[i]==col){
			   return false
		   }
	   }


	}
return true

}
func nqueens(board [] int, row int){

     var n int = len(board)
	 if(row==n){

	 }else{

          for i := 0;i<n;i++{

			if(valid(board,row,i)){
				board[row]=i
				nqueens(board,row+1)
			}
			  

		  }


	 }




}
func nqueenscocnurrent(board [] int, row int){

	var n int = len(board)
	if(row==n){

	}else{

		 for i := 0;i<n;i++{

		   if(valid(board,row,i)){
			   board[row]=i
			    go nqueenscocnurrent(board,row+1)
		   }
			 

		 }


	}




}


func main()  {
	
  board :=[]int{-1,-1.-1,-1}
  nqueens(board,0)
  nqueenscocnurrent(board,0)

}