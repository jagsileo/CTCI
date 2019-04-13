public class Multiply {
	public int multiply(int n, int m) {
	
		
		if(n==1)
			return m;
		if(m==1)
			return n;
		if(n == 0 || m == 0)
			return 0;
		int result = 0;
		int sign = ((n > 0 && m < 0) || (n<0 && m > 0)) ? -1 : 1;	
		if(Math.abs(m) < Math.abs(n)) 
			result = Math.abs(m) + multiply(Math.abs(n), Math.abs(m-1));
		else if(Math.abs(n) < Math.abs(m))
			result = Math.abs(n) + multiply(Math.abs(m), Math.abs(n-1));
		return result/sign;		
	
	}

}
