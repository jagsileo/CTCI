 import java.util.*;


public class FindKtoN {
		
	private static void findkToN(LinkedList ll, int k) {
	    int size = ll.size();
	    if(k > size) {
	        System.out.println("Invalid k value");
	    }
	    for(int i = 0; i<size; i++) {
	        if(i>=k-1) {
	            System.out.println(ll.get(i));
	    }
	}
	}
	
	
	public static void main(String args[]){
		Scanner sc = new Scanner(System.in);
		int k = sc.nextInt();
		LinkedList<Integer> ll = new LinkedList<Integer>();
		ll.addLast(1);
		ll.addLast(2);
		ll.addLast(10);
		ll.addLast(10);
		ll.addLast(1);
		ll.addLast(5);
		ll.addLast(19);
		
		System.out.println(ll);
		findkToN(ll,k);
		
		
	}
}
