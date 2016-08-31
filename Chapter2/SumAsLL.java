import myLinkedList.*;
import java.util.*;
public class SumAsLL {

	private MyNode sumLL(MyNode n1, MyNode n2, int carry) {
		if(n1 == null && n2 == null && carry == 0) {
			return null;
		}
		MyNode result = new MyNode();
		int sum = carry;
		if(n1 != null) {
			sum+= n1.data;
		}
		if(n2 != null) {
			sum += n2.data;
		}
		result.data = sum%10;
		if(n1 != null && n2 != null) {
			MyNode nextNode = sumLL(n1 == null ? null : n1.next,
					n2 == null ? null : n2.next,
					sum >= 10 ? 1 : 0);
			result.next = nextNode;
					
		}
		return result;

	}

	public static void main(String[] args) {
		MyLinkedList ll1 = new MyLinkedList();
		MyLinkedList ll2 = new MyLinkedList();
		SumAsLL s = new SumAsLL();
		ll1.addLast(9);
		ll1.addLast(1);
		ll1.addLast(5);
		ll2.addLast(5);
		ll2.addLast(7);
		ll2.addLast(3);
		ll2.addLast(8);	
		ll1.printList(ll1.head);
		System.out.println("");
		ll2.printList(ll2.head);
		System.out.println("");
		MyNode res = s.sumLL(ll1.head, ll2.head, 0);		
		while(res != null) {
			System.out.println(res.data);
			res = res.next;
		}
	}
}


































