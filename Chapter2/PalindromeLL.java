import myLinkedList.*;
import java.util.*;
public class PalindromeLL {
	int len = 0;
	private HashMap buildMap(MyNode n) {
		HashMap<Integer,Boolean> m = new HashMap<Integer, Boolean>();

		while(n != null) {
			len++;
			if(!m.containsKey(n.data)) {
				m.put(n.data, true);
			}
			else {
				Boolean val = m.get(n.data);
				val = !val;
				m.remove(n.data);
				m.put(n.data, val);
			}
			n = n.next;
		}
		return m;
	}

	private Boolean isPalindrome(HashMap h)  {
		Boolean result = true;
		int buf = 0;
		// Get an iterator
		Set set = h.entrySet();
		Iterator i = set.iterator();
		// Display elements
		while(i.hasNext()) {

			Map.Entry me = (Map.Entry)i.next();
//			System.out.println(me.getValue());
//			System.out.println(me.getKey());
			if((Boolean)me.getValue() == true) {
				buf++;
				if(len % 2 == 0) {
					result = false;
				}
				else if(buf > 1){
					result = false;	
				}
				else {
					result = true;
				}
			}			
		}
		return result;
	}
	public static void main(String[] args) {
		MyLinkedList ll1 = new MyLinkedList();
		PalindromeLL pl = new PalindromeLL();
		ll1.addLast(9);
		ll1.addLast(1);
		ll1.addLast(5);
		ll1.addLast(6);
		ll1.addLast(1);
		ll1.addLast(9);
		Boolean res = pl.isPalindrome(pl.buildMap(ll1.head));
		System.out.println(res);
	}
}










